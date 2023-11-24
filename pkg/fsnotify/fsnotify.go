package fsnotify

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/store"
)

type watcher struct {
	store  store.Store
	logger *logrus.Logger
}

func NewWatcher(store store.Store, logger *logrus.Logger) Watcher {
	return &watcher{
		store:  store,
		logger: logger,
	}
}

type Watcher interface {
	Watch() error
}

func (w *watcher) Watch() error {
	path, err := create("files")
	if err != nil {
		return err
	}

	// Process existing ZIP files in the directory
	existingFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range existingFiles {
		if filepath.Ext(file.Name()) == ".zip" {
			if err := w.readZip(filepath.Join(path, file.Name())); err != nil {
				return err
			}
		}
	}

	// Set up the file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Path to the folder containing the zip folder
	err = watcher.Add(path)
	if err != nil {
		return err
	}

	done := make(chan bool)
	errChan := make(chan error)

	go func() {
		defer close(done)
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				switch {
				case event.Op&fsnotify.Create == fsnotify.Create:
					fmt.Printf("New file detected: %s\n", event.Name)
					if filepath.Ext(event.Name) == ".zip" {
						time.Sleep(3 * time.Second)
						go func(name string) {
							for {
								// Try to open the file in exclusive mode
								file, err := os.OpenFile(name, os.O_RDWR|os.O_EXCL, 0666)
								if err != nil {
									// If open fails, wait for a bit and try again
									time.Sleep(500 * time.Millisecond)
									continue
								}
								file.Close() // Close the file if successfully opened

								// Process the file after successful open
								if err := w.readZip(name); err != nil {
									errChan <- err
								}
								break // Break the loop once the file is processed
							}
						}(event.Name)
					}
				case event.Op&fsnotify.Remove == fsnotify.Remove:
					fmt.Printf("File deleted: %s\n", event.Name)
					if filepath.Ext(event.Name) == ".zip" {
						err := w.isDeleted(event.Name)
						if err != nil {
							errChan <- err
						}
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				errChan <- err
			}
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-done:
		// Watcher stopped
	}

	return nil
}

func (w *watcher) readZip(zipPath string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	fmt.Printf("Contents of %s:\n", zipPath)
	for _, f := range r.File {
		if filepath.Ext(f.Name) == ".csv" {
			fmt.Println(filepath.Base(zipPath) + "/" + filepath.Base(f.Name))
			err := w.store.InsertZipRecord(filepath.Base(zipPath), filepath.Base(f.Name))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *watcher) isDeleted(zipPath string) error {
	fileName := filepath.Base(zipPath)
	fmt.Printf("Zip file deleted: %s\n", fileName)
	err := w.store.DeleteZipRecord(fileName)
	if err != nil {
		return err
	}
	return nil
}

func create(p string) (string, error) {
	files := p
	_, err := os.Stat(files)
	if os.IsNotExist(err) {
		fmt.Println("creating ", files)
		err := os.Mkdir(files, 0755)
		if err != nil {
			return "", err
		}
	}

	zipped := files + "/zip"
	_, err = os.Stat(zipped)
	if os.IsNotExist(err) {
		fmt.Println("creating ", zipped)
		err := os.Mkdir(zipped, 0755)
		if err != nil {
			return "", err
		}
	}

	return zipped, nil
}

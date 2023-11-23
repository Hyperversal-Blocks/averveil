package fsnotify

import (
	"archive/zip"
	"fmt"
	"os"
	"path/filepath"

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
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Printf("New file detected: %s\n", event.Name)
					if filepath.Ext(event.Name) == ".zip" {
						go func(name string) {
							if err := w.readZip(name); err != nil {
								errChan <- err
							}
						}(event.Name)
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
		//
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
		fmt.Println(f.Name)
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

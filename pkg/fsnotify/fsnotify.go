package fsnotify

import (
	"archive/zip"
	"fmt"
	"log"
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
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Printf("New file detected: %s\n", event.Name)
					if filepath.Ext(event.Name) == ".zip" {
						go readZip(event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Path to the folder containing the zip folder
	err = watcher.Add(path)
	if err != nil {
		return err
	}
	<-done
	return nil
}

func readZip(zipPath string) error {
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

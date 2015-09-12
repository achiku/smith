package smith

import (
	"log"
	"os"

	"github.com/howeyc/fsnotify"
)

type Watcher struct {
	watcher  *fsnotify.Watcher
	filePath string
}

func NewWatcher(filePath string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	file, err := NewTargetFile(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Event:
				if event.IsRename() || event.IsDelete() {
					log.Println("file closed")
					file.Close()
				}
				if event.IsModify() {
					log.Println("modified")
					err = file.BufferedLineRead()
					if err != nil {
						log.Printf("%s", err)
					}
				}
				if event.IsCreate() {
					log.Println("created")
				}
			}
		}
	}()

	err = watcher.Watch(filePath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	<-done
	return nil
}

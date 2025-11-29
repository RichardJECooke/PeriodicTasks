package fileLogic

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"

	store "github.com/RichardJECooke/PeriodicTasks/src/3store"
)

var timeDataFileLastRead = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
var pauseFileWatcherChannel = make(chan bool)
var filenameFileWatcherChannel = make(chan string)

func HandleWindowRestored() {
	ReadDataFileIfChangeSinceLastRead(timeDataFileLastRead)
	pauseFileWatcherChannel <- false
}

func HandleWindowMinimized() {
	pauseFileWatcherChannel <- true
}

func WatchDataFileForChangesAndReload() {
	dataFilePath := store.GetStore().Config.DataFilePath
	debouncedReadDataFileIfChangeSinceLastRead, cancelDebouncedReadDataFileIfChangeSinceLastRead := lo.NewDebounce(3000*time.Millisecond, ReadDataFileIfChangeSinceLastRead)

	cancelDebouncedReadDataFileIfChangeSinceLastRead()

	for { // outer loop to restart watcher on path change
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatalf("Error creating file watcher: %v", err)
		}
		if err := watcher.Add(dataFilePath); err != nil {
			log.Fatalf("Error watching data file: %v", err)
		}
		pathChanged := false
		for { // inner loop to watch current path
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					continue
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					if pauseFileWatcher.Load() {
						continue
					}
					if err := ReadDataFile(); err != nil {
						log.Printf("Error reloading data file in file watcher: %v", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if ok {
					log.Printf("Data file watcher error: %v", err)
				}
			case newDataFilePath, ok := <-changeDataFilePathChannel:
				if !ok {
					watcher.Close()
					return
				}
				if dataFilePath != newDataFilePath {
					log.Printf("Data file path changed from %s to %s. Restarting watcher.", dataFilePath, newDataFilePath)
					watcher.Close()
					pathChanged = true
					dataFilePath = newDataFilePath
				}
			}
			if pathChanged {
				break
			}
		}
		if pathChanged {
			continue
		}
		return
	}
}

package fileLogic

import (
	"log"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"

	store "github.com/RichardJECooke/PeriodicTasks/src/3store"
)

var pauseFileWatcherChannel = make(chan bool)
var filenameChangedChannel = make(chan string)
var requestFileReadTimeChannel = make(chan bool)
var sendFileReadTimeChannel = make(chan time.Time)

func HandleWindowRestored() {
	requestFileReadTimeChannel <- true
	timeDataFileLastRead := <-sendFileReadTimeChannel
	ReadDataFileIfChangeSinceLastRead(timeDataFileLastRead)
	PauseFileWatcher(false)
}

func HandleWindowMinimized() {
	PauseFileWatcher(true)
}

func PauseFileWatcher(shouldPause bool) {
	pauseFileWatcherChannel <- shouldPause
}

func SetFileWatcherFilename(filename string) {
	filenameChangedChannel <- filename
}

func WatchDataFileForChangesAndReload() {
	timeDataFileLastRead := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	datafilePath := store.GetStore().Config.DataFilePath
	readDataDebounced, _ := lo.NewDebounce(3000*time.Millisecond, ReadDataFileIgnoreError)
	watcher := createWatcher(datafilePath)
	for {
		select {
		case <-requestFileReadTimeChannel:
			sendFileReadTimeChannel <- timeDataFileLastRead
		case newDatafilePath := <-filenameChangedChannel:
			if newDatafilePath != datafilePath {
				datafilePath = newDatafilePath
				watcher.Close()
				createWatcher(datafilePath)
			}
		case shouldPauseWatcher := <-pauseFileWatcherChannel:
			if shouldPauseWatcher {
				watcher.Close()
			} else {
				if err := watcher.Close(); err != nil {
					// already closed
				}
				watcher = createWatcher(datafilePath)
			}
		case event, ok := <-watcher.Events:
			if !ok {
				continue
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				readDataDebounced()
			}
		case err, ok := <-watcher.Errors:
			if ok {
				log.Printf("Data file watcher error: %v", err)
			}
		}
	}
}

func createWatcher(datafilePath string) *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Error creating file watcher: %v", err)
	}
	if err := watcher.Add(datafilePath); err != nil {
		log.Fatalf("Error watching data file: %v", err)
	}
	return watcher
}

func ReadDataFileIgnoreError() {
	err := ReadDataFile()
	if err == nil {
		log.Printf("Error reloading data file in file watcher: %v", err)
	}
}

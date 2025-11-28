package src

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"

	"github.com/fsnotify/fsnotify"

	types "github.com/RichardJECooke/PeriodicTasks/src/0types"
	constants "github.com/RichardJECooke/PeriodicTasks/src/1constants"
	store "github.com/RichardJECooke/PeriodicTasks/src/3store"
)

var pauseFileWatcher atomic.Bool
var changeDataFilePathChannel = make(chan string)

func Start() {
	setupConfigFile()
	if err := ReadDataFile(); err != nil {
		WriteDataFile()
	}
	store.RegisterForStoreChanged(handleStoreChanged)
	go watchDataFileForChangesAndReload()
}

func watchDataFileForChangesAndReload() {
	dataFilePath := store.GetStore().Config.DataFilePath
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
					watcher.Remove(dataFilePath)
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

func handleStoreChanged() {
	changeDataFilePathChannel <- store.GetStore().Config.DataFilePath
	WriteConfigFile()
	WriteDataFile()
}

func WriteDataFile() {
	if store.GetStore().Config.DataFilePath == "" {
		log.Fatalf("Data file path does not exist when writing data")
	}
	jsonData, err := json.MarshalIndent(store.GetStore().TaskGroups[0], "", "    ")
	if err != nil {
		log.Fatalf("Fatal converting objects to JSON writing data file: %v", err)
	}
	pauseFileWatcher.Store(true)
	if err = os.WriteFile(store.GetStore().Config.DataFilePath, jsonData, constants.Permission_RWX_RX_RX); err != nil {
		log.Fatalf("Fatal error writing data file: %v", err)
	}
	go func() {
		time.Sleep(1 * time.Second)
		pauseFileWatcher.Store(false)
	}()
}

func WriteConfigFile() {
	jsonData, err := json.MarshalIndent(store.GetStore().Config, "", "    ")
	if err != nil {
		log.Fatalf("Fatal error converting objects to JSON writing config file: %v", err)
	}
	if err = os.WriteFile(store.GetStore().Config.ConfigFilePath, jsonData, constants.Permission_RWX_RX_RX); err != nil {
		log.Fatalf("Fatal error writing config file: %v", err)
	}
}

func ReadDataFile() error {
	dataText, err := os.ReadFile(store.GetStore().Config.DataFilePath)
	if err != nil {
		return errors.New("data file path does not exist when reading data")
	}
	if len(dataText) == 0 {
		return nil
	}
	var taskGroup types.TaskGroup
	if err := json.Unmarshal(dataText, &taskGroup); err != nil {
		log.Fatalf("Data file has invalid JSON: %v", err)
	}
	store.SetTaskGroup(taskGroup)
	return nil
}

func setupConfigFile() {
	configFolderPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Fatal error getting config path: %v", err)
	}
	tempConfig := store.GetStore().Config
	tempConfig.ConfigFilePath = filepath.Join(configFolderPath, constants.ConfigPathExtensionAndFileName)
	tempConfig.DataFilePath = filepath.Join(configFolderPath, constants.DefaultDataPathExtensionAndFileName)
	if !DoesFolderExist(configFolderPath) {
		if err = os.MkdirAll(configFolderPath, constants.Permission_RWX_RX_RX); err != nil {
			log.Fatalf("Fatal error creating config directory: %v", err)
		}
	}
	if !DoesFileExist(tempConfig.ConfigFilePath) {
		WriteConfigFile()
	}
	configDataText, err := os.ReadFile(tempConfig.ConfigFilePath)
	if err != nil {
		log.Fatalf("Fatal error reading config JSON: %v", err)
	}
	if err = json.Unmarshal(configDataText, &tempConfig); err != nil {
		log.Fatalf("Fatal error parsing config JSON: %v", err)
	}
	store.SetConfig(tempConfig)
}

func DoesFolderExist(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && fileInfo.IsDir()
}

func DoesFileExist(path string) bool {
	_, err := os.ReadFile(path)
	return err == nil
}

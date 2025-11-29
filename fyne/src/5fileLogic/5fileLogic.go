package fileLogic

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"sync/atomic"
	"time"

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
	go WatchDataFileForChangesAndReload()
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

func ReadDataFileIfChangeSinceLastRead(timeDataFileLastRead time.Time) {
	err, timeFileChanged := getFileChangedTime(store.GetStore().Config.DataFilePath)
	if err != nil {
		if timeFileChanged.After(timeDataFileLastRead) {
			ReadDataFile()
		}
	}
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

func getFileChangedTime(filePath string) (error, time.Time) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Println("Error getting time updated of file:", err)
		return err, time.Time{}
	}
	return nil, fileInfo.ModTime()
}

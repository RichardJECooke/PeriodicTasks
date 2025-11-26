package src

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"

	types "github.com/RichardJECooke/PeriodicTasks/src/0types"
	constants "github.com/RichardJECooke/PeriodicTasks/src/1constants"
)

func Startup() {
	setupConfigFile()
	// TODO  _watch(() => _store.config, async (config) => { await writeConfigFile(); }, { deep: true });
	if err := ReadDataFile(); err != nil {
		WriteDataFile()
	}
	// TODO  _watch(() => _store.taskGroups,  async (tasks)  => { await writeDataFile();   }, { deep: true });
}

func WriteConfigFile() {
	jsonData, err := json.MarshalIndent(Store.Config, "", "    ")
	if err != nil {
		log.Fatalf("Fatal error converting objects to JSON writing config file: %v", err)
	}
	if err = os.WriteFile(Store.Config.ConfigFilePath, jsonData, constants.Permission_RWX_RX_RX); err != nil {
		log.Fatalf("Fatal error writing config file: %v", err)
	}
}

func ReadDataFile() error {
	dataText, err := os.ReadFile(Store.Config.DataFilePath)
	if err != nil {
		return errors.New("data file path does not exist when reading data")
	}
	if len(dataText) == 0 {
		return nil
	}
	var taskGroup types.TTaskGroup
	if err := json.Unmarshal(dataText, &taskGroup); err != nil {
		log.Fatalf("Data file has invalid JSON: %v", err)
	}
	SetTaskGroup(taskGroup)
	return nil
}

func WriteDataFile() {
	if Store.Config.DataFilePath != "" {
		log.Fatalf("Data file path does not exist when writing data")
	}
	jsonData, err := json.MarshalIndent(Store.TaskGroups[0], "", "    ")
	if err != nil {
		log.Fatalf("Fatal converting objects to JSON writing data file: %v", err)
	}
	if err = os.WriteFile(Store.Config.DataFilePath, jsonData, constants.Permission_RWX_RX_RX); err != nil {
		log.Fatalf("Fatal error writing data file: %v", err)
	}
	// TODO watch file for changes and reload
}

func setupConfigFile() {
	configFolderPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Fatal error getting config path: %v", err)
	}
	Store.Config.ConfigFilePath = filepath.Join(configFolderPath, constants.ConfigPathExtensionAndFileName)
	Store.Config.DataFilePath = filepath.Join(configFolderPath, constants.DefaultDataPathExtensionAndFileName)
	if !DoesFolderExist(configFolderPath) {
		if err = os.MkdirAll(configFolderPath, constants.Permission_RWX_RX_RX); err != nil {
			log.Fatalf("Fatal error creating config directory: %v", err)
		}
	}
	if !DoesFileExist(Store.Config.ConfigFilePath) {
		WriteConfigFile()
	}
	configDataText, err := os.ReadFile(Store.Config.ConfigFilePath)
	if err != nil {
		log.Fatalf("Fatal error reading config JSON: %v", err)
	}
	if err = json.Unmarshal(configDataText, &Store.Config); err != nil {
		log.Fatalf("Fatal error parsing config JSON: %v", err)
	}
}

func DoesFolderExist(path string) bool {
	fileInfo, err := os.Stat(path)
	return err == nil && fileInfo.IsDir()
}

func DoesFileExist(path string) bool {
	_, err := os.ReadFile(path)
	return err == nil
}

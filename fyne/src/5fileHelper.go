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

const Permission_RWX_RX_RX os.FileMode = 0755

var _configFilePath string

func Startup() {
	setupConfigFile()
	// TODO  _watch(() => _store.config, async (config) => { await writeConfigFile(); }, { deep: true });
	err := ReadDataFile()
	if err != nil {
		WriteDataFile()
	}
	// TODO  _watch(() => _store.taskGroups,  async (tasks)  => { await writeDataFile();   }, { deep: true });
}

func WriteConfigFile() {
	//  writeFile(_configFolderPath + '/' + _constants.configFilePath, JSON.stringify(_store.config, null, 4));
	jsonData, err := json.MarshalIndent(Store.Config, "", "    ")
	if err != nil {
		log.Fatalf("Fatal converting objects to JSON writing config file: %v", err)
	}
	err = os.WriteFile(Store.Config.DataFilePath, jsonData, Permission_RWX_RX_RX)
	if err != nil {
		log.Fatalf("Fatal error writing config file: %v", err)
	}
}

func ReadDataFile() error {
	dataText, err := os.ReadFile(Store.Config.DataFilePath)
	if err != nil {
		return errors.New("Data file path does not exist when reading data")
	}
	if len(dataText) == 0 {
		return nil
	}
	var taskGroup types.TTaskGroup
	err := json.Unmarshal(dataText, &taskGroup)
	if err != nil {
		log.Fatalf("Data file has invalid JSON: %v", err)
	}
	TaskHelper.setTaskGroup(taskGroup)
}

func WriteDataFile() {
	if Store.Config.DataFilePath != "" {
		log.Fatalf("Data file path does not exist when writing data")
	}
	jsonData, err := json.MarshalIndent(Store.TaskGroups[0], "", "    ")
	if err != nil {
		log.Fatalf("Fatal converting objects to JSON writing data file: %v", err)
	}
	err = os.WriteFile(Store.Config.DataFilePath, jsonData, Permission_RWX_RX_RX)
	if err != nil {
		log.Fatalf("Fatal error writing data file: %v", err)
	}
	// TODO watch file for changes and reload
}

func SetDataFilePath(path string) {
	Store.Config.DataFilePath = path
}

func setupConfigFile() {
	configFolderPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Fatal error getting config path: %v", err)
	}
	_configFilePath = filepath.Join(configFolderPath, constants.ConfigPathExtensionAndFileName)
	SetDataFilePath(filepath.Join(configFolderPath, constants.DefaultDataPathExtensionAndFileName))
	if !DoesFolderExist(configFolderPath) {
		err = os.MkdirAll(configFolderPath, Permission_RWX_RX_RX)
		if err != nil {
			log.Fatalf("Fatal error creating config directory: %v", err)
		}
	}
	if !DoesFileExist(_configFilePath) {
		WriteConfigFile()
	}
	configDataText, err := os.ReadFile(_configFilePath)
	if err != nil {
		log.Fatalf("Fatal error reading config JSON: %v", err)
	}
	err = json.Unmarshal(configDataText, &Store.Config)
	if err != nil {
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

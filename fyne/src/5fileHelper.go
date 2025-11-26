package src

import (
	"encoding/json"
	"errors"
	"io/fs"
  "encoding/json"
	"log"
	"os"
	"path/filepath"

	types "github.com/RichardJECooke/PeriodicTasks/src/0types"
  constants "github.com/RichardJECooke/PeriodicTasks/src/1constants"
)

const Permission_RWX_RX_RX os.FileMode = 0755
var configFolderPath string

func Startup() nil {
  setupConfigFile();
  // TODO  _watch(() => _store.config, async (config) => { await writeConfigFile(); }, { deep: true });
  catch (e) { console.error('Error message: ' + JSON.stringify(e)); throw e; }
  try { await readDataFile(); }
  catch (e) { await writeDataFile(); }
  _watch(() => _store.taskGroups,  async (tasks)  => { await writeDataFile();   }, { deep: true });
}

func WriteConfigFile() nil {
  writeFile(_configFolderPath + '/' + _constants.configFilePath, JSON.stringify(_store.config, null, 4));
  panic console.error('Error message: ' + JSON.stringify(e)); throw e; }
}

func ReadDataFile() error {
  if Store.Config.DataFilePath == nil {
    log.Fatalf("Data file path does not exist when reading data: %v", err)
  }
  dataText, err := os.ReadFile(Store.Config.DataFilePath)
  if err != nil {
    log.Fatalf("Fatal error reading data file: %v", err)
  }
  err := json.Unmarshal(dataText, &Store.TaskGroups[0])
  if err != nil {
    log.Fatalf("Fatal error parsing data file JSON: %v", err)
  }

    const fileContent = await Neutralino.filesystem.readFile(_store.config.dataFilePath);
    if (!fileContent) return;
    const taskFile: _types.tTaskGroup = JSON.parse(fileContent);
    taskFile.tasks = taskFile.tasks.map((task: any) => {
      task.datesDone = task.datesDone.map((date: _types.tIdAndDate) => ({ id: date.id, date: new Date(date.date) }));
      return task;
    });
    _taskHelper.setTaskFile(taskFile);
  }
  catch (e) { console.log('Cannot read data file, but is not an error at first startup: ' + JSON.stringify(e)); throw e; }
}

func WriteDataFile() {
  if !Store.Config.DataFilePath) {
    log.Fatalf("Data file path does not exist when writing data: %v", err)
  }
  await Neutralino.filesystem.writeFile(Store.Config.DataFilePath, JSON.stringify(Store.TaskGroups[0], null, 4));
  jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("Fatal converting objects to JSON writing data file: %v", err)
	}
	err = os.WriteFile(Store.Config.DataFilePath, jsonData, mode)
	if err != nil {
    log.Fatalf("Fatal error writing data file: %v", err)
  }
  // TODO watch file for changes and reload
}

func SetDataFilePath(path string) {
  Store.Config.DataFilePath = path;
}

func setupConfigFile() {
	configFolderPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Fatal error getting config path: %v", err)
	}
	configFolderPath = filepath.Join(configFolderPath, constants.ConfigPathExtensionAndFileName)
	SetDataFilePath(filepath.Join(configFolderPath, constants.DefaultDataPathExtensionAndFileName))
	if !DoesFolderExist(configFolderPath) {
			err = os.MkdirAll(configFolderPath, Permission_RWX_RX_RX)
			if err != nil {
				log.Fatalf("Fatal error creating config directory: %v", err)
			}
	}
  if !DoesFileExist(configFilePath) {
    writeConfigFile()
  }
	configDataText, err := os.ReadFile(configFilePath)
  if err != nil {
    log.Fatalf("Fatal error reading config JSON: %v", err)
  }}
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
  _, err = os.ReadFile(path)
  return err == nil
}
package src

import (
	"encoding/json"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	types "github.com/RichardJECooke/PeriodicTasks/src/0types"
  constants "github.com/RichardJECooke/PeriodicTasks/src/1constants"
)

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

func ReadDataFile() nil {
  try {
    if (!_store.config.dataFilePath) throw new Error('No data file path specified');
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

func WriteDataFile() nil {
  try {
    if (!_store.config.dataFilePath) throw new Error('No file path specified');
    await Neutralino.filesystem.writeFile(_store.config.dataFilePath, JSON.stringify(_store.taskGroups[0], null, 4));

    // TODO ERROR
    let watcherId = await Neutralino.filesystem.createWatcher(_store.config.dataFilePath); // '/home/me/.config/info.richardcooke.PeriodicTasks/tasks.json'
    // Neutralino.events.on('watchFile', (evt) => {if (watcherId == evt.detail.id) {console.log(evt.detail);}});
    // await Neutralino.filesystem.removeWatcher(watcherId);
  }
  catch (e) { console.error('Error message: ' + JSON.stringify(e)); throw e; }
}

func SetDataFilePath(path string) {
  Store.Config.DataFilePath = path;
}

func setupConfigFile() {
  try {
    _configFolderPath = await Neutralino.os.getPath('config');
    setDataFilePath(_configFolderPath + '/' + _constants.defaultDataFilePath);
    try { await Neutralino.filesystem.readDirectory(_configFolderPath + '/' + _constants.configFolderPath); }
    catch (e) { await Neutralino.filesystem.createDirectory(_configFolderPath + '/' + _constants.configFolderPath); }
    let configDataText: string | null = null;
    try { configDataText = await Neutralino.filesystem.readFile(_configFolderPath + '/' + _constants.configFilePath); }
    catch(e) { await writeConfigFile(); }
    if (configDataText) _store.config = JSON.parse(configDataText);
  }
  catch (e) { console.error('Error message: ' + JSON.stringify(e)); throw e; }
}

func setupConfigFile() {
	configFolderPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Fatal error getting config path: %v", err)
	}
	configFolderPath = filepath.Join(configFolderPath, constants.ConfigFileName)
	SetDataFilePath(filepath.Join(configFolderPath, constants.defaultDataFilePath))
	configDirPath := filepath.Join(configFolderPath, constants.configFolderPath)
	configFilePath := filepath.Join(configFolderPath, constants.configFilePath)
	_, err = os.Stat(configDirPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			err = os.MkdirAll(configDirPath, 0755)
			if err != nil {
				log.Fatalf("Fatal error creating config directory: %v", err)
			}
		} else {
			log.Fatalf("Fatal error checking config directory: %v", err)
		}
	}
	var configDataText []byte
	configDataText, err = os.ReadFile(configFilePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			if writeErr := writeConfigFile(); writeErr != nil {
				log.Fatalf("Fatal error writing default config file: %v", writeErr)
			}
		} else {
			log.Fatalf("Fatal error reading config file: %v", err)
		}
	} else {
		if len(configDataText) > 0 {
			err = json.Unmarshal(configDataText, &_store.config)
			if err != nil {
				log.Fatalf("Fatal error parsing config JSON: %v", err)
			}
		}
	}
}
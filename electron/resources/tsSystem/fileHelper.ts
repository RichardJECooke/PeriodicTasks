import { watch as _watch} from 'vue';
import * as _constants from './1constants.ts';
import * as _types from './0types.ts';
import * as _taskHelper from './4taskHelper.ts';
import {store as _store} from './3store.ts';

let _configFolderPath: string = '';

export async function startup(): Promise<void> {
  try {
    await setupConfigFile();
    _watch(() => _store.config, async (config) => { await writeConfigFile(); }, { deep: true });
  }
  catch (e) { console.error('Error message: ' + JSON.stringify(e)); throw e; }
  try { await readDataFile(); }
  catch (e) { await writeDataFile(); }
  _watch(() => _store.taskGroups,  async (tasks)  => { await writeDataFile();   }, { deep: true });
}

export async function writeConfigFile(): Promise<void> {
  try {
    await Neutralino.filesystem.writeFile(_configFolderPath + '/' + _constants.configFilePath, JSON.stringify(_store.config, null, 4));
  }
  catch (e) { console.error('Error message: ' + JSON.stringify(e)); throw e; }
}

export async function readDataFile(): Promise<void> {
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

export async function writeDataFile(): Promise<void> {
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

export function setDataFilePath(path: string) {
  _store.config.dataFilePath = path;
}

async function setupConfigFile(): Promise<void> {
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
import { watch as _watch} from 'vue';
import * as _constants from './1constants.ts';
import * as _types from './0types.ts';
import * as _taskHelper from './4taskHelper.ts';
import {store as _store} from './3store.ts';

export async function startup(): Promise<void> {
  try {
    await setupConfigFile();
    await readDataFile();
    _watch(() => _store.tasks,  async (tasks)  => { await writeDataFile();   }, { deep: true });
    _watch(() => _store.config, async (config) => { await writeConfigFile(); }, { deep: true });
  }
  catch (e) {
    console.error(e.message);
    console.error(e.stack);
    throw e;
  }
}

export async function readDataFile() {
    await loadTasksFromFile();
}

export async function writeConfigFile(): Promise<void> {
    await Neutralino.filesystem.writeFile(_constants.configFile, JSON.stringify(_store.config, null, 4));
}

export async function writeDataFile(): Promise<void> {
    if (!_store.config.dataFilePath) throw new Error('No file path specified');
    await Neutralino.filesystem.writeFile(_store.config.dataFilePath, JSON.stringify(_store.tasks, null, 4));
}

export async function setDataFilePath(path: string) {
    _store.config.dataFilePath = path;
}

async function loadTasksFromFile(): Promise<void> {
    if (!_store.config.dataFilePath) throw new Error('No data file path specified');
    const fileContent = await Neutralino.filesystem.readFile(_store.config.dataFilePath);
    if (!fileContent) return;
    const tasks: _types.ttask[] = JSON.parse(fileContent);
    const tasksWithDates = tasks.map((task: any) => {
        task.datesDone = task.datesDone.map((date: _types.tidAndDate) => ({ id: date.id, date: new Date(date.date) }));
        return task;
    });
    _taskHelper.setTasks(tasksWithDates);
}

async function setupConfigFile(): Promise<void> {
  await Neutralino.filesystem.createDirectory(_constants.configFolder);
  let configDataText: string | null = null;
  try { configDataText = await Neutralino.filesystem.readFile(_constants.configFile); }
  catch(e) { await writeConfigFile(); }
  if (configDataText) _store.config = JSON.parse(configDataText);
}
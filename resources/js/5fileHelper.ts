import { watch as _watch} from 'vue';
import * as _constants from './1constants.ts';
import * as _types from './0types.ts';
import * as _taskHelper from './4taskHelper.ts';
import {store as _store} from './3store.ts';

const _dataFilePathName = 'dataFilePath';

export async function startup(): Promise<string> {
    try {
    const err = await Neutralino.filesystem.createDirectory(_constants.configFolder);
    todo load config file
    const err2 = await loadPreviouslyUsedDataFile();
    return err + err2;
    }
    catch (e) {
        console.dir(e);
        return (e as Error).message;
    }
}

export async function loadPreviouslyUsedDataFile(): Promise<string> {
    try {
        const dataFilePath = await Neutralino.storage.getData(_dataFilePathName);
        if (!dataFilePath) return;
        _store.dataFilePath = dataFilePath;
        await loadTasksFromFile();
        return '';
    }
    catch (e) {
        console.dir(e);
        return (e as Error).message;
    }
}

export async function openDataFile() {
    try {
        loadTasksFromFile();
    }
    catch (e) {
        console.dir(e);
        return (e as Error).message;
    }
}

export async function saveDataFile() {
    try {
        if (!_store.dataFilePath) throw new Error('No file path specified');
        await Neutralino.filesystem.writeFile(_store.dataFilePath, JSON.stringify(_store.tasks, null, 4));
    }
    catch (e) {
        console.dir(e);
        return (e as Error).message;
    }
}

export async function setDataFilePath(path: string) {
    _store.dataFilePath = path;
    await Neutralino.storage.setData(_dataFilePathName, _store.dataFilePath);
}

async function loadTasksFromFile() {
    if (!_store.dataFilePath) throw new Error('No file path specified');
    const fileContent = await Neutralino.filesystem.readFile(_store.dataFilePath);
    if (!fileContent) return;
    const tasks: _types.task[] = JSON.parse(fileContent);
    const tasksWithDates = tasks.map((task: any) => {
        task.datesDone = task.datesDone.map((date: _types.idAndDate) => ({ id: date.id, date: new Date(date.date) }));
        return task;
    });
    _taskHelper.setTasks(tasksWithDates);
}

_watch(_store, async (newStore, oldStore) => { await saveDataFile(); });
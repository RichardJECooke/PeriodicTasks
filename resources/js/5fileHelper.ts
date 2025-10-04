import * as _constants from './1constants.ts';
import * as _types from './0types.ts';
import * as _taskHelper from './4taskHelper.ts';
import {store as _store} from './3store.ts';

const _dataFilePathName = 'dataFilePath';

export async function loadPreviouslyUsedDataFile() {
    try {
        const dataFilePath = await Neutralino.storage.getData(_dataFilePathName);
        if (!dataFilePath) return;
        _store.dataFilePath = dataFilePath;
        await loadTasksFromFile();
    }
    catch (e) { console.dir(e); }
}

export async function openDataFile() {
    try {
        loadTasksFromFile();
    }
    catch (e) { console.dir(e); }
}

export async function saveDataFile() {
    try {
        if (!_store.dataFilePath) throw new Error('No file path specified');
        await Neutralino.filesystem.writeFile(_store.dataFilePath, JSON.stringify(_store.tasks));
    }
    catch (e) { console.dir(e); }
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
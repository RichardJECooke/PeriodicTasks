import * as _constants from './constants.ts';
import * as _store from './store.ts';

const _dataFilePathName = 'dataFilePath';
let _dataFilePath: string | null = null;

export async function loadPreviouslyUsedDataFile() {
    try {
        _dataFilePath = await Neutralino.storage.getData(_dataFilePathName);
        if (!_dataFilePath) return;
        const fileContent = await Neutralino.filesystem.readFile(_dataFilePath);
        if (!fileContent) return;
        _store.tasks.tasks = JSON.parse(fileContent)?.tasks || [];
    }
    catch (e) { console.dir(e); }
}

export async function openDataFile() {
    try {
        const fileContent = await Neutralino.filesystem.readFile(_dataFilePath);
        _store.tasks.tasks = (JSON.parse(fileContent) as typeof _store.tasks).tasks;
    }
    catch (e) { console.dir(e); }
}

export async function saveDataFile() {
    try {
        await Neutralino.filesystem.writeFile(_dataFilePath, JSON.stringify(_store.tasks));
    }
    catch (e) { console.dir(e); }
}

export async function setDataFilePath(path: string) {
    _dataFilePath = path;
    await Neutralino.storage.setData(_dataFilePathName, _dataFilePath);
}
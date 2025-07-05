import * as _constants from './constants.ts';
import * as _types from './types.ts';
import {store as _store} from './store.ts';

const _dataFilePathName = 'dataFilePath';
let _dataFilePath: string | null = null;

export async function loadPreviouslyUsedDataFile() {
    try {
        _dataFilePath = await Neutralino.storage.getData(_dataFilePathName);
        if (!_dataFilePath) return;
        const fileContent = await Neutralino.filesystem.readFile(_dataFilePath);
        if (!fileContent) return;
        _store.tasks = JSON.parse(fileContent)?.tasks || [];
    }
    catch (e) { console.dir(e); }
}

export async function openDataFile() {
    try {
        if (!_store.dataFilePath) throw new Error('No file path specified');
        const fileContent = await Neutralino.filesystem.readFile(_store.dataFilePath);
        _store.tasks = JSON.parse(fileContent) as _types.task[];
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
    _dataFilePath = path;
    await Neutralino.storage.setData(_dataFilePathName, _dataFilePath);
}
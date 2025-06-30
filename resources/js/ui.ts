import { v4 as _uuid } from 'uuid';
import * as _constants from './constants.ts';
import * as _taskHelper from './taskHelper.ts';
import {store as _store} from './store.ts';
import * as _fileHelper from './fileHelper.ts';

export async function addAddTaskButton() {
    const button = document.getElementById('addTaskButton') as HTMLButtonElement;
    button.onclick = addTask;
}

export async function addTask() {
    _taskHelper.addTask();
    await refreshTasks();
}

export async function saveDataFile() {
    try {
        if (!_store.dataFilePath) chooseSaveFileLocation();
        if (!_store.dataFilePath) return;
        await _fileHelper.saveDataFile();
    }
    catch (e) { console.dir(e); }
}

export async function chooseSaveFileLocation() {
    try {
        const filePath = await Neutralino.os.showSaveDialog('Save your tasks', { "defaultPath": _constants.defaultSavePath });
        _store.dataFilePath = filePath;
        _fileHelper.setDataFilePath(filePath);
    }
    catch (e) { console.dir(e); }
}

export async function openDataFile() {
    try {
        await chooseOpenFileLocation();
        if (!_store.dataFilePath) return;
        _fileHelper.setDataFilePath(_store.dataFilePath);
        await _fileHelper.openDataFile()
        await refreshTasks();
    }
    catch (e) { console.dir(e); }
}

export async function chooseOpenFileLocation() {
    try {
        let entries = await Neutralino.os.showOpenDialog('Load your tasks', {
            defaultPath: _constants.defaultSavePath
            // ,filters: [ {name: 'JSON files', extensions: ['json', 'js']}  ] - does not work, bug
        });
        if (!entries || entries.length == 0) return;
        _store.dataFilePath = entries[0];
    }
    catch (e) { console.dir(e); }
}

async function refreshTasks() {
    const tasksUi = document.getElementById('tasks') as HTMLUListElement;
    tasksUi.innerHTML = ''; 
    for (let task of _store.tasks) {
        const li = document.createElement('li');
        li.id = task.id
        li.textContent = task.name;
        tasksUi.appendChild(li);
    }
}
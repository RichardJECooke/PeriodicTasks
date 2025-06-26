import {v4 as _uuid} from 'uuid';
import * as _constants from './constants.mjs';
import * as _taskHelper from './taskHelper.mjs';
import * as _store from './store.mjs';

export async function addAddTaskButton() {
    const button = document.createElement('button');
    button.textContent = 'Add task';
    button.id = 'addTaskButton';
    button.onclick = addTask;
    document.getElementById('app').appendChild(button);
}

export async function chooseSaveFileLocation() {
    try {
        const filePath = await Neutralino.os.showSaveDialog('Save your tasks', { defaultPath: _defaultPath });
        _fileHelper.setDataFilePath = filePath;
        await Neutralino.storage.setData(_dataFilePathName, _dataFilePath);
    }
    catch (e) { console.dir(e); }
}

export async function chooseOpenFileLocation() {
    try {
        let entries = await Neutralino.os.showOpenDialog('Load your tasks', {
            defaultPath: _constants.defaultPath
            // ,filters: [ {name: 'JSON files', extensions: ['json', 'js']}  ]
        });
        if (!entries || entries.length == 0) return;
        _dataFilePath = entries[0];
    }
    catch (e) { console.dir(e); }
}

async function addTask() {
    _taskHelper.addTask();
    await refreshTasks();
}

async function refreshTasks() {
    for (let task of _store.tasks.tasks){

    }
}
<script setup lang="ts">

import * as _constants from '../constants.ts';
import { store as _store } from '../store.ts';
import * as _fileHelper from '../fileHelper.ts';

async function chooseSaveFileLocation() {
    try {
        const filePath = await Neutralino.os.showSaveDialog('Save your tasks', { "defaultPath": _constants.defaultSavePath });
        _store.dataFilePath = filePath;
        _fileHelper.setDataFilePath(filePath);
    }
    catch (e) { console.dir(e); }
}

async function saveDataFile() {
    try {
        if (!_store.dataFilePath) chooseSaveFileLocation();
        if (!_store.dataFilePath) return;
        await _fileHelper.saveDataFile();
    }
    catch (e) { console.dir(e); }
}

</script>

<template>
  <button @click="saveDataFile">Save tasks to file</button>
</template>
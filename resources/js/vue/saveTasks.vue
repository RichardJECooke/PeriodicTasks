<script setup lang="ts">

import * as _constants from '../1constants.ts';
import { store as _store } from '../3store.ts';
import * as _fileHelper from '../5fileHelper.ts';

async function chooseSaveFileLocation() {
    try {
        const filePath = await Neutralino.os.showSaveDialog('Save your tasks', { "defaultPath": _constants.defaultDataFilePath });
        if (!filePath) return;
        _fileHelper.setDataFilePath(filePath);
    }
    catch (e) { console.dir(e); }
}

async function saveDataFile() {
    try {
        if (!_store.config.dataFilePath) chooseSaveFileLocation();
        if (!_store.config.dataFilePath) return;
        await _fileHelper.writeDataFile();
    }
    catch (e) { console.dir(e); }
}

</script>

<template>
  <button @click="saveDataFile">Save tasks to file</button>
</template>
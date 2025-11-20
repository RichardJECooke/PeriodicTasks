<script setup lang="ts">

import * as _constants from '../1constants.ts';
import { store as _store } from '../3store.ts';
import * as _fileHelper from '../5fileHelper.ts';

async function openDataFile() {
  try {
    await chooseOpenFileLocation();
    if (!_store.config.dataFilePath) return;
    _fileHelper.setDataFilePath(_store.config.dataFilePath);
    await _fileHelper.readDataFile();
  }
  catch (e) { console.dir(e); }
}

async function chooseOpenFileLocation() {
  try {
    let entries = await Neutralino.os.showOpenDialog('Load your tasks', {
      defaultPath: _constants.configFolderPath
      // ,filters: [ {name: 'JSON files', extensions: ['json', 'js']}  ] - does not work, bug
    });
    if (!entries || entries.length == 0) return;
    _store.config.dataFilePath = entries[0];
  }
  catch (e) { console.dir(e); }
}

</script>

<template>
  <button @click="openDataFile">Open tasks from file</button>
</template>
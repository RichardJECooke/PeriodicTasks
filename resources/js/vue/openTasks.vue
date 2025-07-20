<script setup lang="ts">

import * as _constants from '../constants.ts';
import { store as _store } from '../store.ts';
import * as _fileHelper from '../fileHelper.ts';

async function openDataFile() {
  try {
    await chooseOpenFileLocation();
    if (!_store.dataFilePath) return;
    _fileHelper.setDataFilePath(_store.dataFilePath);
    await _fileHelper.openDataFile();
  }
  catch (e) { console.dir(e); }
}

async function chooseOpenFileLocation() {
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

</script>

<template>
  <button @click="openDataFile">Open tasks from file</button>
</template>
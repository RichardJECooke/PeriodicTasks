<script setup lang="ts">

import * as _taskHelper from '../taskHelper.ts';

async function addTask() {
  _taskHelper.addTask();
}

</script>

<template>
    <button @click="addTask">Add Task</button>
</template>
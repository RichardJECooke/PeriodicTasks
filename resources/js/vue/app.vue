<script setup lang="ts">
import SaveTasks from './saveTasks.vue';
import AddTask from './addTask.vue';
import TaskList from './taskList.vue';
import OpenTasks from './openTasks.vue';
</script>

<template>
  <h1>Periodic Task Tracker</h1>
  <div><OpenTasks /></div>
  <div><SaveTasks /></div>
  <div><AddTask /></div>
  <div><TaskList /></div>
</template>

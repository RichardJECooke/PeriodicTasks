<script setup lang="ts">
import SaveTasks from './saveTasks.vue';
import OpenTasks from './openTasks.vue';
import AddTask from './addTask.vue';
import TaskList from './taskList.vue';
import ArchivedTaskList from './archivedTaskList.vue';
</script>

<template>
  <h1>Periodic Task Tracker</h1>
  <div><OpenTasks /></div>
  <div><SaveTasks /></div>

  <h2>Tasks</h2>
  <div><AddTask /></div>
  <div><TaskList /></div>

  <h2>Archived Tasks</h2>
  <div><ArchivedTaskList /></div>
</template>

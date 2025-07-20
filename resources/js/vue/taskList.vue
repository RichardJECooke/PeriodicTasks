<script setup lang="ts">

import { store as _store } from '../store.ts';
import TaskListItem from './taskListItem.vue';

</script>

<template>
  <ul id="tasks">
    <TaskListItem v-for="task in _store.tasks" :key="task.id" :task="task" />
  </ul>
</template>
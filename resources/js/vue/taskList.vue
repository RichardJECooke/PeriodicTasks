<script setup lang="ts">

import { store as _store } from '../3store.ts';
import TaskListItem from './taskListItem.vue';

</script>

<template>
  <ul id="tasks">
    <li v-for="task in _store.tasks" :key="task.id">
      <TaskListItem :task="task" />
    </li>
    <br />
  </ul>
</template>
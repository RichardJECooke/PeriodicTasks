<script setup lang="ts">
import { computed } from 'vue';
import { store as _store } from '../3store.ts';
import * as _taskHelper from '../4taskHelper.ts';
import TaskListItem from './taskListItem.vue';

const getNumDaysUntilDue = _taskHelper.getNumDaysUntilDue;
const sortedTasks = computed(() => _store.tasks.sort((a, b) => getNumDaysUntilDue(a) - getNumDaysUntilDue(b)));
</script>

<template>
  <ul id="tasks">
    <li v-for="task in sortedTasks" :key="task.id">
      <TaskListItem :task="task" />
    </li>
    <br />
  </ul>
</template>
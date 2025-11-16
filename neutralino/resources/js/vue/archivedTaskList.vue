<script setup lang="ts">
import { computed } from 'vue';
import * as _types from '../0types.ts';
import { store as _store } from '../3store.ts';
import * as _taskHelper from '../4taskHelper.ts';
import ArchivedTaskListItem from './archivedTaskListItem.vue';

const props = defineProps<{taskGroup: _types.tTaskGroup}>();

const sortedTasks = computed(() => props.taskGroup.tasks
  .filter((task) => task.isArchived)
  .sort((a, b) => a.name < b.name ? -1 : 1));
</script>

<template>
  <ul id="tasks">
    <li v-for="task in sortedTasks" :key="task.id">
      <ArchivedTaskListItem :task="task" />
    </li>
    <br />
  </ul>
</template>
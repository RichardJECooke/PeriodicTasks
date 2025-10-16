<script setup lang="ts">
import * as vue from 'vue';
import * as types from '../0types.ts';
import * as _taskHelper from '../4taskHelper.ts';
import TaskListItemEdit from './taskListItemEdit.vue';

const props = defineProps<{task: types.task}>();

function deleteTask() {
  _taskHelper.deleteTask(props.task);
}
</script>

<template>
  <span>{{ props.task.name }}</span>
  ·
  <input type="button" @click="props.task.isArchived = false" value="Unarchive this task">
  ·
  <input type="button" @click="deleteTask" value="Delete this task">
</template>
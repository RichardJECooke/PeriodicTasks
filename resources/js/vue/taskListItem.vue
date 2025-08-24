<script setup lang="ts">
import * as vue from 'vue';
import * as types from '../types.ts';
import * as _taskHelper from '../taskHelper.ts';
import TaskListItemEdit from './taskListItemEdit.vue';

const props = defineProps<{task: types.task}>();
let isEditing: vue.Ref<boolean> = vue.ref(false);
const getNumDaysUntilDue = _taskHelper.getNumDaysUntilDue; // entire file imports are not available in the template

function startEditing() { isEditing.value = true; }
function stopEditing() { isEditing.value = false; }
</script>

<template>
<li>
<!-- Readonly -->
  <div class="readonly" v-if="!isEditing"  v-on:click="startEditing">
    <span>{{ props.task.name }}</span>
    ·
    Due
    <span v-if="getNumDaysUntilDue(props.task) == 0">today</span>
    <span v-if="getNumDaysUntilDue(props.task) <-1">{{ getNumDaysUntilDue(props.task) }} days ago</span>
    <span v-if="getNumDaysUntilDue(props.task) <0 && getNumDaysUntilDue(props.task) >-1">{{ getNumDaysUntilDue(props.task) }} day ago</span>
    <span v-if="getNumDaysUntilDue(props.task) >0 && getNumDaysUntilDue(props.task) <1">in {{ getNumDaysUntilDue(props.task) }} day</span>
    <span v-if="getNumDaysUntilDue(props.task) >1">in {{ getNumDaysUntilDue(props.task) }} days</span>
  </div>
<!-- Editing   -->
  <TaskListItemEdit v-if="isEditing" :key="props.task.id" :task="props.task"
     :isEditing="isEditing" @stopEditingEvent="stopEditing" />
</li>
</template>
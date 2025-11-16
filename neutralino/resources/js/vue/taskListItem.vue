<script setup lang="ts">
import * as vue from 'vue';
import * as types from '../0types.ts';
import * as _taskHelper from '../4taskHelper.ts';
import TaskListItemEdit from './taskListItemEdit.vue';

const props = defineProps<{task: types.tTask}>();
let isEditing: vue.Ref<boolean> = vue.ref(false);
const getNumDaysUntilDue = _taskHelper.getNumDaysUntilDue; // entire file imports are not available in the template

function handleMarkDone(event: Event) {
  const target = event.target as HTMLInputElement;
  if (target.checked) _taskHelper.addDoneToday(props.task);
  else _taskHelper.removeDoneToday(props.task);
}
const isDoneToday = vue.computed(() => {
  const today = _taskHelper.getToday();
  return props.task.datesDone.some((date) => _taskHelper.isSameDay(new Date(date.date), today));
});

function startEditing() { isEditing.value = true; }
function stopEditing() { isEditing.value = false; }
</script>

<template>
<!-- Readonly -->
  <div v-if="!isEditing"  >
    <input type="checkbox" id="doneToday" v-bind:checked="isDoneToday" v-on:change="handleMarkDone"/>
    <label for="doneToday">Done today</label>
    &nbsp;&nbsp;
    <span>{{ props.task.name }}</span>
    ·
    Due
    <span v-if="getNumDaysUntilDue(props.task) == 0">today</span>
    <span v-if="getNumDaysUntilDue(props.task) <-1">{{ -1*getNumDaysUntilDue(props.task) }} days ago</span>
    <span v-if="getNumDaysUntilDue(props.task) <0 && getNumDaysUntilDue(props.task) >-1">{{ -1*getNumDaysUntilDue(props.task) }} day ago</span>
    <span v-if="getNumDaysUntilDue(props.task) >0 && getNumDaysUntilDue(props.task) <=1">in {{ getNumDaysUntilDue(props.task) }} day</span>
    <span v-if="getNumDaysUntilDue(props.task) >1">in {{ getNumDaysUntilDue(props.task) }} days</span>
    &nbsp;&nbsp;
    <input type="button" v-on:click="startEditing" value="Edit" />
  </div>
<!-- Editing   -->
  <TaskListItemEdit v-if="isEditing" :key="props.task.id" :task="props.task"
     :isEditing="isEditing" @stopEditingEvent="stopEditing" />
</template>
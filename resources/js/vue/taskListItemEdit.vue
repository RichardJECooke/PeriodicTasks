<script setup lang="ts">
import * as types from '../0types.ts';
import * as _taskHelper from '../4taskHelper.ts';
import DateListItem from './dateListItem.vue';

const props = defineProps<{ task: types.task, isEditing: boolean }>();
const emit = defineEmits(['stopEditingEvent']);

function stopEditing() { emit('stopEditingEvent'); }
function allowOnlyDigits(event: Event) {
  const input = event.target as HTMLInputElement;
  input.value = input.value.replace(/[^0-9]/g, '');
  props.task.days = Number(input.value) || 1;
}
function deleteDate(index: number) { props.task.datesDone.splice(index, 1); }
</script>

<template>
  <input  type="text" v-model="props.task.name" />
  <br />
  <span>Repeats every
    <input type="number" step="1" min="1" v-bind:value="props.task.days" v-on:input="allowOnlyDigits" />
    days
  </span>
  <br />
  <label>
    Archive this task:
    <input type="checkbox" v-bind:checked="props.task.isArchived" />
  </label>
  <br /><br />
  Dates completed:
  <ul id="datesList">
    <li v-for="(date, index) in props.task.datesDone">
      <DateListItem :key="date.id" :date="date.date" :index=index v-on:deleteDateEvent="deleteDate" />
    </li>
  </ul>
  <br />
  <input type="button" v-on:click="stopEditing" value="Done" />
</template>
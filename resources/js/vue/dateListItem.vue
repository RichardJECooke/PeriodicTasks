<script setup lang="ts">
import { computed } from 'vue';
import * as _types from '../0types.ts';
import * as _taskHelper from '../4taskHelper.ts';
import * as _baseHelper from '../2baseHelper.ts';

const props = defineProps<{"date": Date, "index": number}>();
// const dateString = computed(() => { return props.date.toISOString().slice(0, 10); });
const dateString = computed(() => { return _baseHelper.getDateFromDate(props.date); });
const emit = defineEmits(['deleteDateEvent', 'updateDateEvent']);

function deleteDate() { emit('deleteDateEvent', props.index); }
function updateDate(event: Event) {
  const newDateString = (event.target as HTMLInputElement).value;
  const newDate = new Date(newDateString);
  emit('updateDateEvent', props.index, newDate);
}
</script>

<template>
  <span class="dayWidth">{{ _baseHelper.getDayFromDate(props.date) }}</span>
  <span>{{ _baseHelper.getDateFromDate(props.date) }}</span>
  &nbsp;
  <input type="date" v-bind:value="dateString" v-on:input="updateDate" />
  &nbsp;
  <button v-on:click="deleteDate">Delete</button>
</template>
<script setup lang="ts">
import * as _types from '../0types.ts';
import * as _taskHelper from '../4taskHelper.ts';
import * as _baseHelper from '../2baseHelper.ts';

const props = defineProps<{"date": Date, "index": number}>();

const emit = defineEmits(['deleteDateEvent']);

function deleteDate() { emit('deleteDateEvent', props.index); }
</script>

<template>
  <span class="dayWidth">{{ _baseHelper.getDayFromDate(props.date) }}</span>
  <span>{{ _baseHelper.getDateFromDate(props.date) }}</span>
  &nbsp;
  <button v-on:click="deleteDate">Delete</button>
</template>
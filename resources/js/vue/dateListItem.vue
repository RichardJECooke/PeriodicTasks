<script setup lang="ts">
import * as types from '../0types.ts';
import * as _taskHelper from '../4taskHelper.ts';
import * as _baseHelper from '../2baseHelper.ts';

const props = defineProps<{ date: Date }>();

</script>

<template>
  <span class="dayWidth">{{ _baseHelper.getDayFromDate(props.date) }}</span><span>{{ _baseHelper.getDateFromDate(props.date) }}</span>
</template>
<script setup lang="ts">
import * as types from '../types.ts';
import * as _taskHelper from '../taskHelper.ts';
import * as _baseHelper from '../baseHelper.ts';

const props = defineProps<{ date: Date }>();

</script>

<template>
  <span class="dayWidth">{{ _baseHelper.getDayFromDate(props.date) }}</span><span>{{ _baseHelper.getDateFromDate(props.date) }}</span>
</template>
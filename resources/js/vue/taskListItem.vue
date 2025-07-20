<script setup lang="ts">
import { defineProps } from 'vue';
import type { task } from '../types.ts';

const props = defineProps<{task: task}>();
</script>

<template>
<li>
  <span>{{ props.task.name }} &nbsp;&nbsp;&nbsp;</span>
  <span>Repeats every {{ props.task.days }} days &nbsp;&nbsp;&nbsp;</span>
  <label>
    Next due date depends on last completion date:
    <input type="checkbox" :checked="props.task.dependsOnLastCompletion" disabled />
  </label>
</li>
</template>
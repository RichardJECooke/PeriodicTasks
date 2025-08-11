<script setup lang="ts">
import type { task } from '../types.ts';
import { ref, type Ref } from 'vue';

const props = defineProps<{task: task}>();
let isEditing: Ref<boolean> = ref(false);

function toggleEditing() { isEditing.value = !isEditing.value; }
</script>

<template>
<li @click="toggleEditing">
<!-- Readonly -->  
  <div class="readonly" v-if="!isEditing">
    <span>{{ props.task.name }}</span>
    · Due in 5 days
  </div>
<!-- Editing   -->
  <div class="readonly" v-if="isEditing">
    <input  type="text" v-model="props.task.name" />
    · 
    <span>Repeats every {{ props.task.days }} days</span>
    · 
    <label>
      Next due date depends on last completion date:
      <input type="checkbox" :checked="props.task.dependsOnLastCompletion" />
    </label>
  </div>
</li>
</template>
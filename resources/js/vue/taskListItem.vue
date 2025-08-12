<script setup lang="ts">
import type { task } from '../types.ts';
import { ref, type Ref } from 'vue';

const props = defineProps<{task: task}>();
let isEditing: Ref<boolean> = ref(false);

function startEditing() { isEditing.value = true; }
function stopEditing() { isEditing.value = false; }
function allowOnlyDigits(event: Event) {
  const input = event.target as HTMLInputElement;
  input.value = input.value.replace(/[^0-9]/g, '');
  props.task.days = Number(input.value) || 1;
}
</script>

<template>
<li>
<!-- Readonly -->
  <div class="readonly" v-if="!isEditing"  v-on:click="startEditing">
    <span>{{ props.task.name }}</span>
    · Due in 5 days
  </div>
<!-- Editing   -->
  <div class="editing" v-if="isEditing">
    <input  type="text" v-model="props.task.name" />
    <br />
    <span>Repeats every
      <input type="number" step="1" min="1" v-bind:value="props.task.days" v-on:input="allowOnlyDigits" />
      days
    </span>
    <br />
    <label>
      Next due date depends on last completion date:
      <input type="checkbox" v-bind:checked="props.task.dependsOnLastCompletion" />
    </label>
    <br />
    <input type="button" v-on:click="stopEditing" value="Done" />
  </div>
</li>
</template>
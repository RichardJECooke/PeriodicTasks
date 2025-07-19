<script setup lang="ts">
import { ref } from 'vue';
const count = ref(0);
</script>

<template>
  <h1>Periodic Task Tracker</h1>
  <div><a href="#" onclick="_ui.saveDataFile();">Save tasks to file</a></div>
  <div><a href="#" onclick="_ui.openDataFile();">Open tasks from file</a></div>
  <button id="addTaskButton">Add task</button>
  <div >
    <ul id="tasks"></ul>
  </div>
  <button @click="count++">Count is: {{ count }}</button>
</template>

<style scoped>
button {
  font-weight: bold;
}
</style>
// <script setup lang="ts">
// import { computed } from 'vue';
// import { store as _store } from '../3store.ts';
// import * as _taskHelper from '../4taskHelper.ts';
// import * as _types from '../0types.ts';
// import TaskListItem from './taskListItem.vue';

// const props = defineProps<{taskGroup: _types.tTaskGroup}>();

// const sortedTasks = computed(() => props.taskGroup.tasks
//   .filter((task) => !task.isArchived)
//   .sort((a, b) => _taskHelper.getNumDaysUntilDue(a) - _taskHelper.getNumDaysUntilDue(b)));
// </script>

// <template>
//   <ul id="tasks">
//     <li v-for="task in sortedTasks" :key="task.id">
//       <TaskListItem :task="task" />
//     </li>
//     <br />
//   </ul>
// </template>

package ui

import (
	"fyne.io/fyne/v2/widget"
	"github.com/RichardJECooke/PeriodicTasks/src"
)

// TODO
func CreateTaskList() *widget.Button {
	result := widget.NewButton("Task list", func() {
		src.AddTask()
	})
	return result
}

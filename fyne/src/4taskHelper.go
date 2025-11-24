package src

import (
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

func AddTask() {
	task := TTask{
		Id:         uuid.New().String(),
		Name:       "newtask",
		Days:       1,
		IsArchived: false,
		DatesDone:  []TIdAndDate{},
	}
	Store.TaskGroups[0].Tasks = append(Store.TaskGroups[0].Tasks, task)
	sortTasks()
}

func DeleteTask(taskToDelete TTask) {
	Store.TaskGroups[0].Tasks = lo.Filter(Store.TaskGroups[0].Tasks, func(taskItem TTask, _ int) bool {
		return taskItem.Id != taskToDelete.Id
	})
}

func SetTaskFile(taskFile TTaskGroup) {
	Store.TaskGroups[0] = taskFile
	sortTasks()
}

func GetNumDaysUntilDue(task TTask) int {
	if len(task.DatesDone) == 0 {
		return 0
	}
	lastDone := task.DatesDone[0].Date
	lastDoneMidnight := time.Date(lastDone.Year(), lastDone.Month(), lastDone.Day(), 0, 0, 0, 0, lastDone.Location())
	today := GetToday()
	daysSinceLastDone := int(today.Sub(lastDoneMidnight).Hours() / 24)
	return task.Days - daysSinceLastDone
}

func GetToday() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func AddDoneToday(task *TTask) {
	task.DatesDone = append(task.DatesDone, TIdAndDate{Id: uuid.New().String(), Date: GetToday()})
	sortTasks()
}

func IsSameDay(first, second time.Time) bool {
	return first.Year() == second.Year() && first.Month() == second.Month() && first.Day() == second.Day()
}

func RemoveDoneToday(task *TTask) {
	today := GetToday()
	task.DatesDone = lo.Filter(task.DatesDone, func(doneEntry TIdAndDate, _ int) bool {
		return !IsSameDay(doneEntry.Date, today)
	})
	sortTasks()
}

func RemoveDate(dateID string) {
	Store.TaskGroups[0].Tasks = lo.Map(Store.TaskGroups[0].Tasks, func(taskItem TTask, _ int) TTask {
		taskItem.DatesDone = lo.Filter(taskItem.DatesDone, func(doneEntry TIdAndDate, _ int) bool {
			return doneEntry.Id != dateID
		})
		return taskItem
	})
	sortTasks()
}

func sortTasks() {
	lo.ForEach(Store.TaskGroups[0].Tasks, func(task TTask, _ int) {
		sort.Slice(task.DatesDone, func(a, b int) bool {
			return task.DatesDone[a].Date.After(task.DatesDone[b].Date)
		})
	})
}

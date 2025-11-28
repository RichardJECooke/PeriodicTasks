package src

import (
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	types "github.com/RichardJECooke/PeriodicTasks/src/0types"
	store "github.com/RichardJECooke/PeriodicTasks/src/3store"
)

func AddTask() {
	task := types.Task{
		Id:         uuid.New().String(),
		Name:       "newtask",
		Days:       1,
		IsArchived: false,
		DatesDone:  []types.IdAndDate{},
	}
	store.SetTasks(append(store.GetStore().TaskGroups[0].Tasks, task))
}

func DeleteTask(taskToDelete types.Task) {
	store.SetTasks(lo.Filter(store.GetStore().TaskGroups[0].Tasks, func(taskItem types.Task, _ int) bool {
		return taskItem.Id != taskToDelete.Id
	}))
}

func GetNumDaysUntilDue(task types.Task) int {
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

func IsSameDay(first, second time.Time) bool {
	return first.Year() == second.Year() && first.Month() == second.Month() && first.Day() == second.Day()
}

func AddDoneToday(task *types.Task) {
	task.DatesDone = append(task.DatesDone, types.IdAndDate{Id: uuid.New().String(), Date: GetToday()})
	store.SetTasks(store.GetStore().TaskGroups[0].Tasks)
}

func RemoveDoneToday(task *types.Task) {
	today := GetToday()
	task.DatesDone = lo.Filter(task.DatesDone, func(doneEntry types.IdAndDate, _ int) bool {
		return !IsSameDay(doneEntry.Date, today)
	})
	store.SetTasks(store.GetStore().TaskGroups[0].Tasks)
}

func RemoveDate(dateID string) {
	for _, task := range store.GetStore().TaskGroups[0].Tasks {
		task.DatesDone = lo.Filter(task.DatesDone, func(date types.IdAndDate, _ int) bool {
			return date.Id != dateID
		})
	}
	store.SetTasks(store.GetStore().TaskGroups[0].Tasks)
}

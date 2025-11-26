package src

import (
	"slices"
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
	sortTasksDatesDone()
}

func DeleteTask(taskToDelete TTask) {
	Store.TaskGroups[0].Tasks = lo.Filter(Store.TaskGroups[0].Tasks, func(taskItem TTask, _ int) bool {
		return taskItem.Id != taskToDelete.Id
	})
}

func SetTaskGroup(taskFile TTaskGroup) {
	Store.TaskGroups[0] = taskFile
	sortTasksDatesDone()
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
	sortTasksDatesDone()
}

func IsSameDay(first, second time.Time) bool {
	return first.Year() == second.Year() && first.Month() == second.Month() && first.Day() == second.Day()
}

func RemoveDoneToday(task *TTask) {
	today := GetToday()
	task.DatesDone = lo.Filter(task.DatesDone, func(doneEntry TIdAndDate, _ int) bool {
		return !IsSameDay(doneEntry.Date, today)
	})
	sortTasksDatesDone()
}

func RemoveDate(dateID string) {
	for _, task := range Store.TaskGroups[0].Tasks {
		task.DatesDone = lo.Filter(task.DatesDone, func(date TIdAndDate, _ int) bool {
			return date.Id != dateID
		})
	}
	sortTasksDatesDone()
}

func sortTasksDatesDone() {
	for _, task := range Store.TaskGroups[0].Tasks {
		slices.SortFunc(task.DatesDone, func(a, b TIdAndDate) int {
			return a.Date.Compare(b.Date)
		})
	}
}

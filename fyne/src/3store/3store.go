package store

import (
	"slices"

	"github.com/google/uuid"

	types "github.com/RichardJECooke/PeriodicTasks/src/0types"
	constants "github.com/RichardJECooke/PeriodicTasks/src/1constants"
)

var observers []func()

var store = types.Store{
	TaskGroups: []types.TaskGroup{
		{
			Id:      uuid.NewString(),
			Version: constants.TaskFileVersion,
			Tasks:   []types.Task{},
		},
	},
	Config: types.Config{
		DataFilePath:               constants.DefaultDataPathExtensionAndFileName,
		ConfigFilePath:             "",
		ShouldNotify:               true,
		ShouldMinimizeToTrayOnQuit: true,
	},
}

func GetStore() types.Store {
	return store
}

func SetConfig(config types.Config) {
	store.Config = config
	notifyStoreChanged()
}

func SetTaskGroup(taskGroup types.TaskGroup) {
	store.TaskGroups[0] = taskGroup
	notifyStoreChanged()
}

func SetTasks(tasks []types.Task) {
	store.TaskGroups[0].Tasks = tasks
	notifyStoreChanged()
}

func RegisterForStoreChanged(callback func()) {
	observers = append(observers, callback)
}

func notifyStoreChanged() {
	for _, notify := range observers {
		notify()
	}
}

func sortTasksDatesDone() {
	for _, task := range store.TaskGroups[0].Tasks {
		slices.SortFunc(task.DatesDone, func(a, b types.IdAndDate) int {
			return a.Date.Compare(b.Date)
		})
	}
}

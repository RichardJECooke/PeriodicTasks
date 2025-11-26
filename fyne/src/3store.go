package src

import (
	"github.com/google/uuid"

	types "github.com/RichardJECooke/PeriodicTasks/src/0types"
	constants "github.com/RichardJECooke/PeriodicTasks/src/1constants"
)

var Store = types.TStore{
	TaskGroups: []types.TTaskGroup{
		{
			Id:      uuid.NewString(),
			Version: constants.TaskFileVersion,
			Tasks:   []types.TTask{},
		},
	},
	Config: types.TConfig{
		DataFilePath:               constants.DefaultDataPathExtensionAndFileName,
		ShouldNotify:               true,
		ShouldMinimizeToTrayOnQuit: true,
	},
}

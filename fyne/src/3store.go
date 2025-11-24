package src

import (
	"github.com/google/uuid"
)

var Store TStore = TStore{
	TaskGroups: []TTaskGroup{
		{
			Id:      uuid.New().String(),
			Version: TaskFileVersion,
			Tasks:   []TTask{},
		},
	},
	Config: TConfig{
		DataFilePath:               DefaultDataFilePath,
		ShouldNotify:               true,
		ShouldMinimizeToTrayOnQuit: true,
	},
}

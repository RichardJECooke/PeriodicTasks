package src

import "github.com/google/uuid"

var Store = TStore{
	TaskGroups: []TTaskGroup{
		{
			Id:      uuid.NewString(),
			Version: TaskFileVersion,
			Tasks:   []TTask{},
		},
	},
	Config: TConfig{
		DataFilePath:               DefaultDataFileName,
		ShouldNotify:               true,
		ShouldMinimizeToTrayOnQuit: true,
	},
}

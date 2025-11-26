package types

import "time"

type TIdAndDate struct {
	Id   string
	Date time.Time
}

type TTask struct {
	Id         string
	Name       string
	Days       int
	IsArchived bool
	DatesDone  []TIdAndDate
}

type TTaskGroup struct {
	Id      string
	Version int
	Tasks   []TTask
}

type TConfig struct {
	DataFilePath               string
	ConfigFilePath             string
	ShouldNotify               bool
	ShouldMinimizeToTrayOnQuit bool
}

type TStore struct {
	TaskGroups []TTaskGroup
	Config     TConfig
}

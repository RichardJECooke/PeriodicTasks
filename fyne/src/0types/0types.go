package types

import "time"

type IdAndDate struct {
	Id   string
	Date time.Time
}

type Task struct {
	Id         string
	Name       string
	Days       int
	IsArchived bool
	DatesDone  []IdAndDate
}

type TaskGroup struct {
	Id      string
	Version int
	Tasks   []Task
}

type Config struct {
	DataFilePath               string
	ConfigFilePath             string
	ShouldNotify               bool
	ShouldMinimizeToTrayOnQuit bool
}

type Store struct {
	TaskGroups []TaskGroup
	Config     Config
}

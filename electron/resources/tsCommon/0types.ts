export type tIdAndDate = { id: string; date: Date }

export type tTaskGroup = {
  id: string;
  version: number;
  tasks: tTask[];
}

export type tTask = {
  id: string;
  name: string;
  days: number;
  isArchived: boolean;
  datesDone: tIdAndDate[];
}

export type tConfig = {
  dataFilePath: string
  shouldNotify: boolean;
  shouldMinimizeToTrayOnQuit: boolean
}

export type tStore = {
  taskGroups: tTaskGroup[];
  config: tConfig;
}
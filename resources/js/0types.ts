export type tidAndDate = { id: string; date: Date }

export type ttask = {
  id: string;
  name: string;
  days: number;
  isArchived: boolean;
  datesDone: tidAndDate[];
}

export type tconfig = {
  dataFilePath: string
  shouldNotify: boolean;
  shouldMinimizeToTrayOnQuit: boolean
}

export type tstore = {
  tasks: ttask[];
  config: tconfig;
}
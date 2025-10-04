export type idAndDate = { id: string; date: Date }

export type task = {
  id: string;
  name: string;
  days: number;
  isArchived: boolean;
  datesDone: idAndDate[];
}

export type store = {
  tasks: task[],
  dataFilePath: string | null
}
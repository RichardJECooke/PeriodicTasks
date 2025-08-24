export type task = {
  id: string;
  name: string;
  days: number;
  isArchived: boolean;
  datesDone: Date[];
}

export type store = {
  tasks: task[],
  dataFilePath: string | null
}
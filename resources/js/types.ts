export type task = {
  id: string;
  name: string;
  days: number;
  dependsOnLastCompletion: boolean;
}

export type store = {
  tasks: task[],
  dataFilePath: string | null
}
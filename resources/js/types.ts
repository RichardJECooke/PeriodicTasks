export type task = {
  id: string;
  name: string;
  days: number;
  dependsOnLastCompletion: boolean;
}
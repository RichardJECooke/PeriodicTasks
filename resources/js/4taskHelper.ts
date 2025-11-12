import * as _constants from './1constants.js';
import * as _types from './0types.ts';
import {store as _store} from './3store.ts';

export async function addTask(): Promise<void> {
  const task: _types.tTask = {'id': crypto.randomUUID(), 'name': 'newtask', 'days': 1, isArchived: false, datesDone: []};
  _store.taskGroups[0].tasks.push(task);
  _store.taskGroups[0].tasks.forEach(task => { task.datesDone.sort((a, b) => b.date.getTime() - a.date.getTime()); });
}

export function deleteTask(taskToDelete: _types.tTask): void {
  _store.taskGroups[0].tasks = _store.taskGroups[0].tasks.filter(task => task.id !== taskToDelete.id);
}

export function setTaskFile(taskFile: _types.tTaskGroup): void {
  taskFile.tasks.forEach(task => { task.datesDone.sort((a, b) => b.date.getTime() - a.date.getTime()); });
  _store.taskGroups[0] = taskFile;
}

export function getNumDaysUntilDue(task : _types.tTask): number {
  if (task.datesDone.length == 0) return 0;
  const lastDoneDate = task.datesDone[0];
  lastDoneDate.date.setHours(0, 0, 0, 0);
  const msPerDay = 1000 * 60 * 60 * 24;
  const daysSinceLastDone = Math.floor((getToday().getTime() - lastDoneDate.date.getTime()) / msPerDay);
  const result = task.days - daysSinceLastDone;
  return result;
}

export function getToday(): Date {
  const today = new Date();
  today.setHours(0, 0, 0, 0);
  return today;
}

export function addDoneToday(task: _types.tTask): void {
  task.datesDone.push({id: crypto.randomUUID(), date: getToday()});
  _store.taskGroups[0].tasks.forEach(task => { task.datesDone.sort((a, b) => b.date.getTime() - a.date.getTime()); });
}

export function isSameDay (first: Date, second: Date): boolean {
  return (first.getFullYear() === second.getFullYear() &&
  first.getMonth() === second.getMonth() &&
  first.getDate() === second.getDate());
}

export function removeDoneToday(task: _types.tTask): void {
  const today = getToday();
  task.datesDone = task.datesDone.filter(date => !isSameDay(date.date, today));
  _store.taskGroups[0].tasks.forEach(task => { task.datesDone.sort((a, b) => b.date.getTime() - a.date.getTime()); });
}

export function removeDate(dateId: string): void {
  _store.taskGroups[0].tasks.forEach(task => {
    task.datesDone = task.datesDone.filter(date => date.id !== dateId);
  });
  _store.taskGroups[0].tasks.forEach(task => { task.datesDone.sort((a, b) => b.date.getTime() - a.date.getTime()); });
}
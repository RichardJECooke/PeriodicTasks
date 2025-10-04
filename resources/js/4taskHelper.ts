import {v4 as _uuid} from 'uuid';
import * as _constants from './1constants.js';
import * as _types from './0types.ts';
import {store as _store} from './3store.ts';

export async function addTask(): Promise<void> {
  const task: _types.task = {'id': _uuid(), 'name': 'newtask', 'days': 1, isArchived: false, datesDone: []};
  _store.tasks.push(task);
  _store.tasks.forEach(task => { task.datesDone.sort((a, b) => b.date.getTime() - a.date.getTime()); });
}

export function setTasks(tasks: _types.task[]) {
  tasks.forEach(task => { task.datesDone.sort((a, b) => b.date.getTime() - a.date.getTime()); });
  _store.tasks = tasks;
}

export function getNumDaysUntilDue(task : _types.task): number {
  if (task.datesDone.length == 0) return 0;
  const lastDoneDate = task.datesDone[0];
  lastDoneDate.date.setHours(0, 0, 0, 0);
  const msPerDay = 1000 * 60 * 60 * 24;
  const daysSinceLastDone = Math.floor((getToday().getTime() - lastDoneDate.date.getTime()) / msPerDay);
  return task.days - daysSinceLastDone;
}

export function getToday(): Date {
  const today = new Date();
  today.setHours(0, 0, 0, 0);
  return today;
}

export function addDoneToday(task: _types.task): void {
  task.datesDone.push({id: _uuid(), date: getToday()});
}

export function isSameDay (first: Date, second: Date): boolean {
  return (first.getFullYear() === second.getFullYear() &&
  first.getMonth() === second.getMonth() &&
  first.getDate() === second.getDate());
}

export function removeDoneToday(task: _types.task): void {
  const today = getToday();
  task.datesDone = task.datesDone.filter(date => !isSameDay(date.date, today));
}

export function removeDate(dateId: string): void {
  _store.tasks.forEach(task => {
    task.datesDone = task.datesDone.filter(date => date.id !== dateId);
  });
}
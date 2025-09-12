import {v4 as _uuid} from 'uuid';
import * as _constants from './constants.js';
import * as types from './types.ts';
import {store as _store} from './store.ts';

export async function addTask(): Promise<void> {
  const task: types.task = {'id': _uuid(), 'name': 'newtask', 'days': 1, isArchived: false, datesDone: []};
  _store.tasks.push(task);
}

export function getNumDaysUntilDue(task : types.task): number {
  if (task.datesDone.length == 0) return 0;
  task.datesDone.map((d) => console.log(d));
  const lastDoneDate = task.datesDone.toSorted((a, b) => b.getTime() - a.getTime())[0];
  lastDoneDate.setHours(0, 0, 0, 0);
  const msPerDay = 1000 * 60 * 60 * 24;
  const daysSinceLastDone = Math.floor((getToday().getTime() - lastDoneDate.getTime()) / msPerDay);
  return task.days - daysSinceLastDone;
}

export function getToday(): Date {
  const today = new Date();
  today.setHours(0, 0, 0, 0);
  return today;
}

export function addDoneToday(task: types.task): void {
  task.datesDone.push(getToday());
}

export function isSameDay (first: Date, second: Date): boolean {
  return (first.getFullYear() === second.getFullYear() &&
  first.getMonth() === second.getMonth() &&
  first.getDate() === second.getDate());
}

export function removeDoneToday(task: types.task): void {
  const today = getToday();
  task.datesDone = task.datesDone.filter(date => !isSameDay(date, today));
}
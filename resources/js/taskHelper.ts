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
  const lastDoneDate = task.datesDone.sort((a, b) => b.getTime() - a.getTime())[0];
  const today = new Date();
  lastDoneDate.setHours(0, 0, 0, 0);
  today.setHours(0, 0, 0, 0);
  const msPerDay = 1000 * 60 * 60 * 24;
  const daysSinceLastDone = Math.floor((today.getTime() - lastDoneDate.getTime()) / msPerDay);
  return task.days - daysSinceLastDone;
}
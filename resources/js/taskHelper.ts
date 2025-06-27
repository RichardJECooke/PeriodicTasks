import {v4 as _uuid} from 'uuid';
import * as _constants from './constants.js';
import * as _store from './store.ts';

export async function addTask(): Promise<void> {
    const task = {'id': _uuid(), 'name': 'newtask', 'days': 1, dependsOnLastCompletion: true};
    _store.tasks.tasks.push(task);
}

import * as _constants from './constants.mjs';
import {v4 as _uuid} from 'uuid';

export async function addTask() {
    const task = {'id': _uuid(), 'name': '', 'days': 1, dependsOnLastCompletion: true};
    console.log(task);
    _tasks.tasks.push(task);
    await refreshTasks();
}

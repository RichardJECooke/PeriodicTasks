import {v4 as _uuid} from 'uuid';
import * as _constants from './constants.mjs';
import * as _store from './store.mjs';


export async function addTask() {
    const task = {'id': _uuid(), 'name': '', 'days': 1, dependsOnLastCompletion: true};
    console.log(task);
    _store.tasks.tasks.push(task);
}

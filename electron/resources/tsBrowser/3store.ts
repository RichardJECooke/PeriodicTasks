import { reactive } from 'vue';
import * as _types from './0types.ts';
import * as _constants from './1constants.ts';


export const store = reactive<_types.tStore>({
    'taskGroups': [{
      'id': crypto.randomUUID(),
      'version': _constants.taskFileVersion,
      'tasks': []
    }],
    'config': {
        'dataFilePath': _constants.defaultDataFilePath,
        'shouldNotify': true,
        'shouldMinimizeToTrayOnQuit': true,
    }
});
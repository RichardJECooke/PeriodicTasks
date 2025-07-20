import { reactive } from 'vue';
import * as _types from './types.ts';

export const store = reactive<_types.store>({
  "tasks": [],
  "dataFilePath": null
});

import { reactive } from 'vue';
import * as _types from './0types.ts';

export const store = reactive<_types.store>({
  "tasks": [],
  "dataFilePath": null
});
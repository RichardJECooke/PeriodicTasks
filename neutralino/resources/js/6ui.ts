globalThis.__VUE_OPTIONS_API__ = false;
globalThis.__VUE_PROD_DEVTOOLS__ = true;
globalThis.__VUE_PROD_HYDRATION_MISMATCH_DETAILS__ = true;

import * as _vue from 'vue';
import _app from './vue/app.vue';

export async function start() {
  const app = _vue.createApp(_app);
  app.config.errorHandler = (e: unknown): void => { console.log(e); }
  app.mount('#app');
}
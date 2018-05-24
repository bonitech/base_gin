import Vue from 'vue';
import Vuex from 'vuex';

import Mod from './module_name';

Vue.use(Vuex);

const debug = process.env.NODE_ENV !== 'production';

export default new Vuex.Store({
  modules: {
    Mod,
  },
  strict: debug,
});

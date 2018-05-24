const state = {
  count: 0,
};

const mutations = {
  INCREMENT(state, payload) {
    state.count += payload;
  },
  DE_INCREMENT(state, payload) {
    state.count -= payload;
  },
};

const actions = {
  increment: ({ commit }, amount) => commit('INCREMENT', amount),
  deincrement: ({ commit }, amount) => commit('DE_INCREMENT', amount),
};

const getters = {
  getCounter: state => state.count,
};

export default {
  state,
  getters,
  mutations,
  actions,
};

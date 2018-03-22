import Vue from 'vue'
import Vuex from 'vuex'
import mutations from './mutations'

Vue.use(Vuex)

const state = {
  user: {},
  location: {}
}

export default new Vuex.Store({
  state,
  mutations,
  getters: {}
})
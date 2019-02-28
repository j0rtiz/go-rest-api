import Vue from 'vue'
import Vuex from 'vuex'
import session from './session'

Vue.use(Vuex)

export default function () {
  const Store = new Vuex.Store({
    modules: {
      session
    }
  })

  return Store
}

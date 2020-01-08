import Vue from 'vue'
import Vuex from 'vuex'   
import crudtable from './crudtable' 

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
  }  
})
{{range .Entities}}
store.registerModule('{{CamelCase .Name}}', crudtable()) {{end}}

export default store
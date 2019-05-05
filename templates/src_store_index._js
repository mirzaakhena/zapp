import Vue from 'vue'
import Vuex from 'vuex' {{range .Classes}}  
import {{CamelCase .Name}} from './modules/{{LowerCase .Name}}' {{end}}

Vue.use(Vuex)

export default new Vuex.Store({
  modules: { {{range .Classes}}  
    {{CamelCase .Name}}, {{end}}
  }
})
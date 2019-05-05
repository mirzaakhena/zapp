import Vue from 'vue'
import Router from 'vue-router' {{range .Classes}}
import {{CamelCase .Name}} from './modules/{{LowerCase .Name}}' {{end}}

Vue.use(Router)

const router = new Router({
  routes: [ {{range .Classes}}
    ...{{CamelCase .Name}}, {{end}}
  ]
})

export default router
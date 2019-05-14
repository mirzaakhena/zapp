import Vue from 'vue'
import Router from 'vue-router' {{range .Classes}}
import {{CamelCase .Name}} from './modules/{{LowerCase .Name}}' {{end}}

Vue.use(Router)

const router = new Router({
  routes: [ 
    {
      path: '/login',
      component: () => import('@/pages/login.vue'),
      meta: {
        authorities: 'GUEST'
      },
    },

    {
      path: '/register',
      component: () => import('@/pages/register.vue'),
      meta: {
        authorities: 'GUEST'
      },
    }, 
    
    {
      path: '/forgotpassword',
      component: () => import('@/pages/forgotpassword.vue'),
      meta: {
        authorities: 'GUEST'
      },
    },  
    
    {
      path: '/notfound',
      component: () => import('@/pages/notfound.vue'),
      meta: {
      },
    },      
    
    {
      path: '/',
      component: () => import('@/pages/home.vue'),
      meta: {
        authorities: 'USER'
      },
      children: [ {{range .Classes}}
        ...{{CamelCase .Name}}, {{end}}
      ],
    },        
  ],
})

export default router

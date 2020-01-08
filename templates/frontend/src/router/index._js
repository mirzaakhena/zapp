import Vue from 'vue'
import Router from 'vue-router' 
import Auth from '@/utils/auth' {{range .Entities}}
import {{CamelCase .Name}} from './modules/{{LowerCase .Name}}' {{end}}

Vue.use(Router)
Vue.use(Auth)

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
    // 
    // {
    //   path: '/forgotpassword',
    //   component: () => import('@/pages/forgotpassword.vue'),
    //   meta: {
    //     authorities: 'GUEST'
    //   },
    // },  
    // 
    // {
    //   path: '/notfound',
    //   component: () => import('@/pages/notfound.vue'),
    //   meta: {
    //   },
    // },
    // {
    //   path: '/successregister',
    //   component: () => import('@/pages/successregister.vue'),
    //   meta: {
    //     authorities: 'GUEST'
    //   },
    // },   
    // 
    {
      path: '/',
      component: () => import('@/pages/home.vue'),
      meta: {
        authorities: 'USER'
      },
      children: [ {{range .Entities}}
        ...{{CamelCase .Name}}, {{end}}
      ],
    },

    { path: '*', redirect: '/' }
  ],
})

router.beforeEach((to, from, next) => {  
  if (to.matched.some(record => record.meta.authorities === 'GUEST')) {
    if (Vue.auth.isAuthenticated()) {
      next({
        path: '/'
      })
    } else {
     next()
    }
  } else if (to.matched.some(record => record.meta.authorities === 'USER')) {
    if (!Vue.auth.isAuthenticated()) {
      next({
        path: '/login'
      })
    } else {
      next()
    }
  } else {
    next()
  }
})


export default router

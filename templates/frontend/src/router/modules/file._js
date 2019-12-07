const route = [          

  {
    path: '/{{LowerCase .Name}}',
    name: '{{CamelCase .Name}}',
    component: () => import('@/pages/{{LowerCase .Name}}/list.vue'),
  },

]

export default route
const route = [          

  {
    path: '/{{LowerCase .Name}}',
    name: '{{CamelCase .Name}}',
    component: () => import('@/pages/{{LowerCase .Name}}/table.vue'),
  },

]

export default route
import api from '@/api';

const store = {
  state: {
    {{CamelCase .Name}}s: [],
    totalItems: 0,
  },

  mutations: {
    SET_{{UpperCase .Name}}: (state, payload) => {
      state.{{CamelCase .Name}}s = payload
    },
    SET_TOTAL_ITEMS: (state, payload) => {
      state.totalItems = payload
    },   
  },

  getters: {
    TotalItems: state => state.totalItems,    
    GetAll{{PascalCase .Name}}: state => state.{{CamelCase .Name}}s,
    GetOne{{PascalCase .Name}}: state => (id) => state.{{CamelCase .Name}}s.find({{CamelCase .Name}} => {{CamelCase .Name}}.id === id)
  },

  actions: {
   
    async GetAll{{PascalCase .Name}} ({commit}, payload) {      
      const result = await api.GetAll{{PascalCase .Name}}(payload)
      commit('SET_TOTAL_ITEMS', result.headers["data-length"])
      commit('SET_{{UpperCase .Name}}', result.data.data)
      return result.data
    },

    async Create{{PascalCase .Name}} ({dispatch}, payload) {
      const {data} = await api.Create{{PascalCase .Name}}(payload)
      dispatch('GetAll{{PascalCase .Name}}')
      return data
    },

    async Update{{PascalCase .Name}} ({dispatch}, payload) {
      const {data} = await api.Update{{PascalCase .Name}}(payload.id, payload)
      dispatch('GetAll{{PascalCase .Name}}')
      return data
    },

    async Delete{{PascalCase .Name}} ({dispatch}, id) {
      const {data} = await api.Delete{{PascalCase .Name}}(id)
      dispatch('GetAll{{PascalCase .Name}}')
      return data
    },        
  },
}

export default store
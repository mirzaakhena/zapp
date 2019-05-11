import api from '@/api';

const store = {

  state: {
    data{{PascalCase .Name}}: [],
    paging{{PascalCase .Name}}: {
      page: 1,
      size: 20,
      total: 0,
    },
    sorting{{PascalCase .Name}}: {
      byField: '',
      isDesc: false,
    },
    filtering{{PascalCase .Name}}: { {{range .Fields}}
      {{CamelCase .Name}}: null, {{end}}
    }     
  },

  mutations: {
    SET_{{UpperCase .Name}}: (state, payload) => {
      state.data{{PascalCase .Name}} = payload
    },
    SET_PAGING_PAGE_{{UpperCase .Name}}: (state, payload) => {
      state.paging{{PascalCase .Name}}.page = payload
    }, 
    SET_PAGING_TOTAL_{{UpperCase .Name}}: (state, payload) => {
      state.paging{{PascalCase .Name}}.total = payload
    },  
    SET_SORTING_{{UpperCase .Name}}: (state, payload) => {
      state.sorting{{PascalCase .Name}} = payload
    }, 
    SET_FILTERING_{{UpperCase .Name}}: (state, payload) => {
      state.filtering{{PascalCase .Name}} = payload
    },     
  },

  getters: {
    Sorting{{PascalCase .Name}}: state => state.sorting{{PascalCase .Name}},  
    Filtering{{PascalCase .Name}}: state => state.filtering{{PascalCase .Name}},  
    Paging{{PascalCase .Name}}: state => state.paging{{PascalCase .Name}},  
    GetAll{{PascalCase .Name}}: state => state.data{{PascalCase .Name}},
    GetOne{{PascalCase .Name}}: state => (id) => state.data{{PascalCase .Name}}.find(data{{PascalCase .Name}} => data{{PascalCase .Name}}.id === id)
  },

  actions: {

    async UpdateSorting{{PascalCase .Name}} ({commit, dispatch}, payload) {
      commit('SET_SORTING_{{UpperCase .Name}}', {byField:payload.byField, isDesc:payload.isDesc})
      dispatch('GetAll{{PascalCase .Name}}')
    }, 

    async UpdateFiltering{{PascalCase .Name}} ({commit, dispatch}, isResetField) {
      if (isResetField) {
        var emptyFilter = { 
          cond: null, 
          code: null, 
          description: null, 
        }  
        commit('SET_FILTERING_{{UpperCase .Name}}', emptyFilter)
      }
      commit('SET_PAGING_PAGE_{{UpperCase .Name}}', 1)
      dispatch('GetAll{{PascalCase .Name}}')
    },      
    
    async GetAll{{PascalCase .Name}} ({commit, state}) {
      const result = await api.GetAll{{PascalCase .Name}}(state.paging{{PascalCase .Name}}, state.sorting{{PascalCase .Name}}, state.filtering{{PascalCase .Name}})
      commit('SET_PAGING_TOTAL_{{UpperCase .Name}}', result.headers["data-length"])
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
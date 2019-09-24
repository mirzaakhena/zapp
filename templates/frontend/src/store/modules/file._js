import api from '@/api';

const store = {

  namespaced: true,

  state: {
    items: [],
    paging: {
      page: 1,
      size: 20,
      total: 0,
    },
    sorting: {
      byField: '',
      isDesc: false,
    },
    filtering: initializeObject(),
  },

  mutations: {
    SET_ITEMS: (state, payload) => {
      state.items = payload
    },
    SET_PAGING_PAGE: (state, payload) => {
      state.paging.page = payload
    }, 
    SET_PAGING_TOTAL: (state, payload) => {
      state.paging.total = payload
    },  
    SET_SORTING: (state, payload) => {
      state.sorting = payload
    }, 
    SET_FILTERING: (state, payload) => {
      state.filtering = payload
    },     
  },

  getters: {
    Sorting: state => state.sorting,  
    Filtering: state => state.filtering,  
    Paging: state => state.paging,  
    GetAll: state => state.items,
    GetOne: state => (id) => state.items.find(item => item.id === id)
  },

  actions: {

    async UpdateSorting ({commit, dispatch}, payload) {
      commit('SET_SORTING', {byField:payload.byField, isDesc:payload.isDesc})
      dispatch('GetAll')
    }, 

    async UpdateFiltering ({commit, dispatch}) { 
      commit('SET_FILTERING', initializeObject())      
      commit('SET_PAGING_PAGE', 1)
      dispatch('GetAll')
    },      
    
    async GetAll ({commit, state}) {

      let filterWithPrefix = {}
      Object.keys(state.filtering).forEach(key => {
        filterWithPrefix['f_' +key] = state.filtering[key]
      })

      const result = await api.{{LowerCase .Name}}.GetAll(state.paging, state.sorting, filterWithPrefix)
      commit('SET_PAGING_TOTAL', result.data.data.totalCount)
      commit('SET_ITEMS', result.data.data.items)
      return result.data
    },    

    async Create ({dispatch}, payload) {
      const {data} = await api.{{LowerCase .Name}}.Create(payload)
      dispatch('GetAll')
      return data
    },

    async Update ({dispatch}, payload) {
      const {data} = await api.{{LowerCase .Name}}.Update(payload.id, payload)
      dispatch('GetAll')
      return data
    },

    async Delete ({dispatch}, id) {
      const {data} = await api.{{LowerCase .Name}}.Delete(id)
      dispatch('GetAll')
      return data
    },        
  },
}

function initializeObject() {
  return { {{range .Fields}}
    {{CamelCase .Name}}: null, {{end}}
  }
}

export default store
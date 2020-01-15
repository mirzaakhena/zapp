import to from 'await-to-js';
import NewRestApi from "@/api/restapi";

function NewStore(urlContextPath) {

  const restapi = NewRestApi(urlContextPath)

  return {

    namespaced: true,
  
    state: {
      items: [],
      totalItems: 0,
      inputedItem: null,
      paging: {
        page: 1,
        size: 20,
      },
      sorting: {
        sortBy: '',
        sortDesc: false,
      },
      filtering: [],
    },
  
    mutations: {
  
      SET_ITEMS: (state, {items, totalItems}) => {
        state.items = items
        state.totalItems = totalItems
      },
  
      SET_PAGING_PAGE: (state, payload) => {
        state.paging.page = payload
      },
  
      SET_SORTING: (state, payload) => {
        state.sorting = payload
      }, 
  
      SET_FILTERING: (state, payload) => {
        state.filtering = payload
      },
  
      RESET_PAGING_PAGE: (state) => {
        state.paging.page = 1
      },
  
      RESET_FILTERING: (state) => {
        state.filtering = []
      },
      
      RESET_SORTING: (state) => {
        state.sorting.sortBy = ''
        state.sorting.sortDesc = false
      },
  
      SET_INPUTED_ITEM: (state, payload) => {
        state.inputedItem = payload
      },
  
    },
  
    getters: {
      getItems: state => state.items,
      getTotalItems: state => state.totalItems,
      getInputedItem: state => state.inputedItem,
      getPaging: state => state.paging,
      getSorting: state => state.sorting,
      getFiltering: state => state.filtering,    
    },
  
    actions: {
  
      resetPaging ({state, commit, dispatch}) {
        if (state.paging.page > 1) {
          commit('RESET_PAGING_PAGE')
          return
        }
        dispatch('queryItems')
      },
  
      resetFiltering ({commit, dispatch}) {
        commit('RESET_FILTERING')
        dispatch('resetPaging')
      },
  
      resetSorting ({commit, dispatch}) {
        commit('RESET_SORTING')
        dispatch('resetPaging')
      },
  
      onFiltering ({dispatch}) {
        dispatch('resetPaging')
      },
  
      onSorting ({dispatch}) {
        dispatch('resetPaging')
      },
  
      onPaging ({dispatch}) {
        dispatch('queryItems')
      },
  
      async queryItems ({state, commit}) {
  
        Object.keys(state.filtering).find(key => {
          if (state.filtering[key] === '') {
            delete state.filtering[key]
          }
        })
  
        if (state.sorting.sortBy === '') {
          delete state.sorting['sortBy']
          delete state.sorting['sortDesc']
        }
  
        if (!state.sorting.sortDesc) {
          delete state.sorting['sortDesc']
        }

        const [error, response] = await to(restapi.getAll(state.paging, state.sorting, state.filtering))
        if (error) {
          commit('SET_ITEMS', {
            items: [], 
            totalItems: 0
          })
          return Promise.reject(error)
        }

        commit('SET_ITEMS', {
          items: response.items, 
          totalItems: response.totalCount
        })

        return Promise.resolve(response)
        
      },
  
      async createItem ({dispatch}, {inputedItem}) {
        const [error, response] = await to(restapi.create(inputedItem))
        if (error) {
          return Promise.reject(error)
        }
        dispatch('queryItems')
        return Promise.resolve(response)
      },
  
      async getOneItem ({commit}, {itemId}) {
        const [error, response] = await to(restapi.getOne(itemId))
        if (error) {
          return Promise.reject(error)
        }
        commit('SET_INPUTED_ITEM', response.data)
        return Promise.resolve(response)
      },
  
      async updateItem ({state, dispatch}) {
        const [error, response] = await to(restapi.update(state.inputedItem, state.inputedItem.id))
        if (error) {
          return Promise.reject(error)
        }
        dispatch('queryItems')
        return Promise.resolve(response)
      },
  
      async deleteItem({dispatch}, {itemId}) {
        const [error, response] = await to(restapi.delete(itemId))
        if (error) {
          return Promise.reject(error)
        }
        dispatch('queryItems')
        return Promise.resolve(response)
      },
         
    },
  }
}

export default NewStore

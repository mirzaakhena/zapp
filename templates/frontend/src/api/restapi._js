import to from 'await-to-js';
import request from "@/utils/httprequest";

function NewRestApi(context) {
  return {
    getAll (paging, filtering, sorting) {
      return new Promise(async (resolve, reject) => {
        const [error, response] = await to(request({
          method: 'get',
          url: `api/${context}`,
          params: {
            ...paging,
            ...sorting,
            ...filtering,
          }
        }))
        if (error) {
          reject(error)
          return
        }
        resolve({
          items: response.data.data.items, 
          totalCount: response.data.data.totalCount
        })
      })
    },

    getOne (itemId) {
      return new Promise(async (resolve, reject) => {
        const [error, response] = await to(request({
          method: 'get',
          url: `api/${context}/${itemId}`,
        }))
        if (error) {
          reject(error)
          return
        }
        resolve(response.data)
      })
    },

    create (item) {
      return new Promise(async (resolve, reject) => {
        const [error, response] = await to(request({
          data: item,
          method: 'post',
          url: `api/${context}`,
        }))
        if (error) {
          reject(error)
          return
        }
        resolve(response.data)
      })
    },

    update (item, itemId) {
      return new Promise(async (resolve, reject) => {
        const [error, response] = await to(request({
          data: item,
          method: 'put',
          url: `api/${context}/${itemId}`,
        }))
        if (error) {
          reject(error)
          return
        }
        resolve(response.data)
      })
    },

    delete (itemId) {
      return new Promise(async (resolve, reject) => {
        const [error, response] = await to(request({
          method: 'delete',
          url: `api/${context}/${itemId}`,
        }))
        if (error) {
          reject(error)
          return
        }
        resolve(response.data)
      })
    },
  }
}

export default NewRestApi
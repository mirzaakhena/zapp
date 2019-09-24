import request from "@/utils/httprequest";

const api = {

  {{LowerCase .Name}}: {

    Create(payload) {
      return request({
        url: `/{{CamelCase .Name}}`,
        method: 'post',
        data: payload,
      })
    },

    GetOne({{CamelCase .Name}}ID) {
      return request({
        url: `/{{CamelCase .Name}}/${{"{"}}{{CamelCase .Name}}ID}`,
        method: 'get',
      })
    },  

    GetAll(paging, sorting, filtering) {    
      return request({
        url: `/{{CamelCase .Name}}`,
        method: 'get',
        params: {
          page: paging.page,
          size: paging.size,
          sortBy: sorting.byField,
          sortDir: sorting.isDesc? 'desc': 'asc', 
          ...filtering,
        }
      })
    },    

    Update({{CamelCase .Name}}ID, payload) {
      return request({
        url: `/{{CamelCase .Name}}/${{"{"}}{{CamelCase .Name}}ID}`,
        method: 'put',
        data: payload,
      })
    },

    Delete({{CamelCase .Name}}ID) {
      return request({
        url: `/{{CamelCase .Name}}/${{"{"}}{{CamelCase .Name}}ID}`,
        method: 'delete',
      })
    }, 

  } 

}

export default api
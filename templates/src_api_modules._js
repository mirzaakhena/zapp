import request from "@/utils/httprequest";

const api = {

  Create{{PascalCase .Name}}(payload) {
    return request({
      url: `/api/{{CamelCase .Name}}`,
      method: 'post',
      data: payload,
    })
  },

  GetOne{{PascalCase .Name}}({{CamelCase .Name}}ID) {
    return request({
      url: `/api/{{CamelCase .Name}}/${{"{"}}{{CamelCase .Name}}ID}`,
      method: 'get',
    })
  },  

  GetAll{{PascalCase .Name}}(paging, sorting, filtering) {    
    return request({
      url: `/api/{{CamelCase .Name}}`,
      method: 'get',
      params: {
        page: paging.page,
        size: paging.size,
        sortBy: sorting.byField,
        sortDir: sorting.isDesc? 'desc': 'asc', {{range .Fields}}
        f_{{CamelCase .Name}}: filtering.{{CamelCase .Name}}, {{end}}
      }
    })
  },    

  Update{{PascalCase .Name}}({{CamelCase .Name}}ID, payload) {
    return request({
      url: `/api/{{CamelCase .Name}}/${{"{"}}{{CamelCase .Name}}ID}`,
      method: 'put',
      data: payload,
    })
  },

  Delete{{PascalCase .Name}}({{CamelCase .Name}}ID) {
    return request({
      url: `/api/{{CamelCase .Name}}/${{"{"}}{{CamelCase .Name}}ID}`,
      method: 'delete',
    })
  },  

}

export default api
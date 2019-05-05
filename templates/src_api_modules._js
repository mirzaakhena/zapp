import request from "@/utils/httprequest";

const api = {

  Create{{PascalCase .Name}}(payload) {
    return request({
      url: `/{{CamelCase .Name}}`,
      method: 'post',
      data: payload,
    })
  },

  GetOne{{PascalCase .Name}}({{CamelCase .Name}}ID) {
    return request({
      url: `/{{CamelCase .Name}}/${ {{CamelCase .Name}}ID}`,
      method: 'get',
    })
  },  

  GetAll{{PascalCase .Name}}(payload) {    
    return request({
      url: `/{{CamelCase .Name}}`,
      method: 'get',
      params: {
        page: payload.page,
        size: payload.size,
        sortBy: payload.sort.sortBy,
        sortDir: payload.sort.sortDir, {{range .Fields}}
        f_{{CamelCase .Name}}: payload.filters.{{CamelCase .Name}}, {{end}}
      }
    })
  },    

  Update{{PascalCase .Name}}({{CamelCase .Name}}ID, payload) {
    return request({
      url: `/{{CamelCase .Name}}/${ {{CamelCase .Name}}ID}`,
      method: 'put',
      data: payload,
    })
  },

  Delete{{PascalCase .Name}}({{CamelCase .Name}}ID) {
    return request({
      url: `/{{CamelCase .Name}}/${ {{CamelCase .Name}}ID}`,
      method: 'delete',
    })
  },  

}

export default api
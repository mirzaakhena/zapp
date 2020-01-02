{{range .Entities}}
import {{CamelCase .Name}} from "./modules/{{LowerCase .Name}}"; {{end}}

const api = { {{range .Entities}}  
  ...{{CamelCase .Name}}, {{end}}
}

export default api
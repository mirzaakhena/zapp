{{range .Classes}}
import {{CamelCase .Name}} from "./modules/{{LowerCase .Name}}"; {{end}}

const api = { {{range .Classes}}  
  ...{{CamelCase .Name}}, {{end}}
}

export default api
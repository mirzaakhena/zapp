package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"unicode"

	"github.com/alecthomas/template"
	"gopkg.in/yaml.v2"
)

func main() {
	// content, err := ioutil.ReadFile("skrip.yaml")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// tp := ThePackage{}
	// err = yaml.Unmarshal(content, &tp)
	// if err != nil {
	// 	log.Fatalf("error: %+v", err)
	// }
	// fmt.Printf("%+v\n", tp)

	processIt()
}

func processIt() {

	content, err := ioutil.ReadFile("skrip.yaml")
	if err != nil {
		log.Fatal(err)
	}

	tp := ThePackage{}
	err = yaml.Unmarshal(content, &tp)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	enumsMap := map[string]TheEnum{}

	for _, e := range tp.Enums {
		enumsMap[e.Name] = e

		for i := 0; i < len(e.Values); i++ {
			if len(strings.TrimSpace(e.Values[i].Value)) == 0 {
				text := e.Values[i].Text
				e.Values[i] = TextAndValue{
					Text:  text,
					Value: text,
				}
			}
		}
	}

	for i := 0; i < len(tp.Entities); i++ {
		for j := 0; j < len(tp.Entities[i].Fields); j++ {
			if tp.Entities[i].Fields[j].DataType == "entity" {
				tp.Entities[i].HasAutocomplete = true
				break
			}
		}

		for j := 0; j < len(tp.Entities[i].Fields); j++ {
			if tp.Entities[i].Fields[j].DataType == "enum" {
				tp.Entities[i].Fields[j].EnumValues = enumsMap[tp.Entities[i].Fields[j].EnumReference].Values
			}
		}
	}

	tp.Run()

	{
		fmt.Println("go fmt")
		cmd := exec.Command("go", "fmt", fmt.Sprintf("%s/...", tp.PackagePath))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}

	{
		fmt.Println("go get")
		cmd := exec.Command("go", "get", fmt.Sprintf("./..."))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}

}

// TextAndValue is
type TextAndValue struct {
	Text  string `yaml:"text"`
	Value string `yaml:"value"`
}

// TheField is
type TheField struct {
	Name            string `yaml:"name"`
	DataType        string `yaml:"dataType"`
	EnumReference   string `yaml:"enumReference"`
	EntityReference string `yaml:"entityReference"`
	EntityField     string `yaml:"entityField"`
	DefaultValue    string `yaml:"defaultValue"`
	Sortable        string `yaml:"sortable"`
	Filterable      string `yaml:"filterable"`
	Regex           string `yaml:"regex"`

	EnumValues []TextAndValue
}

// TheClass is
type TheClass struct {
	Name            string     `yaml:"name"`
	Fields          []TheField `yaml:"fields"`
	TableName       string     `yaml:"tableName"`
	HasAutocomplete bool
}

// TheEnum is
type TheEnum struct {
	Name   string         `yaml:"name"`
	Values []TextAndValue `yaml:"values"`
}

// ThePackage is
type ThePackage struct {
	ApplicationName string     ``
	PackagePath     string     `yaml:"packagePath"`
	Entities        []TheClass `yaml:"entities"`
	Enums           []TheEnum  `yaml:"enums"`
}

// CamelCase is
func CamelCase(name string) string {

	// force it!
	if name == "IPAddress" {
		return "ipAddress"
	}

	if name == "ID" {
		return "id"
	}

	out := []rune(name)
	out[0] = unicode.ToLower([]rune(name)[0])
	return string(out)
}

// UpperCase is
func UpperCase(name string) string {
	return strings.ToUpper(name)
}

// LowerCase is
func LowerCase(name string) string {
	return strings.ToLower(name)
}

// PascalCase is
func PascalCase(name string) string {
	return name
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// SnakeCase is
func SnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// HasTime is
func HasTime(dataTypes []TheField) bool {
	for _, tm := range dataTypes {
		if tm.DataType == "time.Time" {
			return true
		}
	}
	return false
}

// Run is
func (tp *ThePackage) Run() {

	s := strings.Split(tp.PackagePath, "/")
	tp.ApplicationName = s[len(s)-1]

	// single directory generated
	{

		// backend
		{
			{
				dir := fmt.Sprintf("../../../../%s/server/app", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/server/controller", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/server/model", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/server/service", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/server/repository", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}

			{
				dir := ""

				dir = fmt.Sprintf("../../../../%s/server/utils/common", tp.PackagePath)
				os.MkdirAll(dir, 0777)

				dir = fmt.Sprintf("../../../../%s/server/utils/token", tp.PackagePath)
				os.MkdirAll(dir, 0777)

				dir = fmt.Sprintf("../../../../%s/server/utils/config", tp.PackagePath)
				os.MkdirAll(dir, 0777)

				dir = fmt.Sprintf("../../../../%s/server/utils/log", tp.PackagePath)
				os.MkdirAll(dir, 0777)

				dir = fmt.Sprintf("../../../../%s/server/utils/transaction", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}
		}

		// frontend
		{
			{
				dir := fmt.Sprintf("../../../../%s/client/public", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/client/src/assets", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/client/src/store/modules", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/client/src/router/modules", tp.PackagePath)
				os.MkdirAll(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/client/src/pages", tp.PackagePath)
				os.Mkdir(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/client/dist", tp.PackagePath)
				os.Mkdir(dir, 0777)
			}

			{
				dir := fmt.Sprintf("../../../../%s/client/src/utils", tp.PackagePath)
				os.Mkdir(dir, 0777)
			}
		}

	}

	// single file generated
	{
		// backend
		{
			{
				templateFile := fmt.Sprintf("../templates/backend/app/main._go")
				outputFile := fmt.Sprintf("../../../../%s/server/app/main.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/model/model_basic._go")
				outputFile := fmt.Sprintf("../../../../%s/server/model/system.model.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/model/user-model._go")
				outputFile := fmt.Sprintf("../../../../%s/server/model/system.user.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/repository/user-repository._go")
				outputFile := fmt.Sprintf("../../../../%s/server/repository/system.user.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/service/user-service._go")
				outputFile := fmt.Sprintf("../../../../%s/server/service/system.user.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/controller/user-controller._go")
				outputFile := fmt.Sprintf("../../../../%s/server/controller/system.user.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/controller/router._go")
				outputFile := fmt.Sprintf("../../../../%s/server/controller/system.router.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/utils/common/identifier._go")
				outputFile := fmt.Sprintf("../../../../%s/server/utils/common/identifier.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/utils/common/password._go")
				outputFile := fmt.Sprintf("../../../../%s/server/utils/common/password.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/utils/common/json._go")
				outputFile := fmt.Sprintf("../../../../%s/server/utils/common/json.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/utils/common/strings._go")
				outputFile := fmt.Sprintf("../../../../%s/server/utils/common/strings.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/utils/token/jwt._go")
				outputFile := fmt.Sprintf("../../../../%s/server/utils/token/jwt.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/utils/config/config._go")
				outputFile := fmt.Sprintf("../../../../%s/server/utils/config/config.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/utils/log/log._go")
				outputFile := fmt.Sprintf("../../../../%s/server/utils/log/log.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/utils/transaction/transaction._go")
				outputFile := fmt.Sprintf("../../../../%s/server/utils/transaction/transaction.go", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/._gitignore")
				outputFile := fmt.Sprintf("../../../../%s/.gitignore", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/config._toml")
				outputFile := fmt.Sprintf("../../../../%s/server/config.toml", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/build._sh")
				outputFile := fmt.Sprintf("../../../../%s/server/build.sh", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0755)
			}

			{
				templateFile := fmt.Sprintf("../templates/README._md")
				outputFile := fmt.Sprintf("../../../../%s/README.md", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}
		}

		// frontend
		{
			{
				templateFile := fmt.Sprintf("../templates/frontend/public/favicon._ico")
				outputFile := fmt.Sprintf("../../../../%s/client/public/favicon.ico", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/public/index._html")
				outputFile := fmt.Sprintf("../../../../%s/client/public/index.html", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/babel.config._js")
				outputFile := fmt.Sprintf("../../../../%s/client/babel.config.js", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/vue.config._js")
				outputFile := fmt.Sprintf("../../../../%s/client/vue.config.js", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

				templateFile := fmt.Sprintf("../templates/frontend/package._json")
				outputFile := fmt.Sprintf("../../../../%s/client/package.json", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/App._vue")
				outputFile := fmt.Sprintf("../../../../%s/client/src/App.vue", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/assets/style._css")
				outputFile := fmt.Sprintf("../../../../%s/client/src/assets/style.css", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/dist/index._html")
				outputFile := fmt.Sprintf("../../../../%s/client/dist/index.html", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/pages/forgotpassword._vue")
				outputFile := fmt.Sprintf("../../../../%s/client/src/pages/forgotpassword.vue", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/pages/home._vue")
				outputFile := fmt.Sprintf("../../../../%s/client/src/pages/home.vue", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/pages/login._vue")
				outputFile := fmt.Sprintf("../../../../%s/client/src/pages/login.vue", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/pages/notfound._vue")
				outputFile := fmt.Sprintf("../../../../%s/client/src/pages/notfound.vue", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/pages/register._vue")
				outputFile := fmt.Sprintf("../../../../%s/client/src/pages/register.vue", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/pages/successregister._vue")
				outputFile := fmt.Sprintf("../../../../%s/client/src/pages/successregister.vue", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/main._js")
				outputFile := fmt.Sprintf("../../../../%s/client/src/main.js", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/store/index._js")
				outputFile := fmt.Sprintf("../../../../%s/client/src/store/index.js", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/store/crudtable._js")
				outputFile := fmt.Sprintf("../../../../%s/client/src/store/crudtable.js", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/router/index._js")
				outputFile := fmt.Sprintf("../../../../%s/client/src/router/index.js", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/utils/httprequest._js")
				outputFile := fmt.Sprintf("../../../../%s/client/src/utils/httprequest.js", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/utils/auth._js")
				outputFile := fmt.Sprintf("../../../../%s/client/src/utils/auth.js", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/frontend/src/utils/filter._js")
				outputFile := fmt.Sprintf("../../../../%s/client/src/utils/filter.js", tp.PackagePath)
				basic(tp, templateFile, outputFile, tp, 0664)
			}
		}

	}

	// multiple generated file
	for _, et := range tp.Entities {

		// backend
		{
			{
				templateFile := fmt.Sprintf("../templates/backend/repository/repository._go")
				outputFile := fmt.Sprintf("../../../../%s/server/repository/%s.go", tp.PackagePath, LowerCase(et.Name))
				basic(tp, templateFile, outputFile, et, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/controller/controller._go")
				outputFile := fmt.Sprintf("../../../../%s/server/controller/%s.go", tp.PackagePath, LowerCase(et.Name))
				basic(tp, templateFile, outputFile, et, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/model/model._go")
				outputFile := fmt.Sprintf("../../../../%s/server/model/%s.go", tp.PackagePath, LowerCase(et.Name))
				basic(tp, templateFile, outputFile, et, 0664)
			}

			{
				templateFile := fmt.Sprintf("../templates/backend/service/service._go")
				outputFile := fmt.Sprintf("../../../../%s/server/service/%s.go", tp.PackagePath, LowerCase(et.Name))
				basic(tp, templateFile, outputFile, et, 0664)
			}
		}

		// frontend
		{

			// create store modules
			// {
			// 	templateFile := fmt.Sprintf("../templates/frontend/src/store/modules/file._js")
			// 	outputFile := fmt.Sprintf("../../../../%s/webapp/src/store/modules/%s.js", tp.PackagePath, LowerCase(et.Name))
			// 	basic(tp, templateFile, outputFile, et, 0664)
			// }

			// create router modules
			{
				templateFile := fmt.Sprintf("../templates/frontend/src/router/modules/file._js")
				outputFile := fmt.Sprintf("../../../../%s/client/src/router/modules/%s.js", tp.PackagePath, LowerCase(et.Name))
				basic(tp, templateFile, outputFile, et, 0664)
			}

			{
				// create folder pages
				f := fmt.Sprintf("../../../../%s/client/src/pages/%s", tp.PackagePath, LowerCase(et.Name))
				os.Mkdir(f, 0777)

				// create file table under folder
				{
					templateFile := fmt.Sprintf("../templates/frontend/src/pages/folder/list._vue")
					outputFile := fmt.Sprintf("../../../../%s/client/src/pages/%s/list.vue", tp.PackagePath, LowerCase(et.Name))
					basic(tp, templateFile, outputFile, et, 0664)
				}

				// create file input under folder
				{
					templateFile := fmt.Sprintf("../templates/frontend/src/pages/folder/input._vue")
					outputFile := fmt.Sprintf("../../../../%s/client/src/pages/%s/input.vue", tp.PackagePath, LowerCase(et.Name))
					basic(tp, templateFile, outputFile, et, 0664)
				}

			}

		}

	}

	fmt.Printf(">>>>> done Run\n")
}

func basic(pkg *ThePackage, templateFile, outputFile string, object interface{}, perm os.FileMode) {

	fmt.Println(outputFile)
	file, err := os.Open(templateFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var buffer bytes.Buffer

	for scanner.Scan() {
		row := scanner.Text()
		buffer.WriteString(row)
		buffer.WriteString("\n")
	}

	FuncMap := template.FuncMap{
		"HasTime":     HasTime,
		"CamelCase":   CamelCase,
		"PascalCase":  PascalCase,
		"SnakeCase":   SnakeCase,
		"UpperCase":   UpperCase,
		"LowerCase":   LowerCase,
		"AppName":     func() string { return pkg.ApplicationName },
		"PackagePath": func() string { return pkg.PackagePath },
	}

	t, err := template.
		New("todos").
		Funcs(FuncMap).
		Parse(buffer.String())

	if err != nil {
		panic(err)
	}

	var bf bytes.Buffer
	err = t.Execute(&bf, object)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(outputFile, bf.Bytes(), perm)
	if err != nil {
		panic(err)
	}

}

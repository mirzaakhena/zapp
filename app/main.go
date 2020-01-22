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
	Name   string     `yaml:"name"`
	Fields []TheField `yaml:"fields"`

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

	// create app folder
	{
		dir := fmt.Sprintf("../../../../%s/app", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	// create controller folder
	{
		dir := fmt.Sprintf("../../../../%s/controller", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	// create model folder
	{
		dir := fmt.Sprintf("../../../../%s/model", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	// create service folder
	{
		dir := fmt.Sprintf("../../../../%s/service", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	// create repository folder
	{
		dir := fmt.Sprintf("../../../../%s/repository", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	// create utils folder
	{
		dir := ""

		dir = fmt.Sprintf("../../../../%s/utils/common", tp.PackagePath)
		os.MkdirAll(dir, 0777)

		dir = fmt.Sprintf("../../../../%s/utils/token", tp.PackagePath)
		os.MkdirAll(dir, 0777)

		dir = fmt.Sprintf("../../../../%s/utils/config", tp.PackagePath)
		os.MkdirAll(dir, 0777)

		dir = fmt.Sprintf("../../../../%s/utils/log", tp.PackagePath)
		os.MkdirAll(dir, 0777)

		dir = fmt.Sprintf("../../../../%s/utils/transaction", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/webapp/public", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/webapp/src/api", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/webapp/src/assets", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/webapp/src/store/modules", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/webapp/src/router/modules", tp.PackagePath)
		os.MkdirAll(dir, 0777)
	}

	{
		f := fmt.Sprintf("../../../../%s/webapp/src/pages", tp.PackagePath)
		os.Mkdir(f, 0777)
	}

	// {
	// 	f := fmt.Sprintf("../../../../%s/webapp/src/components", tp.PackagePath)
	// 	os.Mkdir(f, 0777)
	// }

	{
		f := fmt.Sprintf("../../../../%s/webapp/src/utils", tp.PackagePath)
		os.Mkdir(f, 0777)
	}

	for _, et := range tp.Entities {

		// create repository
		{
			templateFile := fmt.Sprintf("../templates/backend/repository._go")
			outputFile := fmt.Sprintf("../../../../%s/repository/%s.go", tp.PackagePath, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create controller
		{
			templateFile := fmt.Sprintf("../templates/backend/controller._go")
			outputFile := fmt.Sprintf("../../../../%s/controller/%s.go", tp.PackagePath, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create model
		{
			templateFile := fmt.Sprintf("../templates/backend/model._go")
			outputFile := fmt.Sprintf("../../../../%s/model/%s.go", tp.PackagePath, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create service
		{
			templateFile := fmt.Sprintf("../templates/backend/service._go")
			outputFile := fmt.Sprintf("../../../../%s/service/%s.go", tp.PackagePath, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create common api
		{
			templateFile := fmt.Sprintf("../templates/frontend/src/api/restapi._js")
			outputFile := fmt.Sprintf("../../../../%s/webapp/src/api/restapi.js", tp.PackagePath)
			basic(tp, templateFile, outputFile, et)
		}

		// create api modules
		// {
		// 	templateFile := fmt.Sprintf("../templates/frontend/src/api/modules/file._js")
		// 	outputFile := fmt.Sprintf("../../../../%s/webapp/src/api/modules/%s.js", tp.PackagePath, LowerCase(et.Name))
		// 	basic(tp, templateFile, outputFile, et)
		// }

		// create store modules
		// {
		// 	templateFile := fmt.Sprintf("../templates/frontend/src/store/modules/file._js")
		// 	outputFile := fmt.Sprintf("../../../../%s/webapp/src/store/modules/%s.js", tp.PackagePath, LowerCase(et.Name))
		// 	basic(tp, templateFile, outputFile, et)
		// }

		// create router modules
		{
			templateFile := fmt.Sprintf("../templates/frontend/src/router/modules/file._js")
			outputFile := fmt.Sprintf("../../../../%s/webapp/src/router/modules/%s.js", tp.PackagePath, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		{
			// create folder pages
			f := fmt.Sprintf("../../../../%s/webapp/src/pages/%s", tp.PackagePath, LowerCase(et.Name))
			os.Mkdir(f, 0777)

			// create file table under folder
			{
				templateFile := fmt.Sprintf("../templates/frontend/src/pages/folder/list._vue")
				outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/%s/list.vue", tp.PackagePath, LowerCase(et.Name))
				basic(tp, templateFile, outputFile, et)
			}

			// create file input under folder
			{
				templateFile := fmt.Sprintf("../templates/frontend/src/pages/folder/input._vue")
				outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/%s/input.vue", tp.PackagePath, LowerCase(et.Name))
				basic(tp, templateFile, outputFile, et)
			}

		}

	}

	{
		templateFile := fmt.Sprintf("../templates/backend/main._go")
		outputFile := fmt.Sprintf("../../../../%s/app/main.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/model_basic._go")
		outputFile := fmt.Sprintf("../../../../%s/model/model.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/user-model._go")
		outputFile := fmt.Sprintf("../../../../%s/model/user.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/user-repository._go")
		outputFile := fmt.Sprintf("../../../../%s/repository/user.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/user-service._go")
		outputFile := fmt.Sprintf("../../../../%s/service/user.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/user-controller._go")
		outputFile := fmt.Sprintf("../../../../%s/controller/user.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/router._go")
		outputFile := fmt.Sprintf("../../../../%s/controller/router.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/utils/common/identifier._go")
		outputFile := fmt.Sprintf("../../../../%s/utils/common/identifier.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/utils/common/password._go")
		outputFile := fmt.Sprintf("../../../../%s/utils/common/password.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/utils/common/json._go")
		outputFile := fmt.Sprintf("../../../../%s/utils/common/json.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/utils/common/strings._go")
		outputFile := fmt.Sprintf("../../../../%s/utils/common/strings.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/utils/token/jwt._go")
		outputFile := fmt.Sprintf("../../../../%s/utils/token/jwt.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/utils/config/config._go")
		outputFile := fmt.Sprintf("../../../../%s/utils/config/config.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/utils/log/log._go")
		outputFile := fmt.Sprintf("../../../../%s/utils/log/log.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/utils/transaction/transaction._go")
		outputFile := fmt.Sprintf("../../../../%s/utils/transaction/transaction.go", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/._gitignore")
		outputFile := fmt.Sprintf("../../../../%s/.gitignore", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/config._toml")
		outputFile := fmt.Sprintf("../../../../%s/config.toml", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/README._md")
		outputFile := fmt.Sprintf("../../../../%s/README.md", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/public/favicon._ico")
		outputFile := fmt.Sprintf("../../../../%s/webapp/public/favicon.ico", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/public/index._html")
		outputFile := fmt.Sprintf("../../../../%s/webapp/public/index.html", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/babel.config._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/babel.config.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/vue.config._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/vue.config.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/package._json")
		outputFile := fmt.Sprintf("../../../../%s/webapp/package.json", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/App._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/App.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/assets/style._css")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/assets/style.css", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/forgotpassword._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/forgotpassword.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/home._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/home.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/login._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/login.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/notfound._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/notfound.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/register._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/register.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/successregister._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/successregister.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/main._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/main.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	// {
	// 	templateFile := fmt.Sprintf("../templates/frontend/src/api/index._js")
	// 	outputFile := fmt.Sprintf("../../../../%s/webapp/src/api/index.js", tp.PackagePath)
	// 	basic(tp, templateFile, outputFile, tp)
	// }

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/store/index._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/store/index.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/store/crudtable._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/store/crudtable.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/router/index._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/router/index.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/utils/httprequest._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/utils/httprequest.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/utils/auth._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/utils/auth.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/utils/filter._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/utils/filter.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp)
	}

	fmt.Printf(">>>>> done Run\n")
}

func basic(pkg *ThePackage, templateFile, outputFile string, object interface{}) {

	fmt.Println(templateFile)
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

	err = ioutil.WriteFile(outputFile, bf.Bytes(), 0664)
	if err != nil {
		panic(err)
	}

}

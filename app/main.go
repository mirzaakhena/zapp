package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/alecthomas/template"
)

type TheField struct {
	Name     string
	DataType string
}

type TheClass struct {
	Name   string
	Fields []TheField
}

type ThePackage struct {
	packageName string
	datatypeid  string
	Classes     []TheClass
}

func CamelCase(name string) string {

	// force it!
	if name == "IPAddress" {
		return "ipAddress"
	}

	out := []rune(name)
	out[0] = unicode.ToLower([]rune(name)[0])
	return string(out)
}

func UpperCase(name string) string {
	return strings.ToUpper(name)
}

func LowerCase(name string) string {
	return strings.ToLower(name)
}

func PascalCase(name string) string {
	return name
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func SnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func HasTime(dataTypes []TheField) bool {
	for _, tm := range dataTypes {
		if tm.DataType == "time.Time" {
			return true
		}
	}
	return false
}

func main() {

	filename := "skrip.txt"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tp := ThePackage{}

	var theClass *TheClass

	for scanner.Scan() {
		row := scanner.Text()

		if strings.TrimSpace(row) == "" {
			continue
		}

		if strings.HasPrefix(row, "package") {
			pkgs := strings.Split(row, ",")
			tp.packageName = strings.TrimSpace(pkgs[1])
			tp.Classes = []TheClass{}
			continue
		}

		if strings.HasPrefix(row, "datatypeid") {
			dtt := strings.Split(row, ",")
			tp.datatypeid = strings.TrimSpace(dtt[1])
			continue
		}

		if strings.HasPrefix(row, "class") {
			clName := strings.Split(row, " ")
			theClass = &TheClass{
				Name:   strings.TrimSpace(clName[1]),
				Fields: []TheField{},
			}
			continue
		}

		if strings.HasPrefix(row, "endclass") {
			tp.Classes = append(tp.Classes, *theClass)
			theClass = nil
			continue
		}

		if strings.HasPrefix(row, "field") {
			flField := strings.Split(row, ",")
			theClass.Fields = append(theClass.Fields, TheField{
				Name:     strings.TrimSpace(flField[1]),
				DataType: strings.TrimSpace(flField[2]),
			})
			continue
		}

	}

	tp.Run()
}

func (tp *ThePackage) Run() {

	// create app folder
	{
		dir := fmt.Sprintf("../../../../%s/app", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	// create config folder
	{
		dir := fmt.Sprintf("../../../../%s/config", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	// create controller folder
	{
		dir := fmt.Sprintf("../../../../%s/controller", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	// create dao folder
	{
		dir := fmt.Sprintf("../../../../%s/dao", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	// create database folder
	{
		dir := fmt.Sprintf("../../../../%s/database", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	// create logger folder
	{
		dir := fmt.Sprintf("../../../../%s/logger", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	// create model folder
	{
		dir := fmt.Sprintf("../../../../%s/model", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	// create service folder
	{
		dir := fmt.Sprintf("../../../../%s/service", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/utils", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/webapp/public", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/webapp/src/api/modules", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/webapp/src/store/modules", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	{
		dir := fmt.Sprintf("../../../../%s/webapp/src/router/modules", tp.packageName)
		os.MkdirAll(dir, 0777)
	}

	{
		f := fmt.Sprintf("../../../../%s/webapp/src/components", tp.packageName)
		os.Mkdir(f, 0777)
	}

	{
		f := fmt.Sprintf("../../../../%s/webapp/src/utils", tp.packageName)
		os.Mkdir(f, 0777)
	}

	for _, et := range tp.Classes {

		// create dao
		{
			templateFile := fmt.Sprintf("../templates/dao._go")
			outputFile := fmt.Sprintf("../../../../%s/dao/%s.go", tp.packageName, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create controller
		{
			templateFile := fmt.Sprintf("../templates/controller._go")
			outputFile := fmt.Sprintf("../../../../%s/controller/%s.go", tp.packageName, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create model
		{
			templateFile := fmt.Sprintf("../templates/model._go")
			outputFile := fmt.Sprintf("../../../../%s/model/%s.go", tp.packageName, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create service
		{
			templateFile := fmt.Sprintf("../templates/service._go")
			outputFile := fmt.Sprintf("../../../../%s/service/%s.go", tp.packageName, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create api modules
		{
			templateFile := fmt.Sprintf("../templates/src_api_modules._js")
			outputFile := fmt.Sprintf("../../../../%s/webapp/src/api/modules/%s.js", tp.packageName, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create store modules
		{
			templateFile := fmt.Sprintf("../templates/src_store_modules._js")
			outputFile := fmt.Sprintf("../../../../%s/webapp/src/store/modules/%s.js", tp.packageName, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create router modules
		{
			templateFile := fmt.Sprintf("../templates/src_router_modules._js")
			outputFile := fmt.Sprintf("../../../../%s/webapp/src/router/modules/%s.js", tp.packageName, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create file table under folder
		{
			templateFile := fmt.Sprintf("../templates/src_pages_folders_table._vue")
			outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/%s/table.vue", tp.packageName, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

		// create file input under folder
		{
			templateFile := fmt.Sprintf("../templates/src_pages_folders_input._vue")
			outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/%s/input.vue", tp.packageName, LowerCase(et.Name))
			basic(tp, templateFile, outputFile, et)
		}

	}

	{
		templateFile := fmt.Sprintf("../templates/main._go")
		outputFile := fmt.Sprintf("../../../../%s/app/main.go", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/config._go")
		outputFile := fmt.Sprintf("../../../../%s/config/config.go", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/transaction._go")
		outputFile := fmt.Sprintf("../../../../%s/database/transaction.go", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/justlog._go")
		outputFile := fmt.Sprintf("../../../../%s/logger/justlog.go", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/utils._go")
		outputFile := fmt.Sprintf("../../../../%s/utils/utils.go", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/._gitignore")
		outputFile := fmt.Sprintf("../../../../%s/.gitignore", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/config._toml")
		outputFile := fmt.Sprintf("../../../../%s/config.toml", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/README._md")
		outputFile := fmt.Sprintf("../../../../%s/README.md", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/public_favicon._ico")
		outputFile := fmt.Sprintf("../../../../%s/webapp/public/favicon.ico", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/public_index._html")
		outputFile := fmt.Sprintf("../../../../%s/webapp/public/index.html", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/babel.config._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/babel.config.js", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/vue.config._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/vue.config.js", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/package._json")
		outputFile := fmt.Sprintf("../../../../%s/webapp/package.json", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_App._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/App.vue", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_pages_forgotpassword._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/forgotpassword.vue", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_pages_home._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/home.vue", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_pages_login._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/login.vue", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_pages_notfound._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/notfound.vue", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_pages_register._vue")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/pages/register.vue", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_main._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/main.js", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_api_index._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/api/index.js", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_store_index._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/store/index.js", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_router_index._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/router/index.js", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

	{
		templateFile := fmt.Sprintf("../templates/src_utils_httprequest._js")
		outputFile := fmt.Sprintf("../../../../%s/webapp/src/utils/httprequest.js", tp.packageName)
		basic(tp, templateFile, outputFile, tp)
	}

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
		"PackageName": func() string { return pkg.packageName },
		"DataTypeId":  func() string { return pkg.datatypeid },
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

	ioutil.WriteFile(outputFile, bf.Bytes(), 0664)
	if err != nil {
		panic(err)
	}
}

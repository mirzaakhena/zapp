package process

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/mirzaakhena/zapp/app/model"
)

func basic(pkg *model.ThePackage, templateFile, outputFile string, object interface{}, perm os.FileMode) {

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
		"UniqueFields": GetUniqueFields,
		"HasTime":      HasTime,
		"CamelCase":    CamelCase,
		"PascalCase":   PascalCase,
		"SnakeCase":    SnakeCase,
		"UpperCase":    UpperCase,
		"LowerCase":    LowerCase,
		"AppName":      func() string { return pkg.ApplicationName },
		"PackagePath":  func() string { return pkg.PackagePath },
	}

	t, err := template.
		New("todos").
		Funcs(FuncMap).
		Parse(buffer.String())

	if err != nil {
		panic(err)
	}

	var bf bytes.Buffer
	if err := t.Execute(&bf, object); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(outputFile, bf.Bytes(), perm); err != nil {
		panic(err)
	}

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
func HasTime(dataTypes []model.TheField) bool {
	for _, tm := range dataTypes {
		if tm.DataType == "time.Time" {
			return true
		}
	}
	return false
}

// GetUniqueFields is
func GetUniqueFields(dataTypes []model.TheField, currentEntityName string) []model.TheField {
	uniqueFields := []model.TheField{}
	existing := map[string]model.TheField{}
	for _, tm := range dataTypes {
		_, exist := existing[tm.EntityReference]

		if currentEntityName != "" && tm.EntityReference == currentEntityName {
			continue
		}

		if !exist {
			existing[tm.EntityReference] = tm
			uniqueFields = append(uniqueFields, tm)
		}
	}
	return uniqueFields
}

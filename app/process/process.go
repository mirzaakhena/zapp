package process

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/mirzaakhena/zapp/app/model"
	"gopkg.in/yaml.v2"
)

// RunProcess is
func RunProcess(file string) {

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// prepare root object
	tp := model.ThePackage{}

	// read yaml file
	if err = yaml.Unmarshal(content, &tp); err != nil {
		log.Fatalf("error: %+v", err)
	}

	// prepare enum maps
	enumsMap := map[string]model.TheEnum{}

	// read enum
	for _, e := range tp.Enums {
		enumsMap[e.Name] = e

		for i := 0; i < len(e.Values); i++ {

			// check if value is empty
			if len(strings.TrimSpace(e.Values[i].Value)) == 0 {
				text := e.Values[i].Text

				// then we will use text as value
				e.Values[i] = model.TextAndValue{
					Text:  text,
					Value: text,
				}
			}
		}
	}

	// read entity
	for i := 0; i < len(tp.Entities); i++ {

		// read every field
		for j := 0; j < len(tp.Entities[i].Fields); j++ {

			// used to insert "import" autocomplete component in src/pages/<entity> if entity use entityReference
			if tp.Entities[i].Fields[j].DataType == "entity" {
				tp.Entities[i].HasAutocomplete = true

				// no need to break because we need to include all of component value
			}

			// used to insert "import" in service.go if entity use enumReference
			if tp.Entities[i].Fields[j].DataType == "enum" {
				tp.Entities[i].HasEnum = true

				// break because we only need to import once
				break
			}
		}

		// separate this loop because it must apply to all fields
		// insert all enumMap into enumValues object.
		for j := 0; j < len(tp.Entities[i].Fields); j++ {
			if tp.Entities[i].Fields[j].DataType == "enum" {
				tp.Entities[i].Fields[j].EnumValues = enumsMap[tp.Entities[i].Fields[j].EnumReference].Values
			}
		}
	}

	// set the application name from package path
	{
		s := strings.Split(tp.PackagePath, "/")
		tp.ApplicationName = s[len(s)-1]
	}

	// single directory generated
	{
		makeBackendDirectory(&tp)
		makeFrontendDirectory(&tp)
	}

	// single static file generated
	{
		createStaticBackendFile(&tp)
		createStaticFrontendFile(&tp)
	}

	// multiple generated file for entity
	for _, et := range tp.Entities {
		createDynamicBackendFile(&tp, &et)
		createDynamicFrontendFile(&tp, &et)
	}

	// multiple generated file for enum
	for _, en := range tp.Enums {
		createEnumFile(&tp, &en)
	}

	fmt.Printf(">>>>> done Run\n")

	goFormat(tp.PackagePath)

	goGet()

}

// do go format
func goFormat(path string) {
	fmt.Println("go fmt")
	cmd := exec.Command("go", "fmt", fmt.Sprintf("%s/...", path))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

// call go get
func goGet() {
	fmt.Println("go get")
	cmd := exec.Command("go", "get", fmt.Sprintf("./..."))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

package process

import (
	"fmt"
	"strings"

	"github.com/mirzaakhena/zapp/app/model"
)

func createDynamicBackendFile(tp *model.ThePackage, et *model.TheEntity) {

	{
		templateFile := fmt.Sprintf("../templates/backend/repository/repository._go")
		outputFile := fmt.Sprintf("../../../../%s/backend/repository/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/controller/restapi/controller._go")
		outputFile := fmt.Sprintf("../../../../%s/backend/controller/restapi/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/model/model._go")
		outputFile := fmt.Sprintf("../../../../%s/backend/model/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/service/crud/service._go")
		outputFile := fmt.Sprintf("../../../../%s/backend/service/crud/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

}

func createEnumFile(tp *model.ThePackage, en *model.TheEnum) {
	{
		templateFile := fmt.Sprintf("../templates/backend/model/enum/enum._go")
		outputFile := fmt.Sprintf("../../../../%s/backend/model/enum/%s.go", tp.PackagePath, strings.ToLower(en.Name))
		basic(tp, templateFile, outputFile, en, 0664)
	}
}

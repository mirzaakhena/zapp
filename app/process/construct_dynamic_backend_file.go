package process

import (
	"fmt"
	"strings"

	"github.com/mirzaakhena/zapp/app/model"
)

func createDynamicBackendFile(tp *model.ThePackage, et *model.TheEntity) {

	{
		templateFile := fmt.Sprintf("../templates/backend/repository/repository._go")
		outputFile := fmt.Sprintf("../../../../%s/server/repository/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/controller/controller._go")
		outputFile := fmt.Sprintf("../../../../%s/server/controller/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/model/model._go")
		outputFile := fmt.Sprintf("../../../../%s/server/model/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/backend/service/service._go")
		outputFile := fmt.Sprintf("../../../../%s/server/service/%s.go", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

}

func createEnumFile(tp *model.ThePackage, en *model.TheEnum) {
	{
		templateFile := fmt.Sprintf("../templates/backend/model/enum/enum._go")
		outputFile := fmt.Sprintf("../../../../%s/server/model/enum/%s.go", tp.PackagePath, strings.ToLower(en.Name))
		basic(tp, templateFile, outputFile, en, 0664)
	}
}

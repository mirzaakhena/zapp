package process

import (
	"fmt"

	"github.com/mirzaakhena/zapp/app/model"
)

func createStaticBackendFile(tp *model.ThePackage) {
	// backend
	{
		{
			templateFile := fmt.Sprintf("../templates/backend/app/main._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/app/main.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/model/model_basic._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/model/system-model.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/model/user-model._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/model/system-user.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/repository/user-repository._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/repository/system-user.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/service/auth/user-service._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/service/auth/system-user.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/controller/restapi/user-controller._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/controller/restapi/system-user.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/controller/restapi/router._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/controller/restapi/system-router.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/utils/identifier._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/utils/identifier.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/utils/password._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/utils/password.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/utils/json._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/utils/json.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/utils/strings._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/utils/strings.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/token/contract._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/token/contract.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/token/implementation._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/token/implementation.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/token/public._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/token/public.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/config/contract._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/config/contract.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/config/implementation._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/config/implementation.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/config/public._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/config/public.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/log/contract._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/log/contract.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/log/implementation._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/log/implementation.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/log/public._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/log/public.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/transaction/contract._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/transaction/contract.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/transaction/implementation._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/transaction/implementation.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/transaction/public._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/transaction/public.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/messagebroker/implementation-producer._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/implementation-producer.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/messagebroker/implementation-consumer._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/implementation-consumer.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/messagebroker/contract._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/contract.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/messagebroker/public._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/messagebroker/public.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/httpclient/contract._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/httpclient/contract.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/httpclient/implementation._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/httpclient/implementation.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/httpclient/public._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/httpclient/public.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/error/error_code._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/error/error_code.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/error/error._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/error/error.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/error/error_test._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/error/error_test.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/constant/constant._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/constant/constant.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/extractor/extractor._go")
			outputFile := fmt.Sprintf("../../../../%s/backend/shared/extractor/extractor.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/._gitignore")
			outputFile := fmt.Sprintf("../../../../%s/.gitignore", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/config._toml")
			outputFile := fmt.Sprintf("../../../../%s/backend/config.toml", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/build._sh")
			outputFile := fmt.Sprintf("../../../../%s/backend/build.sh", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0755)
		}

		{
			templateFile := fmt.Sprintf("../templates/README._md")
			outputFile := fmt.Sprintf("../../../../%s/README.md", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}
	}
}

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
			outputFile := fmt.Sprintf("../../../../%s/server/app/main.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/model/model_basic._go")
			outputFile := fmt.Sprintf("../../../../%s/server/model/system-model.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/model/user-model._go")
			outputFile := fmt.Sprintf("../../../../%s/server/model/system-user.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/repository/user-repository._go")
			outputFile := fmt.Sprintf("../../../../%s/server/repository/system-user.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/service/user-service._go")
			outputFile := fmt.Sprintf("../../../../%s/server/service/system-user.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/controller/user-controller._go")
			outputFile := fmt.Sprintf("../../../../%s/server/controller/system-user.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/controller/router._go")
			outputFile := fmt.Sprintf("../../../../%s/server/controller/system-router.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/utils/identifier._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/utils/identifier.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/utils/password._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/utils/password.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/utils/json._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/utils/json.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/utils/strings._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/utils/strings.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/token/jwt._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/token/jwt.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/config/config._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/config/config.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/log/log._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/log/log.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/transaction/transaction._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/transaction/transaction.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/messagebroker/producer._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/messagebroker/producer.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/messagebroker/consumer._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/messagebroker/consumer.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/error/error_code._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/error/error_code.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/error/error._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/error/error.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/error/error_test._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/error/error_test.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/constant/constant._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/constant/constant.go", tp.PackagePath)
			basic(tp, templateFile, outputFile, tp, 0664)
		}

		{
			templateFile := fmt.Sprintf("../templates/backend/shared/converter/converter._go")
			outputFile := fmt.Sprintf("../../../../%s/server/shared/converter/converter.go", tp.PackagePath)
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
}

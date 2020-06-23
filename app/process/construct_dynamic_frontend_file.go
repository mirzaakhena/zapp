package process

import (
	"fmt"
	"os"
	"strings"

	"github.com/mirzaakhena/zapp/app/model"
)

func createDynamicFrontendFile(tp *model.ThePackage, et *model.TheEntity) {

	// create store modules
	// {
	// 	templateFile := fmt.Sprintf("../templates/frontend/src/store/modules/file._js")
	// 	outputFile := fmt.Sprintf("../../../../%s/webapp/src/store/modules/%s.js", tp.PackagePath, strings.ToLower(et.Name))
	// 	basic(tp, templateFile, outputFile, et, 0664)
	// }

	// create router modules
	{
		templateFile := fmt.Sprintf("../templates/frontend/src/router/modules/file._js")
		outputFile := fmt.Sprintf("../../../../%s/client/src/router/modules/%s.js", tp.PackagePath, strings.ToLower(et.Name))
		basic(tp, templateFile, outputFile, et, 0664)
	}

	{
		// create folder pages
		f := fmt.Sprintf("../../../../%s/client/src/pages/%s", tp.PackagePath, strings.ToLower(et.Name))
		os.Mkdir(f, 0777)

		// create file table under folder
		{
			templateFile := fmt.Sprintf("../templates/frontend/src/pages/folder/list._vue")
			outputFile := fmt.Sprintf("../../../../%s/client/src/pages/%s/list.vue", tp.PackagePath, strings.ToLower(et.Name))
			basic(tp, templateFile, outputFile, et, 0664)
		}

		// create file input under folder
		{
			templateFile := fmt.Sprintf("../templates/frontend/src/pages/folder/input._vue")
			outputFile := fmt.Sprintf("../../../../%s/client/src/pages/%s/input.vue", tp.PackagePath, strings.ToLower(et.Name))
			basic(tp, templateFile, outputFile, et, 0664)
		}

	}

}

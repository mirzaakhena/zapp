package process

import (
	"fmt"

	"github.com/mirzaakhena/zapp/app/model"
)

func createStaticFrontendFile(tp *model.ThePackage) {
	// frontend
	{
		templateFile := fmt.Sprintf("../templates/frontend/public/favicon._ico")
		outputFile := fmt.Sprintf("../../../../%s/frontend/public/favicon.ico", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/public/index._html")
		outputFile := fmt.Sprintf("../../../../%s/frontend/public/index.html", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/babel.config._js")
		outputFile := fmt.Sprintf("../../../../%s/frontend/babel.config.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/vue.config._js")
		outputFile := fmt.Sprintf("../../../../%s/frontend/vue.config.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/package._json")
		outputFile := fmt.Sprintf("../../../../%s/frontend/package.json", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/App._vue")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/App.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/assets/style._css")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/assets/style.css", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/dist/index._html")
		outputFile := fmt.Sprintf("../../../../%s/frontend/dist/index.html", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/forgotpassword._vue")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/pages/forgotpassword.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/home._vue")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/pages/home.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/login._vue")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/pages/login.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/notfound._vue")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/pages/notfound.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/register._vue")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/pages/register.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/pages/successregister._vue")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/pages/successregister.vue", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/main._js")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/main.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/store/index._js")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/store/index.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/store/crudtable._js")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/store/crudtable.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/router/index._js")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/router/index.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/utils/httprequest._js")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/utils/httprequest.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/utils/auth._js")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/utils/auth.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}

	{
		templateFile := fmt.Sprintf("../templates/frontend/src/utils/filter._js")
		outputFile := fmt.Sprintf("../../../../%s/frontend/src/utils/filter.js", tp.PackagePath)
		basic(tp, templateFile, outputFile, tp, 0664)
	}
}

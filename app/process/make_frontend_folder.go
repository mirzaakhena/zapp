package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/zapp/app/model"
)

func makeFrontendDirectory(tp *model.ThePackage) {

	var dir string

	dir = fmt.Sprintf("../../../../%s/client/public", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/client/src/assets", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/client/src/store/modules", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/client/src/router/modules", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/client/src/pages", tp.PackagePath)
	os.Mkdir(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/client/dist", tp.PackagePath)
	os.Mkdir(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/client/src/utils", tp.PackagePath)
	os.Mkdir(dir, 0777)

}

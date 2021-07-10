package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/zapp/app/model"
)

func makeFrontendDirectory(tp *model.ThePackage) {

	var dir string

	dir = fmt.Sprintf("../../../%s/frontend/public", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../%s/frontend/src/assets", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../%s/frontend/src/store/modules", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../%s/frontend/src/router/modules", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../%s/frontend/src/pages", tp.PackagePath)
	os.Mkdir(dir, 0777)

	dir = fmt.Sprintf("../../../%s/frontend/dist", tp.PackagePath)
	os.Mkdir(dir, 0777)

	dir = fmt.Sprintf("../../../%s/frontend/src/utils", tp.PackagePath)
	os.Mkdir(dir, 0777)

}

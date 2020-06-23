package process

import (
	"fmt"
	"os"

	"github.com/mirzaakhena/zapp/app/model"
)

func makeBackendDirectory(tp *model.ThePackage) {

	var dir string

	dir = fmt.Sprintf("../../../../%s/server/app", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/controller", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/model", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/model/enum", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/service", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/repository", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/shared/utils", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/shared/token", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/shared/config", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/shared/log", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/shared/transaction", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/shared/messagebroker", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/shared/constant", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/shared/converter", tp.PackagePath)
	os.MkdirAll(dir, 0777)

	dir = fmt.Sprintf("../../../../%s/server/shared/error", tp.PackagePath)
	os.MkdirAll(dir, 0777)

}

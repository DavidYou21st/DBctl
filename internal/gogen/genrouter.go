package gogen

import (
	_ "embed"
	"fctl/pkg/helper/stringx"
	"fctl/pkg/pathx"
	"fctl/pkg/tool"
	"os"
	"path"
)

//go:embed template/router.tpl
var routerTemplate string

func GenRouter(dir string, fileName, appName string) error {
	dtoFilename := fileName + ".go"
	filename := path.Join(dir, routerDir(appName), dtoFilename)
	os.Remove(filename)

	text, err := pathx.LoadTemplate(category, routerTemplateFile, routerTemplate)
	if err != nil {
		return err
	}
	err = tool.With("router").GoFmt(true).Parse(text).SaveTo(map[string]interface{}{
		"appName":      stringx.From(appName).Lower(),
		"Name":         stringx.From(fileName).ToCamel(),
		"appNameLower": fileName,
	}, filename, false)
	return err
}

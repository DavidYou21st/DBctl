package gogen

import (
	_ "embed"
	"fctl/pkg/helper/stringx"
	"fctl/pkg/pathx"
	"fctl/pkg/tool"
	"fmt"
	"os"
	"path"
)

//go:embed template/controller.tpl
var controllerTemplate string

func GenController(dir, fileName, appName string) error {
	dtoFilename := fileName + ".go"
	fileRelName := path.Join(dir, controllerDir(appName), dtoFilename)
	errs := os.Remove(fileRelName)
	if errs != nil {
		fmt.Println("no file remove")
	}

	text, err := pathx.LoadTemplate("model", controllerTemplateFile, controllerTemplate)
	if err != nil {
		return err
	}
	err = tool.With("controller").GoFmt(true).Funcs(funcsMap).Parse(text).SaveTo(map[string]interface{}{
		"appName":        stringx.From(appName).ToCamel(),
		"controllerName": stringx.From(fileName).ToCamel(),
	}, fileRelName, false)
	return err
}

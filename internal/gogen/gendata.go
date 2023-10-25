package gogen

import (
	_ "embed"
	"fctl/internal/types"
	"fctl/pkg/helper/stringx"
	"fctl/pkg/pathx"
	"fctl/pkg/tool"
	"fmt"
	"os"
	"path"
)

//go:embed template/data.tpl
var dataTemplate string

func GenData(dir string, fileName, appName, tableName string) error {
	dtoFilename := fileName + ".go"
	filename := path.Join(dir, dataDir(appName), dtoFilename)
	errs := os.Remove(filename)
	if errs != nil {
		fmt.Println("no file remove")
	}

	text, err := pathx.LoadTemplate(category, dataTemplateFile, dataTemplate)
	if err != nil {
		return err
	}
	fieldNames, err := types.GenTypeNameByMysql(tableName)
	if err != nil {
		return err
	}

	err = tool.With("data").GoFmt(true).Funcs(funcsMap).Parse(text).SaveTo(map[string]interface{}{
		"fieldNames": fieldNames,
		"appName":    stringx.From(appName).ToCamel(),
		"dataName":   stringx.From(fileName).ToCamel(),
	}, filename, false)
	return err
}

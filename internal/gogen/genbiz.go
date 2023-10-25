package gogen

import (
	_ "embed"
	"fctl/internal/types"
	"fctl/pkg/global"
	"fctl/pkg/helper/config"
	"fctl/pkg/helper/format"
	"fctl/pkg/helper/stringx"
	"fctl/pkg/pathx"
	"fctl/pkg/tool"
	"fmt"
	"os"
	"path"
)

//go:embed template/biz.tpl
var bizTemplate string

func GenBiz(dir, fileName, appName, tableName string) error {
	cfg, err := config.NewConfig(global.CFG.File.Style)
	if err != nil {
		return err
	}
	filename, err := format.FileNamingFormat(cfg.NamingFormat, fileName)
	if err != nil {
		return err
	}

	dtoFilename := filename + ".go"
	fileRelName := path.Join(dir, bizDir(appName), dtoFilename)
	errs := os.Remove(fileRelName)
	if errs != nil {
		fmt.Println("no file remove")
	}

	text, err := pathx.LoadTemplate(category, bizTemplateFile, bizTemplate)
	if err != nil {
		return err
	}
	code, _ := types.GenTypeByMysql(tableName)
	err = tool.With("biz").GoFmt(true).Funcs(funcsMap).Parse(text).SaveTo(map[string]interface{}{
		"typeCode": code,
		"appName":  stringx.From(appName).ToCamel(),
		"bizName":  stringx.From(fileName).ToCamel(),
	}, fileRelName, false)
	return err
}

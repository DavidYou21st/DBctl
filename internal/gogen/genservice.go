package gogen

import (
	_ "embed"
	"fctl/internal/types"
	"fctl/pkg/global"
	"fctl/pkg/helper/config"
	"fctl/pkg/helper/format"
	"fctl/pkg/helper/stringx"
	"fctl/pkg/parser"
	"fctl/pkg/pathx"
	"fctl/pkg/spec"
	"fctl/pkg/tool"
	"fmt"
	"github.com/samber/lo"
	"os"
	"path"
	"strings"
)

//go:embed template/service.tpl
var serviceTemplate string

func GenService(apiPath, dir, fileName, appName, tableName string) error {
	serviceAbbreviation := string([]rune(tableName)[0])
	serviceStr := serviceAbbreviation + "c" //组合缩写名称
	api, err := parser.Parse(apiPath)
	if err != nil {
		return err
	}

	cfg, err := config.NewConfig(global.CFG.File.Style)
	if err != nil {
		return err
	}
	var getReply, ListReply []string

	getReplyName := "Get" + stringx.From(fileName).ToCamel() + "Reply"
	ListReplyName := "List" + stringx.From(fileName).ToCamel() + "Reply"

	for _, tp := range api.Types {
		if ok := lo.Contains[string]([]string{getReplyName, ListReplyName}, tp.Name()); !ok { // 过滤
			continue
		}
		structType, ok := tp.(spec.DefineStruct)
		if !ok {
			return fmt.Errorf("unspport struct type: %s", tp.Name())
		}
		for _, member := range structType.Members {
			if member.IsInline {
				continue
			}
			member.Name = strings.Replace(member.Name, "Id", "ID", -1)
			if tp.Name() == getReplyName {
				getReply = append(getReply, member.Name)
			}

			if tp.Name() == ListReplyName {
				ListReply = append(ListReply, member.Name)
			}
		}
	}
	filename, err := format.FileNamingFormat(cfg.NamingFormat, fileName)
	if err != nil {
		return err
	}
	dtoFilename := filename + ".go" // 文件名
	fileRelName := path.Join(dir, serviceDir(appName), dtoFilename)
	errs := os.Remove(fileRelName)
	if errs != nil {
		fmt.Println("no file remove")
	}

	text, err := pathx.LoadTemplate(category, serviceTemplateFile, serviceTemplate)
	if err != nil {
		return err
	}
	fieldNames, err := types.GenTypeNameByMysql(tableName)
	if err != nil {
		return err
	}
	err = tool.With("service").GoFmt(false).Funcs(funcsMap).Parse(text).SaveTo(map[string]interface{}{
		"serviceStr":  serviceStr,
		"getReply":    getReply,
		"listReply":   ListReply,
		"fieldNames":  fieldNames,
		"appName":     stringx.From(appName).ToCamel(),
		"serviceName": stringx.From(fileName).ToCamel(),
	}, fileRelName, false)
	return err
}

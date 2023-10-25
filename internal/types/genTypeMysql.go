package types

import (
	"fctl/internal/model/sql/model"
	"fctl/internal/model/sql/parser"
	"fctl/pkg/global"
	"fctl/pkg/pathx"
	"fctl/pkg/tool"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

type Table struct {
	parser.Table
}

func GenTypeByMysql(tableName string) (string, error) {
	dsnUrl := global.CFG.MySQL.DSNURL()
	dsn, err := mysql.ParseDSN(dsnUrl)
	if err != nil {
		return "", err
	}
	logx.Disable()
	databaseSource := strings.TrimSuffix(dsnUrl, "/"+dsn.DBName) + "/information_schema"
	db := sqlx.NewMysql(databaseSource)
	im := model.NewInformationSchemaModel(db)

	tables, err := im.GetAllTables(dsn.DBName)
	if err != nil {
		return "", err
	}
	matchTables := make(map[string]*model.Table)
	for _, item := range tables {
		if item != tableName {
			continue
		}

		columnData, err := im.FindColumns(dsn.DBName, item)
		if err != nil {
			return "", err
		}
		table, err := columnData.Convert()
		if err != nil {
			return "", err
		}
		matchTables[item] = table
	}
	if len(matchTables) == 0 {
		return "", err
	}

	for _, each := range matchTables {
		table, err := parser.ConvertDataType(each)
		if err != nil {
			return "", err
		}

		code, err := genModel(*table)
		if err != nil {
			return "", err
		}

		return code, nil
	}
	return "", fmt.Errorf("error gen")
}

func GenTypeNameByMysql(tableName string) ([]string, error) {
	dsnUrl := global.CFG.MySQL.DSNURL()
	dsn, err := mysql.ParseDSN(dsnUrl)
	if err != nil {
		return nil, err
	}
	logx.Disable()
	databaseSource := strings.TrimSuffix(dsnUrl, "/"+dsn.DBName) + "/information_schema"
	db := sqlx.NewMysql(databaseSource)
	im := model.NewInformationSchemaModel(db)

	tables, err := im.GetAllTables(dsn.DBName)
	if err != nil {
		return nil, err
	}
	matchTables := make(map[string]*model.Table)
	for _, item := range tables {
		if item != tableName {
			continue
		}

		columnData, err := im.FindColumns(dsn.DBName, item)
		if err != nil {
			return nil, err
		}
		table, err := columnData.Convert()
		if err != nil {
			return nil, err
		}
		matchTables[item] = table
	}
	if len(matchTables) == 0 {
		return nil, err
	}

	for _, each := range matchTables {
		table, err := parser.ConvertDataType(each)
		if err != nil {
			return nil, err
		}
		fields := table.Fields
		names, err := genFieldNames(fields)
		if err != nil {
			return nil, err
		}

		return names, nil
	}
	return nil, fmt.Errorf("error gen")
}

func genModel(in parser.Table) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}

	var table Table
	table.Table = in

	typesCode, err := genTypes(table)
	if err != nil {
		return "", err
	}
	return typesCode, nil
}

func genTypes(table Table) (string, error) {
	fields := table.Fields
	fieldsString, err := genFields(table, fields)
	if err != nil {
		return "", err
	}

	text, err := pathx.LoadTemplate("model", "types.tpl", Types)
	if err != nil {
		return "", err
	}

	output, err := tool.With("types").
		Parse(text).
		GoFmt(true).
		Execute(map[string]interface{}{
			"upperStartCamelObject": table.Name.ToCamel(),
			"fields":                fieldsString,
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

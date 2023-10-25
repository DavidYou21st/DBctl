package types

import (
	"fctl/internal/model/sql/parser"
	"fctl/pkg/tool"
	"strings"

	"fctl/pkg/pathx"
)

func genFieldNames(fields []*parser.Field) ([]string, error) {
	var list []string
	for _, field := range fields {
		fieldName := tool.SafeString(field.Name.ToPascal())
		list = append(list, fieldName)
	}
	return list, nil
}

func genFields(table Table, fields []*parser.Field) (string, error) {
	var list []string

	for _, field := range fields {
		result, err := genField(table, field)
		if err != nil {
			return "", err
		}

		list = append(list, result)
	}

	return strings.Join(list, "\n"), nil
}

func genField(table Table, field *parser.Field) (string, error) {
	text, err := pathx.LoadTemplate("model", "field.tpl", Field)
	if err != nil {
		return "", err
	}
	fieldName := tool.SafeString(field.Name.ToPascal())
	output, err := tool.With("types").
		Parse(text).
		Execute(map[string]interface{}{
			"name": fieldName,
			"type": field.DataType,
			//"comment": field.Comment,
			"data": table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

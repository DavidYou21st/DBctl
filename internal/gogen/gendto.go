package gogen

import (
	_ "embed"
	"fctl/pkg/parser"
	"fctl/pkg/tool"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"fctl/pkg/spec"
)

//go:embed template/dto.tpl
var dtoTemplate string

// BuildTypes gen types to string
func BuildTypes(types []spec.Type) (string, error) {
	var builder strings.Builder
	first := true
	for _, tp := range types {
		if first {
			first = false
		} else {
			builder.WriteString("\n\n")
		}
		if err := writeType(&builder, tp); err != nil {
			return "", tool.WrapErr(err, "Type "+tp.Name()+" gorm error")
		}
	}

	return builder.String(), nil
}

func GenDto(apiFile, dir, name, appName string) error {
	api, err := parser.Parse(apiFile)
	if err != nil {
		return err
	}
	val, err := BuildTypes(api.Types)
	if err != nil {
		return err
	}

	dtoFilename := name + ".go"
	filename := path.Join(dir, dtoDir(appName), dtoFilename)
	errs := os.Remove(filename)
	if errs != nil {
		fmt.Println("no file remove")
	}

	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          dtoDir(appName),
		filename:        dtoFilename,
		templateName:    "typesTemplate",
		category:        category,
		templateFile:    dtoTemplateFile,
		builtinTemplate: dtoTemplate,
		data: map[string]interface{}{
			"types":        val,
			"containsTime": false,
		},
	})
}

func writeType(writer io.Writer, tp spec.Type) error {
	structType, ok := tp.(spec.DefineStruct)
	if !ok {
		return fmt.Errorf("unspport struct type: %s", tp.Name())
	}

	fmt.Fprintf(writer, "type %s struct {\n", tool.Title(tp.Name()))
	for _, member := range structType.Members {
		if member.IsInline {
			if _, err := fmt.Fprintf(writer, "%s\n", strings.Title(member.Type.Name())); err != nil {
				return err
			}

			continue
		}

		if err := writeProperty(writer, member.Name, member.Tag, member.GetComment(), member.Type, 1); err != nil {
			return err
		}
	}
	fmt.Fprintf(writer, "}")
	return nil
}

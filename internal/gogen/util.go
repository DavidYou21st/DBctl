package gogen

import (
	"bytes"
	"fctl/pkg/tool"
	"fmt"
	"io"
	"strings"
	"text/template"

	"fctl/pkg/golang"
	"fctl/pkg/pathx"
	"fctl/pkg/spec"
)

type fileGenConfig struct {
	dir             string
	subdir          string
	filename        string
	templateName    string
	category        string
	templateFile    string
	builtinTemplate string
	data            interface{}
}

func genFile(c fileGenConfig) error {
	fp, created, err := tool.MaybeCreateFile(c.dir, c.subdir, c.filename)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	var text string
	if len(c.category) == 0 || len(c.templateFile) == 0 {
		text = c.builtinTemplate
	} else {
		text, err = pathx.LoadTemplate(c.category, c.templateFile, c.builtinTemplate)
		if err != nil {
			return err
		}
	}

	t := template.Must(template.New(c.templateName).Parse(text))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, c.data)
	if err != nil {
		return err
	}

	code := golang.FormatCode(buffer.String())
	_, err = fp.WriteString(code)
	return err
}

func writeProperty(writer io.Writer, name, tag, comment string, tp spec.Type, indent int) error {
	tool.WriteIndent(writer, indent)
	var err error
	if len(comment) > 0 {
		comment = strings.TrimPrefix(comment, "//")
		comment = "//" + comment
		_, err = fmt.Fprintf(writer, "%s %s %s %s\n", strings.Title(name), tp.Name(), tag, comment)
	} else {
		_, err = fmt.Fprintf(writer, "%s %s %s\n", strings.Title(name), tp.Name(), tag)
	}

	return err
}

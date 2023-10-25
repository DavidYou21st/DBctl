package format

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	flagZEUS  = "ZEUS"
	flagCLOUD = "CLOUD"

	unknown style = iota
	title
	lower
	upper
)

// ErrNamingFormat defines an error for unknown format
var ErrNamingFormat = errors.New("unsupported format")

type (
	styleFormat struct {
		before     string
		through    string
		after      string
		zeusStyle  style
		cloudStyle style
	}

	style int
)

// FileNamingFormat is used to format the file name. You can define the format style
// through the zeus and cloud formatting characters. For example, you can define the snake
// format as zeus_cloud, and the camel case format as zeusCloud. You can even specify the split
// character, such as zeus#Cloud, theoretically any combination can be used, but the prerequisite
// must meet the naming conventions of each operating system file name.
// Note: Formatting is based on snake or camel string
func FileNamingFormat(format, content string) (string, error) {
	upperFormat := strings.ToUpper(format)
	indexZeus := strings.Index(upperFormat, flagZEUS)
	indexCloud := strings.Index(upperFormat, flagCLOUD)

	if indexZeus < 0 || indexCloud < 0 || indexZeus > indexCloud {
		return "", ErrNamingFormat
	}
	var (
		before, through, after string
		flagZeus, flagCloud    string
		zeusStyle, cloudStyle  style
		err                    error
	)
	before = format[:indexZeus]
	flagZeus = format[indexZeus : indexZeus+4]
	through = format[indexZeus+4 : indexCloud]
	flagCloud = format[indexCloud : indexCloud+5]
	after = format[indexCloud+5:]

	zeusStyle, err = getStyle(flagZeus)
	if err != nil {
		return "", err
	}

	cloudStyle, err = getStyle(flagCloud)
	if err != nil {
		return "", err
	}
	var formatStyle styleFormat
	formatStyle.zeusStyle = zeusStyle
	formatStyle.cloudStyle = cloudStyle
	formatStyle.before = before
	formatStyle.through = through
	formatStyle.after = after
	return doFormat(formatStyle, content)
}

func doFormat(f styleFormat, content string) (string, error) {
	splits, err := split(content)
	if err != nil {
		return "", err
	}
	var join []string
	for index, split := range splits {
		if index == 0 {
			join = append(join, transferTo(split, f.zeusStyle))
			continue
		}
		join = append(join, transferTo(split, f.cloudStyle))
	}
	joined := strings.Join(join, f.through)
	return f.before + joined + f.after, nil
}

func transferTo(in string, style style) string {
	switch style {
	case upper:
		return strings.ToUpper(in)
	case lower:
		return strings.ToLower(in)
	case title:
		return strings.Title(in)
	default:
		return in
	}
}

func split(content string) ([]string, error) {
	var (
		list   []string
		reader = strings.NewReader(content)
		buffer = bytes.NewBuffer(nil)
	)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				if buffer.Len() > 0 {
					list = append(list, buffer.String())
				}
				return list, nil
			}
			return nil, err
		}
		if r == '_' {
			if buffer.Len() > 0 {
				list = append(list, buffer.String())
			}
			buffer.Reset()
			continue
		}

		if r >= 'A' && r <= 'Z' {
			if buffer.Len() > 0 {
				list = append(list, buffer.String())
			}
			buffer.Reset()
		}
		buffer.WriteRune(r)
	}
}

func getStyle(flag string) (style, error) {
	compare := strings.ToLower(flag)
	switch flag {
	case strings.ToLower(compare):
		return lower, nil
	case strings.ToUpper(compare):
		return upper, nil
	case strings.Title(compare):
		return title, nil
	default:
		return unknown, fmt.Errorf("unexpected format: %s", flag)
	}
}

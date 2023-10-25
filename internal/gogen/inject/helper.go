package inject

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const (
	delimiter = "\n"
)

type CheckExist func(line string)

// 执行go代码格式化
func execGoFmt(name string) error {
	cmd := exec.Command("gofmt", "-w", name, name)
	return cmd.Run()
}

// 读取文件内容
func readFile(name string) (*bytes.Buffer, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, file)
	return buf, nil
}

// 写入文件
func writeFile(name string, buf *bytes.Buffer) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, _ = io.Copy(file, buf)
	return nil
}

// insertContent 插入文件内容
func insertContent(name, injectContent, startContainStr, endContainStr string) error {
	injectStart := 0
	insertFn := func(line string) (data string, flag int, ok bool) { // fn 回调当前行数据，返回-1为插入当前行之前，1为插入当前行之后
		if injectStart == 0 && strings.Contains(line, startContainStr) {
			injectStart = 1
			return
		}

		if injectStart == 1 && strings.Contains(line, endContainStr) {
			injectStart = -1
			data = injectContent
			flag = -1
			ok = true
			return
		}
		return "", 0, false
	}

	buf, err := readFile(name)
	if err != nil {
		return err
	}

	nbuf := new(bytes.Buffer)
	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		cline := scanner.Text()
		data, flag, ok := insertFn(cline)
		if !ok {
			nbuf.WriteString(cline)
			nbuf.WriteString(delimiter)
			continue
		}
		if flag == -1 {
			nbuf.WriteString(data)
			nbuf.WriteString(delimiter)
			nbuf.WriteString(cline)
			nbuf.WriteString(delimiter)
			continue
		}
		nbuf.WriteString(cline)
		nbuf.WriteString(delimiter)
		nbuf.WriteString(data)
		nbuf.WriteString(delimiter)
	}

	return writeFile(name, nbuf)
}

// 检查是否已经有插入
func checkContentExists(name, injectContent string) (bool, error) {
	checkFn := func(line string) (exist bool) {
		if strings.Contains(line, injectContent) {
			return true
		}
		return false
	}

	buf, err := readFile(name)
	if err != nil {
		return false, err
	}
	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		cline := scanner.Text()
		exist := checkFn(cline)
		if exist {
			return true, fmt.Errorf("已经存在不必重复插入")
		}
	}
	return false, nil
}

func getFullFileName(dir, appName, module string) string {
	return fmt.Sprintf("%s/internal/%s/%s/%s.go", dir, appName, module, module)
}

func getServiceAbbreviation(name string) string {
	serviceAbbreviation := string([]rune(name)[0])
	serviceStr := serviceAbbreviation + "c" //组合缩写名称
	return strings.ToLower(serviceStr)
}

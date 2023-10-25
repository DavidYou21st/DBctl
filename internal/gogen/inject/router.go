package inject

import (
	"fmt"
)

// InsertRouterInject 注入文件
func InsertRouterInject(dir, name, appName string) error {
	fullName := getFullFileName(dir, appName, "router")

	injectContent := fmt.Sprintf("%sApi *controller.%s", name, name) //插入的内容

	exist, errs := checkContentExists(fullName, injectContent) // 检查是否已经插入 true 标识已经存在
	if errs != nil || exist {
		return errs
	}
	err := insertContent(fullName, injectContent, "type Router struct {", "}")
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullName)
	fmt.Printf("完成router 模块生成")
	return execGoFmt(fullName)
}

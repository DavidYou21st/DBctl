package inject

import (
	"fmt"
)

// InsertDataInject  插入data注入文件
func InsertDataInject(dir, name, appName string) error {
	fullName := getFullFileName(dir, appName, "data")

	injectContent := fmt.Sprintf("New%sRepo,", name) //插入的内容

	exist, errs := checkContentExists(fullName, injectContent) // 检查是否已经插入 true 标识已经存在
	if errs != nil || exist {
		return errs
	}
	err := insertContent(fullName, injectContent, "var ProviderSet = wire.NewSet(", ")")
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullName)
	fmt.Printf("完成data 模块生成")
	return execGoFmt(fullName)
}

package inject

import (
	"fmt"
)

// InsertControllerInjectEnt 注入文件
func InsertControllerInjectEnt(dir, name, appName string) error {
	fullName := getFullFileName(dir, appName, "controller")

	injectContent := fmt.Sprintf("%sSet,", name) //插入的内容

	exist, errs := checkContentExists(fullName, injectContent) // 检查是否已经插入 true 标识已经存在
	if errs != nil || exist {
		return errs
	}
	err := insertContent(fullName, injectContent, "var ControllerSet = wire.NewSet(", ")")
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullName)
	fmt.Println("完成biz 模块生成")
	return execGoFmt(fullName)
}

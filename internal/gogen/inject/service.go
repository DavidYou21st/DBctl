package inject

import (
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
)

// InsertServiceInject  注入文件
func InsertServiceInject(dir, name, appName string) error {
	serviceStr := getServiceAbbreviation(name)

	fullName := getFullFileName(dir, appName, "service")

	injectContent1 := fmt.Sprintf("%s  *biz.%sUseCase", serviceStr, name) //插入的内容

	exist, err1 := checkContentExists(fullName, injectContent1) // 检查是否已经插入 true 标识已经存在
	if err1 != nil || exist {
		return err1
	}
	fmt.Println(injectContent1)
	startContainStr1 := fmt.Sprintf("type %sService struct {", strutil.UpperFirst(appName))
	err2 := insertContent(fullName, injectContent1, startContainStr1, "}")
	if err2 != nil {
		return err2
	}
	fmt.Println(startContainStr1)
	injectContent2 := fmt.Sprintf("%s  *biz.%sUseCase,", serviceStr, name) //插入的内容
	startContainStr2 := fmt.Sprintf("func New%sService(", strutil.UpperFirst(appName))
	err3 := insertContent(fullName, injectContent2, startContainStr2, ")")
	if err1 != nil {
		return err3
	}

	injectContent3 := fmt.Sprintf("%s:  %s,", serviceStr, serviceStr) //插入的内容
	startContainStr3 := fmt.Sprintf("return &%sService{", strutil.UpperFirst(appName))
	err4 := insertContent(fullName, injectContent3, startContainStr3, "\t}")
	if err4 != nil {
		return err4
	}

	fmt.Printf("文件[%s]写入成功\n", fullName)
	fmt.Printf("完成service 模块生成")
	return execGoFmt(fullName)
}

package inject

import (
	"fctl/internal/gogen/inject"
	"fctl/pkg/cobrax"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	fileName  string // 文件名
	appName   string // app 微服务名称
	dir       string // 生成代码目录
	tableName string // 表名
	config    string // 配置文件

	InjectCmd = cobrax.NewCommand("inject", cobrax.WithRunE(runAll))
)

func init() {
	InjectCmdFlags := InjectCmd.Flags()

	InjectCmdFlags.StringVarPWithDefaultValue(&dir, "dir", "d", "./")
	InjectCmdFlags.StringVarPWithDefaultValue(&fileName, "file", "f", "")
	InjectCmdFlags.StringVarPWithDefaultValue(&appName, "appName", "n", "console")
}

func runAll(_ *cobra.Command, _ []string) error {
	logx.Must(inject.InsertBizInject(dir, fileName, appName))
	logx.Must(inject.InsertRouterInject(dir, fileName, appName))
	logx.Must(inject.InsertDataInject(dir, fileName, appName))
	logx.Must(inject.InsertServiceInject(dir, fileName, appName))
	logx.Must(inject.InsertRouterInject(dir, fileName, appName))
	return nil
}

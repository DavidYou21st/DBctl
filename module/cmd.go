package module

import (
	"fctl/pkg/cobrax"
)

var (
	fileName     string // 文件名
	appName      string // app 微服务名称
	dir          string // 生成代码目录
	tableName    string // 表名
	config       string // 配置文件
	name         string
	table        string
	OutPath      string
	OutFile      string
	modelPkgPath string
	apiPath      string
	style        string

	ControllerCmd = cobrax.NewCommand("controller", cobrax.WithRunE(runController))
	BizCmd        = cobrax.NewCommand("gorm", cobrax.WithRunE(runBiz))
	ServiceCmd    = cobrax.NewCommand("service", cobrax.WithRunE(runService))
	RouterCmd     = cobrax.NewCommand("router", cobrax.WithRunE(runRouter))
	DtoCmd        = cobrax.NewCommand("dto", cobrax.WithRunE(runGenDto))
	DataCmd       = cobrax.NewCommand("dto", cobrax.WithRunE(runData))
)

func init() {

	ControllerFlags := ControllerCmd.Flags()
	BizCmdFlags := BizCmd.Flags()
	ServiceCmdFlags := ServiceCmd.Flags()
	RouterCmdFlags := RouterCmd.Flags()
	DtoCmdFlags := DtoCmd.Flags()
	DataCmdFlags := DataCmd.Flags()

	ControllerFlags.StringVarPWithDefaultValue(&dir, "dir", "d", "./")
	ControllerFlags.StringVarPWithDefaultValue(&fileName, "file", "f", "")
	ControllerFlags.StringVarPWithDefaultValue(&appName, "appName", "n", "")

	BizCmdFlags.StringVarPWithDefaultValue(&dir, "dir", "d", "./")
	BizCmdFlags.StringVarPWithDefaultValue(&fileName, "file", "f", "")
	BizCmdFlags.StringVarPWithDefaultValue(&appName, "appName", "n", "")
	BizCmdFlags.StringVarPWithDefaultValue(&config, "config", "c", "config/config.yaml")
	BizCmdFlags.StringVarPWithDefaultValue(&tableName, "table", "t", "")

	ServiceCmdFlags.StringVarPWithDefaultValue(&apiPath, "api", "a", "")
	ServiceCmdFlags.StringVarPWithDefaultValue(&dir, "dir", "d", "./")
	ServiceCmdFlags.StringVarPWithDefaultValue(&fileName, "file", "f", "")
	ServiceCmdFlags.StringVarPWithDefaultValue(&appName, "appName", "n", "")
	ServiceCmdFlags.StringVarPWithDefaultValue(&config, "config", "c", "config/config.yaml")
	ServiceCmdFlags.StringVarPWithDefaultValue(&tableName, "tableName", "t", "")

	RouterCmdFlags.StringVarPWithDefaultValue(&dir, "dir", "d", "")
	RouterCmdFlags.StringVarPWithDefaultValue(&fileName, "file", "f", "")
	RouterCmdFlags.StringVarPWithDefaultValue(&appName, "appName", "n", "")

	DtoCmdFlags.StringVarPWithDefaultValue(&apiPath, "api", "a", "")
	DtoCmdFlags.StringVarPWithDefaultValue(&dir, "dir", "d", "")
	DtoCmdFlags.StringVarPWithDefaultValue(&fileName, "file", "f", "")
	DtoCmdFlags.StringVarPWithDefaultValue(&appName, "appName", "n", "")

	DataCmdFlags.StringVarPWithDefaultValue(&dir, "dir", "d", "./")
	DataCmdFlags.StringVarPWithDefaultValue(&fileName, "fileName", "f", "")
	DataCmdFlags.StringVarPWithDefaultValue(&appName, "appName", "n", "")
	DataCmdFlags.StringVarPWithDefaultValue(&config, "config", "c", "config/config.yaml")
	DataCmdFlags.StringVarPWithDefaultValue(&tableName, "tableName", "t", "")

}

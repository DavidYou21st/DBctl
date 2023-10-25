package gorm

import (
	"fctl/pkg/cobrax"
	"fctl/pkg/tool"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	table        string
	config       string
	OutPath      string
	OutFile      string
	modelPkgPath string
	Cmd          = cobrax.NewCommand("gorm", cobrax.WithRunE(run))
)

func init() {
	gormCmdFlags := Cmd.Flags()
	gormCmdFlags.StringVarWithDefaultValue(&table, "table", "")
	gormCmdFlags.StringVarWithDefaultValue(&config, "config", "config/config.yaml")
	gormCmdFlags.StringVar(&OutPath, "path")
	gormCmdFlags.StringVarWithDefaultValue(&OutFile, "file", "gen.go")
	gormCmdFlags.StringVarWithDefaultValue(&modelPkgPath, "mpath", "")
}

func run(_ *cobra.Command, _ []string) error {
	if ok, _ := tool.PathExists(config); !ok {
		fmt.Println("配置不存在")
		return nil
	}
	return Run(
		SetTables(table),
		SetConfigFile(config),
		SetModelPkgPath(modelPkgPath),
		SetOutFile(OutFile),
		SetOutPath(OutPath))
}

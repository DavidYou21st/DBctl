package module

import (
	"fctl/internal/gogen"
	"fctl/pkg/global"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func runData(_ *cobra.Command, _ []string) error {
	if "" == tableName {
		return fmt.Errorf("表名不能为空")
	}
	if "" == appName {
		return fmt.Errorf("模块名不能为空")
	}

	global.VIPER = global.CFG.Viper.Viper(&global.CFG, global.ConfigPaths, config)
	err := gogen.GenData(dir, fileName, appName, tableName)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("done!")
	}
	return nil
}

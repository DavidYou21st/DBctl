package module

import (
	"fctl/internal/gogen"
	"fctl/pkg/global"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func runBiz(_ *cobra.Command, _ []string) error {
	global.VIPER = global.CFG.Viper.Viper(&global.CFG, global.ConfigPaths, config)
	err := gogen.GenBiz(dir, fileName, appName, tableName)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("done!")
	}
	return nil
}

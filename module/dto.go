package module

import (
	"fctl/internal/gogen"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func runGenDto(_ *cobra.Command, _ []string) error {
	if fileName == "" {
		fmt.Println("文件名不能为空")
		return nil
	}
	err := gogen.GenDto(apiPath, dir, fileName, appName)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("done!")
	}
	return nil
}

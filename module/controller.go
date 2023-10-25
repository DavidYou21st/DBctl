package module

import (
	"fctl/internal/gogen"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func runController(_ *cobra.Command, _ []string) error {
	err := gogen.GenController(dir, fileName, appName)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("done!")
	}
	return nil
}

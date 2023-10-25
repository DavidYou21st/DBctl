package module

import (
	"fctl/internal/gogen"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func runRouter(_ *cobra.Command, _ []string) error {
	err := gogen.GenRouter(dir, fileName, appName)
	if err != nil {
		fmt.Println(err)
	} else {
		log.Printf("done!")
	}
	return nil
}

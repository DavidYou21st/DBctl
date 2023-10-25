package gogen

import "fmt"

const (
	internal = "internal/"
)

func dtoDir(appName string) string {
	return fmt.Sprintf(internal+"%s/dto", appName)
}

func bizDir(appName string) string {
	return fmt.Sprintf(internal+"%s/biz", appName)
}

func serviceDir(appName string) string {
	return fmt.Sprintf(internal+"%s/service", appName)
}

func dataDir(appName string) string {
	return fmt.Sprintf(internal+"%s/data", appName)
}

func controllerDir(appName string) string {
	return fmt.Sprintf(internal+"%s/controller", appName)
}

func routerDir(appName string) string {
	return fmt.Sprintf(internal+"%s/router", appName)
}

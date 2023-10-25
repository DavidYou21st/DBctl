package main

import (
	"embed"
	"fctl/cmd"
	"fctl/pkg/global"
)

//go:embed config/*
var ConfigPaths embed.FS

func main() {
	global.ConfigPaths = ConfigPaths // 打包配置初始化
	cmd.Execute()
}

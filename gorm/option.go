package gorm

import (
	"fctl/pkg/global"
)

type options struct {
	ConfigFile   string
	OutPath      string
	OutFile      string
	ModelPkgPath string
	Table        string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetTables 设定配置文件
func SetTables(s string) Option {
	return func(o *options) {
		o.Table = s
	}
}

// SetOutPath   设定OutPath
func SetOutPath(s string) Option {
	return func(o *options) {
		o.OutPath = s
	}
}

// SetOutFile   设定OutFile
func SetOutFile(s string) Option {
	return func(o *options) {
		o.OutFile = s
	}
}

// SetModelPkgPath   设定ModelPkgPath
func SetModelPkgPath(s string) Option {
	return func(o *options) {
		o.ModelPkgPath = s
	}
}

// Init 初始化
func Init(opts ...Option) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	if v := o.ConfigFile; v != "" {
		global.VIPER = global.CFG.Viper.Viper(&global.CFG, global.ConfigPaths, v)
	}

	if v := o.Table; v != "" {
		global.CFG.Dal.Table = v
	}

	if v := o.OutPath; v != "" {
		global.CFG.Dal.OutPath = v
	}

	if v := o.OutFile; v != "" {
		global.CFG.Dal.OutFile = v
	}

	if v := o.ModelPkgPath; v != "" {
		global.CFG.Dal.ModelPkgPath = v
	}
}

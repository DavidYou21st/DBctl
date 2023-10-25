package global

import (
	"embed"
	"fctl/pkg/config"
	"github.com/spf13/viper"
	"sync"
)

var (
	CFG         = new(config.Config)
	ONCE        sync.Once
	VIPER       *viper.Viper
	ConfigPaths embed.FS
)

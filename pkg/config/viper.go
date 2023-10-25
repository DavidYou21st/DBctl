package config

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ViperConfig struct {
	Path string
}

func (v ViperConfig) Viper(cfg interface{}, ConfigPaths embed.FS, path string) *viper.Viper {
	var config string
	config = path
	fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	vv := viper.New()
	vv.SetConfigFile(config)

	err := v.ReadInConfig(vv, config, ConfigPaths)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	vv.WatchConfig()

	//监听文件修改
	vv.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := vv.Unmarshal(cfg); err != nil {
			fmt.Println(err)
		}
	})

	if err := vv.Unmarshal(cfg); err != nil {
		fmt.Println("config Unmarshal err:", err)
	}
	return vv
}

// ReadInConfig 读取配置文件 (支持Embed打包包含在二进制文件)
func (v ViperConfig) ReadInConfig(vv *viper.Viper, config string, ConfigPaths embed.FS) error {
	if err := vv.ReadInConfig(); err == nil {
		return err
	}

	file, err := ConfigPaths.ReadFile(config)

	if err != nil {
		return err
	}

	if err := vv.ReadConfig(bytes.NewReader(file)); err != nil {
		return err
	}

	fmt.Printf("您正在使用Emabed,config的路径为%v\n", config)
	return nil
}

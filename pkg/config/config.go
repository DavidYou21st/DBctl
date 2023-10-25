package config

// Config 配置参数
type Config struct {
	MySQL MySQLConfig `mapstructure:"MySQL" json:"mysql"`
	Viper ViperConfig `mapstructure:"Viper" json:"viper"`
	Dal   DalConfig   `mapstructure:"Dal" json:"dal"`
	File  FileConfig  `mapstructure:"File" json:"File"`
}

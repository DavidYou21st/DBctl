package config

import (
	"fmt"
)

// MySQLConfig MySQL mysql配置参数
type MySQLConfig struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

// DSN 数据库连接串
func (a MySQLConfig) DSN() string {
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s",
		a.User, a.Password, a.Host, a.Port, a.DBName)
}

// DSNURL 数据库连接串
func (a MySQLConfig) DSNURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		a.User, a.Password, a.Host, a.Port, a.DBName)
}

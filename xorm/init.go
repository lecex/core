package xorm

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

// Config 数据库默认配置
type Config struct {
	// Driver 主机连接方式
	Driver string
	// Host 主机地址
	Host string
	// Port 主机端口
	Port string
	// User 用户名
	User string
	// Password 密码
	Password string
	// DbName 数据库名称
	DbName string
	// Charset 数据库编码
	Charset string
}

// Connection 根据驱动创建连接
func Connection(conf *Config) (engine *xorm.Engine, err error) {
	if conf.Driver == "odbc" {
		return odbcConnection(conf)
	}
	return engine, fmt.Errorf(" '%v' driver doesn't exist. ", conf.Driver)
}

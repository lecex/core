package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// mysqlConnection 创建数据库连接
func mysqlConnection(conf *Config) (*gorm.DB, error) {
	db, _ := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s:%s)/?charset=%s&parseTime=True&loc=Local",
			conf.User, conf.Password, conf.Host, conf.Port, conf.Charset,
		),
	)
	db.Exec("CREATE DATABASE IF NOT EXISTS " + conf.DbName + ";")
	return gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			conf.User, conf.Password, conf.Host, conf.Port, conf.DbName, conf.Charset,
		),
	)
}

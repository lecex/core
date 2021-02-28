package xorm

import (
	"fmt"

	_ "github.com/alexbrainman/odbc"
	"github.com/go-xorm/xorm"
)

// odbcConnection 创建数据库连接
func odbcConnection(conf *Config) (*xorm.Engine, error) {
	return xorm.NewEngine(
		"odbc",
		fmt.Sprintf(
			"driver=freetds;server=%s;port=%s;database=%s;uid=%s;pwd=%s;TDS_Version=8.0;clientcharset=%s",
			conf.Host, conf.Port, conf.DbName, conf.User, conf.Password, conf.Charset,
		),
	)
}

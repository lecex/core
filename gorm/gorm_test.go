package gorm

import (
	"fmt"
	"testing"
)

func TestGorm(t *testing.T) {
	_, err := Connection(&Config{
		Driver: "mysql",
		// Host 主机地址
		Host: "127.0.0.1",
		// Port 主机端口
		Port: "3306",
		// User 用户名
		User: "root",
		// Password 密码
		Password: "123456",
		// DbName 数据库名称
		DbName: "user_service",
		// Charset 数据库编码
		Charset: "utf8",
	})
	fmt.Println(1)
	if err != nil {
		t.Errorf("Database connection failed, %v!", err)
	}
}

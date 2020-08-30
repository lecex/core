package pligins

import (
	// 公共引入
	_ "github.com/lecex/core/plugins/registry/not"            // 不注册插件
	_ "github.com/micro/go-plugins/client/selector/static/v2" // 静态选择器插件
)

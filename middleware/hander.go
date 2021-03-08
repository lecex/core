package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"

	client "github.com/lecex/core/client"
	"github.com/lecex/core/env"

	authPb "github.com/lecex/user/proto/auth"
	casbinPb "github.com/lecex/user/proto/casbin"
	PB "github.com/lecex/user/proto/permission"
)

// Handler 处理器
// 包含一些高阶函数比如中间件常用的 token 验证等
type Handler struct {
	UserService string
	Permissions []*PB.Permission
}

// Wrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func (srv *Handler) Wrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) (err error) {
		meta, ok := metadata.FromContext(ctx)
		// 去掉命名空间的 service 名称
		meta["Service"] = strings.Replace(req.Service(), env.Getenv("MICRO_API_NAMESPACE", "go.micro.api."), "", -1)
		meta["Method"] = req.Method()
		// 白名单黑名单拦截
		if meta["X-Real-Ip"] != "" {
			// 白名单 需要设置允许方法和对应白名单ip 对应方法只允许白名单ip通过
			if find := strings.Contains(env.Getenv("IP_WHITELIST", ""), meta["Method"]); find {
				if find := strings.Contains(env.Getenv("IP_WHITELIST", ""), meta["X-Real-Ip"]); !find {
					return errors.New("The interface is disabled")
				}
			}
			if find := strings.Contains(env.Getenv("IP_BLACKLIST", ""), meta["X-Real-Ip"]); find {
				return errors.New("Intercept illegal access")
			}
		}
		if srv.IsAuth(req) {
			if !ok {
				return errors.New("no auth meta-data found in request")
			}
			if token, ok := meta["X-Csrf-Token"]; ok {
				// Note this is now uppercase (not entirely sure why this is...)
				// token := strings.Split(meta["authorization"], "Bearer ")[1]
				// Auth here
				authRes := &authPb.Response{}
				err := client.Call(ctx, srv.UserService, "Auth.ValidateToken", &authPb.Request{
					Token: token,
				}, authRes)
				if err != nil || authRes.Valid == false {
					return err
				}
				// 设置用户 id
				meta["Userid"] = authRes.User.Id
				meta["Username"] = authRes.User.Username
				meta["Email"] = authRes.User.Email
				meta["Mobile"] = authRes.User.Mobile
				ctx = metadata.NewContext(ctx, meta)
				if srv.IsPolicy(req) {
					// 通过 meta user_id 验证权限
					casbinRes := &authPb.Response{}
					err := client.Call(ctx, srv.UserService, "Casbin.Validate", &casbinPb.Request{}, casbinRes)
					if err != nil {
						return err
					}
					if casbinRes.Valid == false {
						return errors.New("Permission denied")
					}
				}
			} else {
				return errors.New("Empty Authorization")
			}
		}
		ctx = metadata.NewContext(ctx, meta)
		err = fn(ctx, req, resp)
		return err
	}
}

// IsAuth 检测用户授权
func (srv *Handler) IsAuth(req server.Request) bool {
	for _, p := range srv.Permissions {
		if p.Method == req.Method() && p.Auth == true {
			return true
		}
	}
	return false
}

// IsPolicy 检查用户权限
func (srv *Handler) IsPolicy(req server.Request) bool {
	for _, p := range srv.Permissions {
		if p.Method == req.Method() && p.Policy == true {
			return true
		}
	}
	return false
}

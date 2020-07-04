package client

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/util/log"
)

func init() {
	cmd.Init()
}

// Call 使用默认客户端对服务进行同步调用
func Call(ctx context.Context, service string, endpoint string, req interface{}, rsp interface{}, opts ...client.CallOption) error {
	request := client.NewRequest(service, endpoint, req)
	if err := client.Call(ctx, request, rsp, opts...); err != nil {
		return err
	}
	// Stream Close 关闭数据流
	if defer err := client.Stream.Close(); err != nil {
		log.Log(err)
	}
	return nil
}

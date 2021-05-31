package client

import (
	"context"

	"github.com/micro/go-micro/v2/client"
)

// Call 使用默认客户端对服务进行同步调用
func Call(ctx context.Context, service string, endpoint string, req interface{}, rsp interface{}, opts ...client.RequestOption) error {
	request := client.NewRequest(service, endpoint, req, opts...)
	return client.Call(ctx, request, rsp)
}

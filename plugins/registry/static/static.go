// Package static provides a multicast dns registry
package static

import (
	"fmt"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/mdns"
)

func init() {
	cmd.DefaultRegistries["static"] = NewRegistry
}
func main() {
	a := NewRegistry
	fmt.Println(a)
}

type staticRegistry struct {
	r registry.Registry
}

func (m *staticRegistry) Init(opts ...registry.Option) error {
	return m.r.Init(opts...)
}

func (m *staticRegistry) Options() registry.Options {
	return m.r.Options()
}

func (m *staticRegistry) Register(service *registry.Service, opts ...registry.RegisterOption) error {
	return m.r.Register(service, opts...)
}

func (m *staticRegistry) Deregister(service *registry.Service) error {
	return m.r.Deregister(service)
}

func (m *staticRegistry) GetService(service string) ([]*registry.Service, error) {
	if selector.DefaultSelector.String() == "static" { // 静态选择器是直接返回服务
		return []*registry.Service{
			&registry.Service{Name: service},
		}, nil
	}
	return m.r.GetService(service)
}

func (m *staticRegistry) ListServices() ([]*registry.Service, error) {
	return m.r.ListServices()
}

func (m *staticRegistry) Watch(opts ...registry.WatchOption) (registry.Watcher, error) {
	return m.r.Watch(opts...)
}

func (m *staticRegistry) String() string {
	return "static"
}

// NewRegistry returns a new static registry
func NewRegistry(opts ...registry.Option) registry.Registry {
	return &staticRegistry{
		r: registry.NewRegistry(opts...),
	}
}

// Domain sets the mdnsDomain
func Domain(d string) registry.Option {
	return mdns.Domain(d)
}

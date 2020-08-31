// Package mdns provides a multicast dns registry
package mdns

import (
	"fmt"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/mdns"
)

func init() {
	cmd.DefaultRegistries["mdns"] = NewRegistry
}
func main() {
	a := NewRegistry
	fmt.Println(a)
}

type mdnsRegistry struct {
	r registry.Registry
}

func (m *mdnsRegistry) Init(opts ...registry.Option) error {
	return m.r.Init(opts...)
}

func (m *mdnsRegistry) Options() registry.Options {
	return m.r.Options()
}

func (m *mdnsRegistry) Register(service *registry.Service, opts ...registry.RegisterOption) error {
	return m.r.Register(service, opts...)
}

func (m *mdnsRegistry) Deregister(service *registry.Service) error {
	return m.r.Deregister(service)
}

func (m *mdnsRegistry) GetService(service string) ([]*registry.Service, error) {
	if selector.DefaultSelector.String() == "static" { // 静态选择器是直接返回服务
		return []*registry.Service{
			&registry.Service{Name: service},
		}, nil
	}
	return m.r.GetService(service)
}

func (m *mdnsRegistry) ListServices() ([]*registry.Service, error) {
	return m.r.ListServices()
}

func (m *mdnsRegistry) Watch(opts ...registry.WatchOption) (registry.Watcher, error) {
	return m.r.Watch(opts...)
}

func (m *mdnsRegistry) String() string {
	return m.r.String()
}

// NewRegistry returns a new mdns registry
func NewRegistry(opts ...registry.Option) registry.Registry {
	return &mdnsRegistry{
		r: registry.NewRegistry(opts...),
	}
}

// Domain sets the mdnsDomain
func Domain(d string) registry.Option {
	return mdns.Domain(d)
}

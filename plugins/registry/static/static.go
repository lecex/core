// Package static is a registry plugin for the micro static
package static

import (
	"github.com/micro/go-micro/v2/config/cmd"
	"github.com/micro/go-micro/v2/registry"
)

type static struct {
	opts registry.Options
}

func init() {
	cmd.DefaultRegistries["static"] = NewRegistry
}

func configure(s *static, opts ...registry.Option) error {
	return nil
}

func newRegistry(opts ...registry.Option) registry.Registry {
	s := &static{
		opts: registry.Options{},
	}
	configure(s, opts...)
	return s
}

func (s *static) Init(opts ...registry.Option) error {
	return configure(s, opts...)
}

func (s *static) Options() registry.Options {
	return s.opts
}

func (s *static) Register(service *registry.Service, opts ...registry.RegisterOption) error {
	return nil
}

func (s *static) Deregister(service *registry.Service) error {
	return nil
}

func (s *static) GetService(service string) (services []*registry.Service, err error) {
	services = append(services, &registry.Service{Name: service})
	return services, nil
}

func (s *static) ListServices() ([]*registry.Service, error) {
	return nil, nil
}

func (s *static) Watch(opts ...registry.WatchOption) (registry.Watcher, error) {
	return &staticWatcher{}, nil
}

func (s *static) String() string {
	return "static"
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return newRegistry(opts...)
}

// Package not is a registry plugin for the micro not
package not

import (
	"github.com/micro/go-micro/v2/cmd"
	"github.com/micro/go-micro/v2/registry"
)

type not struct {
	opts registry.Options
}

func init() {
	cmd.DefaultRegistries["not"] = NewRegistry
}

func configure(s *not, opts ...registry.Option) error {
	return nil
}

func newRegistry(opts ...registry.Option) registry.Registry {
	s := &not{
		opts: registry.Options{},
	}
	configure(s, opts...)
	return s
}

func (s *not) Init(opts ...registry.Option) error {
	return configure(s, opts...)
}

func (s *not) Options() registry.Options {
	return s.opts
}

func (s *not) Register(service *registry.Service, opts ...registry.RegisterOption) error {
	return nil
}

func (s *not) Deregister(service *registry.Service, opts ...registry.DeregisterOption) error {
	return nil
}

func (s *not) GetService(service string, opts ...registry.GetOption) ([]*registry.Service, error) {
	return nil, nil
}

func (s *not) ListServices(opts ...registry.ListOption) ([]*registry.Service, error) {
	return nil, nil
}

func (s *not) Watch(opts ...registry.WatchOption) (registry.Watcher, error) {
	return nil, nil
}

func (s *not) String() string {
	return "not"
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return newRegistry(opts...)
}

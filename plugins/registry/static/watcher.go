package static

import (
	"github.com/micro/go-micro/v2/registry"
)

type staticWatcher struct {
}

func (n *staticWatcher) Next() (*registry.Result, error) {
	return nil, nil
}

func (n *staticWatcher) Stop() {
}

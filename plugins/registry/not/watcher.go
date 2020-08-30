package not

import (
	"github.com/micro/go-micro/v2/registry"
)

type notWatcher struct {
}

func (n *notWatcher) Next() (*registry.Result, error) {
	return nil, nil
}

func (n *notWatcher) Stop() {
}

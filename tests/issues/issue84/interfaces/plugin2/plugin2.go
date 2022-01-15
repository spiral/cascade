package plugin2

import (
	"github.com/roadrunner-server/endure/tests/issues/issue84/interfaces/plugin3"
)

type Plugin2 struct{}

func (p *Plugin2) Init(_ plugin3.Fooer) error {
	return nil
}

func (p *Plugin2) Serve() chan error {
	errCh := make(chan error, 1)
	return errCh
}

func (p *Plugin2) Stop() error {
	return nil
}

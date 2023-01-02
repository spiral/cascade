package plugin1

import (
	"github.com/roadrunner-server/endure/v2/tests/happy_scenarios/plugin6"
)

type Plugin1 struct {
}

// No deps
func (s *Plugin1) Init(foow plugin6.FooWriter) error {
	foow.Fooo()
	return nil
}

func (s *Plugin1) Serve() chan error {
	errCh := make(chan error, 1)
	return errCh
}

func (s *Plugin1) Stop() error {
	return nil
}

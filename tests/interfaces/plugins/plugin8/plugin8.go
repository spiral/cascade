package plugin8

import (
	"fmt"

	"github.com/spiral/endure"
	"github.com/spiral/endure/tests/interfaces/plugins/plugin10"
)

type SomeInterface interface {
	Boom()
}

type Plugin8 struct {
}

// No deps
func (s *Plugin8) Init() error {
	return nil
}

func (s *Plugin8) Serve() chan error {
	errCh := make(chan error, 1)
	return errCh
}

func (s *Plugin8) Name() string {
	return "plugin8"
}

func (s *Plugin8) Collects() []interface{} {
	return []interface{}{
		s.SomeCollects,
	}
}

func (s *Plugin8) SomeCollects(named endure.Named, b SomeInterface, p10 *plugin10.Plugin10) error {
	fmt.Println(named.Name())
	b.Boom()
	return nil
}

func (s *Plugin8) Stop() error {
	return nil
}

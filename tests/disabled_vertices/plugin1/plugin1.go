package plugin1

import "github.com/roadrunner-server/errors"

type Plugin1 struct {
}

func (p *Plugin1) Init() error {
	const op = errors.Op("init")
	return errors.E(op, errors.Disabled)
}

package milv

import "github.com/magicmatatjahu/milv/cli"

type Milv struct {
	config 	*Config
	hooks 	*Hooks
}

func New(commands cli.Commands, config *Config) (*Milv, error) {
	config, err := newConfig(commands, config)
	if err != nil {
		return nil, err
	}

	milv := new(Milv)
	milv.config = config

	return milv, nil
}

func (m *Milv) Run() error {
	return nil
}

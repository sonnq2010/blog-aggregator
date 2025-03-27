package command

import (
	"errors"

	"github.com/sonnq2010/blog-aggregator/internal/state"
)

type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	c.Handlers[name] = f
}

func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, ok := c.Handlers[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}

	return handler(s, cmd)
}

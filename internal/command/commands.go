package command

import (
	"fmt"

	"github.com/mecebeci/blog-aggregator/internal/state"
)

type Commands struct {
	Handlers map[string]func(*state.State, Command) error
}

func (c *Commands) Register(name string, f func(*state.State, Command) error) {
	if c.Handlers == nil {
		c.Handlers = make(map[string]func(*state.State, Command) error)
	}
	c.Handlers[name] = f
}


func (c *Commands) Run(s *state.State, cmd Command) error {
	handler, exists := c.Handlers[cmd.Name]
	if !exists{
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
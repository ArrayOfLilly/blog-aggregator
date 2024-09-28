package main

import (
	"errors"
)

func (c *commands) register(name string, f func(s *state, cmd command) error)  {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.handlers[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}


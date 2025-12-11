package main

import (
	"errors"
	"fmt"

	"github.com/aott33/gator/internal/config"
)

type state struct {
	cfg		*config.Config
}

type command struct {
	name	string
	args	[]string	
}


type commands struct {
	m		map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Set user to %s\n",cmd.args[0])
	
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.m[cmd.name]

	if !ok {
		return errors.New("Could not find command")
	}

	err := value(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.m == nil {
		c.m = make(map[string]func(*state, command) error)
	}

	c.m[name] = f	
}

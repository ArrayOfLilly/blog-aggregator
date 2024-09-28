package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}

	name := cmd.Args[0]

	err := s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("Logged in user set to %s\n", s.cfg.CurrentUserName)
	return nil
}
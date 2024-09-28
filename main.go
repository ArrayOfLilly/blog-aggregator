package main

import (
	"fmt"
	"os"

	"github.com/ArrayOfLilly/blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

var (
	s state
	cmd command
	c commands
)
	
func main() {
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments were provided")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	s = state{
		cfg: &cfg,
	}
	
	cmd = command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	c = commands{
		handlers: make(map[string]func(s *state, cmd command) error),
	}

	c.register("login", handlerLogin)

	err = c.run(&s, cmd)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}


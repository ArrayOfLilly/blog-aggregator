package main

type command struct {
	Name	string	 `json:"name"`
	Args	[]string `json:"args"`
}

type commands struct {
	handlers map[string]func(s *state, cmd command) error
}


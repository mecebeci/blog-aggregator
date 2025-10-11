package handlers

import (
	"errors"
	"fmt"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleLogin(s *state.State, cmd command.Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("usage: gator login <username>")
	}

	username := cmd.Args[0]

	err := s.Config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User set to: ", username)
	return nil
}

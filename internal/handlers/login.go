package handlers

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleLogin(s *state.State, cmd command.Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("usage: gator login <username>")
	}

	username := cmd.Args[0]

	user, err := s.DB.GetUserByName(context.Background(), username)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: user does not exist")
		os.Exit(1)
	}

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("Logged in as: %s (id: %s)\n", user.Name, user.ID.String())
	return nil
}

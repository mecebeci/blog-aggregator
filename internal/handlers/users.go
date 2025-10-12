package handlers

import (
	"context"
	"fmt"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleUsers(s *state.State, cmd command.Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not fetch users: %w", err)
	}

	if len(users) < 1 {
		fmt.Println("No users found.")
		return nil
	}

	for _, user := range users {
		if user.Name != s.Config.CurrentUserName {
			fmt.Printf("* %s\n", user.Name)
		} else {
			fmt.Printf("* %s (current)\n", user.Name)
		}
	}
	return nil
}

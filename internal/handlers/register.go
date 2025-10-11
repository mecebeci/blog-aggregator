package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/database"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleRegister(s *state.State, cmd command.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("missing username: usage `register <username>`")
	}

	username := cmd.Args[0]
	id := uuid.New()
	now := time.Now()

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      username,
	})
	if err != nil {
		if err.Error() != "" && strings.Contains(err.Error(), "duplicate key") {
			fmt.Fprintln(os.Stderr, "User already exists")
			os.Exit(1)
		}
		return err
	}

	if err := s.Config.SetUser(username); err != nil {
		return err
	}
	fmt.Printf("User %s successfully registered with id: %s\n", user.Name, user.ID.String())
	return nil
}

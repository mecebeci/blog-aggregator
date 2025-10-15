package middleware

import (
	"context"
	"errors"
	"fmt"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/database"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

var ErrNotLoggedIn = errors.New("no user is currently logged in")

type LoggedInHandler func(s *state.State, cmd command.Command, user database.User) error

func MiddlewareLoggedIn(h LoggedInHandler) func(*state.State, command.Command) error {
	return func(s *state.State, cmd command.Command) error {
		if s.Config.CurrentUserName == "" {
			return ErrNotLoggedIn
		}

		u, err := s.DB.GetUserByName(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("failed to find logged-in user %q: %w", s.Config.CurrentUserName, err)
		}
		return h(s, cmd, u)
	}
}

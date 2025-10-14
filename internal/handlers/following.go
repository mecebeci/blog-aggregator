package handlers

import (
	"context"
	"fmt"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleFollowing(s *state.State, cmd command.Command) error {
	user, err := s.DB.GetUserByName(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("could not get current user: %w", err)
	}

	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could not get follows: %w", err)
	}

	if len(follows) == 0 {
		fmt.Printf("%s is not following any feeds yet.\n", user.Name)
		return nil
	}

	for _, follow := range follows {
		fmt.Println(follow.FeedName)
	}
	return nil
}
package handlers

import (
	"context"
	"fmt"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/database"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleFollowing(s *state.State, cmd command.Command, user database.User) error {
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
package handlers

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/database"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleFollow(s *state.State, cmd command.Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: follow <feed-url>")
	}

	feedURL := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("feed not found URL %s: %w", feedURL, err)
	}

	follow, err := s.DB.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:     uuid.New(),
			UserID: user.ID,
			FeedID: feed.ID,
		})
	if err != nil {
		return fmt.Errorf("could not follow feed: %w", err)
	}

	fmt.Printf("User %s followed feed %s\n", follow.UserName, follow.FeedName)
	return nil
}

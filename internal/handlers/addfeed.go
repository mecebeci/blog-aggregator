package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/database"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleAddFeed(s *state.State, cmd command.Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]
	now := time.Now()

	feed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	follow, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Println("Feed created successfully:")
	fmt.Printf("- ID: %s\n", feed.ID.String())
	fmt.Printf("- Name: %s\n", feed.Name)
	fmt.Printf("- URL: %s\n", feed.Url)
	fmt.Printf("- UserID: %s\n", feed.UserID.String())
	fmt.Printf("- CreatedAt: %s\n", feed.CreatedAt.Format(time.RFC3339))

	fmt.Printf("\nUser %s is now following feed %s\n", follow.UserName, follow.FeedName)

	return nil
}

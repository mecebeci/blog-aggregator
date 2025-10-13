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

func HandleAddFeed(s *state.State, cmd command.Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	user, err := s.DB.GetUserByName(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to find logged-in user: %w", err)
	}

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

	fmt.Println("Feed created successfully:")
	fmt.Printf("- ID: %s\n", feed.ID.String())
	fmt.Printf("- Name: %s\n", feed.Name)
	fmt.Printf("- URL: %s\n", feed.Url)
	fmt.Printf("- UserID: %s\n", feed.UserID.String())
	fmt.Printf("- CreatedAt: %s\n", feed.CreatedAt.Format(time.RFC3339))
	return nil
}

package handlers

import (
	"context"
	"fmt"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleFeeds(s *state.State, cmd command.Command) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("could not fetch feeds: %w", err)
	}

	if len(feeds) < 1 {
		fmt.Println("No feeds found")
		return nil
	}

	for _, feed := range feeds {
		user, err := s.DB.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			fmt.Printf("Could not get user for feed %s\n", feed.Name)
		}

		fmt.Printf("Name of the feed: %s\n", feed.Name)
		fmt.Printf("URL of the feed: %s\n", feed.Url)
		fmt.Printf("The name of the user: %s\n", user.Name)
	}
	return nil
}

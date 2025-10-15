package handlers

import (
	"context"
	"fmt"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/database"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleUnfollow(s *state.State, cmd command.Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: unfollow <feed-url>")
	}

	feedUrl := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("feed not found URL %s: %w", feedUrl, err)
	}

	err = s.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("follow failed %w", err)
	}

	fmt.Println("Feed unfollowed successfully!")
	return nil
}

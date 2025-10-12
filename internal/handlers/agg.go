package handlers

import (
	"context"
	"fmt"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/feed"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleAgg(s *state.State, cmd command.Command) error {
	feedUrl := "https://www.wagslane.dev/index.xml"

	rss, err := feed.FetchFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch feed: %w: ", err)
	}

	fmt.Printf("Fetched feed: \n%+v\n", rss)
	return nil
}

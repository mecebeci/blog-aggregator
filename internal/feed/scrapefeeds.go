package feed

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mecebeci/blog-aggregator/internal/state"
)

func ScrapeFeeds(s *state.State) error {
	ctx := context.Background()

	nextFeed, err := s.DB.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("failed to get next feed: %w", err)
	}

	fmt.Printf("Fetching feed: %s (%s)\n", nextFeed.Name, nextFeed.Url)

	err = s.DB.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark feed fetched: %w", err)
	}

	rss, err := FetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch RSS feed: %w", err)
	}

	for _, item := range rss.Items {
		fmt.Printf("• %s\n", item.Title)
	}
	log.Printf("Finished fetching feed: %s at %s\n", nextFeed.Name, time.Now().Format(time.RFC3339))
	return nil
}

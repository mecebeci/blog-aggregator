package feed

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mecebeci/blog-aggregator/internal/database"
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
		publishedAt, err := time.Parse(time.RFC1123Z, item.Published)
		if err != nil {
			log.Printf("failed to parse published time for %s: %v", item.Title, err)
		}

		_, err = s.DB.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value"){
				continue
			}
			log.Printf("failed to insert post (%s): %v", err)
			continue
		}
	}

	log.Printf("Finished fetching feed: %s at %s\n", nextFeed.Name, time.Now().Format(time.RFC3339))
	return nil
}

package handlers

import (
	"fmt"
	"time"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/feed"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func HandleAgg(s *state.State, cmd command.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: agg <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}

	fmt.Printf("Collecting feeds every %s...\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	if err := feed.ScrapeFeeds(s); err != nil {
		fmt.Println("initial scrape error: ", err)
	}

	for range ticker.C {
		if err := feed.ScrapeFeeds(s); err != nil {
			fmt.Println("scrape error: ", err)
		}
	}
	return nil
}

package feed

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Title       string `xml:"channel>title"`
	Description string `xml:"channel>description"`
	Link        string `xml:"channel>link"`
	Items       []Item `xml:"channel>item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Published   string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	if resp.Body == nil {
		return nil, errors.New("response body is nil")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal feed XML: %w", err)
	}

	feed.Title = html.UnescapeString(feed.Title)
	feed.Description = html.UnescapeString(feed.Description)
	for i := range feed.Items {
		feed.Items[i].Title = html.UnescapeString(feed.Items[i].Title)
		feed.Items[i].Description = html.UnescapeString(feed.Items[i].Description)
	}

	return &feed, nil
}

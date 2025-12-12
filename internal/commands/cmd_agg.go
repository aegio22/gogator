package commands

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/database"
	"github.com/aegio22/gogator/internal/rss"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func HandlerAgg(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("agg takes one argument: a time between requests string (eg. 1m, 1s)")
	}
	time_between_reqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error parsing time argument: %w", err)
	}

	fmt.Println("Collecting feeds every", time_between_reqs.String())

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *config.State) error {
	ctx := context.Background()

	nextFeed, err := s.DbQueries.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}

	err = s.DbQueries.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		ID:        nextFeed.ID,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	fetchedFeed, err := rss.FetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed by url: %w", err)
	}
	for _, item := range fetchedFeed.Channel.Item {

		pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			// Try RFC1123
			pubTime, err = time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				// If all else fails, use current time
				pubTime = time.Now()
			}
		}
		//fmt.Println("Item:", item.Title)
		err = s.DbQueries.CreatePost(ctx, database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: item.Title, Valid: item.Title != ""},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
			PublishedAt: pubTime,
			FeedID:      nextFeed.ID,
		})

		if err != nil {
			// lib/pq example:
			if pgErr, ok := err.(*pq.Error); ok {
				if pgErr.Code == "23505" { // unique_violation
					continue

				}
			}
			log.Printf("error inserting post: %v", err)
			continue

		}
	}

	return nil
}

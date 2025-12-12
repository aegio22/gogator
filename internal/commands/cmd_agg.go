package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/database"
	"github.com/aegio22/gogator/internal/rss"
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

	fmt.Println()
	fmt.Println()
	fmt.Println("Feed Name:", fetchedFeed.Channel.Title)
	fmt.Println("-----------")

	fmt.Println("Feed Items:")
	fmt.Println("-----------")
	for _, item := range fetchedFeed.Channel.Item {
		fmt.Println("Item:", item.Title)
	}

	return nil
}

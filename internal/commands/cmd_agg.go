package commands

import (
	"context"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/rss"
)

func HandlerAgg(s *config.State, cmd Command) error {
	ctx := context.Background()
	feed, err := rss.FetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

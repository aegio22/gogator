package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
)

func HandlerFeeds(s *config.State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("the 'feeds' command takes no arguments")
	}

	ctx := context.Background()

	feeds, err := s.DbQueries.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println("FeedName:", feed.FeedName)
		fmt.Println("Url:", feed.Url)
		fmt.Println("UserName:", feed.UserName)
	}
	return nil
}

package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/database"
)

func HandlerUnfollow(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("the unfollow command takes one URL argument")
	}

	ctx := context.Background()

	feed, err := s.DbQueries.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error fetching feed from url '%s': %w", cmd.Args[0], err)
	}
	err = s.DbQueries.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: feed.UserID,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}

	fmt.Println("Feed unfollowed successfully!")
	return nil
}

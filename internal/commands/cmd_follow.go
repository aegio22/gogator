package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return errors.New("follow command takes a single url argument")
	}

	ctx := context.Background()

	feed, err := s.DbQueries.GetFeedByURL(ctx, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	follow, err := s.DbQueries.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Printf("feed name: %v\nuser name: %v\n", follow.FeedName, follow.UserName)

	return nil
}

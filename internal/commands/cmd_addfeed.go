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

func HandlerAddFeed(s *config.State, cmd Command,user database.User) error {
	if len(cmd.Args) != 2 {
		return errors.New("add feed command takes 2 arguments, a feed name and feed url")
	}
	ctx := context.Background()



	feed, err := s.DbQueries.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error adding feed: %w", err)
	}
	_, err = s.DbQueries.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed after creation: %w", err)
	}

	fmt.Println("feed created successfully!")
	fmt.Printf("ID: %v\nCreatedAt: %v\nUpdatedAt: %v\nName: %v\nUrl: %v\nUserID: %v\n",
		feed.ID, feed.CreatedAt, feed.UpdatedAt, feed.Name, feed.Url, feed.UserID)
	return nil

}

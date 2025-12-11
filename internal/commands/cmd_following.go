package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
)

func HandlerFollowing(s *config.State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return errors.New("follows command takes no arguments")
	}
	ctx := context.Background()
	currUser, err := s.DbQueries.GetUser(ctx, s.CfgPointer.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching user: %w", err)
	}
	follows, err := s.DbQueries.GetFeedFollowsForUser(ctx, currUser.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows for user %s: %w", s.CfgPointer.CurrentUserName, err)
	}

	for _, follow := range follows {
		fmt.Println("Feed Name:", follow.FeedName)
		fmt.Println("User Name:", follow.UserName)
		fmt.Println("-----")
	}

	return nil
}

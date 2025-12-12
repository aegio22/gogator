package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/database"
)

func HandlerFollowing(s *config.State, cmd Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return errors.New("follows command takes no arguments")
	}
	ctx := context.Background()

	follows, err := s.DbQueries.GetFeedFollowsForUser(ctx, user.ID)
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

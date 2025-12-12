package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/database"
)

func MiddlewareLoggedIn(handler func(s *config.State, cmd Command, user database.User) error) func(*config.State, Command) error {
	ctx := context.Background()

	formatted_func := func(s *config.State, cmd Command) error {
		currUser := s.CfgPointer.CurrentUserName
		if currUser == "" {
			return errors.New("no current user found")
		}
		user, err := s.DbQueries.GetUser(ctx, currUser)
		if err != nil {
			return fmt.Errorf("error fetching user: %w", err)
		}
		return handler(s, cmd, user)
	}
	return formatted_func
}

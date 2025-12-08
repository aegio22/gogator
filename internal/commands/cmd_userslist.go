package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
)

func HandlerListUsers(s *config.State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("users command takes no arguments")
	}
	users, err := s.DbQueries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting list of users: %v", err)
	}
	if len(users) == 0 {

		return errors.New("no users found")
	}

	currUser := s.CfgPointer.CurrentUserName
	for _, user := range users {
		line := fmt.Sprintf("* %s ", user)
		if user == currUser {
			line += "(current)"
		}
		fmt.Println(line)
	}
	return nil
}

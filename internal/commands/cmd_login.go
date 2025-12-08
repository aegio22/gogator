package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/database"
)

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("no arguments provided- login command expects a username")
	}

	if len(cmd.Args) > 1 {
		return errors.New("too many arguments provided- login command expects a username")
	}
	userVerification, _ := s.DbQueries.GetUser(context.Background(), cmd.Args[0])
	if userVerification == (database.User{}) {
		return fmt.Errorf("user '%s' not found in database", cmd.Args[0])
	}

	//Debug
	//fmt.Printf("args given to login handler: %v", cmd.Args)

	arg := cmd.Args[0]
	err := s.CfgPointer.SetUser(arg)

	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Println("User has been set")
	return nil
}

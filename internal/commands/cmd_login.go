package commands

import (
	"errors"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
)

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("no arguments provided- login command expects a username")
	}

	if len(cmd.Args) > 1 {
		return errors.New("too many arguments provided- login command expects a username")
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

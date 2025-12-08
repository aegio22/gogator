package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/aegio22/gogator/internal/config"
)

func HandlerReset(s *config.State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return errors.New("reset command takes no arguments")
	}

	err := s.DbQueries.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting database: %v", err)

	}

	return nil
}

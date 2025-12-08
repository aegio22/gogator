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

func HandlerRegister(s *config.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("no arguments provided- register command expects a name")
	}

	if len(cmd.Args) > 1 {
		return errors.New("too many arguments provided- register command expects a name")
	}

	ctx := context.Background()

	userVerification, _ := s.DbQueries.GetUser(ctx, cmd.Args[0])
	if userVerification != (database.User{}) {

		return fmt.Errorf("user '%s' already in database", cmd.Args[0])
	}

	user, err := s.DbQueries.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	err = s.CfgPointer.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}
	fmt.Println("User created successfully:")
	fmt.Printf("User ID:%v  CreatedAt:%v  UpdatedAt:%v Name: %s\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)

	return nil
}

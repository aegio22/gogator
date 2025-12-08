package commands

import (
	"fmt"
	"log"

	"github.com/aegio22/gogator/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CommandMap map[string]func(*config.State, Command) error
}

func (c *Commands) Run(s *config.State, cmd Command) error {
	cmdFunc, exists := c.CommandMap[cmd.Name]
	if !exists {
		return fmt.Errorf("command doesn't exist: %v", cmd.Name)
	}
	err := cmdFunc(s, cmd)
	if err != nil {
		return fmt.Errorf("error running command %s: %v", cmd.Name, err)
	}
	return nil
}

func (c *Commands) Register(name string, f func(*config.State, Command) error) {
	if name == "" {
		log.Fatal("no name provided")
	}
	c.CommandMap[name] = f
}

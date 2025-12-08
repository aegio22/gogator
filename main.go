package main

import (
	"log"
	"os"

	"github.com/aegio22/gogator/internal/commands"
	"github.com/aegio22/gogator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatalln("not enough arguments given")

	}

	state := config.State{CfgPointer: &cfg}

	cmds := commands.Commands{CommandMap: map[string]func(*config.State, commands.Command) error{}}
	cmds.Register("login", commands.HandlerLogin)

	commandName := os.Args[1]
	args := os.Args[2:]
	cmdToRun := commands.Command{Name: commandName, Args: args}

	err = cmds.Run(&state, cmdToRun)
	if err != nil {
		log.Fatalln(err)

	}
}

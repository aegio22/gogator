package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/aegio22/gogator/internal/commands"
	"github.com/aegio22/gogator/internal/config"
	"github.com/aegio22/gogator/internal/database"
	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", state.CfgPointer.DbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	state.DbQueries = dbQueries

	cmds := commands.Commands{CommandMap: map[string]func(*config.State, commands.Command) error{}}
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegister)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerListUsers)
	cmds.Register("agg", commands.HandlerAgg)
	cmds.Register("addfeed", commands.MiddlewareLoggedIn(commands.HandlerAddFeed))
	cmds.Register("feeds", commands.HandlerFeeds)
	cmds.Register("follow", commands.MiddlewareLoggedIn(commands.HandlerFollow))
	cmds.Register("following", commands.MiddlewareLoggedIn(commands.HandlerFollowing))
	cmds.Register("unfollow", commands.MiddlewareLoggedIn(commands.HandlerUnfollow))
	cmds.Register("browse", commands.MiddlewareLoggedIn(commands.HandlerBrowse))
	commandName := os.Args[1]
	args := os.Args[2:]
	cmdToRun := commands.Command{Name: commandName, Args: args}

	err = cmds.Run(&state, cmdToRun)
	if err != nil {
		log.Fatalln(err)

	}
}

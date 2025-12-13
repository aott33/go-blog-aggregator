package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/aott33/gator/internal/config"
	"github.com/aott33/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	var state state
	var cmds commands

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error!\n%v\n", err)
		os.Exit(1)
	}

	state.cfg = &cfg

	db, err := sql.Open("postgres", state.cfg.DbURL)
	if err != nil {
		fmt.Printf("Error!\n%v\n",err)
		os.Exit(1)
	}

	dbQueries := database.New(db)
	state.db = dbQueries

	cmds.register("reset", handlerReset)
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerCreateFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollowFeed))
	cmds.register("following", middlewareLoggedIn(handlerGetFeedFollows))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))


	args := os.Args

	if len(args) < 2 {
		fmt.Println("Error! Not enough args")
		os.Exit(1)
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	cmd := command{name: cmdName, args: cmdArgs}

	err = cmds.run(&state, cmd)
	if err != nil {
		fmt.Printf("Error!\n%v\n",err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Read Error: %v\n", err)
	}	
}

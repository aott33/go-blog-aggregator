package main

import (
	"fmt"
	"os"

	"github.com/aott33/gator/internal/config"
)

func main() {
	var state state
	var cmds commands

	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Read Error: %v\n", err)
	}

	state.cfg = &cfg

	cmds.register("login", handlerLogin)

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

	fmt.Printf("%+v\n",cfg)
}

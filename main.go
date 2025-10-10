package main

import (
	"log"
	"os"

	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/config"
	"github.com/mecebeci/blog-aggregator/internal/handlers"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func main(){
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config", err)
	}

	appState := &state.State{
		Config: &cfg,
	}

	cmds := command.Commands{}
	cmds.Register("login", handlers.HandleLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("No command provided. Example usage: go run . login <username>")
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	cmd := command.Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	err = cmds.Run(appState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
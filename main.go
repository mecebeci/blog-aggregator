package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mecebeci/blog-aggregator/internal/command"
	"github.com/mecebeci/blog-aggregator/internal/config"
	"github.com/mecebeci/blog-aggregator/internal/database"
	"github.com/mecebeci/blog-aggregator/internal/handlers"
	"github.com/mecebeci/blog-aggregator/internal/state"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config", err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("Failed to connect to databse:", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	appState := &state.State{
		Config: &cfg,
		DB:     dbQueries,
	}

	cmds := command.Commands{}
	cmds.Register("login", handlers.HandleLogin)
	cmds.Register("register", handlers.HandleRegister)
	cmds.Register("reset", handlers.HandleReset)
	cmds.Register("users", handlers.HandleUsers)
	cmds.Register("agg", handlers.HandleAgg)

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

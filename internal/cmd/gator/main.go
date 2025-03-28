package gator

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/sonnq2010/blog-aggregator/internal/command"
	"github.com/sonnq2010/blog-aggregator/internal/config"
	"github.com/sonnq2010/blog-aggregator/internal/database"
	"github.com/sonnq2010/blog-aggregator/internal/handler"
	"github.com/sonnq2010/blog-aggregator/internal/state"
)

func Run() {
	commands := &command.Commands{
		Handlers: make(map[string]func(*state.State, command.Command) error),
	}
	commands.Register("login", handler.LoginHandler)

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	state := &state.State{
		Config: &cfg,
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error opening db: %v", err)
	}

	dbQueries := database.New(db)
	state.DB = dbQueries

	args := os.Args
	if len(args) < 2 {
		log.Fatal("invalid arguments")
	}

	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := command.Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	err = commands.Run(state, cmd)
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}

}

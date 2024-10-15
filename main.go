package main

import (
	"database/sql"
	"example.com/username/blog-aggregator/cheese/internal/database"
	_ "github.com/lib/pq"
)
import (
	"example.com/username/blog-aggregator/cheese/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// open a connection to the database
	db, err2 := sql.Open("postgres", cfg.DBURL)

	if err2 != nil {
		log.Fatalf("error opening database connection: %v", err)
	}

	// create a new query
	dbQueries := database.New(db)

	// set the config and dbQuery to the programState
	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAggregate)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", handlerFollowFeed)
	cmds.register("following", handlerFollowing)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}

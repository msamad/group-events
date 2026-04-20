package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	internalmigrate "github.com/msamad/group-events/backend/internal/migrate"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "migrations", "path to migration directory")
	flag.Parse()

	command := "up"
	if flag.NArg() > 0 {
		command = flag.Arg(0)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/group_events_test?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}
	defer db.Close()

	switch command {
	case "up":
		if err := internalmigrate.Up(db, dir); err != nil {
			log.Fatalf("run migration command %q: %v", command, err)
		}
	case "down":
		if err := internalmigrate.Down(db, dir); err != nil {
			log.Fatalf("run migration command %q: %v", command, err)
		}
	default:
		log.Fatalf("unsupported migration command %q", command)
	}

	fmt.Printf("migration command %q completed\n", command)
}

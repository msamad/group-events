package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func DefaultDatabaseURL() string {
	if value := os.Getenv("DATABASE_URL"); value != "" {
		return value
	}

	return "postgres://postgres:postgres@localhost:5432/group_events_test?sslmode=disable"
}

func OpenTestDB(t *testing.T) *sql.DB {
	t.Helper()

	dsn := DefaultDatabaseURL()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("open test db: %v", err)
	}

	db.SetMaxOpenConns(2)
	db.SetConnMaxLifetime(30 * time.Second)

	if err := db.Ping(); err != nil {
		_ = db.Close()
		t.Skipf("skipping postgres-backed test, database unavailable: %v", err)
	}

	t.Cleanup(func() {
		_ = db.Close()
	})

	return db
}

func ResetSchema(t *testing.T, db *sql.DB, schema string) {
	t.Helper()

	if _, err := db.Exec(fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schema)); err != nil {
		t.Fatalf("drop schema %s: %v", schema, err)
	}

	if _, err := db.Exec(fmt.Sprintf("CREATE SCHEMA %s", schema)); err != nil {
		t.Fatalf("create schema %s: %v", schema, err)
	}
}

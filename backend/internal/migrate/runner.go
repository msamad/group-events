package migrate

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
)

func Up(db *sql.DB, dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*.up.sql"))
	if err != nil {
		return fmt.Errorf("find up migrations: %w", err)
	}
	sort.Strings(files)

	for _, file := range files {
		if err := execSQLFile(db, file); err != nil {
			return fmt.Errorf("apply migration %s: %w", filepath.Base(file), err)
		}
	}

	return nil
}

func Down(db *sql.DB, dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*.down.sql"))
	if err != nil {
		return fmt.Errorf("find down migrations: %w", err)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	for _, file := range files {
		if err := execSQLFile(db, file); err != nil {
			return fmt.Errorf("rollback migration %s: %w", filepath.Base(file), err)
		}
	}

	return nil
}

func execSQLFile(db *sql.DB, file string) error {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	if _, err := db.Exec(string(contents)); err != nil {
		return fmt.Errorf("execute sql: %w", err)
	}

	return nil
}

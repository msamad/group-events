package migrations

import (
	"path/filepath"
	"runtime"
	"testing"

	internalmigrate "github.com/msamad/group-events/backend/internal/migrate"
	"github.com/msamad/group-events/backend/internal/testutil"
)

func TestMigrationsApplyAndRollback(t *testing.T) {
	t.Parallel()

	db := testutil.OpenTestDB(t)

	_, currentFile, _, _ := runtime.Caller(0)
	migrationsDir := filepath.Clean(filepath.Join(filepath.Dir(currentFile), "..", "..", "migrations"))

	if err := internalmigrate.Up(db, migrationsDir); err != nil {
		t.Fatalf("apply migrations: %v", err)
	}

	expectedTables := []string{"groups", "group_members", "events", "polls", "poll_options", "poll_votes"}
	for _, table := range expectedTables {
		var exists bool
		if err := db.QueryRow(`SELECT EXISTS (
			SELECT 1 FROM information_schema.tables WHERE table_schema='public' AND table_name=$1
		)`, table).Scan(&exists); err != nil {
			t.Fatalf("check table %s: %v", table, err)
		}
		if !exists {
			t.Fatalf("expected table %s to exist after migration", table)
		}
	}

	if err := internalmigrate.Down(db, migrationsDir); err != nil {
		t.Fatalf("rollback migrations: %v", err)
	}
}

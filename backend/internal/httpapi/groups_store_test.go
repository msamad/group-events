package httpapi

import (
	"net/http/httptest"
	"testing"

	"github.com/msamad/group-events/backend/internal/domain"
)

func TestGroupStoreAndRoleHelpers(t *testing.T) {
	t.Parallel()

	store := newGroupStore()
	created := store.create("Core", "core", "desc")
	if created.ID == "" {
		t.Fatal("expected created group id")
	}

	listed := store.list(10, 0)
	if len(listed) != 1 {
		t.Fatalf("expected one group, got %d", len(listed))
	}

	if _, err := store.get(created.ID); err != nil {
		t.Fatalf("expected group to exist: %v", err)
	}

	updated, err := store.update(created.ID, "Core Team", "core-team", "updated")
	if err != nil {
		t.Fatalf("expected update to succeed: %v", err)
	}
	if updated.Name != "Core Team" {
		t.Fatalf("expected updated name, got %q", updated.Name)
	}

	if err := store.delete(created.ID); err != nil {
		t.Fatalf("expected delete to succeed: %v", err)
	}

	if err := store.delete(created.ID); err == nil {
		t.Fatal("expected deleting missing group to fail")
	}

	missingID := domain.GroupID("missing")
	if _, err := store.get(missingID); err == nil {
		t.Fatal("expected get on missing group to fail")
	}
	if _, err := store.update(missingID, "x", "x", "x"); err == nil {
		t.Fatal("expected update on missing group to fail")
	}
}

func TestRoleMappingAndDescriptorBuilder(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest("GET", "/api/v1/groups", nil)
	req.Header.Set("X-Group-Role", "role_assigned")
	role := roleFromRequest(req)
	if role != domain.RoleOrganizer {
		t.Fatalf("expected organizer role, got %q", role)
	}

	descriptor := buildGroupUIDescriptor(domain.RoleViewer, false)
	if descriptor.Screen == "" {
		t.Fatal("expected descriptor screen")
	}

	for _, action := range descriptor.Actions {
		if action.ID == "create_event" {
			t.Fatal("viewer should not see create_event")
		}
	}
}

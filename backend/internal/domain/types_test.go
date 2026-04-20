package domain

import (
	"testing"
	"time"
)

func TestRolePermissions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		role            Role
		manageGroup     bool
		manageMembers   bool
		createEvents    bool
		respond         bool
	}{
		{name: "owner", role: RoleOwner, manageGroup: true, manageMembers: true, createEvents: true, respond: true},
		{name: "admin", role: RoleAdmin, manageGroup: true, manageMembers: true, createEvents: true, respond: true},
		{name: "organizer", role: RoleOrganizer, createEvents: true, respond: true},
		{name: "member", role: RoleMember, respond: true},
		{name: "viewer", role: RoleViewer},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.role.CanManageGroup(); got != test.manageGroup {
				t.Fatalf("CanManageGroup() = %t, want %t", got, test.manageGroup)
			}

			if got := test.role.CanManageMembership(); got != test.manageMembers {
				t.Fatalf("CanManageMembership() = %t, want %t", got, test.manageMembers)
			}

			if got := test.role.CanCreateEvents(); got != test.createEvents {
				t.Fatalf("CanCreateEvents() = %t, want %t", got, test.createEvents)
			}

			if got := test.role.CanRespond(); got != test.respond {
				t.Fatalf("CanRespond() = %t, want %t", got, test.respond)
			}
		})
	}
}

func TestMembershipDelegatesRolePermissions(t *testing.T) {
	t.Parallel()

	membership := Membership{GroupID: "group-1", UserID: "user-1", Role: RoleOrganizer}
	if !membership.CanCreateEvents() {
		t.Fatal("expected organizer membership to create events")
	}

	if !membership.CanRespond() {
		t.Fatal("expected organizer membership to respond")
	}

	if !membership.IsValid() {
		t.Fatal("expected membership with role and identities to be valid")
	}
}

func TestEventHasValidSchedule(t *testing.T) {
	t.Parallel()

	start := time.Date(2026, time.April, 20, 18, 0, 0, 0, time.UTC)
	end := start.Add(2 * time.Hour)

	event := Event{StartsAt: start, EndsAt: end}
	if !event.HasValidSchedule() {
		t.Fatal("expected event schedule to be valid")
	}

	invalid := Event{StartsAt: end, EndsAt: start}
	if invalid.HasValidSchedule() {
		t.Fatal("expected reversed event schedule to be invalid")
	}
}

func TestGroupAndMemberValidation(t *testing.T) {
	t.Parallel()

	group := Group{ID: "group-1", Slug: "platform", Name: "Platform"}
	if !group.IsValid() {
		t.Fatal("expected group with required fields to be valid")
	}

	member := Member{ID: "user-1", DisplayName: "Mustansar"}
	if !member.IsValid() {
		t.Fatal("expected member with required fields to be valid")
	}

	invalidMember := Member{ID: "user-2"}
	if invalidMember.IsValid() {
		t.Fatal("expected member without display name to be invalid")
	}
}

func TestPollHasValidConfiguration(t *testing.T) {
	t.Parallel()

	poll := Poll{
		Options: []PollOption{
			{ID: "one", Label: "One"},
			{ID: "two", Label: "Two"},
		},
		MinSelections: 1,
		MaxSelections: 1,
	}

	if !poll.HasValidConfiguration() {
		t.Fatal("expected poll configuration to be valid")
	}

	poll.MaxSelections = 3
	if poll.HasValidConfiguration() {
		t.Fatal("expected poll with too many max selections to be invalid")
	}
}

func TestEventAndPollValidation(t *testing.T) {
	t.Parallel()

	start := time.Date(2026, time.April, 20, 18, 0, 0, 0, time.UTC)
	end := start.Add(90 * time.Minute)

	event := Event{
		ID:        "event-1",
		GroupID:   "group-1",
		Title:     "Quarterly planning",
		StartsAt:  start,
		EndsAt:    end,
		CreatedBy: "user-1",
	}
	if !event.IsValid() {
		t.Fatal("expected event with required fields and valid schedule to be valid")
	}

	poll := Poll{
		ID:       "poll-1",
		GroupID:  "group-1",
		Question: "Pick a venue",
		Options: []PollOption{
			{ID: "one", Label: "Berlin"},
			{ID: "two", Label: "London"},
		},
		MinSelections: 1,
		MaxSelections: 1,
		CreatedBy:     "user-1",
	}
	if !poll.IsValid() {
		t.Fatal("expected poll with valid configuration to be valid")
	}

	if !poll.Options[0].IsValid() {
		t.Fatal("expected poll option with id and label to be valid")
	}
}
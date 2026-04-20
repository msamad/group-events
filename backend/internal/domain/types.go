package domain

import "time"

type UserID string
type GroupID string
type EventID string
type PollID string
type AckID string

type Role string

const (
	RoleOwner     Role = "owner"
	RoleAdmin     Role = "admin"
	RoleOrganizer Role = "organizer"
	RoleMember    Role = "member"
	RoleViewer    Role = "viewer"
)

func (r Role) CanManageGroup() bool {
	return r == RoleOwner || r == RoleAdmin
}

func (r Role) CanManageMembership() bool {
	return r.CanManageGroup()
}

func (r Role) CanCreateEvents() bool {
	return r.CanManageGroup() || r == RoleOrganizer
}

func (r Role) CanRespond() bool {
	return r == RoleOwner || r == RoleAdmin || r == RoleOrganizer || r == RoleMember
}

// Group is the top-level collaboration boundary.
// Validation rule: ID, slug, and name must all be non-empty.
type Group struct {
	ID          GroupID    `json:"id"`
	Slug        string     `json:"slug"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Archived    bool       `json:"archived"`
	CreatedAt   time.Time  `json:"createdAt"`
	ArchivedAt  *time.Time `json:"archivedAt,omitempty"`
}

func (g Group) IsValid() bool {
	return g.ID != "" && g.Slug != "" && g.Name != ""
}

// Member represents an identity that can join one or more groups.
// Validation rule: ID and display name must be non-empty.
type Member struct {
	ID          UserID `json:"id"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarUrl,omitempty"`
}

func (m Member) IsValid() bool {
	return m.ID != "" && m.DisplayName != ""
}

type Membership struct {
	GroupID  GroupID   `json:"groupId"`
	UserID   UserID    `json:"userId"`
	Role     Role      `json:"role"`
	JoinedAt time.Time `json:"joinedAt"`
}

func (m Membership) IsValid() bool {
	return m.GroupID != "" && m.UserID != "" && m.Role != ""
}

func (m Membership) CanCreateEvents() bool {
	return m.Role.CanCreateEvents()
}

func (m Membership) CanRespond() bool {
	return m.Role.CanRespond()
}

type Event struct {
	ID          EventID   `json:"id"`
	GroupID     GroupID   `json:"groupId"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Location    string    `json:"location,omitempty"`
	StartsAt    time.Time `json:"startsAt"`
	EndsAt      time.Time `json:"endsAt"`
	CreatedBy   UserID    `json:"createdBy"`
	Cancelled   bool      `json:"cancelled"`
}

func (e Event) IsValid() bool {
	if e.ID == "" || e.GroupID == "" || e.Title == "" || e.CreatedBy == "" {
		return false
	}

	return e.HasValidSchedule()
}

func (e Event) HasValidSchedule() bool {
	if e.StartsAt.IsZero() || e.EndsAt.IsZero() {
		return false
	}

	return e.EndsAt.After(e.StartsAt)
}

type ParticipationStatus string

const (
	ParticipationInvited  ParticipationStatus = "invited"
	ParticipationGoing    ParticipationStatus = "going"
	ParticipationMaybe    ParticipationStatus = "maybe"
	ParticipationDeclined ParticipationStatus = "declined"
)

type Participation struct {
	EventID     EventID             `json:"eventId"`
	UserID      UserID              `json:"userId"`
	Status      ParticipationStatus `json:"status"`
	RespondedAt *time.Time          `json:"respondedAt,omitempty"`
}

type Poll struct {
	ID             PollID        `json:"id"`
	GroupID        GroupID       `json:"groupId"`
	EventID        EventID       `json:"eventId,omitempty"`
	Question       string        `json:"question"`
	Options        []PollOption  `json:"options"`
	MinSelections  int           `json:"minSelections"`
	MaxSelections  int           `json:"maxSelections"`
	ClosesAt       *time.Time    `json:"closesAt,omitempty"`
	CreatedBy      UserID        `json:"createdBy"`
	AllowsRevoting bool          `json:"allowsRevoting"`
}

func (p Poll) IsValid() bool {
	if p.ID == "" || p.GroupID == "" || p.Question == "" || p.CreatedBy == "" {
		return false
	}

	return p.HasValidConfiguration()
}

func (p Poll) HasValidConfiguration() bool {
	if len(p.Options) == 0 {
		return false
	}

	if p.MinSelections < 0 || p.MaxSelections < 1 {
		return false
	}

	if p.MinSelections > p.MaxSelections {
		return false
	}

	return p.MaxSelections <= len(p.Options)
}

type PollOption struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

func (o PollOption) IsValid() bool {
	return o.ID != "" && o.Label != ""
}

type Vote struct {
	PollID     PollID     `json:"pollId"`
	OptionID   string     `json:"optionId"`
	UserID     UserID     `json:"userId"`
	SubmittedAt time.Time `json:"submittedAt"`
}

type AcknowledgementKind string

const (
	AcknowledgementAnnouncement AcknowledgementKind = "announcement"
	AcknowledgementReminder     AcknowledgementKind = "reminder"
	AcknowledgementDecision     AcknowledgementKind = "decision"
)

type Acknowledgement struct {
	ID             AckID               `json:"id"`
	GroupID        GroupID             `json:"groupId"`
	EventID        EventID             `json:"eventId,omitempty"`
	UserID         UserID              `json:"userId"`
	Kind           AcknowledgementKind `json:"kind"`
	Message        string              `json:"message"`
	SeenAt         *time.Time          `json:"seenAt,omitempty"`
	AcknowledgedAt *time.Time          `json:"acknowledgedAt,omitempty"`
}
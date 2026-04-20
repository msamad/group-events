package httpapi

import (
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/msamad/group-events/backend/internal/domain"
)

var errGroupNotFound = errors.New("group not found")

type groupStore struct {
	mu     sync.RWMutex
	groups map[domain.GroupID]domain.Group
}

func newGroupStore() *groupStore {
	return &groupStore{groups: make(map[domain.GroupID]domain.Group)}
}

func (s *groupStore) list(limit, offset int) []domain.Group {
	s.mu.RLock()
	defer s.mu.RUnlock()

	items := make([]domain.Group, 0, len(s.groups))
	for _, group := range s.groups {
		items = append(items, group)
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].CreatedAt.Equal(items[j].CreatedAt) {
			return items[i].ID < items[j].ID
		}
		return items[i].CreatedAt.Before(items[j].CreatedAt)
	})

	if offset >= len(items) {
		return []domain.Group{}
	}

	end := offset + limit
	if end > len(items) {
		end = len(items)
	}

	result := make([]domain.Group, end-offset)
	copy(result, items[offset:end])
	return result
}

func (s *groupStore) create(name, slug, description string) domain.Group {
	now := time.Now().UTC()
	group := domain.Group{
		ID:          domain.GroupID(uuid.NewString()),
		Name:        name,
		Slug:        slug,
		Description: description,
		CreatedAt:   now,
		Archived:    false,
	}

	s.mu.Lock()
	s.groups[group.ID] = group
	s.mu.Unlock()

	return group
}

func (s *groupStore) get(id domain.GroupID) (domain.Group, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	group, ok := s.groups[id]
	if !ok {
		return domain.Group{}, errGroupNotFound
	}

	return group, nil
}

func (s *groupStore) update(id domain.GroupID, name, slug, description string) (domain.Group, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	group, ok := s.groups[id]
	if !ok {
		return domain.Group{}, errGroupNotFound
	}

	group.Name = name
	group.Slug = slug
	group.Description = description
	s.groups[id] = group

	return group, nil
}

func (s *groupStore) delete(id domain.GroupID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.groups[id]; !ok {
		return errGroupNotFound
	}

	delete(s.groups, id)
	return nil
}

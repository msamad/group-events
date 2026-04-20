package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/msamad/group-events/backend/internal/domain"
	"github.com/msamad/group-events/backend/internal/sdui"
)

type groupsHandler struct {
	store *groupStore
}

type groupPayload struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description,omitempty"`
}

func newGroupsHandler(store *groupStore) groupsHandler {
	return groupsHandler{store: store}
}

func (h groupsHandler) createGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "only POST is supported")
		return
	}

	var payload groupPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "request body must be valid JSON")
		return
	}

	if strings.TrimSpace(payload.Name) == "" || strings.TrimSpace(payload.Slug) == "" {
		writeError(w, http.StatusBadRequest, "invalid_input", "name and slug are required")
		return
	}

	group := h.store.create(payload.Name, payload.Slug, payload.Description)
	writeJSON(w, http.StatusCreated, Response{Data: group, UI: buildGroupUIDescriptor(roleFromRequest(r), true)})
}

func (h groupsHandler) listGroups(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "only GET is supported")
		return
	}

	limit, err := parsePositiveInt(r.URL.Query().Get("limit"), 20)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_query", "limit must be a positive integer")
		return
	}

	offset, err := parseNonNegativeInt(r.URL.Query().Get("offset"), 0)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_query", "offset must be a non-negative integer")
		return
	}

	groups := h.store.list(limit, offset)
	writeJSON(w, http.StatusOK, Response{Data: groups, UI: buildGroupUIDescriptor(roleFromRequest(r), false)})
}

func (h groupsHandler) getGroup(w http.ResponseWriter, r *http.Request, id domain.GroupID) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "only GET is supported")
		return
	}

	group, err := h.store.get(id)
	if err != nil {
		if errors.Is(err, errGroupNotFound) {
			writeError(w, http.StatusNotFound, "not_found", "group not found")
			return
		}

		writeError(w, http.StatusInternalServerError, "internal_error", "unable to read group")
		return
	}

	writeJSON(w, http.StatusOK, Response{Data: group, UI: buildGroupUIDescriptor(roleFromRequest(r), false)})
}

func (h groupsHandler) updateGroup(w http.ResponseWriter, r *http.Request, id domain.GroupID) {
	if r.Method != http.MethodPut {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "only PUT is supported")
		return
	}

	var payload groupPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "request body must be valid JSON")
		return
	}

	if strings.TrimSpace(payload.Name) == "" || strings.TrimSpace(payload.Slug) == "" {
		writeError(w, http.StatusBadRequest, "invalid_input", "name and slug are required")
		return
	}

	group, err := h.store.update(id, payload.Name, payload.Slug, payload.Description)
	if err != nil {
		if errors.Is(err, errGroupNotFound) {
			writeError(w, http.StatusNotFound, "not_found", "group not found")
			return
		}

		writeError(w, http.StatusInternalServerError, "internal_error", "unable to update group")
		return
	}

	writeJSON(w, http.StatusOK, Response{Data: group, UI: buildGroupUIDescriptor(roleFromRequest(r), true)})
}

func (h groupsHandler) deleteGroup(w http.ResponseWriter, r *http.Request, id domain.GroupID) {
	if r.Method != http.MethodDelete {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "only DELETE is supported")
		return
	}

	if err := h.store.delete(id); err != nil {
		if errors.Is(err, errGroupNotFound) {
			writeError(w, http.StatusNotFound, "not_found", "group not found")
			return
		}

		writeError(w, http.StatusInternalServerError, "internal_error", "unable to delete group")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func roleFromRequest(r *http.Request) domain.Role {
	raw := strings.ToLower(strings.TrimSpace(r.Header.Get("X-Group-Role")))
	switch raw {
	case string(domain.RoleOwner):
		return domain.RoleOwner
	case string(domain.RoleAdmin):
		return domain.RoleAdmin
	case string(domain.RoleOrganizer), "role-assigned", "role_assigned":
		return domain.RoleOrganizer
	case string(domain.RoleMember):
		return domain.RoleMember
	default:
		return domain.RoleViewer
	}
}

func buildGroupUIDescriptor(role domain.Role, mutating bool) sdui.UIDescriptor {
	actions := []sdui.UIAction{
		{ID: "create_event", Label: "Create event", Type: http.MethodPost, Endpoint: "/api/v1/events", Visible: role.CanCreateEvents()},
		{ID: "create_poll", Label: "Create poll", Type: http.MethodPost, Endpoint: "/api/v1/polls", Visible: role.CanCreateEvents()},
		{ID: "manage_members", Label: "Manage members", Type: http.MethodPatch, Endpoint: "/api/v1/groups/{id}/members", Visible: role.CanManageMembership()},
		{ID: "join_event", Label: "Join event", Type: http.MethodPost, Endpoint: "/api/v1/events/{id}/responses", Visible: role.CanRespond()},
		{ID: "vote_poll", Label: "Vote", Type: http.MethodPost, Endpoint: "/api/v1/polls/{id}/votes", Visible: role.CanRespond()},
		{ID: "acknowledge", Label: "Acknowledge", Type: http.MethodPost, Endpoint: "/api/v1/acks", Visible: true},
		{ID: "react", Label: "React", Type: http.MethodPost, Endpoint: "/api/v1/reactions", Visible: true},
	}

	visibleActions := make([]sdui.UIAction, 0, len(actions))
	for _, action := range actions {
		if action.Visible {
			visibleActions = append(visibleActions, action)
		}
	}

	screen := "groups:list"
	if mutating {
		screen = "groups:mutate"
	}

	return sdui.UIDescriptor{
		Screen: screen,
		Components: []sdui.UIComponent{
			{Type: "list", ID: "groups", Visible: true, Data: map[string]interface{}{"emptyLabel": "No groups yet"}},
		},
		Actions: visibleActions,
		Navigation: sdui.UINavigation{
			Route: "/groups",
		},
	}
}

func parsePositiveInt(raw string, fallback int) (int, error) {
	if strings.TrimSpace(raw) == "" {
		return fallback, nil
	}

	value, err := strconv.Atoi(raw)
	if err != nil || value < 1 {
		return 0, errors.New("invalid positive integer")
	}

	return value, nil
}

func parseNonNegativeInt(raw string, fallback int) (int, error) {
	if strings.TrimSpace(raw) == "" {
		return fallback, nil
	}

	value, err := strconv.Atoi(raw)
	if err != nil || value < 0 {
		return 0, errors.New("invalid non-negative integer")
	}

	return value, nil
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, ErrorResponse{Error: APIError{Code: code, Message: message}})
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

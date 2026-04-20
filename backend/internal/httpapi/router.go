package httpapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/msamad/group-events/backend/internal/domain"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	groups := newGroupsHandler(newGroupStore())

	mux.HandleFunc("/health", health)
	mux.HandleFunc("/api/v1/sdui/schema", sduiSchemaHandler)
	mux.HandleFunc("/api/v1/groups", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			groups.listGroups(w, r)
		case http.MethodPost:
			groups.createGroup(w, r)
		default:
			writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "unsupported method")
		}
	})
	mux.HandleFunc("/api/v1/groups/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/v1/groups/")
		if id == "" || strings.Contains(id, "/") {
			writeError(w, http.StatusNotFound, "not_found", "group not found")
			return
		}

		groupID := domain.GroupID(id)
		switch r.Method {
		case http.MethodGet:
			groups.getGroup(w, r, groupID)
		case http.MethodPut:
			groups.updateGroup(w, r, groupID)
		case http.MethodDelete:
			groups.deleteGroup(w, r, groupID)
		default:
			writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "unsupported method")
		}
	})
	return mux
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

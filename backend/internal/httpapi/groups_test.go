package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type groupsEnvelope struct {
	Data json.RawMessage `json:"data"`
	UI   struct {
		Screen  string `json:"screen"`
		Actions []struct {
			ID      string `json:"id"`
			Visible bool   `json:"visible"`
		} `json:"actions"`
	} `json:"ui"`
}

func TestGroupsCRUDLifecycle(t *testing.T) {
	t.Parallel()

	router := NewRouter()

	createBody := []byte(`{"name":"Platform","slug":"platform","description":"Core team"}`)
	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/groups", bytes.NewReader(createBody))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("X-Group-Role", "owner")
	createRes := httptest.NewRecorder()
	router.ServeHTTP(createRes, createReq)

	if createRes.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, createRes.Code)
	}

	var created groupsEnvelope
	if err := json.Unmarshal(createRes.Body.Bytes(), &created); err != nil {
		t.Fatalf("unmarshal create response: %v", err)
	}

	if created.UI.Screen == "" {
		t.Fatal("expected ui.screen in create response")
	}

	var createdGroup struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(created.Data, &createdGroup); err != nil {
		t.Fatalf("unmarshal create group payload: %v", err)
	}

	listReq := httptest.NewRequest(http.MethodGet, "/api/v1/groups?limit=10&offset=0", nil)
	listReq.Header.Set("X-Group-Role", "member")
	listRes := httptest.NewRecorder()
	router.ServeHTTP(listRes, listReq)

	if listRes.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, listRes.Code)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/api/v1/groups/"+createdGroup.ID, nil)
	getReq.Header.Set("X-Group-Role", "member")
	getRes := httptest.NewRecorder()
	router.ServeHTTP(getRes, getReq)

	if getRes.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, getRes.Code)
	}

	updateBody := []byte(`{"name":"Platform Team","slug":"platform-team","description":"Updated"}`)
	updateReq := httptest.NewRequest(http.MethodPut, "/api/v1/groups/"+createdGroup.ID, bytes.NewReader(updateBody))
	updateReq.Header.Set("Content-Type", "application/json")
	updateReq.Header.Set("X-Group-Role", "admin")
	updateRes := httptest.NewRecorder()
	router.ServeHTTP(updateRes, updateReq)

	if updateRes.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, updateRes.Code)
	}

	deleteReq := httptest.NewRequest(http.MethodDelete, "/api/v1/groups/"+createdGroup.ID, nil)
	deleteRes := httptest.NewRecorder()
	router.ServeHTTP(deleteRes, deleteReq)

	if deleteRes.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, deleteRes.Code)
	}
}

func TestGroupsHandlersErrors(t *testing.T) {
	t.Parallel()

	router := NewRouter()

	tests := []struct {
		name       string
		method     string
		path       string
		body       []byte
		wantStatus int
	}{
		{name: "invalid create payload", method: http.MethodPost, path: "/api/v1/groups", body: []byte(`{"name":""}`), wantStatus: http.StatusBadRequest},
		{name: "invalid limit", method: http.MethodGet, path: "/api/v1/groups?limit=-1", wantStatus: http.StatusBadRequest},
		{name: "not found group", method: http.MethodGet, path: "/api/v1/groups/missing", wantStatus: http.StatusNotFound},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(test.method, test.path, bytes.NewReader(test.body))
			req.Header.Set("Content-Type", "application/json")
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			if res.Code != test.wantStatus {
				t.Fatalf("expected status %d, got %d", test.wantStatus, res.Code)
			}

			if test.wantStatus != http.StatusNoContent {
				var payload ErrorResponse
				if err := json.Unmarshal(res.Body.Bytes(), &payload); err != nil {
					t.Fatalf("expected structured json error: %v", err)
				}
				if payload.Error.Code == "" {
					t.Fatal("expected error code in payload")
				}
			}
		})
	}
}

func TestRoleAwareActionSets(t *testing.T) {
	t.Parallel()

	router := NewRouter()

	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/groups", bytes.NewReader([]byte(`{"name":"Ops","slug":"ops"}`)))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("X-Group-Role", "owner")
	createRes := httptest.NewRecorder()
	router.ServeHTTP(createRes, createReq)
	if createRes.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, createRes.Code)
	}

	tests := []struct {
		name    string
		role    string
		allowed map[string]bool
	}{
		{name: "leader", role: "owner", allowed: map[string]bool{"create_event": true, "create_poll": true, "manage_members": true, "join_event": true, "vote_poll": true, "acknowledge": true, "react": true}},
		{name: "role assigned", role: "role_assigned", allowed: map[string]bool{"create_event": true, "create_poll": true, "manage_members": false, "join_event": true, "vote_poll": true, "acknowledge": true, "react": true}},
		{name: "member", role: "member", allowed: map[string]bool{"create_event": false, "create_poll": false, "manage_members": false, "join_event": true, "vote_poll": true, "acknowledge": true, "react": true}},
		{name: "read only", role: "viewer", allowed: map[string]bool{"create_event": false, "create_poll": false, "manage_members": false, "join_event": false, "vote_poll": false, "acknowledge": true, "react": true}},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			req := httptest.NewRequest(http.MethodGet, "/api/v1/groups", nil)
			req.Header.Set("X-Group-Role", test.role)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			if res.Code != http.StatusOK {
				t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
			}

			var payload groupsEnvelope
			if err := json.Unmarshal(res.Body.Bytes(), &payload); err != nil {
				t.Fatalf("unmarshal list response: %v", err)
			}

			seen := map[string]bool{}
			for _, action := range payload.UI.Actions {
				seen[action.ID] = true
			}

			for id, allowed := range test.allowed {
				_, ok := seen[id]
				if ok != allowed {
					t.Fatalf("role %q action %q mismatch: got visible=%t want=%t", test.role, id, ok, allowed)
				}
			}
		})
	}
}

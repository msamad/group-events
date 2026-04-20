package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseIntHelpers(t *testing.T) {
	t.Parallel()

	if value, err := parsePositiveInt("5", 20); err != nil || value != 5 {
		t.Fatalf("expected positive parse to return 5, got value=%d err=%v", value, err)
	}

	if _, err := parsePositiveInt("0", 20); err == nil {
		t.Fatal("expected parsePositiveInt to reject zero")
	}

	if value, err := parsePositiveInt("", 20); err != nil || value != 20 {
		t.Fatalf("expected fallback value for empty positive input, got value=%d err=%v", value, err)
	}

	if value, err := parseNonNegativeInt("2", 0); err != nil || value != 2 {
		t.Fatalf("expected non-negative parse to return 2, got value=%d err=%v", value, err)
	}

	if _, err := parseNonNegativeInt("-1", 0); err == nil {
		t.Fatal("expected parseNonNegativeInt to reject negative values")
	}

	if value, err := parseNonNegativeInt("", 7); err != nil || value != 7 {
		t.Fatalf("expected fallback value for empty non-negative input, got value=%d err=%v", value, err)
	}

}

func TestMethodGuardsAndSchemaEndpoint(t *testing.T) {
	t.Parallel()

	router := NewRouter()

	tests := []struct {
		method string
		path   string
		status int
	}{
		{method: http.MethodPost, path: "/health", status: http.StatusMethodNotAllowed},
		{method: http.MethodPost, path: "/api/v1/sdui/schema", status: http.StatusMethodNotAllowed},
		{method: http.MethodPatch, path: "/api/v1/groups", status: http.StatusMethodNotAllowed},
		{method: http.MethodPatch, path: "/api/v1/groups/invalid", status: http.StatusMethodNotAllowed},
		{method: http.MethodGet, path: "/api/v1/groups/invalid/path", status: http.StatusNotFound},
		{method: http.MethodGet, path: "/api/v1/groups?offset=-2", status: http.StatusBadRequest},
	}

	for _, test := range tests {
		test := test
		t.Run(test.method+" "+test.path, func(t *testing.T) {
			res := httptest.NewRecorder()
			req := httptest.NewRequest(test.method, test.path, nil)
			router.ServeHTTP(res, req)

			if res.Code != test.status {
				t.Fatalf("expected status %d, got %d", test.status, res.Code)
			}
		})
	}
}

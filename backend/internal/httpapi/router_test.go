package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouterHealth(t *testing.T) {
	t.Parallel()

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	NewRouter().ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	const expected = "{\"status\":\"ok\"}\n"
	if recorder.Body.String() != expected {
		t.Fatalf("expected body %q, got %q", expected, recorder.Body.String())
	}

	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Fatalf("expected content type %q, got %q", "application/json", contentType)
	}
}
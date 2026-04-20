package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

func TestSduiContractForGroupEndpoints(t *testing.T) {
	t.Parallel()

	router := NewRouter()
	schemaLoader := schemaLoaderFromRepo(t)

	createReq := httptest.NewRequest(http.MethodPost, "/api/v1/groups", bytes.NewReader([]byte(`{"name":"Contract","slug":"contract"}`)))
	createReq.Header.Set("Content-Type", "application/json")
	createReq.Header.Set("X-Group-Role", "owner")
	createRes := httptest.NewRecorder()
	router.ServeHTTP(createRes, createReq)
	if createRes.Code != http.StatusCreated {
		t.Fatalf("create group expected %d, got %d", http.StatusCreated, createRes.Code)
	}

	var createPayload struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
		UI map[string]interface{} `json:"ui"`
	}
	if err := json.Unmarshal(createRes.Body.Bytes(), &createPayload); err != nil {
		t.Fatalf("decode create payload: %v", err)
	}
	validateSchema(t, schemaLoader, createPayload.UI)

	paths := []string{
		"/api/v1/groups",
		"/api/v1/groups/" + createPayload.Data.ID,
	}

	for _, path := range paths {
		path := path
		t.Run(path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, path, nil)
			req.Header.Set("X-Group-Role", "member")
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			if res.Code != http.StatusOK {
				t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
			}

			var payload struct {
				UI map[string]interface{} `json:"ui"`
			}
			if err := json.Unmarshal(res.Body.Bytes(), &payload); err != nil {
				t.Fatalf("decode payload: %v", err)
			}

			validateSchema(t, schemaLoader, payload.UI)
		})
	}
}

func TestSduiSchemaEndpoint(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/sdui/schema", nil)
	res := httptest.NewRecorder()
	NewRouter().ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}

	if contentType := res.Header().Get("Content-Type"); contentType != "application/schema+json" {
		t.Fatalf("unexpected content type %q", contentType)
	}
}

func schemaLoaderFromRepo(t *testing.T) gojsonschema.JSONLoader {
	t.Helper()

	_, currentFile, _, _ := runtime.Caller(0)
	schemaPath := filepath.Clean(filepath.Join(filepath.Dir(currentFile), "..", "..", "..", "tests", "contracts", "sdui-descriptor.schema.json"))
	return gojsonschema.NewReferenceLoader("file://" + filepath.ToSlash(schemaPath))
}

func validateSchema(t *testing.T, schemaLoader gojsonschema.JSONLoader, payload map[string]interface{}) {
	t.Helper()

	result, err := gojsonschema.Validate(schemaLoader, gojsonschema.NewGoLoader(payload))
	if err != nil {
		t.Fatalf("validate schema: %v", err)
	}

	if !result.Valid() {
		t.Fatalf("schema validation failed: %v", result.Errors())
	}
}

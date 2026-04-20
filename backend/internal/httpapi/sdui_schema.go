package httpapi

import "net/http"

const sduiSchema = `{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://group-events/sdui-descriptor.schema.json",
  "title": "UIDescriptor",
  "type": "object",
  "required": ["screen", "components", "actions"],
  "properties": {
    "screen": {"type": "string"},
    "components": {
      "type": "array",
      "items": {
        "type": "object",
        "required": ["type", "visible"],
        "properties": {
          "id": {"type": "string"},
          "type": {"type": "string"},
          "visible": {"type": "boolean"},
          "data": {"type": "object"}
        }
      }
    },
    "actions": {
      "type": "array",
      "items": {
        "type": "object",
        "required": ["id", "label", "type", "visible"],
        "properties": {
          "id": {"type": "string"},
          "label": {"type": "string"},
          "type": {"type": "string"},
          "endpoint": {"type": "string"},
          "confirm": {"type": "boolean"},
          "visible": {"type": "boolean"}
        }
      }
    },
    "navigation": {
      "type": "object",
      "properties": {
        "route": {"type": "string"},
        "params": {"type": "object"}
      }
    }
  }
}`

func sduiSchemaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "only GET is supported")
		return
	}

	w.Header().Set("Content-Type", "application/schema+json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(sduiSchema))
}

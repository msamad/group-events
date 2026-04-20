package sdui

// UIDescriptor defines the server-driven UI payload attached to API responses.
// Validation rule: screen and all nested visible entries must be valid.
type UIDescriptor struct {
	Screen     string        `json:"screen"`
	Components []UIComponent `json:"components,omitempty"`
	Actions    []UIAction    `json:"actions,omitempty"`
	Navigation UINavigation  `json:"navigation,omitempty"`
}

func (d UIDescriptor) IsValid() bool {
	if d.Screen == "" {
		return false
	}

	for _, component := range d.Components {
		if !component.IsValid() {
			return false
		}
	}

	for _, action := range d.Actions {
		if !action.IsValid() {
			return false
		}
	}

	return true
}

type UINavigation struct {
	Route  string                 `json:"route,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
}

// UIComponent describes a renderable UI block.
// Validation rule: type must be non-empty.
type UIComponent struct {
	Type    string                 `json:"type"`
	ID      string                 `json:"id,omitempty"`
	Visible bool                   `json:"visible"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (c UIComponent) IsValid() bool {
	return c.Type != "" && c.Visible
}

// UIAction describes an interaction surface emitted by the backend.
// Validation rule: type must be non-empty.
type UIAction struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Type     string `json:"type"`
	Endpoint string `json:"endpoint,omitempty"`
	Confirm  bool   `json:"confirm"`
	Visible  bool   `json:"visible"`
}

func (a UIAction) IsValid() bool {
	return a.ID != "" && a.Label != "" && a.Type != "" && a.Visible
}

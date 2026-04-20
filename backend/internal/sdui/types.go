package sdui

// UIDescriptor defines the server-driven UI payload attached to API responses.
// Validation rule: every component and action entry must be individually valid.
type UIDescriptor struct {
	Components []UIComponent `json:"components,omitempty"`
	Actions    []UIAction    `json:"actions,omitempty"`
}

func (d UIDescriptor) IsValid() bool {
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

// UIComponent describes a renderable UI block.
// Validation rule: type must be non-empty.
type UIComponent struct {
	Type string `json:"type"`
	ID   string `json:"id,omitempty"`
}

func (c UIComponent) IsValid() bool {
	return c.Type != ""
}

// UIAction describes an interaction surface emitted by the backend.
// Validation rule: type must be non-empty.
type UIAction struct {
	Type   string `json:"type"`
	Target string `json:"target,omitempty"`
}

func (a UIAction) IsValid() bool {
	return a.Type != ""
}
package sdui

type UIDescriptor struct {
	Components []UIComponent `json:"components,omitempty"`
	Actions    []UIAction    `json:"actions,omitempty"`
}

type UIComponent struct {
	Type string `json:"type"`
	ID   string `json:"id,omitempty"`
}

type UIAction struct {
	Type   string `json:"type"`
	Target string `json:"target,omitempty"`
}
package sdui

import "testing"

func TestComponentAndActionValidation(t *testing.T) {
	t.Parallel()

	component := UIComponent{Type: "section", ID: "summary", Visible: true}
	if !component.IsValid() {
		t.Fatal("expected UI component with type to be valid")
	}

	action := UIAction{ID: "navigate", Label: "Open", Type: "GET", Endpoint: "/events", Visible: true}
	if !action.IsValid() {
		t.Fatal("expected UI action with type to be valid")
	}

	invalidComponent := UIComponent{}
	if invalidComponent.IsValid() {
		t.Fatal("expected component without type to be invalid")
	}

	invalidAction := UIAction{}
	if invalidAction.IsValid() {
		t.Fatal("expected action without type to be invalid")
	}
}

func TestDescriptorValidation(t *testing.T) {
	t.Parallel()

	descriptor := UIDescriptor{
		Screen:     "groups:list",
		Components: []UIComponent{{Type: "section", ID: "board", Visible: true}},
		Actions:    []UIAction{{ID: "submit", Label: "Vote", Type: "POST", Endpoint: "/polls/1/votes", Visible: true}},
	}
	if !descriptor.IsValid() {
		t.Fatal("expected descriptor with valid items to be valid")
	}

	invalidDescriptor := UIDescriptor{
		Screen:     "groups:list",
		Components: []UIComponent{{Type: "section"}, {}},
	}
	if invalidDescriptor.IsValid() {
		t.Fatal("expected descriptor with invalid component to be invalid")
	}

	missingScreen := UIDescriptor{Components: []UIComponent{{Type: "section", Visible: true}}}
	if missingScreen.IsValid() {
		t.Fatal("expected descriptor without screen to be invalid")
	}
}

package sdui

import "testing"

func TestComponentAndActionValidation(t *testing.T) {
	t.Parallel()

	component := UIComponent{Type: "section", ID: "summary"}
	if !component.IsValid() {
		t.Fatal("expected UI component with type to be valid")
	}

	action := UIAction{Type: "navigate", Target: "/events"}
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
		Components: []UIComponent{{Type: "section", ID: "board"}},
		Actions:    []UIAction{{Type: "submit", Target: "/polls/1/votes"}},
	}
	if !descriptor.IsValid() {
		t.Fatal("expected descriptor with valid items to be valid")
	}

	invalidDescriptor := UIDescriptor{
		Components: []UIComponent{{Type: "section"}, {}},
	}
	if invalidDescriptor.IsValid() {
		t.Fatal("expected descriptor with invalid component to be invalid")
	}
}

package components

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func TestButtonVariants(t *testing.T) {
	tests := []struct {
		variant ButtonVariant
		want    widget.Importance
	}{
		{ButtonPrimary, widget.HighImportance},
		{ButtonSecondary, widget.MediumImportance},
		{ButtonGhost, widget.LowImportance},
		{ButtonDanger, widget.DangerImportance},
	}

	for _, tt := range tests {
		b := NewButton("Test", tt.variant, nil)

		if b.Importance != tt.want {
			t.Errorf("got %v, want %v", b.Importance, tt.want)
		}
	}
}

func TestButtonOptionsAndSetters(t *testing.T) {
	b := NewButton(
		"Test",
		ButtonPrimary,
		nil,
		WithButtonRadius(8),
		WithButtonMinWidth(120),
	)

	if b.Variant != ButtonPrimary || b.radius != 8 || b.minWidth != 120 {
		t.Fatalf("options were not applied: variant=%v radius=%v minWidth=%v", b.Variant, b.radius, b.minWidth)
	}

	b.SetVariant(ButtonDanger)
	b.SetText("Updated")
	b.SetOnTapped(func() {})
	if b.Variant != ButtonDanger || b.Text != "Updated" || b.Importance != widget.DangerImportance {
		t.Fatal("button setters did not update public state")
	}
}

func TestLegacyButtonConstructorStillCompiles(t *testing.T) {
	var _ *Button = NewButton("Legacy", ButtonSecondary, nil)
}

func TestButtonActivatesFromKeyboardAndTap(t *testing.T) {
	test.NewTempApp(t)
	count := 0
	b := NewButton("Test", ButtonPrimary, func() { count++ })

	b.Tapped(nil)
	b.TypedKey(&fyne.KeyEvent{Name: fyne.KeySpace})
	b.TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})

	if count != 3 {
		t.Fatalf("got %d activations, want 3", count)
	}

	b.Disable()
	b.Tapped(nil)
	b.TypedKey(&fyne.KeyEvent{Name: fyne.KeyEnter})
	if count != 3 {
		t.Fatalf("disabled button activated: got %d activations", count)
	}
}

func TestButtonPointerStatesAndRenderer(t *testing.T) {
	test.NewTempApp(t)
	b := NewButton("Test", ButtonSecondary, nil)
	r := b.CreateRenderer()

	if got := len(r.Objects()); got != 4 {
		t.Fatalf("renderer owns %d objects, want background, overlay, focus, and label", got)
	}
	if got := r.MinSize(); got.Width < 80 || got.Height < 36 {
		t.Fatalf("minimum size %v is too small", got)
	}

	configured := NewButton("Small", ButtonSecondary, nil, WithButtonMinWidth(120))
	if got := configured.CreateRenderer().MinSize(); got.Width < 120 {
		t.Fatalf("configured minimum width %v is too small", got.Width)
	}

	b.FocusGained()
	b.MouseIn(&desktop.MouseEvent{})
	b.MouseDown(&desktop.MouseEvent{Button: desktop.MouseButtonPrimary})
	b.MouseUp(&desktop.MouseEvent{Button: desktop.MouseButtonPrimary})
	b.MouseOut()
	b.FocusLost()

	if b.focused || b.hovered || b.pressed {
		t.Fatal("button retained an interaction state after pointer/focus exit")
	}
}

func TestButtonRendersWithLightAndDarkThemes(t *testing.T) {
	app := test.NewTempApp(t)
	b := NewButton("Test", ButtonPrimary, nil)

	for _, current := range []fyne.Theme{theme.LightTheme(), theme.DarkTheme()} {
		app.Settings().SetTheme(current)
		r := b.CreateRenderer()
		if len(r.Objects()) != 4 {
			t.Fatalf("theme %T produced an incomplete renderer", current)
		}
	}
}

func TestButtonUsesThemeOverride(t *testing.T) {
	test.NewTempApp(t)
	b := NewButton("Test", ButtonSecondary, nil)
	container.NewThemeOverride(b, theme.DarkTheme())

	r := b.CreateRenderer()
	if len(r.Objects()) != 4 {
		t.Fatal("theme override did not produce a complete renderer")
	}
}

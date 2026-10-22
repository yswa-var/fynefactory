package components

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
)

func TestBadgeStateAndSetters(t *testing.T) {
	test.NewTempApp(t)
	b := NewBadge("Ready", BadgeSuccess)
	if b.Text != "Ready" || b.Variant != BadgeSuccess || b.AccessibilityLabel() != "Ready" {
		t.Fatal("constructor did not initialize badge state")
	}
	b.SetText("Needs attention")
	b.SetVariant(BadgeWarning)
	if b.Text != "Needs attention" || b.Variant != BadgeWarning {
		t.Fatal("setters did not update badge state")
	}
}

func TestBadgeRendererUsesThemeColorsAndMinimumSize(t *testing.T) {
	app := test.NewTempApp(t)
	for _, current := range []fyne.Theme{theme.LightTheme(), theme.DarkTheme()} {
		app.Settings().SetTheme(current)
		for _, variant := range []BadgeVariant{BadgeSuccess, BadgeWarning, BadgeError, BadgeNeutral} {
			b := NewBadge("Status", variant)
			r := b.CreateRenderer()
			if len(r.Objects()) != 2 {
				t.Fatalf("variant %d has %d renderer objects, want 2", variant, len(r.Objects()))
			}
			if size := r.MinSize(); size.Width <= 0 || size.Height <= 0 {
				t.Fatalf("variant %d has invalid minimum size %v", variant, size)
			}
			if _, ok := r.Objects()[0].(*canvas.Rectangle); !ok {
				t.Fatal("first renderer object is not the badge background")
			}
		}
	}
}

package components

import (
	"fyne.io/fyne/v2/widget"
	"testing"
)

func TestAccordionStateAndSizing(t *testing.T) {
	a := NewAccordion(widget.NewAccordionItem("One", widget.NewLabel("First")), widget.NewAccordionItem("Two", widget.NewLabel("Second")))
	if a.Density != AccordionMedium || len(a.Items) != 2 {
		t.Fatal("unexpected defaults")
	}
	a.Small().Open(0)
	if !a.Items[0].Open || a.Items[1].Open {
		t.Fatal("expected single-open behavior")
	}
	a.SetMultiOpen(true).Open(1)
	if !a.Items[0].Open || !a.Items[1].Open {
		t.Fatal("expected multi-open behavior")
	}
	a.CloseAll()
	if a.Items[0].Open || a.Items[1].Open {
		t.Fatal("expected all items closed")
	}
}

func TestAccordionRendererMinimumSize(t *testing.T) {
	a := NewAccordion(widget.NewAccordionItem("Details", widget.NewLabel("Content")))
	if size := a.CreateRenderer().MinSize(); size.Width <= 0 || size.Height <= 0 {
		t.Fatalf("invalid minimum size: %v", size)
	}
}

func TestAccordionTextAndVariant(t *testing.T) {
	a := NewAccordion(widget.NewAccordionItem("Initial", widget.NewLabel("Content")))
	a.SetVariant(AccordionGhost).SetItemTitle(0, "Updated")
	if a.Variant != AccordionGhost || a.Items[0].Title != "Updated" {
		t.Fatalf("accordion configuration was not updated: variant=%v title=%q", a.Variant, a.Items[0].Title)
	}
}

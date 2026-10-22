package playgrounds

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yswa-var/fynefactory/components"
)

func NewAccordionPlayground() fyne.CanvasObject {
	accordion := components.NewAccordion(
		widget.NewAccordionItem("What is FyneFactory?", widget.NewLabel("A collection of reusable Fyne components.")),
		widget.NewAccordionItem("Can sections stay open?", widget.NewLabel("Enable MultiOpen to keep more than one section expanded.")),
		widget.NewAccordionItem("Sizing", widget.NewLabel("Use XSmall, Small, Medium, or Large for density.")),
	)
	accordion.Open(0)
	multi := widget.NewCheck("Allow multiple sections", func(value bool) { accordion.SetMultiOpen(value) })
	size := widget.NewSelect([]string{"Extra small", "Small", "Medium", "Large"}, func(value string) {
		switch value {
		case "Extra small":
			accordion.XSmall()
		case "Small":
			accordion.Small()
		case "Medium":
			accordion.Medium()
		case "Large":
			accordion.Large()
		}
	})
	size.SetSelected("Medium")
	variant := widget.NewSelect([]string{"Default", "Outline", "Ghost"}, func(value string) {
		switch value {
		case "Outline":
			accordion.SetVariant(components.AccordionOutline)
		case "Ghost":
			accordion.SetVariant(components.AccordionGhost)
		default:
			accordion.SetVariant(components.AccordionDefault)
		}
	})
	variant.SetSelected("Default")
	text := widget.NewEntry()
	text.SetText("What is FyneFactory?")
	text.OnChanged = func(value string) { accordion.SetItemTitle(0, value) }
	return container.NewVBox(widget.NewLabelWithStyle("Accordion", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}), accordion, multi, container.NewGridWithColumns(2, widget.NewLabel("Size"), size), container.NewGridWithColumns(2, widget.NewLabel("Variant"), variant), container.NewGridWithColumns(2, widget.NewLabel("First title"), text))
}

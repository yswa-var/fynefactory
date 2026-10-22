package playgrounds

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yswa-var/fynefactory/components"
)

func NewSurfacePlayground() fyne.CanvasObject {
	title := widget.NewLabel("System Performance")
	title.Wrapping = fyne.TextWrapWord
	description := widget.NewLabel("All systems operational. This content wraps as the preview becomes narrower.")
	description.Wrapping = fyne.TextWrapWord
	surface := components.NewSurface(container.NewVBox(title, widget.NewLabelWithStyle("99.8%", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}), description))
	code := widget.NewMultiLineEntry()
	code.Wrapping = fyne.TextWrapWord
	code.TextStyle = fyne.TextStyle{Monospace: true}
	code.SetText("surface := components.NewSurface(\n    container.NewVBox(\n        widget.NewLabel(\"System Performance\"),\n        widget.NewLabel(\"99.8%\"),\n        widget.NewLabel(\"All systems operational\"),\n    ),\n)")
	return container.NewVBox(
		widget.NewLabelWithStyle("Surface", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Theme-aware content surface with responsive text."),
		container.NewPadded(surface),
		widget.NewLabelWithStyle("Go example · select and copy", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}), code,
	)
}

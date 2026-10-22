package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("FyneForge")

	title := widget.NewLabelWithStyle(
		"FyneForge",
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	content := container.NewVBox(
		widget.NewLabel("Components / Button"),
		widget.NewSeparator(),
		widget.NewButton("Primary Button", func() {}),
	)

	sidebar := container.NewVBox(
		title,
		widget.NewSeparator(),
		widget.NewButton("Buttons", func() {}),
		widget.NewButton("Cards", func() {}),
		widget.NewButton("Inputs", func() {}),
	)

	layout := container.NewBorder(
		nil, nil,
		container.NewPadded(sidebar),
		nil,
		container.NewPadded(content),
	)

	w.SetContent(layout)
	w.Resize(fyne.NewSize(900, 600))
	w.ShowAndRun()
}

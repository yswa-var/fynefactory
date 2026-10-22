package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yswa-var/fynefactory/theme"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.ForgeTheme{})
	w := a.NewWindow("FyneFactory · Showcase")
	entries, err := loadComponentEntries()
	if err != nil {
		panic(err)
	}

	title := widget.NewLabelWithStyle("FyneFactory", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	subtitle := widget.NewLabel("Showcase")
	subtitle.Wrapping = fyne.TextWrapWord

	w.SetContent(container.NewBorder(
		container.NewPadded(container.NewVBox(title, subtitle)),
		nil, nil, nil,
		newGallery(entries),
	))
	w.Resize(fyne.NewSize(1000, 700))
	w.ShowAndRun()
}

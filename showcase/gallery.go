package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func newGallery(entries []ComponentEntry) fyne.CanvasObject {
	filtered := append([]ComponentEntry(nil), entries...)
	preview := container.NewMax()
	show := func(entry ComponentEntry) {
		preview.Objects = []fyne.CanvasObject{entry.Build()}
		preview.Refresh()
	}
	if len(filtered) > 0 {
		show(filtered[0])
	}

	search := widget.NewEntry()
	search.SetPlaceHolder("Search components")
	list := widget.NewList(
		func() int { return len(filtered) },
		func() fyne.CanvasObject { return widget.NewLabel("Component") },
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(filtered[id].Name)
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		if id >= 0 && id < len(filtered) {
			show(filtered[id])
		}
	}
	search.OnChanged = func(query string) {
		query = strings.ToLower(strings.TrimSpace(query))
		filtered = filtered[:0]
		for _, entry := range entries {
			if query == "" || strings.Contains(strings.ToLower(entry.Name), query) || strings.Contains(strings.ToLower(entry.Category), query) || strings.Contains(strings.ToLower(entry.Description), query) {
				filtered = append(filtered, entry)
			}
		}
		list.Refresh()
		if len(filtered) > 0 {
			list.Select(0)
		}
	}

	title := widget.NewLabelWithStyle("COMPONENTS", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	sidebar := container.NewBorder(container.NewVBox(title, search), nil, nil, nil, container.NewVScroll(list))
	split := container.NewHSplit(container.NewPadded(sidebar), container.NewVScroll(container.NewPadded(preview)))
	split.SetOffset(0.25)
	return split
}

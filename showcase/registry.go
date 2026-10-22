package main

import (
	"fyne.io/fyne/v2"
	"github.com/yswa-var/fynefactory/components"
	"github.com/yswa-var/fynefactory/showcase/playgrounds"
)

type ComponentEntry struct {
	ID          string
	Name        string
	Category    string
	Description string
	Status      string
	Variants    []string
	Properties  []components.ManifestProperty
	Build       func() fyne.CanvasObject
}

var componentFactories = map[string]func() fyne.CanvasObject{
	"button":    playgrounds.NewButtonPlayground,
	"badge":     playgrounds.NewBadgePlayground,
	"surface":   playgrounds.NewSurfacePlayground,
	"accordion": playgrounds.NewAccordionPlayground,
}

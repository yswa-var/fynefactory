package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// AccordionSize controls the density of an Accordion.
type AccordionSize uint8

const (
	AccordionXSmall AccordionSize = iota
	AccordionSmall
	AccordionMedium
	AccordionLarge
)

// AccordionVariant controls the visual emphasis of accordion headers. The
// native Fyne accordion supplies the theme-aware disclosure treatment; these
// values are retained as part of the compact component API for callers and
// showcase tooling.
type AccordionVariant uint8

const (
	AccordionDefault AccordionVariant = iota
	AccordionOutline
	AccordionGhost
)

// Accordion is a shadcn-inspired disclosure list. It uses Fyne's native
// accordion implementation so content is laid out and collapsed correctly.
type Accordion struct {
	widget.BaseWidget
	Items     []*widget.AccordionItem
	MultiOpen bool
	Density   AccordionSize
	Variant   AccordionVariant

	accordion *widget.Accordion
}

var _ fyne.Widget = (*Accordion)(nil)

// NewAccordion creates an accordion with the supplied items. Medium is the
// default density and only one item is open at a time unless MultiOpen is set.
func NewAccordion(items ...*widget.AccordionItem) *Accordion {
	a := &Accordion{Items: items, Density: AccordionMedium, Variant: AccordionDefault}
	a.accordion = widget.NewAccordion(items...)
	a.ExtendBaseWidget(a)
	return a
}

func (a *Accordion) CreateRenderer() fyne.WidgetRenderer {
	a.sync()
	return widget.NewSimpleRenderer(a.accordion)
}

func (a *Accordion) sync() {
	if a.accordion == nil {
		a.accordion = widget.NewAccordion(a.Items...)
	}
	a.accordion.Items = a.Items
	a.accordion.MultiOpen = a.MultiOpen
}

// Append adds an item and refreshes the accordion.
func (a *Accordion) Append(item *widget.AccordionItem) {
	a.Items = append(a.Items, item)
	a.sync()
	a.Refresh()
}

// Open expands one item, collapsing others unless MultiOpen is enabled.
func (a *Accordion) Open(index int) { a.sync(); a.accordion.Open(index); a.Refresh() }

// Close collapses one item.
func (a *Accordion) Close(index int) { a.sync(); a.accordion.Close(index); a.Refresh() }

// CloseAll collapses every item.
func (a *Accordion) CloseAll() { a.sync(); a.accordion.CloseAll(); a.Refresh() }

// SetMultiOpen changes whether multiple sections may remain open.
func (a *Accordion) SetMultiOpen(multi bool) *Accordion {
	a.MultiOpen = multi
	a.sync()
	a.Refresh()
	return a
}

// SetSize changes the requested density. The native Fyne theme applies the
// final platform-specific metrics; this value is retained for future theme
// overrides and for consumers configuring a consistent component API.
func (a *Accordion) SetSize(size AccordionSize) {
	if a.Density != size {
		a.Density = size
		a.Refresh()
	}
}

// SetVariant changes the header presentation variant and returns the
// accordion for fluent configuration. The underlying Fyne renderer remains
// theme-aware for every variant.
func (a *Accordion) SetVariant(variant AccordionVariant) *Accordion {
	a.Variant = variant
	a.Refresh()
	return a
}

// SetItemTitle changes the visible title of one section.
func (a *Accordion) SetItemTitle(index int, title string) *Accordion {
	if index >= 0 && index < len(a.Items) && a.Items[index] != nil {
		a.Items[index].Title = title
		a.sync()
		a.Refresh()
	}
	return a
}

func (a *Accordion) XSmall() *Accordion { a.SetSize(AccordionXSmall); return a }
func (a *Accordion) Small() *Accordion  { a.SetSize(AccordionSmall); return a }
func (a *Accordion) Medium() *Accordion { a.SetSize(AccordionMedium); return a }
func (a *Accordion) Large() *Accordion  { a.SetSize(AccordionLarge); return a }

package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// BadgeVariant controls the semantic color of a Badge.
type BadgeVariant uint8

const (
	BadgeSuccess BadgeVariant = iota
	BadgeWarning
	BadgeError
	BadgeNeutral
)

// Badge is a compact, theme-aware status label.
type Badge struct {
	widget.BaseWidget
	Text    string
	Variant BadgeVariant
}

var _ fyne.Widget = (*Badge)(nil)
var _ fyne.Accessible = (*Badge)(nil)

// NewBadge constructs a badge with the supplied text and semantic variant.
func NewBadge(text string, variant BadgeVariant) *Badge {
	b := &Badge{Text: text, Variant: variant}
	b.ExtendBaseWidget(b)
	return b
}

func (b *Badge) CreateRenderer() fyne.WidgetRenderer { return newBadgeRenderer(b) }

func (b *Badge) AccessibilityLabel() string { return b.Text }

func (b *Badge) AccessibilityRole() fyne.AccessibleRole { return fyne.AccessibleRoleText }

// SetText updates the visible badge text.
func (b *Badge) SetText(text string) {
	if b.Text == text {
		return
	}
	b.Text = text
	b.Refresh()
}

// SetVariant updates the badge's semantic color.
func (b *Badge) SetVariant(variant BadgeVariant) {
	if b.Variant == variant {
		return
	}
	b.Variant = variant
	b.Refresh()
}

type badgeRenderer struct {
	badge      *Badge
	background *canvas.Rectangle
	label      *canvas.Text
	objects    []fyne.CanvasObject
}

func newBadgeRenderer(badge *Badge) *badgeRenderer {
	r := &badgeRenderer{
		badge:      badge,
		background: canvas.NewRectangle(nil),
		label:      canvas.NewText(badge.Text, nil),
	}
	r.background.CornerRadius = theme.SizeForWidget(theme.SizeNameButtonRadius, badge)
	r.label.Alignment = fyne.TextAlignCenter
	r.objects = []fyne.CanvasObject{r.background, r.label}
	r.Refresh()
	return r
}

func (r *badgeRenderer) Destroy()                     {}
func (r *badgeRenderer) Objects() []fyne.CanvasObject { return r.objects }

func (r *badgeRenderer) MinSize() fyne.Size {
	padding := theme.SizeForWidget(theme.SizeNameInnerPadding, r.badge)
	size := r.label.MinSize()
	return fyne.NewSize(size.Width+padding*2, size.Height+padding)
}

func (r *badgeRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.label.Resize(size)
}

func (r *badgeRenderer) Refresh() {
	background, foreground := badgeColors(r.badge)
	r.background.FillColor = background
	r.label.Text = r.badge.Text
	r.label.Color = foreground
	r.label.Refresh()
	r.background.Refresh()
	r.Layout(r.badge.Size())
}

func badgeColors(badge *Badge) (background, foreground color.Color) {
	backgroundName := theme.ColorNameInputBackground
	foregroundName := theme.ColorNameForeground
	switch badge.Variant {
	case BadgeSuccess:
		backgroundName = theme.ColorNameSuccess
		foregroundName = theme.ColorNameForegroundOnSuccess
	case BadgeWarning:
		backgroundName = theme.ColorNameWarning
		foregroundName = theme.ColorNameForegroundOnWarning
	case BadgeError:
		backgroundName = theme.ColorNameError
		foregroundName = theme.ColorNameForegroundOnError
	}
	return theme.ColorForWidget(backgroundName, badge), theme.ColorForWidget(foregroundName, badge)
}

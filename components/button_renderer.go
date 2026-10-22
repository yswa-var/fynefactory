package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type buttonRenderer struct {
	button         *Button
	background     *canvas.Rectangle
	overlay        *canvas.Rectangle
	label          *canvas.Text
	focusIndicator *canvas.Rectangle
	objects        []fyne.CanvasObject
}

func newButtonRenderer(button *Button) *buttonRenderer {
	background := canvas.NewRectangle(color.Transparent)
	radius := button.radius
	if radius <= 0 {
		radius = theme.SizeForWidget(theme.SizeNameButtonRadius, button)
	}
	background.CornerRadius = radius
	overlay := canvas.NewRectangle(color.Transparent)
	overlay.CornerRadius = background.CornerRadius

	label := canvas.NewText(button.Text, color.Transparent)
	label.Alignment = fyne.TextAlignCenter

	focusIndicator := canvas.NewRectangle(color.Transparent)
	focusIndicator.StrokeWidth = 2

	objects := []fyne.CanvasObject{background, overlay, focusIndicator, label}
	r := &buttonRenderer{
		button:         button,
		background:     background,
		overlay:        overlay,
		label:          label,
		focusIndicator: focusIndicator,
		objects:        objects,
	}
	r.Refresh()
	return r
}

func (r *buttonRenderer) Destroy() {}

func (r *buttonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *buttonRenderer) MinSize() fyne.Size {
	padding := theme.SizeForWidget(theme.SizeNameInnerPadding, r.button)
	size := r.label.MinSize()
	return fyne.NewSize(
		fyne.Max(fyne.Max(size.Width+padding*4, 80), r.button.minWidth),
		fyne.Max(size.Height+padding*2, 36),
	)
}

func (r *buttonRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.overlay.Resize(size)
	r.focusIndicator.Resize(fyne.NewSize(
		fyne.Max(size.Width-2, 0),
		fyne.Max(size.Height-2, 0),
	))
	r.focusIndicator.Move(fyne.NewPos(1, 1))
	r.label.Resize(size)
}

func (r *buttonRenderer) Refresh() {
	r.label.Text = r.button.Text
	r.label.Color = theme.ColorForWidget(theme.ColorNameForeground, r.button)
	r.background.FillColor = theme.ColorForWidget(theme.ColorNameButton, r.button)
	radius := r.button.radius
	if radius <= 0 {
		radius = theme.SizeForWidget(theme.SizeNameButtonRadius, r.button)
	}
	r.background.CornerRadius = radius
	r.overlay.CornerRadius = r.background.CornerRadius
	r.focusIndicator.StrokeColor = theme.ColorForWidget(theme.ColorNameFocus, r.button)
	r.focusIndicator.CornerRadius = r.background.CornerRadius

	switch r.button.Importance {
	case widget.HighImportance:
		r.background.FillColor = theme.ColorForWidget(theme.ColorNamePrimary, r.button)
		r.label.Color = theme.ColorForWidget(theme.ColorNameForegroundOnPrimary, r.button)
	case widget.LowImportance:
		r.background.FillColor = color.Transparent
	case widget.DangerImportance:
		r.background.FillColor = theme.ColorForWidget(theme.ColorNameError, r.button)
		r.label.Color = theme.ColorForWidget(theme.ColorNameForegroundOnError, r.button)
	}

	if r.button.Disabled() {
		r.label.Color = theme.ColorForWidget(theme.ColorNameDisabled, r.button)
		r.background.FillColor = theme.ColorForWidget(theme.ColorNameDisabledButton, r.button)
	}
	r.overlay.FillColor = color.Transparent
	if !r.button.Disabled() {
		switch {
		case r.button.pressed:
			r.overlay.FillColor = theme.ColorForWidget(theme.ColorNamePressed, r.button)
		case r.button.hovered:
			r.overlay.FillColor = theme.ColorForWidget(theme.ColorNameHover, r.button)
		}
	}

	if r.button.focused && !r.button.Disabled() {
		r.focusIndicator.Show()
	} else {
		r.focusIndicator.Hide()
	}

	r.label.Refresh()
	r.background.Refresh()
	r.overlay.Refresh()
	r.focusIndicator.Refresh()
	r.Layout(r.button.Size())
}

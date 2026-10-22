package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type ButtonVariant uint8

const (
	ButtonPrimary ButtonVariant = iota
	ButtonSecondary
	ButtonGhost
	ButtonDanger
)

type Button struct {
	widget.BaseWidget

	Text       string
	Variant    ButtonVariant
	Importance widget.Importance
	OnTapped   func()
	radius     float32
	minWidth   float32

	focused  bool
	hovered  bool
	pressed  bool
	disabled bool
}

var (
	_ fyne.Accessible   = (*Button)(nil)
	_ fyne.Tappable     = (*Button)(nil)
	_ fyne.Focusable    = (*Button)(nil)
	_ desktop.Hoverable = (*Button)(nil)
	_ desktop.Mouseable = (*Button)(nil)
)

func NewButton(text string, variant ButtonVariant, onTapped func(), opts ...ButtonOption) *Button {
	b := &Button{Text: text, Variant: variant, OnTapped: onTapped}
	b.applyVariant(variant)
	for _, opt := range opts {
		if opt != nil {
			opt(b)
		}
	}
	b.ExtendBaseWidget(b)
	return b
}

func (b *Button) applyVariant(v ButtonVariant) {
	b.Variant = v
	switch v {
	case ButtonPrimary:
		b.Importance = widget.HighImportance
	case ButtonSecondary:
		b.Importance = widget.MediumImportance
	case ButtonGhost:
		b.Importance = widget.LowImportance
	case ButtonDanger:
		b.Importance = widget.DangerImportance
	}
}

func (b *Button) CreateRenderer() fyne.WidgetRenderer {
	return newButtonRenderer(b)
}

func (b *Button) AccessibilityLabel() string { return b.Text }

func (b *Button) AccessibilityRole() fyne.AccessibleRole {
	return fyne.AccessibleRoleButton
}

func (b *Button) Disable() {
	if b.disabled {
		return
	}
	b.disabled = true
	b.hovered = false
	b.pressed = false
	b.Refresh()
}

func (b *Button) Enable() {
	if !b.disabled {
		return
	}
	b.disabled = false
	b.Refresh()
}

func (b *Button) Disabled() bool { return b.disabled }

func (b *Button) FocusGained() {
	b.focused = true
	b.Refresh()
}

func (b *Button) FocusLost() {
	b.focused = false
	b.pressed = false
	b.Refresh()
}

func (b *Button) TypedRune(rune) {}

func (b *Button) TypedKey(ev *fyne.KeyEvent) {
	if b.disabled {
		return
	}
	if ev.Name == fyne.KeySpace || ev.Name == fyne.KeyReturn || ev.Name == fyne.KeyEnter {
		b.Tapped(nil)
	}
}

func (b *Button) Tapped(*fyne.PointEvent) {
	if b.disabled {
		return
	}
	b.pressed = false
	b.Refresh()
	if b.OnTapped != nil {
		b.OnTapped()
	}
}

func (b *Button) MouseIn(*desktop.MouseEvent) {
	if b.disabled {
		return
	}
	b.hovered = true
	b.Refresh()
}

func (b *Button) MouseMoved(*desktop.MouseEvent) {}

func (b *Button) MouseOut() {
	b.hovered = false
	b.pressed = false
	b.Refresh()
}

func (b *Button) MouseDown(ev *desktop.MouseEvent) {
	if b.disabled || ev.Button != desktop.MouseButtonPrimary {
		return
	}
	b.pressed = true
	b.Refresh()
}

func (b *Button) MouseUp(ev *desktop.MouseEvent) {
	if ev.Button != desktop.MouseButtonPrimary {
		return
	}
	b.pressed = false
	b.Refresh()
}

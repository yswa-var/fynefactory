package components

import "fyne.io/fyne/v2"

type ButtonOption func(*Button)

func WithButtonRadius(radius float32) ButtonOption {
	return func(b *Button) {
		b.radius = radius
	}
}

func WithButtonMinWidth(width float32) ButtonOption {
	return func(b *Button) {
		b.minWidth = width
	}
}

func (b *Button) SetVariant(v ButtonVariant) {
	b.applyVariant(v)
	b.Refresh()
}

func (b *Button) SetOnTapped(fn func()) {
	b.OnTapped = fn
}

func (b *Button) SetRadius(radius float32) {
	if b.radius == radius {
		return
	}
	b.radius = radius
	b.Refresh()
}

func (b *Button) SetMinWidth(width float32) {
	if b.minWidth == width {
		return
	}
	b.minWidth = width
	b.Refresh()
}

func (b *Button) SetText(text string) {
	if b.Text == text {
		return
	}
	b.Text = text
	b.Refresh()
}

var _ fyne.Widget = (*Button)(nil)

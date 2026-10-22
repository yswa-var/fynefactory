package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Surface struct {
	widget.BaseWidget
	Content fyne.CanvasObject
}

func NewSurface(content fyne.CanvasObject) *Surface {
	s := &Surface{Content: content}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Surface) CreateRenderer() fyne.WidgetRenderer {
	bg := canvas.NewRectangle(
		theme.Color(theme.ColorNameInputBackground),
	)
	bg.CornerRadius = 12

	content := s.Content
	if content == nil {
		content = widget.NewLabel("")
	}

	root := container.NewStack(
		bg,
		container.NewPadded(content),
	)

	return widget.NewSimpleRenderer(root)
}

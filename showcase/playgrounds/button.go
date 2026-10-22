package playgrounds

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yswa-var/fynefactory/components"
)

func NewButtonPlayground() fyne.CanvasObject {
	variant := components.ButtonPrimary
	radius, minWidth := float32(8), float32(120)
	disabled, text := false, "Primary"
	button := components.NewButton(text, variant, func() { fmt.Println("Button clicked") }, components.WithButtonRadius(radius), components.WithButtonMinWidth(minWidth))

	code := widget.NewMultiLineEntry()
	code.Wrapping = fyne.TextWrapWord
	code.SetMinRowsVisible(4)
	code.TextStyle = fyne.TextStyle{Monospace: true}
	updateCode := func() { code.SetText(buttonExample(text, variant, radius, minWidth, disabled)) }
	updateCode()

	variantSelect := widget.NewSelect([]string{"Primary", "Secondary", "Ghost", "Danger"}, func(value string) {
		switch value {
		case "Primary":
			variant = components.ButtonPrimary
		case "Secondary":
			variant = components.ButtonSecondary
		case "Ghost":
			variant = components.ButtonGhost
		case "Danger":
			variant = components.ButtonDanger
		}
		button.SetVariant(variant)
		updateCode()
	})
	variantSelect.SetSelected("Primary")

	radiusSlider := widget.NewSlider(0, 24)
	radiusSlider.Step, radiusSlider.Value = 1, float64(radius)
	radiusSlider.OnChanged = func(value float64) { radius = float32(value); button.SetRadius(radius); updateCode() }
	widthSlider := widget.NewSlider(80, 240)
	widthSlider.Step, widthSlider.Value = 8, float64(minWidth)
	widthSlider.OnChanged = func(value float64) { minWidth = float32(value); button.SetMinWidth(minWidth); updateCode() }
	disabledCheck := widget.NewCheck("Disabled", func(value bool) {
		disabled = value
		if value {
			button.Disable()
		} else {
			button.Enable()
		}
		updateCode()
	})
	textEntry := widget.NewEntry()
	textEntry.SetText(text)
	textEntry.OnChanged = func(value string) { text = value; button.SetText(value); updateCode() }

	controls := container.NewVBox(
		widget.NewLabelWithStyle("Properties", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		container.NewGridWithColumns(2,
			widget.NewLabel("Variant"), variantSelect,
			widget.NewLabel("Radius"), radiusSlider,
			widget.NewLabel("Minimum width"), widthSlider,
			widget.NewLabel("Button text"), textEntry,
		),
		disabledCheck,
	)
	preview := components.NewSurface(container.NewCenter(button))
	return container.NewVBox(
		container.NewVBox(widget.NewLabelWithStyle("Button", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}), widget.NewLabel("Actions, interactions and variants")),
		container.NewPadded(preview), controls,
		widget.NewLabelWithStyle("Go example · select and copy", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}), code,
	)
}

func buttonExample(text string, variant components.ButtonVariant, radius, minWidth float32, disabled bool) string {
	variantName := map[components.ButtonVariant]string{components.ButtonPrimary: "ButtonPrimary", components.ButtonSecondary: "ButtonSecondary", components.ButtonGhost: "ButtonGhost", components.ButtonDanger: "ButtonDanger"}[variant]
	var builder strings.Builder
	fmt.Fprintf(&builder, "button := components.NewButton(\n    %q,\n    components.%s,\n    func() { /* handle click */ },\n    components.WithButtonRadius(%.0f),\n    components.WithButtonMinWidth(%.0f),\n)", text, variantName, radius, minWidth)
	if disabled {
		builder.WriteString("\nbutton.Disable()")
	}
	return builder.String()
}

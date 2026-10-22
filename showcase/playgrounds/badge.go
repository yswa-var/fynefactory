package playgrounds

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yswa-var/fynefactory/components"
)

// NewBadgePlayground creates an interactive Badge demonstration.
func NewBadgePlayground() fyne.CanvasObject {
	text, variant := "Ready", components.BadgeSuccess
	badge := components.NewBadge(text, variant)
	code := widget.NewMultiLineEntry()
	code.Wrapping = fyne.TextWrapWord
	code.TextStyle = fyne.TextStyle{Monospace: true}
	updateCode := func() { code.SetText(badgeExample(text, variant)) }
	updateCode()

	selectVariant := widget.NewSelect([]string{"Success", "Warning", "Error", "Neutral"}, func(value string) {
		variant = map[string]components.BadgeVariant{"Success": components.BadgeSuccess, "Warning": components.BadgeWarning, "Error": components.BadgeError, "Neutral": components.BadgeNeutral}[value]
		badge.SetVariant(variant)
		updateCode()
	})
	selectVariant.SetSelected("Success")
	textEntry := widget.NewEntry()
	textEntry.SetText(text)
	textEntry.OnChanged = func(value string) { text = value; badge.SetText(value); updateCode() }

	controls := container.NewGridWithColumns(2,
		widget.NewLabel("Variant"), selectVariant,
		widget.NewLabel("Badge text"), textEntry,
	)
	return container.NewVBox(
		widget.NewLabelWithStyle("Badge", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel("Compact status labels with semantic, theme-aware variants."),
		container.NewPadded(container.NewCenter(badge)),
		widget.NewLabelWithStyle("Properties", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		controls,
		widget.NewLabelWithStyle("Go example · select and copy", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}), code,
	)
}

func badgeExample(text string, variant components.BadgeVariant) string {
	name := map[components.BadgeVariant]string{components.BadgeSuccess: "BadgeSuccess", components.BadgeWarning: "BadgeWarning", components.BadgeError: "BadgeError", components.BadgeNeutral: "BadgeNeutral"}[variant]
	return fmt.Sprintf("badge := components.NewBadge(%q, components.%s)", text, name)
}

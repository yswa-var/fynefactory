package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	fynetheme "fyne.io/fyne/v2/theme"
)

type ForgeTheme struct{}

var _ fyne.Theme = (*ForgeTheme)(nil)

func (t *ForgeTheme) Color(
	name fyne.ThemeColorName,
	variant fyne.ThemeVariant,
) color.Color {
	switch name {
	case fynetheme.ColorNamePrimary:
		return Accent

	case fynetheme.ColorNameBackground:
		if variant == fynetheme.VariantDark {
			return DarkBackground
		}
		return LightBackground
	}

	return fynetheme.DefaultTheme().Color(name, variant)
}

func (t *ForgeTheme) Font(
	style fyne.TextStyle,
) fyne.Resource {
	return fynetheme.DefaultTheme().Font(style)
}

func (t *ForgeTheme) Icon(
	name fyne.ThemeIconName,
) fyne.Resource {
	return fynetheme.DefaultTheme().Icon(name)
}

func (t *ForgeTheme) Size(
	name fyne.ThemeSizeName,
) float32 {
	return fynetheme.DefaultTheme().Size(name)
}

# Components

## Badge

`components.Badge` is a compact, theme-aware status label:

```go
badge := components.NewBadge("Ready", components.BadgeSuccess)
badge.SetText("Needs attention")
badge.SetVariant(components.BadgeWarning)
```

It supports `BadgeSuccess`, `BadgeWarning`, `BadgeError`, and `BadgeNeutral`.

## Button

`components.Button` is a theme-aware Fyne button with accessible keyboard and
pointer interaction, focus feedback, hover/pressed states, and four intentional
visual variants.

## Installation

From a Go module, import the package:

```bash
go get github.com/yswa-var/fynefactory/components
```

The package requires Go 1.26 or newer and Fyne v2.8 or newer.

```go
import "github.com/yswa-var/fynefactory/components"
```

## Constructor

The constructor keeps the original three-argument form:

```go
button := components.NewButton("Save", components.ButtonPrimary, save)
```

Its arguments are:

1. `text string` — the visible button label.
2. `variant ButtonVariant` — the visual and semantic emphasis.
3. `onTapped func()` — the callback, or `nil` when no callback is needed.
4. Optional `ButtonOption` values — narrowly scoped configuration.

Existing three-argument calls remain source-compatible.

## Variants

- `ButtonPrimary` — the main action.
- `ButtonSecondary` — a less prominent action.
- `ButtonGhost` — a low-emphasis action with no filled background.
- `ButtonDanger` — a destructive action.

## Functional options

Use options only when the default theme sizing is not appropriate for the
specific component:

```go
button := components.NewButton(
	"Delete",
	components.ButtonDanger,
	deleteItem,
	components.WithButtonRadius(8),
	components.WithButtonMinWidth(120),
)
```

- `WithButtonRadius(float32)` sets the corner radius. A non-positive value uses
  the active theme's button radius.
- `WithButtonMinWidth(float32)` sets an explicit minimum width. Otherwise the
  button uses its label size, theme padding, and a compact default minimum.

Theme and layout defaults should remain the first choice; these options exist
for deliberate component-level requirements, not general layout control.

## Dynamic setters

Properties can be changed after construction:

```go
button.SetText("Saved")
button.SetVariant(components.ButtonSecondary)
button.SetOnTapped(showDetails)
button.SetRadius(10)
button.SetMinWidth(140)
```

Visual changes refresh automatically. `SetOnTapped` only replaces the
callback.

## Disabled behavior

Disable a button to prevent pointer and keyboard activation:

```go
button.Disable()
button.Enable()

if button.Disabled() {
	// The button cannot be activated.
}
```

Disabled buttons render with the theme's disabled colors and do not respond to
hover, press, tap, Space, or Enter activation. Focus indication is hidden while
disabled.

## Runnable example

Save this as `main.go` in a Go module and run it with `go run .`:

```go
package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/yswa-var/fynefactory/components"
)

func main() {
	a := app.New()
	w := a.NewWindow("Button example")

	primary := components.NewButton(
		"Primary",
		components.ButtonPrimary,
		func() { fmt.Println("primary clicked") },
		components.WithButtonMinWidth(120),
	)
	secondary := components.NewButton("Secondary", components.ButtonSecondary, nil)
	danger := components.NewButton("Delete", components.ButtonDanger, nil)

	w.SetContent(container.NewVBox(primary, secondary, danger))
	w.Resize(fyne.NewSize(320, 220))
	w.ShowAndRun()
}
```

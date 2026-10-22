# FyneFactory

Reusable, theme-aware UI components for [Fyne](https://fyne.io/) applications.

FyneFactory provides small, composable widgets with a compact Go API, sensible
defaults, and interactive examples in the component showcase. It is intended
for applications that want a consistent visual language without building every
control from scratch.

<img width="1316" height="923" alt="Screenshot 2026-10-22 at 11 02 54 PM" src="https://github.com/user-attachments/assets/49fcbd1e-c510-4477-a754-cd88e7561082" />

## Components

| Component | Category | Description |
| --- | --- | --- |
| [Button](components/button.go) | Inputs | Accessible action button with primary, secondary, ghost, and danger variants. |
| [Badge](components/badge.go) | Display | Compact status label with success, warning, error, and neutral variants. |
| [Surface](components/surface.go) | Layout | Theme-aware content surface with rounded presentation. |
| [Accordion](components/accordion.go) | Layout | Collapsible content sections using Fyne's native accordion behavior. |

All public components are represented by a manifest in `components/*.manifest.json`
and are available in the showcase gallery.

## Installation

FyneFactory requires Go 1.26 or newer and Fyne v2.8 or newer.

```bash
go get github.com/yswa-var/fynefactory/components
```

Import the components package in your application:

```go
import "github.com/yswa-var/fynefactory/components"
```

## Quick start

```go
package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yswa-var/fynefactory/components"
)

func main() {
	a := app.New()
	w := a.NewWindow("FyneFactory")

	button := components.NewButton("Save", components.ButtonPrimary, func() {
		fmt.Println("saved")
	})

	accordion := components.NewAccordion(
		widget.NewAccordionItem("About", widget.NewLabel("A reusable Fyne component library.")),
		widget.NewAccordionItem("Usage", widget.NewLabel("Compose components with native Fyne layouts.")),
	).Medium()

	w.SetContent(container.NewVBox(button, accordion))
	w.Resize(fyne.NewSize(420, 320))
	w.ShowAndRun()
}
```

## Component APIs

Components use constructors for required values and setters for values that may
change while an application is running.

<img width="993" height="558" alt="image" src="https://github.com/user-attachments/assets/19deb5c0-2129-4c23-8332-5017b917a754" />

```go
button.SetText("Saved")
button.SetVariant(components.ButtonSecondary)
button.Disable()
```

<img width="990" height="514" alt="image" src="https://github.com/user-attachments/assets/d0e2508d-69f6-4c83-996a-d9dab9aa3513" />

```go
badge := components.NewBadge("Ready", components.BadgeSuccess)
badge.SetText("Needs attention")
badge.SetVariant(components.BadgeWarning)
```

```go
accordion.SetMultiOpen(true)
accordion.SetItemTitle(0, "Updated title")
accordion.Large()
```

The Accordion supports `XSmall()`, `Small()`, `Medium()`, and `Large()` density
helpers, plus `AccordionDefault`, `AccordionOutline`, and `AccordionGhost`
variants.

## Showcase

Run the interactive gallery from the repository root:

```bash
go run ./showcase
```

The showcase provides searchable component entries and playground controls for
supported variants, properties, and interaction states. It also uses the
repository theme so components can be reviewed as a cohesive set.

## Project layout

```text
components/                 Reusable public components and manifests
showcase/                   Interactive component gallery
showcase/playgrounds/       Isolated examples for each component
theme/                      Shared Fyne theme implementation
templates/component/        Templates for new components
```

## Development

Format and validate changes before submitting them:

```bash
gofmt -w components showcase theme
go test ./...
go vet ./...
```

When adding a component, follow the existing contracts:

1. Add the reusable implementation under `components/`.
2. Add focused unit and interaction tests.
3. Add a matching manifest.
4. Add a playground under `showcase/playgrounds/`.
5. Register the playground in `showcase/registry.go`.
6. Run the full test and vet checks.

Keep components independent of showcase code and prefer native Fyne widgets and
theme tokens wherever possible.

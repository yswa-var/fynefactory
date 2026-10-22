---
name: fyne-component
description: Create, modify, or improve production-quality FyneFactory Go UI components from natural-language descriptions. Use for custom Fyne widgets, component manifests, renderers, interactive playgrounds, tests, and component documentation.
---

# FyneFactory Component Engineer

You are an expert Go and Fyne component engineer.

Your goal is to create beautiful, reusable, accessible,
well-tested native UI components for FyneFactory.

## Repository structure

- `components/` - Public reusable UI components
- `theme/` - Shared themes and tokens
- `showcase/` - Interactive component explorer
- `showcase/playgrounds/` - Component demonstrations
- `templates/component/` - Generation templates
- `.agents/skills/fyne-component/` - This skill

## Phase 1: Inspect

Before generating code:

1. Read `go.mod`.
2. Inspect existing Button and Surface implementations.
3. Read `components/manifest.go`.
4. Read `showcase/registry.go` and `showcase/manifest.go`.
5. Inspect `templates/component/`.
6. Read `references/conventions.md`.

Reuse the existing project contracts and naming conventions.
Never invent a second manifest schema or parallel registry.

## Phase 2: Plan

Determine:

- Component name and purpose
- Public constructor and options
- Visual structure
- Interaction states
- Theme integration
- Keyboard and accessibility behavior
- Layout and minimum-size requirements
- Tests and showcase controls

Inspect related existing components first.

For straightforward requests, implement directly.
Ask questions only when a missing detail blocks a correct implementation.

## Phase 3: Implement

Produce the appropriate files:

1. Component Go implementation
2. Custom renderer, only when necessary
3. Functional options, when useful
4. Manifest matching the existing schema
5. Unit and interaction tests
6. Showcase playground
7. Registry integration
8. Developer documentation

Do not create files that the component does not need.

## Component engineering rules

- Prefer native Fyne widgets and containers.
- Use `widget.BaseWidget` for genuinely custom widgets.
- Call `ExtendBaseWidget` correctly.
- Reuse existing theme tokens.
- Prefer theme-aware colors and sizes.
- Keep reusable components independent of showcase code.
- Avoid mutable global state.
- Avoid unnecessary goroutines and third-party dependencies.
- Refresh visual state when properties change.
- Never duplicate event activation.
- Do not break existing public APIs.
- Handle disabled, hover, focus and keyboard behavior where applicable.
- Account for narrow layouts and long text.
- Avoid hardcoded platform-specific paths.
- Preserve Fyne's UI-thread requirements.

## Phase 4: Showcase

Every public visual component must have a working playground.

The playground should demonstrate:

- Default appearance
- Supported variants
- Important interaction states
- Configurable properties
- Copyable Go example
- Light and dark themes

Integrate with the existing gallery registry.

## Phase 5: Validate

Run:

    gofmt -w <changed-go-files>
    go test ./...
    go vet ./...

Run race tests when relevant.

Use `scripts/validate.sh` as the standard project gate.

Do not claim GUI behavior was verified unless the GUI
was actually tested.

If tests fail:
1. Identify the cause.
2. Correct the implementation.
3. Rerun validation.

Never hide failing tests by weakening assertions.

## Phase 6: Review

Before finishing, check:

- Does the component follow existing APIs?
- Does the public API feel natural in Go?
- Does it render correctly across themes?
- Are disabled and focus states correct?
- Are tests meaningful?
- Is the showcase integrated?
- Is the manifest valid?
- Are docs accurate?

## Final output

Report:

- Files created or changed
- Public API and example
- Tests actually executed
- Remaining limitations
- Any manual GUI checks required

Do not mark unfinished work as complete.

# FyneFactory Engineering Conventions

## API Design

- Use idiomatic exported Go names.
- Prefer `New<Component>(...)` constructors.
- Preserve existing constructors when adding options.
- Use typed functional options for optional configuration.
- Provide setters for mutable public properties.
- Keep callbacks explicit.
- Use zero values sensibly.

## Widget Lifecycle

- Extend BaseWidget when implementing custom widgets.
- Store long-lived canvas objects in the renderer.
- Never allocate new canvas objects on every Refresh.
- Keep Layout and MinSize consistent.
- Avoid recursively triggering Refresh.
- Treat focus, hover, press and disabled as separate states.
- Avoid double-calling OnTapped.

## Theme

- Use shared FyneFactory theme tokens.
- Resolve widget-specific theme colors when supported.
- Support light and dark variants.
- Never assume one background color.
- Ensure foreground/background contrast.

## Responsive Behavior

- Test at normal and narrow window sizes.
- Support text wrapping where appropriate.
- Avoid clipping essential controls.
- Use scrolling for genuinely oversized content.
- Do not assume desktop-only usage.

## Testing

- Test constructors and options.
- Test state transitions.
- Test callback invocation counts.
- Test disabled behavior.
- Test manifest validity.
- Test registry integration.
- Test MinSize and layout edge cases.
- Add benchmarks only where performance matters.

## Compatibility

- Follow the Go and Fyne versions declared in go.mod.
- Do not change existing public APIs without permission.
- Never add a dependency without a clear justification.

## Manifest Contract

The source of truth is components/manifest.go.

Inspect the existing JSON manifests.
Do not infer the schema from examples alone.
Validate using existing test infrastructure.

## Quality

A component is complete only when its implementation,
documentation, manifest, tests and showcase integration
are consistent.

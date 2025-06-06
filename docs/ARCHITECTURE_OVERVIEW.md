# Architecture Overview

This document summarizes the modular handler pattern introduced in the `v1` directory. It serves as a reference for contributors who are exploring the codebase.

## Module Layout

All gameplay code lives under `v1/internal/`. The packages were split out of a
single `game` module and are organized as follows:

- `assets` – fonts, palettes and other asset helpers
- `config` – runtime configuration and defaults
- `core` – timers, points, HUD helpers and other utilities
- `event` – event bus and event type definitions
- `econ` – resource structs and tech tree parsing
- `word` – the global queue and typing statistics
- `entity` – shared interfaces for drawable objects
- `mob` – enemy implementations
- `structure` – buildings, towers and the player base
- `worker` – resource gathering buildings
- `sprite` – image helpers and sprite utilities
- `input` – keyboard input processing
- `phase` – game state enumeration helpers
- `skill` – skill tree logic and persistence
- `tech` – tech tree handler
- `game` – the orchestrator that owns all handlers

See [INTERNAL_RESTRUCTURE.md](INTERNAL_RESTRUCTURE.md) for the current import
rules and planned renames.

Each module exposes a `Handler` with an `Update(dt float64)` method. The `Handler` interface allows the main game engine to update modules in a uniform manner.

## Event Bus

`internal/event` provides a lightweight pub/sub system. Each handler exposes channels for its specific event type (for example `EntityEvents` or `UIEvents`). Handlers subscribe to the events they care about using the shared `EventBus` and communicate by publishing events.
Handlers can stop receiving updates with `Unsubscribe` when they are done.

Example:

```go
// Publishing a notification from the tech handler
bus.Publish("ui", event.UIEvent{Type: "notification", Payload: "Tech unlocked"})

// UI handler subscribes to UI events
bus.Subscribe("ui", uiCh)
```

The `game.Game` struct owns a pointer to every handler and the shared `EventBus`. During `Game.Update` each handler's `Update` method is called. Rendering is also coordinated by `game.Game` using state from these handlers.

## Adding New Features

1. Create a new module under `internal/` if needed and define a `Handler`.
2. Add event types or channels in `internal/event` if the new module needs to communicate with others.
3. Wire the handler and its event channels into `game.Game` (see the existing fields for guidance).
4. Keep event payloads small and decoupled. Prefer passing IDs or simple structs rather than large objects.

This pattern keeps modules loosely coupled while still allowing coordination through events.

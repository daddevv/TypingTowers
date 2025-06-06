# Architecture Overview

This document summarizes the modular handler pattern introduced in the `v1` directory. It serves as a reference for contributors who are exploring the codebase.

## Module Layout

All gameplay code lives under `v1/internal/`. Modules include:

- `entity` – ally and enemy units
- `ui` – HUD, menus and overlays
- `tech` – tech tree and skill tree logic
- `tower` – towers, projectiles and related logic
- `phase` – game state transitions
- `content` – asset loading helpers
- `sprite` – image helpers and sprite utilities
- `event` – event bus and event types

Each module exposes a `Handler` with an `Update(dt float64)` method. The `Handler` interface allows the main game engine to update modules in a uniform manner.

## Event Bus

`internal/event` provides a lightweight pub/sub system. Each handler exposes channels for its specific event type (for example `EntityEvents` or `UIEvents`). Handlers subscribe to the events they care about using the shared `EventBus` and communicate by publishing events.

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

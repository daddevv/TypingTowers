# Internal Package Restructure

This document explains the current layout under `v1/internal` and outlines the permitted import
relationships between packages. The goal is to keep the codebase modular and avoid import cycles as
we continue to break the original monolithic `game` package into smaller components.

## Current Modules

```
assets    config    core     event
entity    mob       structure worker
econ      word      sprite   input
phase     skill     tech     game
```

- **assets** – fonts, palettes and helper functions for loading image/audio data.
- **config** – runtime configuration and defaults.
- **core** – miscellaneous utilities such as timers, points and HUD helpers.
- **event** – the publish/subscribe event bus and event types.
- **econ** – resource values and tech tree parsing.
- **word** – the global queue, typing stats and word types.
- **entity** – shared interfaces for anything that appears on screen.
- **mob** – enemy implementations (currently various "Mob" types).
- **structure** – buildings, towers and the player base.
- **worker** – resource gathering buildings such as the Farmer and Miner.
- **sprite** – helpers for image manipulation.
- **input** – keyboard input processing.
- **phase** – game state enumeration and helpers.
- **skill** – skill tree logic and persistence.
- **tech** – tech tree handler.
- **game** – orchestrates all other handlers and owns the update loop.

## Import Map

The packages form a loose hierarchy. `game` imports every other package but nothing inside
`internal/` should import `game`.

**Base packages** (no internal imports):

- `assets`
- `config`
- `core`
- `event`

**Domain packages** may import the base packages and each other only through defined
interfaces. Avoid circular dependencies by following these guidelines:

- `entity` imports `core` and `assets` only.
- `mob` imports `assets`, `core`, `entity` and `structure`.
- `structure` imports `assets`, `core`, `entity`, `econ` and `word`.
- `worker` imports `assets`, `core`, `econ` and `word`.
- `econ` and `word` do not depend on other domain packages.
- `input`, `phase`, `skill`, `sprite` and `tech` depend only on base packages and optionally
  on `econ` or `word` where required.
- `game` depends on all of the above and ties them together.

This import map ensures that assets, configuration and core utilities remain independent
and that higher level systems communicate via events or well defined interfaces.

## Planned Renames

To better reflect their roles in the modular architecture we plan to rename several packages:

- `mob` → `enemy`
- `structure` → `building`
- `worker` → `gatherer`

These changes will occur once all dependent code is updated to the new package names.

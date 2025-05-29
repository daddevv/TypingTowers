# TypeDefense Architecture Overview

This document provides a high-level overview of the architecture, core systems, and extensibility patterns of TypeDefense.

---

## Directory Structure

```markdown
desktop/
├── cmd/
│   ├── game/           # Main game entry point
│   └── preview/        # Sprite animation preview tool
├── content/            # JSON content configs (levels, mobs, worlds)
├── internal/
│   ├── engine/         # Game engine and state management
│   ├── game/           # Core game logic and loop
│   ├── entity/         # Game entities (Player, Mobs, Projectiles)
│   ├── ui/             # User interface components
│   ├── world/          # Level and world definitions
│   └── utils/          # Utility functions
└── assets/             # Game assets (images, fonts, sounds)
```

---

## Core Systems

### 1. Entity System

- **Entity Interface**: All game objects implement `Entity`:

  ```go
  type Entity interface {
      Draw(screen *ebiten.Image)
      Update() error
      SetPosition(x, y float64)
      GetPosition() ui.Location
  }
  ```

- **Mob Interface**: Specialized for enemies, extends `Entity` and adds letter/behavior methods.
- **Component-Based**: Entities have position, sprites, animations, and behavior components.

### 2. Content System

- **JSON-Driven**: All levels, mobs, and worlds are defined in JSON files under `desktop/content/`.
- **Hot Reload**: Edit JSON and restart the game to see changes; no recompilation needed.
- **Content Types**:
  - `levels.json`: Level definitions (waves, letter pools, etc.)
  - `mobs.json`: Mob/enemy definitions (sprite, animation, letter count, etc.)
  - `worlds.json`: World/biome definitions (backgrounds, themes)

### 3. Rendering System

- **Fixed Canvas**: 1920x1080 internal resolution, scaled to window size.
- **Optimized Drawing**: Letter images are cached globally for performance.
- **Layered Rendering**: Background, entities, projectiles, UI/HUD.

### 4. Input System

- **InputHandler**: Processes keyboard input, manages projectile creation.
- **Targeting**: Automatically targets the closest mob with the matching letter.
- **Rapid Typing**: Supports fast consecutive inputs.

### 5. Game Loop & State Management

- **Game Loop**: Runs at 100 TPS, updates all entities and handles input/collisions.
- **State Machine**: Clean separation between menu, game, pause, and game over states.

### 6. Level & World System

- **Levels**: Composed of waves, each with its own letter pool and mob spawn chances.
- **Worlds/Biomes**: Provide backgrounds and themes; referenced by levels.

---

## Extensibility Patterns

- **Adding Mobs**: Implement a new struct embedding `MobBase`, register with the spawner, and define in `mobs.json`.
- **Adding Worlds/Biomes**: Add new backgrounds and reference them in `worlds.json`.
- **Adding Game Modes**: Implement new state structs and integrate with the engine state machine.
- **Custom Letter Pools**: Implement the `LetterPool` interface for themed or mode-specific letter sets.

---

## Performance & Best Practices

- **Parallel Mob Updates**: Mobs are updated in parallel for better performance.
- **Asset Caching**: Images and fonts are loaded and cached to minimize per-frame allocations.
- **Efficient Update Loops**: Keep entity `Update()` methods fast; avoid allocations in hot paths.

---

## Contributing

- **Content**: Add or edit JSON files in `desktop/content/`.
- **Code**: Follow Go conventions, use interfaces for extensibility, and document new systems.
- **Testing**: Playtest new content and use the sprite preview tool for asset validation.

---

For more details, see `README.md`, `CONTRIBUTING.md`, and `desktop/content/README.md`.

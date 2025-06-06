# Sprint Backlog

> **Convention:** One feature or fix = one line item.  
> Sub-tasks are indented bullet points (checked as they land).

---

## Core Gameplay Loop Demo (Highest Priority)

- [ ] **ARCH-001** Modularize internal codebase with handler/event system
  - [x] **T-001** Define new module structure and create directories:
    - `entity`: All ally/enemy minions, workers, and base logic
    - `ui`: HUD, overlays, menus, and all rendering/UI logic
    - `tech`: Tech tree, skill tree, upgrades, and related systems
    - `tower`: Tower logic, projectiles, and related mechanics
    - `phase`: Game phase/state enums and transitions
    - `econ`: Economy resources and asset loading management
    - `sprite`: Sprite/image providers and ebiten.Image helpers
    - `game`: Core engine, main loop, and glue code
  - [x] **T-002** For each module, define a `Handler` struct (e.g., `EntityHandler`, `UIHandler`, etc.) with basic state and interface
  - [x] **T-003** Refactor `game.Engine` to hold pointers to all handlers as fields
  - [x] **T-004** Implement `Update(dt)` method for each handler; call all handlers from `Engine.Update(dt)`
  - [x] **T-005** Design Go channel-based pub/sub event system:
    - [x] **T-006** Define event types for each module (e.g., `EntityEvent`, `UIEvent`, etc.)
    - [x] **T-007** Each handler exposes channels for publishing/subscribing to events
    - [x] **T-008** Implement event communication between handlers (e.g., UI notification on tech unlock)
  - [x] **T-009** Migrate all existing logic/files into new module structure
  - [x] **T-010** Update all imports and references to match new structure
  - [x] **T-011** Write/adjust tests for new handler/event system
  - [x] **T-012** Document the new architecture and handler/event pattern for contributors
  - [x] **T-013** Ensure `game.Game` acts as main renderer, coordinating rendering using handler state
  - [ ] **T-014** Finalize modular split and package renames
    - [x] **MIG-001** Document package layout and import map (`INTERNAL_RESTRUCTURE.md`)
    - [x] **MIG-002** Rename `mob` package to `enemy`
    - [x] **MIG-003** Rename `structure` package to `building`
    - [x] **MIG-004** Move worker types under `building/gatherer`
    - [ ] **MIG-005** Update all imports and tests after renames
    - [ ] **MIG-006** Remove deprecated package references
- [ ] Ensure all core gameplay loop features are fully playable and testable
  - [ ] Integrate combat and skill tree systems into the main loop
  - [ ] Finalize and document any remaining edge cases in handler/event system
  - [ ] Review and update contributor documentation as needed

---

## Typing Minigames & Metrics

- [ ] **MINIGAME-001** Speed trial mode
  - [ ] Design rules and win/lose conditions
  - [ ] Implement timer and scoring logic
  - [ ] Integrate with main game UI and stats
  - [ ] Playtest and balance difficulty

- [ ] **MINIGAME-002** Accuracy challenge mode
  - [ ] Define accuracy thresholds and penalties
  - [ ] Implement accuracy tracking and feedback
  - [ ] Integrate with stats panel and results screen
  - [ ] Playtest and adjust thresholds

- [ ] **MINIGAME-003** Word puzzle/anagram mode
  - [ ] Design puzzle generation logic
  - [ ] Implement input and validation for anagrams
  - [ ] Add scoring and feedback
  - [ ] Playtest for variety and challenge

- [ ] **MINIGAME-004** Boss practice mode
  - [ ] Define boss mechanics and special rules
  - [ ] Implement boss encounter logic
  - [ ] Integrate with main game and stats
  - [ ] Playtest and balance

---

## Art, Audio & Polish

- [ ] **ART-001** 32x32 farmer, lumberjack, miner idle sprites
  - [ ] Define sprite requirements and animation frames
  - [ ] Create and export spritesheets (PNG)
  - [ ] Integrate into asset pipeline and entity rendering
  - [ ] Test in-game display

- [ ] **ART-002** Orc grunt walk + hit animation
  - [ ] Define animation frames and timing
  - [ ] Create and export spritesheets
  - [ ] Integrate into asset pipeline and entity logic
  - [ ] Test animation playback

- [ ] **ART-003** Tower upgrade visual indicators
  - [ ] Design and create upgrade icons/overlays
  - [ ] Integrate into asset pipeline and tower rendering
  - [ ] Test indicator updates in-game

- [ ] **SFX-001** Key-hit, crit, jam placeholders (chiptune)
  - [ ] Define and source required sound effects
  - [ ] Integrate into asset pipeline and game logic
  - [ ] Test sound playback

- [ ] **SFX-002** Background music for different game states
  - [ ] Define music requirements and source tracks
  - [ ] Integrate into asset pipeline and game state logic
  - [ ] Test music transitions

---

## Advanced Systems (Future Sprints)

- [ ] **MINION-001** Minion summoning via typed commands
- [ ] **MINION-002** Minion AI and unique roles
- [ ] **MINION-003** Minion upgrades and management UI
- [ ] **IDLE-001** Auto-collection and offline progress
- [ ] **IDLE-002** Upgradable idle generators
- [ ] **IDLE-003** Prestige/reset system
- [ ] **FUZZ-002** Stress test for performance and stability
- [ ] **FUZZ-003** Automated regression tests for core systems

---

*Focus all effort on the "Core Gameplay Loop Demo" and minigame integration until fully playable and testable. Archive completed sprints to `TODO_ARCHIVE.md` when merged.*

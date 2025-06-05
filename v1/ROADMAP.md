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
    - `content`: Asset/content loaders and resource management
    - `sprite`: Sprite/image providers and ebiten.Image helpers
    - `game`: Core engine, main loop, and glue code
  - [x] **T-002** For each module, define a `Handler` struct (e.g., `EntityHandler`, `UIHandler`, etc.) with basic state and interface
  - [x] **T-003** Refactor `game.Engine` to hold pointers to all handlers as fields
  - [x] **T-004** Implement `Update(dt)` method for each handler; call all handlers from `Engine.Update(dt)`
  - [x] **T-005** Design Go channel-based pub/sub event system:
    - [x] **T-006** Define event types for each module (e.g., `EntityEvent`, `UIEvent`, etc.)
    - [x] **T-007** Each handler exposes channels for publishing/subscribing to events
    - [ ] **T-008** Implement event communication between handlers (e.g., UI notification on tech unlock)
  - [ ] **T-009** Migrate all existing logic/files into new module structure
  - [ ] **T-010** Update all imports and references to match new structure
  - [ ] **T-011** Write/adjust tests for new handler/event system
  - [ ] **T-012** Document the new architecture and handler/event pattern for contributors
  - [ ] **T-013** Ensure `game.Engine` acts as main renderer, coordinating rendering using handler state
  - [ ] **T-014** (Optional) Add migration/compatibility notes for contributors
  - [ ] **Design Note:** The `game.Engine` will act as the main renderer. Each handler's state (e.g., entities, UI, tech, towers) composes the overall game state. There is no separate/dedicated renderer; rendering is coordinated by the engine using handler state.

---

## Combat & Military

- [x] **TEST-COMBAT-EDGE** Simulation unit tests for all common and edge cases
  - [x] Footman survives after killing a single grunt (verify HP > 0)
  - [x] Footman dies if grunt damage is lethal (verify removal from military)
  - [x] Multiple Footmen vs multiple Grunts: all combinations (1v2, 2v1, 2v2)
  - [x] Simultaneous combat: overlapping units resolve damage correctly
  - [x] No combat occurs if units do not overlap (verify no HP loss)
  - [x] Dead units are removed from the military/orc lists immediately
  - [x] Units with 0 HP cannot attack or be attacked further
  - [x] Combat does not occur if either unit is already dead
  - [x] Test for correct handling of edge cases (e.g., both units die in same tick)
  - [x] Test for no panics or index errors when removing units during iteration
  - [x] Add all tests to CI pipeline

---

## Art, Audio & Polish

- [ ] **ART-001** 32x32 farmer, lumberjack, miner idle sprites
  - [ ] Define sprite requirements and animation frames for each character
  - [ ] Create initial 32x32 pixel art for each idle sprite
  - [ ] Export spritesheets in required format (e.g., PNG)
  - [ ] Integrate sprites into game asset pipeline
  - [ ] Update entity rendering logic to use new idle sprites
  - [ ] Test or visually check correct sprite display

- [ ] **ART-002** Orc grunt walk + hit animation
  - [ ] Define animation frame count and timing for walk and hit actions
  - [ ] Create 32x32 orc grunt walk/hit animation frames
  - [ ] Export animation spritesheets in required format
  - [ ] Integrate orc grunt animations into asset pipeline
  - [ ] Update orc grunt entity logic to trigger walk/hit animations
  - [ ] Test animation playback in-game

- [ ] **ART-003** Tower upgrade visual indicators
  - [ ] Define visual indicator styles for each tower upgrade level
  - [ ] Design and create upgrade icons or overlays for towers
  - [ ] Integrate upgrade indicators into game asset pipeline
  - [ ] Update tower rendering logic to display correct indicator based on upgrade level
  - [ ] Test in-game to verify indicators update correctly

- [ ] **SFX-001** Key-hit, crit, jam placeholders (chiptune)
  - [ ] Define required sound effects (key-hit, crit, jam)
  - [ ] Create/source chiptune placeholder sounds
  - [ ] Integrate sound effects into game asset pipeline
  - [ ] Update game logic to trigger sounds on key-hit, crit, and jam events
  - [ ] Test in-game for correct sound playback

- [ ] **SFX-002** Background music for different game states
  - [ ] Define music requirements for each game state
  - [ ] Create/source chiptune background music tracks
  - [ ] Integrate music tracks into game asset pipeline
  - [ ] Update game state logic to play appropriate music for each state
  - [ ] Test transitions for smooth music changes

---

## Typing Metrics & Minigames

- [x] **MET-001** Capture per-word accuracy & time
  - [x] Define data structures to store per-word accuracy and completion time
  - [x] Update queue processing logic to record accuracy and time for each word
  - [x] Store per-word stats in a history buffer or log
  - [x] Add unit tests to verify per-word stats are captured correctly
  - [x] Expose per-word stats to HUD or stats panel

- [x] **MET-002** Rolling WPM (last 30 s)
  - [x] Implement a time-based buffer to track recent typing events
  - [x] Calculate rolling WPM using only events from the last 30 seconds
  - [x] Add method to TypingStats or a new struct for rolling WPM
  - [x] Write tests to verify rolling WPM calculation
  - [x] Display rolling WPM in the HUD or stats panel

- [x] **UI-MET** Toggle stats panel (`Tab`)
  - [x] Add a boolean field to Game for stats panel visibility
  - [x] Capture `Tab` key in Input handler to toggle stats panel
  - [x] Implement HUD rendering logic for the stats panel
  - [x] Display per-word stats, rolling WPM, and accuracy in the panel
  - [x] Add tests to verify panel toggling and display

- [ ] **MINIGAME-001** Speed trial mode
- [ ] **MINIGAME-002** Accuracy challenge mode
- [ ] **MINIGAME-003** Word puzzle/anagram mode
- [ ] **MINIGAME-004** Boss practice mode

  *(See previous roadmap for detailed sub-tasks for each minigame.)*

---

## Advanced Systems (Future Sprints)

- [ ] **MINION-001** Minion summoning via typed commands
- [ ] **MINION-002** Minion AI and unique roles
- [ ] **MINION-003** Minion upgrades and management UI
- [ ] **IDLE-001** Auto-collection and offline progress
- [ ] **IDLE-002** Upgradable idle generators
- [ ] **IDLE-003** Prestige/reset system
- [x] **FUZZ-001** Engine fuzz tester and robustness checks
- [ ] **FUZZ-002** Stress test for performance and stability
- [ ] **FUZZ-003** Automated regression tests for core systems

---

*Focus all effort on the "Core Gameplay Loop Demo" and combat/skill tree integration until fully playable and testable. Archive completed sprints to `TODO_ARCHIVE.md` when merged.*

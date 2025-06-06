# Sprint Archive

## Core Gameplay Loop Demo

- **CORE-DEMO** Achieved a working demo of the core gameplay loop:
  - Farmer and Barracks buildings enqueue words to the global queue
  - Typing words processes queue, produces resources, and spawns units
  - Shared queue manager with jam/back-pressure mechanics
  - Per-building cooldowns and letter unlocks
  - HUD displays queue, cooldowns, and resources
  - Basic enemy waves and base HP system
  - End-to-end test: survive 5+ waves with resource/typing feedback

---

## Immediate Next Steps (Post-Demo Polish)

- **QUEUE-001** Letter-by-letter queue processing
  - Adjust backlog pressure for letter queues
  - Static word processing location at (400, 900) with conveyor effect
- **UI-001** Tower selection and upgrade system
  - `/` to enter tower selection mode, letter labels, upgrade menu
- **CMD-001** Command mode for power users
  - `:` to enter command mode, basic and advanced commands
- **TITLE-001** Title screen and main menu
  - MainMenuState, logo, background music, animated background, settings
- **PREGAME-001** Pre-game setup and tutorial
  - Character/difficulty selection, tutorial, typing test, mode selection

---

## Resource Loop & HUD

- **R-001** Implement Gold/Wood/Stone/Iron structs
- **R-002** Farmer, Lumberjack, Miner cooldowns produce resources
  - Balance numbers in `config.json`
- **HUD-001** Top bar resource icons (`G`, `W`, `S`, `I`, `M`)
- **HUD-002** Show word processing queue with conveyor belt animation
- **HUD-003** Tower selection overlay with letter labels
- **TEST-RES** Integration test 3 min sim, resources > 0
  - **T-001** Create test file `internal/game/resources_integration_test.go` with basic scaffold
  - **T-002** Initialize a `Game` instance in test using default config
  - **T-003** Simulate update loop for 3 minutes of game time (e.g., step through `Update` calls)
  - **T-004** Access `ResourcePool` after simulation and assert each resource > 0
  - **T-005** Ensure test runs headlessly and integrate into CI pipeline

---

## Tech Tree & Progression

- **T-002** Tech tree parser + in-memory graph
  - Define Go structs matching YAML schema (TechNode, TechTree, effects/prereqs)
  - Implement YAML loader in Go (use yaml.v2)
  - Validate tech graph (detect cycles, missing prereqs)
  - Build in-memory graph (nodes by ID, adjacency)
  - Expose graph API (GetPrerequisites, UnlockOrder)
  - Write unit tests for parser and validation
- **T-003** Keyboard UI for tech purchase (`/` search, `Enter` buy)
  - Add `techMenuOpen`, `searchBuffer`, and selection index fields to `Game`
  - Capture `/` key in `Input.Update` to toggle tech menu mode
  - Render tech menu overlay: list `TechNode.Name`, unlocked letters, and achievements
  - Implement search input handling: append typed chars and backspace to `searchBuffer`
  - Filter `TechTree.nodes` by `searchBuffer` and update displayed list
  - Handle Up/Down arrow keys to move highlight over filtered nodes
  - Handle `Enter` key to purchase selected tech: check prerequisites/resources, call `UnlockNext`
  - Write unit tests for tech menu: toggling, filtering, navigation, and purchase flow

---

## Military Prototype

- **M-001** Barracks building pushes unit words (letter-by-letter)
  - Refactor Barracks to enqueue words to the global queue letter-by-letter
  - Update queue manager to support partial word progress and per-letter validation
  - Ensure Barracks cooldown only resets after full word is typed
  - Add unit tests for Barracks letter-by-letter queue logic
  - Integrate Barracks with HUD to show letter-by-letter progress

- **M-002** Footman entity (HP, dmg, speed)
  - Define Footman struct with HP, damage, and speed fields
  - Implement Footman movement and update logic
  - Add Footman spawn logic to Barracks on word completion
  - Write unit tests for Footman creation and state updates

- **M-003** Combat resolution attacker vs orc grunt
  - Define OrcGrunt struct with HP and damage
  - Implement combat logic between Footman and OrcGrunt
  - Update military system to resolve combat each tick
  - Add tests for combat outcomes and edge cases

---

## Game States & Persistence

- **STATE-001** Proper game state management (MainMenu, PreGame, Playing, Paused, GameOver, Settings)
  - Define a `GamePhase` enum/type covering all major states
  - Refactor main game loop to use `GamePhase` for state transitions
  - Implement state transition logic (e.g., menu → pregame → playing → paused/gameover/settings)
  - Ensure each state has a dedicated update and draw handler
  - Add keyboard navigation and transitions between states (e.g., Esc to pause, Enter to continue)
  - Write unit tests for state transitions and edge cases

- **SAVE-001** Comprehensive save/load system (multiple slots, auto-save)
  - Design a save file structure supporting multiple slots and versioning
  - Implement save/load logic for all core game data (resources, towers, buildings, tech, settings)
  - Add auto-save functionality (e.g., after each wave or major event)
  - Create a save/load menu UI for selecting slots
  - Handle save/load errors and version mismatches gracefully
  - Write integration tests for save/load, including slot switching and auto-save

---

## Gathering & Military Core

- **P-001** Implement Gathering (Farmer) building
  - Farmer cooldown logic
  - Farmer word generation from letter pool
  - Resource output on word completion
- **P-002** Implement Military (Barracks) building
  - Barracks cooldown logic
  - Barracks word generation from letter pool
  - Unit spawn on word completion
- **P-003** Shared queue manager
  - Global FIFO queue structure
  - Enqueue from multiple buildings
  - Dequeue and typing validation
- **P-004** Per-building cooldown timers
  - Timer tick/update logic
  - Cooldown reset on word completion
- **P-005** Playtest word density
  - Simulate 5 min session, measure words/sec
  - Adjust cooldowns/word lengths for 1–1.5 words/sec target
- **P-006** Letter unlock order & cost curves
  - Draft full unlock order for all buildings
  - Define cost progression for each letter unlock
  - Document in `docs/LETTER_UNLOCKS.md`

---

## Queue MVP Hardening

- **Q-001** Refactor global queue to support color-coding per building
  - Add `family` field to `Word` struct
  - Palette map + ANSI tests
- **Q-002** Back-pressure damage when backlog ≥ 5
  - Unit test: enqueue 6 words, expect base HP −1
- **Q-003** Jam state visual
  - Red flash on mistype
  - Audio “clank” SFX placeholder

---

## Integration & Playtest

- **INT-001** Integrate Gathering (Farmer) building with resource system
- **INT-002** Integrate Military (Barracks) building with unit spawning
- **INT-003** Integrate Shared Queue Manager with HUD and building inputs
- **INT-004** Integrate Per-Building Cooldown Timers with UI
- **INT-005** Integrate Back-Pressure Damage mechanic
- **INT-006** Integrate Jam State Visuals and Audio
- **INT-007** Implement Letter Unlock System
- **TEST-CORELOOP** End-to-end playtest of the core loop

# TODO Archive

## Core Gameplay Loop Demo

- **ARCH-001** Modularize internal codebase with handler/event system
  - Defined new module structure and created directories for entity, ui, tech, tower, phase, econ, sprite, game
  - Defined Handler structs for each module
  - Refactored game.Engine to hold pointers to all handlers
  - Implemented Update(dt) method for each handler and coordinated updates
  - Designed Go channel-based pub/sub event system and defined event types
  - Exposed channels for publishing/subscribing to events in each handler
  - Implemented event communication between handlers
  - Migrated all logic/files into new module structure
  - Updated all imports and references
  - Wrote/adjusted tests for new handler/event system
  - Documented new architecture and handler/event pattern
  - Ensured game.Game acts as main renderer using handler state
  - Finalized modular split and package renames
    - Documented package layout and import map
    - Renamed mob to enemy, structure to building
    - Moved worker types under building/gatherer
    - Updated all imports and tests after renames
    - Removed deprecated package references

## Combat & Military

- **TEST-COMBAT-EDGE** Simulation unit tests for all common and edge cases
  - Footman survives after killing a single grunt (verify HP > 0)
  - Footman dies if grunt damage is lethal (verify removal from military)
  - Multiple Footmen vs multiple Grunts: all combinations (1v2, 2v1, 2v2)
  - Simultaneous combat: overlapping units resolve damage correctly
  - No combat occurs if units do not overlap (verify no HP loss)
  - Dead units are removed from the military/orc lists immediately
  - Units with 0 HP cannot attack or be attacked further
  - Combat does not occur if either unit is already dead
  - Correct handling of edge cases (e.g., both units die in same tick)
  - No panics or index errors when removing units during iteration
  - Added all tests to CI pipeline

## Typing Metrics & Minigames

- **MET-001** Capture per-word accuracy & time
  - Defined data structures for per-word accuracy and completion time
  - Updated queue processing logic to record stats
  - Stored per-word stats in a history buffer/log
  - Added unit tests for per-word stats
  - Exposed stats to HUD/stats panel

- **MET-002** Rolling WPM (last 30 s)
  - Implemented time-based buffer for recent typing events
  - Calculated rolling WPM using last 30 seconds
  - Added method for rolling WPM
  - Wrote tests for rolling WPM
  - Displayed rolling WPM in HUD/stats panel

- **UI-MET** Toggle stats panel (`Tab`)
  - Added boolean field for stats panel visibility
  - Captured `Tab` key in Input handler
  - Implemented HUD rendering logic for stats panel
  - Displayed per-word stats, rolling WPM, and accuracy
  - Added tests for panel toggling and display

- **FUZZ-001** Engine fuzz tester and robustness checks
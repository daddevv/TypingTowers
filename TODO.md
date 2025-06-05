# Sprint Backlog

> **Convention:** One feature or fix = one line item.  
> Sub-tasks are indented bullet points (checked as they land).

---

## Core Gameplay Loop Demo (Highest Priority)

- [x] **CORE-DEMO** Achieve a working demo of the core gameplay loop:
  - [x] Farmer and Barracks buildings enqueue words to the global queue
  - [x] Typing words processes queue, produces resources, and spawns units
  - [x] Shared queue manager with jam/back-pressure mechanics
  - [x] Per-building cooldowns and letter unlocks
  - [x] HUD displays queue, cooldowns, and resources
  - [x] Basic enemy waves and base HP system
  - [x] End-to-end test: survive 5+ waves with resource/typing feedback

---

## Immediate Next Steps (Post-Demo Polish)

- [x] **QUEUE-001** Letter-by-letter queue processing
  - [x] Adjust backlog pressure for letter queues
  - [x] Static word processing location at (400, 900) with conveyor effect
- [x] **UI-001** Tower selection and upgrade system
  - [x] `/` to enter tower selection mode, letter labels, upgrade menu
- [x] **CMD-001** Command mode for power users
  - [x] `:` to enter command mode, basic and advanced commands
- [x] **TITLE-001** Title screen and main menu
  - [x] MainMenuState, logo, background music, animated background, settings
- [x] **PREGAME-001** Pre-game setup and tutorial
  - [x] Character/difficulty selection, tutorial, typing test, mode selection

---

## Resource Loop & HUD

- [x] **R-001** Implement Gold/Wood/Stone/Iron structs
- [x] **R-002** Farmer, Lumberjack, Miner cooldowns produce resources
  - [x] Balance numbers in `config.json`
- [x] **HUD-001** Top bar resource icons (`G`, `W`, `S`, `I`, `M`)
- [x] **HUD-002** Show word processing queue with conveyor belt animation
- [x] **HUD-003** Tower selection overlay with letter labels
- [x] **TEST-RES** Integration test 3 min sim, resources > 0
  - [x] **T-001** Create test file `internal/game/resources_integration_test.go` with basic scaffold
  - [x] **T-002** Initialize a `Game` instance in test using default config
  - [x] **T-003** Simulate update loop for 3 minutes of game time (e.g., step through `Update` calls)
  - [x] **T-004** Access `ResourcePool` after simulation and assert each resource > 0
  - [x] **T-005** Ensure test runs headlessly and integrate into CI pipeline

---

## Tech Tree & Progression

- [x] **T-002** Tech tree parser + in-memory graph
  - [x] **T-002.1** Define Go structs matching YAML schema (TechNode, TechTree, effects/prereqs)
  - [x] **T-002.2** Implement YAML loader in Go (use yaml.v2)
  - [x] **T-002.3** Validate tech graph (detect cycles, missing prereqs)
  - [x] **T-002.4** Build in-memory graph (nodes by ID, adjacency)
  - [x] **T-002.5** Expose graph API (GetPrerequisites, UnlockOrder)
  - [x] **T-002.6** Write unit tests for parser and validation
- [x] **T-003** Keyboard UI for tech purchase (`/` search, `Enter` buy)
  - [x] **T-003.1** Add `techMenuOpen`, `searchBuffer`, and selection index fields to `Game`
  - [x] **T-003.2** Capture `/` key in `Input.Update` to toggle tech menu mode
  - [x] **T-003.3** Render tech menu overlay: list `TechNode.Name`, unlocked letters, and achievements
  - [x] **T-003.4** Implement search input handling: append typed chars and backspace to `searchBuffer`
  - [x] **T-003.5** Filter `TechTree.nodes` to
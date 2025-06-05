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
- [ ] **TEST-RES** Integration test 3 min sim, resources > 0
  - [ ] **T-001** Create test file `internal/game/resources_integration_test.go` with basic scaffold
  - [ ] **T-002** Initialize a `Game` instance in test using default config
  - [ ] **T-003** Simulate update loop for 3 minutes of game time (e.g., step through `Update` calls)
  - [ ] **T-004** Access `ResourcePool` after simulation and assert each resource > 0
  - [ ] **T-005** Ensure test runs headlessly and integrate into CI pipeline

---

## Tech Tree & Progression

- [ ] **T-002** Tech tree parser + in-memory graph
  - [ ] **T-002.1** Define Go structs matching YAML schema (TechNode, TechTree, effects/prereqs)
  - [ ] **T-002.2** Implement YAML loader in Go (use yaml.v2)
  - [ ] **T-002.3** Validate tech graph (detect cycles, missing prereqs)
  - [ ] **T-002.4** Build in-memory graph (nodes by ID, adjacency)
  - [ ] **T-002.5** Expose graph API (GetPrerequisites, UnlockOrder)
  - [ ] **T-002.6** Write unit tests for parser and validation
- [ ] **T-003** Keyboard UI for tech purchase (`/` search, `Enter` buy)
  - [ ] **T-003.1** Add `techMenuOpen`, `searchBuffer`, and selection index fields to `Game`
  - [ ] **T-003.2** Capture `/` key in `Input.Update` to toggle tech menu mode
  - [ ] **T-003.3** Render tech menu overlay: list `TechNode.Name`, unlocked letters, and achievements
  - [ ] **T-003.4** Implement search input handling: append typed chars and backspace to `searchBuffer`
  - [ ] **T-003.5** Filter `TechTree.nodes` by `searchBuffer` and update displayed list
  - [ ] **T-003.6** Handle Up/Down arrow keys to move highlight over filtered nodes
  - [ ] **T-003.7** Handle `Enter` key to purchase selected tech: check prerequisites/resources, call `UnlockNext`
  - [ ] **T-003.8** Write unit tests for tech menu: toggling, filtering, navigation, and purchase flow
- [ ] **SKILL-001** Global skill tree UI (offense, defense, typing, automation, utility)
- [ ] **SKILL-002** Integrate skill tree with building/tech systems
- [ ] **SKILL-003** Save/load skill tree state

---

## Military Prototype

- [ ] **M-001** Barracks building pushes unit words (letter-by-letter)
- [ ] **M-002** Footman entity (HP, dmg, speed)
- [ ] **M-003** Combat resolution attacker vs orc grunt
- [ ] **TEST-COMBAT** Unit kills grunt in <8 s with perfect typing

---

## Game States & Persistence

- [ ] **STATE-001** Proper game state management (MainMenu, PreGame, Playing, Paused, GameOver, Settings)
- [ ] **SAVE-001** Comprehensive save/load system (multiple slots, auto-save)

---

## Art, Audio & Polish

- [ ] **ART-001** 16×16 farmer, lumberjack, miner idle sprites
- [ ] **ART-002** Orc grunt walk + hit animation
- [ ] **ART-003** Tower upgrade visual indicators
- [ ] **SFX-001** Key-hit, crit, jam placeholders (chiptune)
- [ ] **SFX-002** Background music for different game states

---

## Typing Metrics & Minigames

- [ ] **MET-001** Capture per-word accuracy & time
- [ ] **MET-002** Rolling WPM (last 30 s)
- [ ] **UI-MET** Toggle stats panel (`Tab`)
- [ ] **MINIGAME-001** Speed trial mode
- [ ] **MINIGAME-002** Accuracy challenge mode
- [ ] **MINIGAME-003** Word puzzle/anagram mode
- [ ] **MINIGAME-004** Boss practice mode

---

## Advanced Systems (Future Sprints)

- [ ] **MINION-001** Minion summoning via typed commands
- [ ] **MINION-002** Minion AI and unique roles
- [ ] **MINION-003** Minion upgrades and management UI
- [ ] **IDLE-001** Auto-collection and offline progress
- [ ] **IDLE-002** Upgradable idle generators
- [ ] **IDLE-003** Prestige/reset system
- [ ] **FUZZ-001** Engine fuzz tester and robustness checks

---

*(Archive completed sprints to `TODO_ARCHIVE.md` when merged. Focus all effort on the "Core Gameplay Loop Demo" until it is fully playable and testable.)*

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
  - [ ] **SKILL-001.1** Define Go structs for skill tree nodes and categories (offense, defense, typing, automation, utility)
  - [ ] **SKILL-001.2** Implement in-memory skill tree structure and sample data
  - [ ] **SKILL-001.3** Add keyboard UI: open skill tree menu, navigate categories/nodes, show node details
  - [ ] **SKILL-001.4** Render skill tree overlay: display branches, highlight selected node, show unlock status
  - [ ] **SKILL-001.5** Implement skill unlock logic: check prerequisites/resources, update state on unlock
  - [ ] **SKILL-001.6** Integrate skill effects with game systems (e.g., global stat boosts, automation unlocks)
  - [ ] **SKILL-001.7** Write unit tests for skill tree navigation, unlocks, and effect application
  - [ ] **SKILL-001.8** Persist skill tree state in save/load system

---

## Military Prototype

- [ ] **M-001** Barracks building pushes unit words (letter-by-letter)
  - [ ] **T-001** Refactor Barracks to enqueue words to the global queue letter-by-letter
  - [ ] **T-002** Update queue manager to support partial word progress and per-letter validation
  - [ ] **T-003** Ensure Barracks cooldown only resets after full word is typed
  - [ ] **T-004** Add unit tests for Barracks letter-by-letter queue logic
  - [ ] **T-005** Integrate Barracks with HUD to show letter-by-letter progress

- [ ] **M-002** Footman entity (HP, dmg, speed)
  - [ ] **T-001** Define Footman struct with HP, damage, and speed fields
  - [ ] **T-002** Implement Footman movement and update logic
  - [ ] **T-003** Add Footman spawn logic to Barracks on word completion
  - [ ] **T-004** Write unit tests for Footman creation and state updates

- [ ] **M-003** Combat resolution attacker vs orc grunt
  - [ ] **T-001** Define OrcGrunt struct with HP and damage
  - [ ] **T-002** Implement combat logic between Footman and OrcGrunt
  - [ ] **T-003** Update military system to resolve combat each tick
  - [ ] **T-004** Add tests for combat outcomes and edge cases

- [ ] **TEST-COMBAT** Unit kills grunt in <8 s with perfect typing
  - [ ] **T-001** Create integration test simulating perfect typing input
  - [ ] **T-002** Spawn Footman and OrcGrunt, simulate combat loop
  - [ ] **T-003** Assert OrcGrunt is defeated in under 8 seconds
  - [ ] **T-004** Add test to CI pipeline

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

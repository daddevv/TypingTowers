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
  - [x] **T-003.5** Filter `TechTree.nodes` by `searchBuffer` and update displayed list
  - [x] **T-003.6** Handle Up/Down arrow keys to move highlight over filtered nodes
  - [x] **T-003.7** Handle `Enter` key to purchase selected tech: check prerequisites/resources, call `UnlockNext`
  - [x] **T-003.8** Write unit tests for tech menu: toggling, filtering, navigation, and purchase flow
- [ ] **SKILL-001** Global skill tree UI (offense, defense, typing, automation, utility)
  - [x] **SKILL-001.1** Define Go structs for skill tree nodes and categories (offense, defense, typing, automation, utility)
  - [x] **SKILL-001.2** Implement in-memory skill tree structure and sample data
  - [x] **SKILL-001.3** Add keyboard UI: open skill tree menu, navigate categories/nodes, show node details
  - [x] **SKILL-001.4** Render skill tree overlay: display branches, highlight selected node, show unlock status
  - [x] **SKILL-001.5** Implement skill unlock logic: check prerequisites/resources, update state on unlock
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
  - [ ] **T-001** Define a `GamePhase` enum/type covering all major states (MainMenu, PreGame, Playing, Paused, GameOver, Settings)
  - [ ] **T-002** Refactor main game loop to use `GamePhase` for state transitions
  - [ ] **T-003** Implement state transition logic (e.g., menu → pregame → playing → paused/gameover/settings)
  - [ ] **T-004** Ensure each state has a dedicated update and draw handler
  - [ ] **T-005** Add keyboard navigation and transitions between states (e.g., Esc to pause, Enter to continue)
  - [ ] **T-006** Write unit tests for state transitions and edge cases

- [ ] **SAVE-001** Comprehensive save/load system (multiple slots, auto-save)
  - [ ] **T-001** Design a save file structure supporting multiple slots and versioning
  - [ ] **T-002** Implement save/load logic for all core game data (resources, towers, buildings, tech, settings)
  - [ ] **T-003** Add auto-save functionality (e.g., after each wave or major event)
  - [ ] **T-004** Create a save/load menu UI for selecting slots
  - [ ] **T-005** Handle save/load errors and version mismatches gracefully
  - [ ] **T-006** Write integration tests for save/load, including slot switching and auto-save

---

## Art, Audio & Polish

- [ ] **ART-001** 16×16 farmer, lumberjack, miner idle sprites
  - [ ] **T-001** Define sprite requirements and animation frames for each character
  - [ ] **T-002** Create initial 16×16 pixel art for farmer idle sprite
  - [ ] **T-003** Create initial 16×16 pixel art for lumberjack idle sprite
  - [ ] **T-004** Create initial 16×16 pixel art for miner idle sprite
  - [ ] **T-005** Export spritesheets in required format (e.g., PNG)
  - [ ] **T-006** Integrate sprites into game asset pipeline
  - [ ] **T-007** Update entity rendering logic to use new idle sprites
  - [ ] **T-008** Write tests or visual checks for correct sprite display

- [ ] **ART-002** Orc grunt walk + hit animation
  - [ ] **T-001** Define animation frame count and timing for walk and hit actions
  - [ ] **T-002** Create 16×16 orc grunt walk animation frames
  - [ ] **T-003** Create 16×16 orc grunt hit animation frames
  - [ ] **T-004** Export animation spritesheets in required format
  - [ ] **T-005** Integrate orc grunt animations into asset pipeline
  - [ ] **T-006** Update orc grunt entity logic to trigger walk/hit animations
  - [ ] **T-007** Test animation playback in-game for smoothness and correctness

- [ ] **ART-003** Tower upgrade visual indicators
  - [ ] **T-001** Define visual indicator styles for each tower upgrade level
  - [ ] **T-002** Design and create upgrade icons or overlays for towers
  - [ ] **T-003** Export indicator assets in required format
  - [ ] **T-004** Integrate upgrade indicators into game asset pipeline
  - [ ] **T-005** Update tower rendering logic to display correct indicator based on upgrade level
  - [ ] **T-006** Test in-game to verify indicators update correctly on upgrade

- [ ] **SFX-001** Key-hit, crit, jam placeholders (chiptune)
  - [ ] **T-001** Define required sound effects (key-hit, crit, jam)
  - [ ] **T-002** Create or source chiptune placeholder sounds for each effect
  - [ ] **T-003** Convert sounds to required audio format (e.g., WAV, OGG)
  - [ ] **T-004** Integrate sound effects into game asset pipeline
  - [ ] **T-005** Update game logic to trigger sounds on key-hit, crit, and jam events
  - [ ] **T-006** Test in-game to ensure correct sound playback for each event

- [ ] **SFX-002** Background music for different game states
  - [ ] **T-001** Define music requirements for each game state (menu, gameplay, pause, etc.)
  - [ ] **T-002** Create or source chiptune background music tracks
  - [ ] **T-003** Convert music tracks to required audio format
  - [ ] **T-004** Integrate music tracks into game asset pipeline
  - [ ] **T-005** Update game state logic to play appropriate music for each state
  - [ ] **T-006** Test transitions between states to ensure smooth music changes

---

## Typing Metrics & Minigames

- [ ] **MET-001** Capture per-word accuracy & time
  - [ ] **T-001** Define data structures to store per-word accuracy and completion time.
  - [ ] **T-002** Update queue processing logic to record accuracy and time for each word.
  - [ ] **T-003** Store per-word stats in a history buffer or log.
  - [ ] **T-004** Add unit tests to verify per-word stats are captured correctly.
  - [ ] **T-005** Expose per-word stats to HUD or stats panel.

- [ ] **MET-002** Rolling WPM (last 30 s)
  - [ ] **T-001** Implement a time-based buffer to track recent typing events.
  - [ ] **T-002** Calculate rolling WPM using only events from the last 30 seconds.
  - [ ] **T-003** Add method to TypingStats or a new struct for rolling WPM.
  - [ ] **T-004** Write tests to verify rolling WPM calculation.
  - [ ] **T-005** Display rolling WPM in the HUD or stats panel.

- [ ] **UI-MET** Toggle stats panel (`Tab`)
  - [ ] **T-001** Add a boolean field to Game for stats panel visibility.
  - [ ] **T-002** Capture `Tab` key in Input handler to toggle stats panel.
  - [ ] **T-003** Implement HUD rendering logic for the stats panel.
  - [ ] **T-004** Display per-word stats, rolling WPM, and accuracy in the panel.
  - [ ] **T-005** Add tests to verify panel toggling and display.

- [ ] **MINIGAME-001** Speed trial mode
  - [ ] **T-001** Define rules and win/lose conditions for speed trial mode
  - [ ] **T-002** Implement a new game state for speed trial mode
  - [ ] **T-003** Add UI to select and start speed trial from main menu or pregame
  - [ ] **T-004** Generate and display a sequence of words for the trial
  - [ ] **T-005** Track time taken and words per minute during the trial
  - [ ] **T-006** Display results and high scores after completion
  - [ ] **T-007** Write unit and integration tests for speed trial logic

- [ ] **MINIGAME-002** Accuracy challenge mode
  - [ ] **T-001** Define accuracy challenge rules and scoring system
  - [ ] **T-002** Implement a new game state for accuracy challenge mode
  - [ ] **T-003** Add UI to select and start accuracy challenge from menu
  - [ ] **T-004** Generate a set of words with varying difficulty
  - [ ] **T-005** Track accuracy and provide feedback after each word
  - [ ] **T-006** Display final accuracy and award in-game rewards
  - [ ] **T-007** Write tests for accuracy tracking and challenge flow

- [ ] **MINIGAME-003** Word puzzle/anagram mode
  - [ ] **T-001** Design word puzzle/anagram rules and mechanics
  - [ ] **T-002** Implement a new game state for word puzzle/anagram mode
  - [ ] **T-003** Add UI to select and start puzzle mode from menu
  - [ ] **T-004** Generate and shuffle words for anagram puzzles
  - [ ] **T-005** Validate player solutions and provide hints or skips
  - [ ] **T-006** Track puzzle completion stats and reward player
  - [ ] **T-007** Write tests for puzzle generation and validation

- [ ] **MINIGAME-004** Boss practice mode
  - [ ] **T-001** Define boss practice rules and selectable bosses
  - [ ] **T-002** Implement a new game state for boss practice mode
  - [ ] **T-003** Add UI to select boss and start practice from menu
  - [ ] **T-004** Simulate boss fight with appropriate word/typing challenges
  - [ ] **T-005** Track player performance and provide feedback
  - [ ] **T-006** Display results and allow retry or exit
  - [ ] **T-007** Write tests for boss practice logic and transitions

---

## Advanced Systems (Future Sprints)

- [ ] **MINION-001** Minion summoning via typed commands
  - [ ] **T-001** Define minion types and their summon keywords
  - [ ] **T-002** Implement command parser for minion summoning (typed input triggers summon)
  - [ ] **T-003** Integrate minion summoning with the global queue and command mode
  - [ ] **T-004** Add resource cost and prerequisite checks for summoning minions
  - [ ] **T-005** Write unit tests for command parsing and minion summoning logic

- [ ] **MINION-002** Minion AI and unique roles
  - [ ] **T-001** Design and document unique roles/abilities for each minion type
  - [ ] **T-002** Implement basic AI behaviors (movement, targeting, action loop)
  - [ ] **T-003** Integrate role-specific logic (e.g., healer, gatherer, defender)
  - [ ] **T-004** Add minion state tracking and update logic to the game loop
  - [ ] **T-005** Write unit tests for AI behaviors and role execution

- [ ] **MINION-003** Minion upgrades and management UI
  - [ ] **T-001** Define upgrade paths and effects for each minion type
  - [ ] **T-002** Implement upgrade logic and resource cost handling
  - [ ] **T-003** Design keyboard-only UI for minion management (view, select, upgrade)
  - [ ] **T-004** Integrate minion management UI with the main HUD and input system
  - [ ] **T-005** Write unit tests for upgrade logic and UI navigation

- [ ] **IDLE-001** Auto-collection and offline progress
  - [ ] **T-001** Define data structures for auto-collection and offline progress tracking
  - [ ] **T-002** Implement auto-collection logic for eligible buildings/resources
  - [ ] **T-003** Add UI indicators for auto-collection status and progress
  - [ ] **T-004** Track last active timestamp and calculate offline gains on load
  - [ ] **T-005** Integrate auto-collection and offline progress with save/load system
  - [ ] **T-006** Write unit and integration tests for auto-collection and offline progress

- [ ] **IDLE-002** Upgradable idle generators
  - [ ] **T-001** Define idle generator building types and upgrade paths
  - [ ] **T-002** Implement generator upgrade logic and resource scaling
  - [ ] **T-003** Add keyboard UI for viewing and upgrading idle generators
  - [ ] **T-004** Integrate generator upgrades with auto-collection and resource systems
  - [ ] **T-005** Write tests for generator upgrades and resource output

- [ ] **IDLE-003** Prestige/reset system
  - [ ] **T-001** Design prestige/reset mechanics and rewards (e.g., permanent bonuses, unlocks)
  - [ ] **T-002** Implement prestige/reset logic and state reset flow
  - [ ] **T-003** Add UI for prestige/reset confirmation and reward display
  - [ ] **T-004** Integrate prestige system with save/load and progression tracking
  - [ ] **T-005** Write tests for prestige/reset flow and reward application

- [ ] **FUZZ-001** Engine fuzz tester and robustness checks
  - [ ] **T-001** Define fuzz testing goals and scenarios (e.g., input validation, edge cases)
  - [ ] **T-002** Implement a fuzz testing framework for the game engine
  - [ ] **T-003** Create test cases for common input errors and unexpected states
  - [ ] **T-004** Run fuzz tests and log any crashes or unexpected behavior
  - [ ] **T-005** Fix identified issues and improve engine robustness
  - [ ] **T-006** Write documentation on fuzz testing methodology and results

- [ ] **FUZZ-002** Stress test for performance and stability
  - [ ] **T-001** Define stress testing scenarios (e.g., high resource generation, many entities)
  - [ ] **T-002** Implement a stress testing framework to simulate heavy loads
  - [ ] **T-003** Run stress tests and monitor performance metrics (FPS, memory usage)
  - [ ] **T-004** Identify bottlenecks and optimize game systems for better performance
  - [ ] **T-005** Write documentation on stress testing methodology and results
  
- [ ] **FUZZ-003** Automated regression tests for core systems
  - [ ] **T-001** Define core systems and features to be covered by regression tests
  - [ ] **T-002** Implement automated regression tests for core systems
  - [ ] **T-003** Run regression tests and log any failures or issues
  - [ ] **T-004** Fix identified issues and improve system stability
  - [ ] **T-005** Write documentation on regression testing methodology and results

---

*(Archive completed sprints to `TODO_ARCHIVE.md` when merged. Focus all effort on the "Core Gameplay Loop Demo" until it is fully playable and testable.)*

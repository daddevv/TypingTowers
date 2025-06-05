# Sprint Backlog

> **Convention:** One feature or fix = one line item.  
> Sub-tasks are indented bullet points (checked as they land).

---

## Immediate Critical Fixes (Current Sprint)

- [ ] **CORE-001** Fix game loop timing and slow down all gameplay significantly
  - [x] Replace manual float64 cooldowns with CooldownTimer throughout
  - [x] Increase all intervals by 3-5x for manageable gameplay
  - [x] Slow down mob movement and spawning dramatically
- [ ] **QUEUE-001** Implement letter-by-letter queue processing
  - [x] Process individual letters instead of whole words
  - [ ] Adjust backlog pressure to accommodate larger letter queues
  - [ ] Static word processing location at (400, 900) with conveyor effect
- [ ] **UI-001** Tower selection and upgrade system
  - [x] Press `/` to enter tower selection mode
  - [x] Generate letter labels for each tower
  - [x] Press corresponding letter to select tower and open upgrade menu
- [ ] **CMD-001** Command mode for power users
  - [x] Press `:` to enter command mode
  - [x] Basic commands: pause, unpause, god, slow, fast
  - [ ] Advanced building/unit commands
- [ ] **TITLE-001** Add proper title screen and main menu
  - [ ] Create MainMenuState with options: New Game, Load Game, Settings, Quit
  - [ ] Title screen with game logo and background music
  - [ ] Animated background with floating letters/particles
  - [ ] Settings menu for audio, difficulty, key bindings
  - [ ] Save/load game selection screen with preview thumbnails
- [ ] **PREGAME-001** Pre-game setup and tutorial
  - [ ] Character/difficulty selection screen
  - [ ] Interactive tutorial covering basic mechanics
  - [ ] Typing test to calibrate difficulty settings
  - [ ] Campaign vs sandbox mode selection
  - [ ] Custom game settings (wave count, resources, etc.)

---

## Expanded Vision & New Systems

- Deep progression via a 100+ node skill tree (offense, defense, typing, automation, utility)
- Autonomous minions & heroes (summoned/managed by typing)
- Incremental & idle mechanics (auto-collection, offline progress, prestige/reset)
- Typing minigames (speed trials, accuracy challenges, word puzzles, boss practice)
- Multiple playstyle support (grind, optimize, idle, chaos)

---

### Next Sprint – Resource Loop & HUD

- [ ] **R-001** Implement Gold/Wood/Stone/Iron structs
- [ ] **R-002** Farmer, Lumberjack, Miner cooldowns produce resources
  - [ ] Balance numbers in `config.json`
- [ ] **HUD-001** Top bar resource icons (`G`, `W`, `S`, `I`, `M`)
- [ ] **HUD-002** Show word processing queue with conveyor belt animation
- [ ] **HUD-003** Tower selection overlay with letter labels
- [ ] **TEST-RES** Integration test 3 min sim, resources > 0

### Backlog #1 – Tech Tree Loader

 - [x] **T-001** YAML schema for node graph
   - [x] `type`, `cost`, `effects`, `prereqs`
- [ ] **T-002** Parser + in-memory graph
- [ ] **T-003** Keyboard UI for tech purchase (`/` search, `Enter` buy)

### Backlog #2 – Military Prototype

- [ ] **M-001** Barracks building pushes unit words (letter-by-letter)
- [ ] **M-002** Footman entity (HP, dmg, speed)
- [ ] **M-003** Combat resolution attacker vs orc grunt
- [ ] **TEST-COMBAT** Unit kills grunt in <8 s with perfect typing

### Backlog #3 – Game States & Flow

- [ ] **STATE-001** Implement proper game state management
  - [ ] States: MainMenu, PreGame, Playing, Paused, GameOver, Settings
  - [ ] Clean transitions between states
  - [ ] State-specific input handling and rendering
- [ ] **SAVE-001** Comprehensive save/load system
  - [ ] Save game state, tower configurations, progress
  - [ ] Multiple save slots with metadata
  - [ ] Auto-save functionality

### Backlog #4 – Art & Audio Pass 1

- [ ] **ART-001** 16×16 farmer, lumberjack, miner idle sprites
- [ ] **ART-002** Orc grunt walk + hit animation
- [ ] **ART-003** Tower upgrade visual indicators
- [ ] **SFX-001** Key-hit, crit, jam placeholders (chiptune)
- [ ] **SFX-002** Background music for different game states

### Backlog #5 – Continuous Typing Metrics

- [ ] **MET-001** Capture per-word accuracy & time
- [ ] **MET-002** Rolling WPM (last 30 s)
- [ ] **UI-MET** Toggle stats panel (`Tab`)

### Backlog #6 – Skill Tree & Progression

- [ ] **SKILL-001** Design and implement global skill tree UI
  - [ ] Node categories: offense, defense, typing, automation, utility
  - [ ] WPM/accuracy gating for advanced nodes
- [ ] **SKILL-002** Integrate skill tree with building/tech systems
- [ ] **SKILL-003** Save/load skill tree state

### Backlog #7 – Minions & Heroes

- [ ] **MINION-001** Implement minion summoning via typed commands
- [ ] **MINION-002** Minion AI and unique roles
- [ ] **MINION-003** Minion upgrades and management UI

### Backlog #8 – Idle & Incremental Mechanics

- [ ] **IDLE-001** Auto-collection and offline progress
- [ ] **IDLE-002** Upgradable idle generators
- [ ] **IDLE-003** Prestige/reset system

### Backlog #9 – Typing Minigames

- [ ] **MINIGAME-001** Speed trial mode
- [ ] **MINIGAME-002** Accuracy challenge mode
- [ ] **MINIGAME-003** Word puzzle/anagram mode
- [ ] **MINIGAME-004** Boss practice mode

### Backlog #10 – Fuzz Testing & Robustness

- [ ] **FUZZ-001** Implement engine fuzz tester
  - [ ] Generate randomized sequences of inputs/events (typing, build/deploy, mob movement, resource updates)
  - [ ] Integrate with Go's fuzzing support or go-fuzz to automate stress runs  
  - [ ] Detect and log boundary conditions, panics, invariant breaches, and unexpected states  
  - [ ] Verify that all bad states are handled gracefully without crashing  
  - [ ] Produce detailed trace output and reproducible minimised cases for debugging

*(Add new sprints at bottom; archive completed ones to `TODO_ARCHIVE.md` when merged.)*

---

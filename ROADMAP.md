# Sprint Backlog

> **Convention:** One feature or fix = one line item.  
> Sub-tasks are indented bullet points (checked as they land).

---

## Expanded Vision & New Systems

- Deep progression via a 100+ node skill tree (offense, defense, typing, automation, utility)
- Autonomous minions & heroes (summoned/managed by typing)
- Incremental & idle mechanics (auto-collection, offline progress, prestige/reset)
- Typing minigames (speed trials, accuracy challenges, word puzzles, boss practice)
- Multiple playstyle support (grind, optimize, idle, chaos)

---

### Completed – Core Gameplay Loop Integration

- [x] **INT-001** Integrate Gathering (Farmer) building with resource system (R-001, R-002)
  - [x] Farmer produces Gold and other relevant resources.
  - [x] Resource output correctly updates global resource pools.
- [x] **INT-002** Integrate Military (Barracks) building with unit spawning (M-001, M-002)
  - [x] Barracks spawns Footman entities upon word completion.
  - [x] Spawned units are tracked by the military system.
- [x] **INT-003** Integrate Shared Queue Manager with HUD and building inputs
  - [x] Display color-coded words (per building `family`) in the typing queue.
  - [x] Ensure words from Farmer and Barracks correctly populate the global queue.
  - [x] Typing validation and word dequeue logic functions as expected.
- [x] **INT-004** Integrate Per-Building Cooldown Timers with UI
  - [x] Display visual cooldown progress for Farmer and Barracks.
  - [x] Cooldowns reset correctly after word completion.
- [x] **INT-005** Integrate Back-Pressure Damage mechanic
  - [x] Player/base health decreases when active word queue exceeds threshold (e.g., ≥ 5 words).
  - [x] Link to player health system.
- [x] **INT-006** Integrate Jam State Visuals and Audio
  - [x] Implement red flash on mistype.
  - [x] Implement "clank" SFX placeholder on mistype.
- [x] **INT-007** Implement Letter Unlock System
  - [x] Create UI for viewing and purchasing letter unlocks as per `docs/LETTER_UNLOCKS.md`.
  - [x] Connect letter unlocks to word generation logic for buildings.
  - [x] Ensure resource costs for unlocks are deducted correctly.
- [x] **TEST-CORELOOP** End-to-end playtest of the core loop
  - [x] Verify resource gathering from Farmer.
  - [x] Verify letter unlocking and its effect on word generation.
  - [x] Verify unit spawning from Barracks.
  - [x] Verify queue mechanics: color-coding, back-pressure damage.
  - [x] Verify jam state feedback (visual and audio).
  - [x] Check overall game balance and flow for a 5-10 min session.

---

### Current Sprint – Resource Loop & HUD

- [ ] **R-001** Implement Gold/Wood/Stone/Iron structs
- [ ] **R-002** Farmer, Lumberjack, Miner cooldowns produce resources
  - [ ] Balance numbers in `config.json`
- [ ] **HUD-001** Top bar resource icons (`G`, `W`, `S`, `I`, `M`)
- [ ] **TEST-RES** Integration test 3 min sim, resources > 0

### Backlog #1 – Tech Tree Loader

- [ ] **T-001** YAML schema for node graph
  - [ ] `type`, `cost`, `effects`, `prereqs`
- [ ] **T-002** Parser + in-memory graph
- [ ] **T-003** Keyboard UI for tech purchase (`/` search, `Enter` buy)

### Backlog #2 – Military Prototype

- [ ] **M-001** Barracks building pushes unit words
- [ ] **M-002** Footman entity (HP, dmg, speed)
- [ ] **M-003** Combat resolution attacker vs orc grunt
- [ ] **TEST-COMBAT** Unit kills grunt in <8 s with perfect typing

### Backlog #3 – Art & Audio Pass 1

- [ ] **ART-001** 16×16 farmer, lumberjack, miner idle sprites
- [ ] **ART-002** Orc grunt walk + hit animation
- [ ] **SFX-001** Key-hit, crit, jam placeholders (chiptune)

### Backlog #4 – Continuous Typing Metrics

- [ ] **MET-001** Capture per-word accuracy & time
- [ ] **MET-002** Rolling WPM (last 30 s)
- [ ] **UI-MET** Toggle stats panel (`Tab`)

### Backlog #5 – Skill Tree & Progression

- [ ] **SKILL-001** Design and implement global skill tree UI
  - [ ] Node categories: offense, defense, typing, automation, utility
  - [ ] WPM/accuracy gating for advanced nodes
- [ ] **SKILL-002** Integrate skill tree with building/tech systems
- [ ] **SKILL-003** Save/load skill tree state

### Backlog #6 – Minions & Heroes

- [ ] **MINION-001** Implement minion summoning via typed commands
- [ ] **MINION-002** Minion AI and unique roles
- [ ] **MINION-003** Minion upgrades and management UI

### Backlog #7 – Idle & Incremental Mechanics

- [ ] **IDLE-001** Auto-collection and offline progress
- [ ] **IDLE-002** Upgradable idle generators
- [ ] **IDLE-003** Prestige/reset system

### Backlog #8 – Typing Minigames

- [ ] **MINIGAME-001** Speed trial mode
- [ ] **MINIGAME-002** Accuracy challenge mode
- [ ] **MINIGAME-003** Word puzzle/anagram mode
- [ ] **MINIGAME-004** Boss practice mode

### Backlog #9 – Fuzz Testing & Robustness

- [ ] **FUZZ-001** Implement engine fuzz tester
  - [ ] Generate randomized sequences of inputs/events (typing, build/deploy, mob movement, resource updates)
  - [ ] Integrate with Go’s fuzzing support or go-fuzz to automate stress runs  
  - [ ] Detect and log boundary conditions, panics, invariant breaches, and unexpected states  
  - [ ] Verify that all bad states are handled gracefully without crashing  
  - [ ] Produce detailed trace output and reproducible minimised cases for debugging

*(Add new sprints at bottom; archive completed ones to `TODO_ARCHIVE.md` when merged.)*

---

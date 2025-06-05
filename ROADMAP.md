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

## Prototype Sprint – Gathering & Military Core

- [x] **P-001** Implement Gathering (Farmer) building
  - [x] Farmer cooldown logic
  - [x] Farmer word generation from letter pool
  - [x] Resource output on word completion
- [ ] **P-002** Implement Military (Barracks) building
  - [ ] Barracks cooldown logic
  - [ ] Barracks word generation from letter pool
  - [ ] Unit spawn on word completion
- [ ] **P-003** Shared queue manager
  - [ ] Global FIFO queue structure
  - [ ] Enqueue from multiple buildings
  - [ ] Dequeue and typing validation
- [ ] **P-004** Per-building cooldown timers
  - [ ] Timer tick/update logic
  - [ ] Cooldown reset on word completion
- [ ] **P-005** Playtest word density
  - [ ] Simulate 5 min session, measure words/sec
  - [ ] Adjust cooldowns/word lengths for 1–1.5 words/sec target
- [ ] **P-006** Letter unlock order & cost curves
  - [ ] Draft full unlock order for all buildings
  - [ ] Define cost progression for each letter unlock
  - [ ] Document in `docs/LETTER_UNLOCKS.md`

---

## Current Sprint – Queue MVP Hardening

- [ ] **Q-001** Refactor global queue to support colour-coding per building
  - [ ] Add `family` field to `Word` struct
  - [ ] Palette map + ANSI tests
- [ ] **Q-002** Back-pressure damage when backlog ≥ 5
  - [ ] Unit test: enqueue 6 words, expect base HP −1
- [ ] **Q-003** Jam state visual
  - [ ] Red flash on mistype
  - [ ] Audio “clank” SFX placeholder

### Sprint #1 – Resource Loop & HUD

- [ ] **R-001** Implement Gold/Wood/Stone/Iron structs
- [ ] **R-002** Farmer, Lumberjack, Miner cooldowns produce resources
  - [ ] Balance numbers in `config.json`
- [ ] **HUD-001** Top bar resource icons (`G`, `W`, `S`, `I`, `M`)
- [ ] **TEST-RES** Integration test 3 min sim, resources > 0

### Sprint #2 – Tech Tree Loader

- [ ] **T-001** YAML schema for node graph
  - [ ] `type`, `cost`, `effects`, `prereqs`
- [ ] **T-002** Parser + in-memory graph
- [ ] **T-003** Keyboard UI for tech purchase (`/` search, `Enter` buy)

### Sprint #3 – Military Prototype

- [ ] **M-001** Barracks building pushes unit words
- [ ] **M-002** Footman entity (HP, dmg, speed)
- [ ] **M-003** Combat resolution attacker vs orc grunt
- [ ] **TEST-COMBAT** Unit kills grunt in <8 s with perfect typing

### Sprint #4 – Art & Audio Pass 1

- [ ] **ART-001** 16×16 farmer, lumberjack, miner idle sprites
- [ ] **ART-002** Orc grunt walk + hit animation
- [ ] **SFX-001** Key-hit, crit, jam placeholders (chiptune)

### Sprint #5 – Continuous Typing Metrics

- [ ] **MET-001** Capture per-word accuracy & time
- [ ] **MET-002** Rolling WPM (last 30 s)
- [ ] **UI-MET** Toggle stats panel (`Tab`)

### Sprint #6 – Skill Tree & Progression

- [ ] **SKILL-001** Design and implement global skill tree UI
  - [ ] Node categories: offense, defense, typing, automation, utility
  - [ ] WPM/accuracy gating for advanced nodes
- [ ] **SKILL-002** Integrate skill tree with building/tech systems
- [ ] **SKILL-003** Save/load skill tree state

### Sprint #7 – Minions & Heroes

- [ ] **MINION-001** Implement minion summoning via typed commands
- [ ] **MINION-002** Minion AI and unique roles
- [ ] **MINION-003** Minion upgrades and management UI

### Sprint #8 – Idle & Incremental Mechanics

- [ ] **IDLE-001** Auto-collection and offline progress
- [ ] **IDLE-002** Upgradable idle generators
- [ ] **IDLE-003** Prestige/reset system

### Sprint #9 – Typing Minigames

- [ ] **MINIGAME-001** Speed trial mode
- [ ] **MINIGAME-002** Accuracy challenge mode
- [ ] **MINIGAME-003** Word puzzle/anagram mode
- [ ] **MINIGAME-004** Boss practice mode

*(Add new sprints at bottom; archive completed ones to `TODO_ARCHIVE.md` when merged.)*

---

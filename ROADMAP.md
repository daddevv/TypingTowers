# Sprint Backlog

> **Convention:** One feature or fix = one line item.  
> Sub-tasks are indented bullet points (checked as they land).

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

*(Add new sprints at bottom; archive completed ones to `TODO_ARCHIVE.md` when merged.)*

---

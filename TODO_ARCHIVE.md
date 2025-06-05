# Sprint Archive

## Gathering & Military Core

- [x] **P-001** Implement Gathering (Farmer) building
  - [x] Farmer cooldown logic
  - [x] Farmer word generation from letter pool
  - [x] Resource output on word completion
- [x] **P-002** Implement Military (Barracks) building
  - [x] Barracks cooldown logic
  - [x] Barracks word generation from letter pool
  - [x] Unit spawn on word completion
- [x] **P-003** Shared queue manager
  - [x] Global FIFO queue structure
  - [x] Enqueue from multiple buildings
  - [x] Dequeue and typing validation
- [x] **P-004** Per-building cooldown timers
  - [x] Timer tick/update logic
  - [x] Cooldown reset on word completion
- [x] **P-005** Playtest word density
  - [x] Simulate 5 min session, measure words/sec
  - [x] Adjust cooldowns/word lengths for 1–1.5 words/sec target
- [x] **P-006** Letter unlock order & cost curves
  - [x] Draft full unlock order for all buildings
  - [x] Define cost progression for each letter unlock
  - [x] Document in `docs/LETTER_UNLOCKS.md`

---

## Queue MVP Hardening

- [x] **Q-001** Refactor global queue to support color-coding per building
  - [x] Add `family` field to `Word` struct
  - [x] Palette map + ANSI tests
- [x] **Q-002** Back-pressure damage when backlog ≥ 5
  - [x] Unit test: enqueue 6 words, expect base HP −1
- [x] **Q-003** Jam state visual
  - [x] Red flash on mistype
  - [x] Audio “clank” SFX placeholder

  ---

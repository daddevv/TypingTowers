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


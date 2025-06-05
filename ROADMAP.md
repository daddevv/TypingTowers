# Sprint Backlog

> **Convention:** One feature or fix = one line item.  
> Sub-tasks are indented bullet points (checked as they land).

---

## Core Gameplay Loop Demo (Highest Priority)

- [x] **SKILL-001** Global skill tree UI (offense, defense, typing, automation, utility)
  - [x] Define Go structs for skill tree nodes and categories
  - [x] Implement in-memory skill tree structure and sample data
  - [x] Add keyboard UI: open skill tree menu, navigate categories/nodes, show node details
  - [x] Render skill tree overlay: display branches, highlight selected node, show unlock status
  - [x] Implement skill unlock logic: check prerequisites/resources, update state on unlock
  - [x] Integrate skill effects with game systems (e.g., global stat boosts, automation unlocks)
  - [x] Write unit tests for skill tree navigation, unlocks, and effect application
  - [x] Persist skill tree state in save/load system

---

## Combat & Military

- [x] **TEST-COMBAT-EDGE** Simulation unit tests for all common and edge cases
  - [x] Footman survives after killing a single grunt (verify HP > 0)
  - [x] Footman dies if grunt damage is lethal (verify removal from military)
  - [x] Multiple Footmen vs multiple Grunts: all combinations (1v2, 2v1, 2v2)
  - [x] Simultaneous combat: overlapping units resolve damage correctly
  - [x] No combat occurs if units do not overlap (verify no HP loss)
  - [x] Dead units are removed from the military/orc lists immediately
  - [x] Units with 0 HP cannot attack or be attacked further
  - [x] Combat does not occur if either unit is already dead
  - [x] Test for correct handling of edge cases (e.g., both units die in same tick)
  - [x] Test for no panics or index errors when removing units during iteration
  - [x] Add all tests to CI pipeline

---

## Art, Audio & Polish

- [ ] **ART-001** 32x32 farmer, lumberjack, miner idle sprites
  - [ ] Define sprite requirements and animation frames for each character
  - [ ] Create initial 32x32 pixel art for each idle sprite
  - [ ] Export spritesheets in required format (e.g., PNG)
  - [ ] Integrate sprites into game asset pipeline
  - [ ] Update entity rendering logic to use new idle sprites
  - [ ] Test or visually check correct sprite display

- [ ] **ART-002** Orc grunt walk + hit animation
  - [ ] Define animation frame count and timing for walk and hit actions
  - [ ] Create 32x32 orc grunt walk/hit animation frames
  - [ ] Export animation spritesheets in required format
  - [ ] Integrate orc grunt animations into asset pipeline
  - [ ] Update orc grunt entity logic to trigger walk/hit animations
  - [ ] Test animation playback in-game

- [ ] **ART-003** Tower upgrade visual indicators
  - [ ] Define visual indicator styles for each tower upgrade level
  - [ ] Design and create upgrade icons or overlays for towers
  - [ ] Integrate upgrade indicators into game asset pipeline
  - [ ] Update tower rendering logic to display correct indicator based on upgrade level
  - [ ] Test in-game to verify indicators update correctly

- [ ] **SFX-001** Key-hit, crit, jam placeholders (chiptune)
  - [ ] Define required sound effects (key-hit, crit, jam)
  - [ ] Create/source chiptune placeholder sounds
  - [ ] Integrate sound effects into game asset pipeline
  - [ ] Update game logic to trigger sounds on key-hit, crit, and jam events
  - [ ] Test in-game for correct sound playback

- [ ] **SFX-002** Background music for different game states
  - [ ] Define music requirements for each game state
  - [ ] Create/source chiptune background music tracks
  - [ ] Integrate music tracks into game asset pipeline
  - [ ] Update game state logic to play appropriate music for each state
  - [ ] Test transitions for smooth music changes

---

## Typing Metrics & Minigames

- [x] **MET-001** Capture per-word accuracy & time
  - [x] Define data structures to store per-word accuracy and completion time
  - [x] Update queue processing logic to record accuracy and time for each word
  - [x] Store per-word stats in a history buffer or log
  - [x] Add unit tests to verify per-word stats are captured correctly
  - [x] Expose per-word stats to HUD or stats panel

- [x] **MET-002** Rolling WPM (last 30 s)
  - [x] Implement a time-based buffer to track recent typing events
  - [x] Calculate rolling WPM using only events from the last 30 seconds
  - [x] Add method to TypingStats or a new struct for rolling WPM
  - [x] Write tests to verify rolling WPM calculation
  - [x] Display rolling WPM in the HUD or stats panel

- [ ] **UI-MET** Toggle stats panel (`Tab`)
  - [ ] Add a boolean field to Game for stats panel visibility
  - [ ] Capture `Tab` key in Input handler to toggle stats panel
  - [ ] Implement HUD rendering logic for the stats panel
  - [ ] Display per-word stats, rolling WPM, and accuracy in the panel
  - [ ] Add tests to verify panel toggling and display

- [ ] **MINIGAME-001** Speed trial mode
- [ ] **MINIGAME-002** Accuracy challenge mode
- [ ] **MINIGAME-003** Word puzzle/anagram mode
- [ ] **MINIGAME-004** Boss practice mode

  *(See previous roadmap for detailed sub-tasks for each minigame.)*

---

## Advanced Systems (Future Sprints)

- [ ] **MINION-001** Minion summoning via typed commands
- [ ] **MINION-002** Minion AI and unique roles
- [ ] **MINION-003** Minion upgrades and management UI
- [ ] **IDLE-001** Auto-collection and offline progress
- [ ] **IDLE-002** Upgradable idle generators
- [ ] **IDLE-003** Prestige/reset system
- [ ] **FUZZ-001** Engine fuzz tester and robustness checks
- [ ] **FUZZ-002** Stress test for performance and stability
- [ ] **FUZZ-003** Automated regression tests for core systems

---

*Focus all effort on the "Core Gameplay Loop Demo" and combat/skill tree integration until fully playable and testable. Archive completed sprints to `TODO_ARCHIVE.md` when merged.*

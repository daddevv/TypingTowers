# TypeDefense v2 - Detailed TODO List

## Phase 1: Core Architecture Refactor - Single Game State

### 1.2. Refactor Core Game Systems

- [x] **Refactor Main Game Loop (`client/src/main.ts` or `GameScene.ts`):**
  - [x] Game loop should fetch current delta time and update it in `gameState`.
  - [x] Game loop should call update functions of various systems, passing `gameState` or relying on them to access it via `StateManager`.
- [x] **Refactor Input Handling (New `InputSystem` - `client/src/systems/InputSystem.ts`):**
  - [x] Centralize all keyboard/mouse event listeners here.
  - [x] On input, `InputSystem` updates `gameState` (e.g., `gameState.player.currentInput`, or triggers `stateManager.setGameStatus('paused')` on Escape key).
  - [x] Remove input handling logic scattered across different scenes/entities.
- [ ] **Refactor Entities (`Player`, `Mob`, `MobSpawner`):**
  - [ ] `Player`: Behavior (e.g., taking damage) driven by `gameState`.
  - [ ] `Mob`:
    - [ ] Data (word, position) stored in `gameState.mobs`.
    - [ ] Movement and logic updates based on its state in `gameState.mobs` and global `gameState` (e.g., delta time).
    - [ ] When a mob is hit/defeated, update its state in `gameState.mobs` or remove it via `StateManager`.
  - [ ] `MobSpawner`:
    - [ ] Logic driven by `gameState.mobSpawnerState` and `gameState.currentLevel.spawnRules`.
    - [ ] When spawning a mob, adds it to `gameState.mobs` via `StateManager`.
- [ ] **Refactor Managers (or integrate into Systems):**
  - [ ] `LevelManager`: Functionality largely moves to `ProgressionSystem` and `StateManager`. Data like unlocked levels stored in `gameState.progression`.
  - [ ] `FingerGroupManager`: Operates on typing data, potentially sourced from `gameState.player.currentInput` or events. Stores its stats within `gameState.curriculum.fingerGroupStats`.
- [ ] **Implement Pause/Unpause Functionality:**
  - [ ] Escape key (via `InputSystem`) toggles `gameState.gameStatus` between `playing` and `paused`.
  - [ ] All time-based updates (mob movement, spawning, timers) must check `gameState.gameStatus` and halt if `paused`.
  - [ ] Rendering of pause menu UI driven by `gameState.gameStatus === 'paused'`.
  - [ ] Resume/Quit options in pause menu update `gameState.gameStatus` accordingly.

### 1.3. Testing & Debugging for Core Architecture

- [ ] **Unit Tests for `StateManager`:** Test state initialization, updates, getters.
- [ ] **Unit Tests for Systems:** Test system logic with mocked `gameState`.
- [ ] **Console Debugging:** Continuously verify that `window.gameState` is accessible and reflects the current state accurately. Use it to debug issues during refactoring.

## Phase 2: Gameplay Enhancements & Content (Leveraging New Architecture)

### 2.1. Mob & Spawning System

- [ ] **Balance & Playtest:**
  - [ ] Playtest and balance mob spawn rates and speeds for all existing and new levels. Adjust configurations in level data files, read by `MobSpawner` via `gameState`.

### 2.2. Gameplay Loop & Feedback

- [ ] **Action-Challenge-Reward Loop:**
  - [ ] Implement audio cue when a mob is defeated (triggered by state change in `gameState.mobs`).
- [ ] **Visual Effects:**
  - [ ] Add camera shake and screen flash effects for wave completion and boss defeat (triggered by relevant `gameState` changes).
- [ ] **Audio:**
  - [ ] Integrate layered audio cues for typing, combos, wave clearances (managed by an `AudioSystem` reacting to `gameState`).
- [ ] **Scene Modularization (Phaser Scenes):**
  - [ ] Ensure all game states (preload, menu, game, game over, etc.) are handled by distinct Phaser Scenes that are activated/deactivated based on `gameState.gameStatus`.

### 2.3. Level & World Progression (World 1 Completion)

- [ ] **Create Level 1-4 (T/Y):**
  - [ ] Update level selection UI (`LevelSelectScene` reading from `gameState`) to include Level 1-4 and handle its unlocking based on `gameState.progression`.
  - [ ] Add tests: Level 1-4 unlocks correctly; uses correct word list (verify `WordGenerator` with level-specific letters from `gameState`).
- [ ] **Create Level 1-5 (V/M):**
  - [ ] Create word list JSON file (`fjghrutyvmWords.json` - already listed in `project_layout`).
  - [ ] Update level selection UI.
  - [ ] Add tests (unlocks, word list).
  - [ ] Playtest Level 1-5.
- [ ] **Create Level 1-6 (B/N):**
  - [ ] Create word list JSON file (`fjghrutyvmbnWords.json` - already listed).
  - [ ] Update level selection UI.
  - [ ] Add tests (unlocks, word list).
  - [ ] Playtest Level 1-6.
- [ ] **Create Level 1-7 (Boss):**
  - [ ] Create word list JSON file (`fjghrutyvmbn_bossWords.json` - already listed).
  - [ ] Update level selection UI.
  - [ ] Add tests (unlocks, word list).
  - [ ] Playtest Level 1-7.
- [ ] **General Level Progression & Win Conditions:**
  - [ ] Add tests to ensure Level 1-2 word generation (via `WordGenerator` using letters from `gameState.currentLevel.allowedLetters`) only uses "f", "j", "g", "h".
  - [ ] Test and verify Level 1-3 unlocks after 1-2 is marked complete in `gameState.progression`.
  - [ ] Implement robust win condition logic (e.g., defeat X enemies, survive Y waves) within `GameScene` or a `GameLogicSystem`, updating `gameState.levelStatus` to `complete`.
  - [ ] Ensure all new code is well-commented and tested, especially interactions with `gameState`.

### 2.4. Menu, World, and Level Selection UI & Logic

- [ ] **Refine UI Scenes:**
  - [ ] `WorldSelectScene`: Displays worlds based on `gameState.curriculum.worldConfig` and `gameState.progression.unlockedWorlds`.
  - [ ] `LevelSelectScene`: Displays levels for the selected world based on `gameState` and `gameState.progression.unlockedLevels`.
  - [ ] Ensure levels are clickable in `LevelSelectScene` to start the game (updates `gameState.gameStatus` to `playing` and loads `gameState.currentLevel`).
- [ ] **Navigation:**
  - [ ] Implement "Back" button in level complete UI (`LevelCompleteScene`) that returns to level selection (updates `gameState.gameStatus`).
  - [ ] Ensure "Continue" button in `LevelCompleteScene` advances to the next level/world or menu as appropriate (updates `gameState`).
  - [ ] Ensure keyboard shortcuts (Enter for continue, Esc for back) in `LevelCompleteScene` work via `InputSystem` updating `gameState`.
  - [ ] Test all navigation flows thoroughly.

## Phase 3: Polish & Expansion Prep

### 3.1. Asset Tasks (Visual & Audio Feedback)

- [ ] **Visual Enhancements:**
  - [ ] Create finger position guidance overlays for tutorials.
  - [ ] Implement letter highlighting system showing which finger should be used.
  - [ ] Design unique visual effects for each world/finger group.
  - [ ] Create thematic boss designs for each world.
- [ ] **Audio Enhancements:**
  - [ ] Implement distinctive sound effects for different finger groups.
  - [ ] Create celebratory animations and sounds for level completion.

### 3.2. Documentation & Testing

- [ ] **README Updates:**
  - [ ] Document new levels and keys as they are completed.
  - [ ] Document the v2 architecture, focusing on the global game state and its benefits.
  - [ ] Update testing instructions for v2.
- [ ] **Project Layout Documentation:**
  - [ ] Keep `.github/instructions/project_layout.instructions.md` updated as v2 evolves.
- [ ] **Comprehensive Testing:**
  - [ ] Aim for high test coverage for `StateManager`, `Systems`, and critical UI interactions.
  - [ ] Write E2E tests for key user flows (starting game, completing level, pausing, etc.).

## Phase 4: Future Expansion (Post v2 Core Stability)

- [ ] **New Content Worlds:**
  - [ ] Create Numbers & Symbols World with specialized levels.
- [ ] **New Game Modes:**
  - [ ] Implement Programming/Coding Mode with syntax exercises.
- [ ] **Advanced Challenges:**
  - [ ] Design advanced challenges combining all character types.

---

This detailed TODO should provide a solid roadmap for developing TypeDefense v2 with a more robust and manageable architecture. Remember to commit frequently and test thoroughly, especially during the core refactoring phase. Good luck!

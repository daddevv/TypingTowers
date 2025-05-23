## Phase 1: Headless/Decoupled Game Logic

### 1.1. Engine Decoupling & Render Abstraction

- [x] Audit all game logic to ensure no direct Phaser dependencies remain in the engine (core update loop, mob/player logic, win/loss, input, etc.).
- [x] Refactor any remaining UI or rendering logic out of the engine and into a dedicated render layer.
- [x] Define a clear interface/contract for the engine’s state and events, so any render system (Phaser, Three.js, etc.) can subscribe and react.
- [x] Add tests to verify the engine can run and be tested in a pure Node.js environment (no DOM, no Phaser).
- [x] Document the engine’s public API and how to integrate it with different renderers (Phaser, Three.js, headless).

### 1.2. Render Manager Abstraction

- [x] Design and implement a `RenderManager` abstraction that acts as the bridge between the engine and the chosen render library.
- [x] Refactor Phaser-specific rendering code in scenes (e.g., `GameScene`, `LevelMenuScene`, etc.) to use the `RenderManager` interface.
- [x] Ensure all rendering (mobs, player, UI, effects) is handled via the `RenderManager`, not directly in the engine or game logic.
- [x] Add tests/mocks for the `RenderManager` to verify correct rendering calls are made based on engine state.
- [x] Document how to implement a new renderer (e.g., Three.js) by providing a new `RenderManager` implementation.

### 1.3. Multi-Renderer Support & Experimentation

- [ ] Create a branch or prototype for a Three.js-based `RenderManager` to validate the abstraction.
- [ ] Add a build/runtime flag to select between Phaser and Three.js rendering backends.
- [ ] Playtest both renderers to ensure feature parity and performance.
- [ ] Update documentation to describe how to switch or extend renderers.

### 1.4. General

- [ ] Ensure all new/changed code is well-commented and covered by tests.
- [ ] Keep project layout and README documentation up to date with these architectural changes.

### 1.5. Headless Game Engine & API

- [x] Refactor core game logic (game loop, mob updates, win/loss conditions, input processing) into a headless, UI-agnostic module or service in `client/src/engine/`.
- [x] Add comprehensive unit and integration tests for headless gameplay (simulate full games, bots, and edge cases).
- [ ] Document how to use the headless engine and API for automated play/testing in the README and project layout docs.

# TypeDefense v2 - TODO List

## Phase 2: Gameplay Enhancements & Content

### 2.1. Mob & Spawning System

- [ ] Playtest and balance mob spawn rates and speeds for all levels. Adjust configurations in level data files, read by `MobSpawner` via `gameState`.

### 2.2. Gameplay Loop & Feedback

- [ ] Implement audio cue when a mob is defeated (triggered by state change in `gameState.mobs`).
- [ ] Add camera shake and screen flash effects for wave completion and boss defeat (triggered by relevant `gameState` changes).
- [x] Ensure keyboard shortcuts (Enter for continue, Esc for back) in `LevelCompleteScene` work via `InputSystem` updating `gameState`.

### 2.3. Level & World Progression (World 1 Completion)

- [ ] Level 1-4 (T/Y):
  - [ ] Update level selection UI (`LevelSelectScene`) to include Level 1-4 and handle unlocking via `gameState.progression`.
  - [ ] Add tests: Level 1-4 unlocks correctly; uses correct word list (verify `WordGenerator` with level-specific letters from `gameState`).
- [ ] Level 1-5 (V/M):
  - [ ] Create word list JSON file (`fjghrutyvmWords.json`).
  - [ ] Update level selection UI.
  - [ ] Add tests (unlocks, word list).
  - [ ] Playtest Level 1-5.
- [ ] Level 1-6 (B/N):
  - [ ] Create word list JSON file (`fjghrutyvmbnWords.json`).
  - [ ] Update level selection UI.
  - [ ] Add tests (unlocks, word list).
  - [ ] Playtest Level 1-6.
- [ ] Level 1-7 (Boss):
  - [ ] Create word list JSON file (`fjghrutyvmbn_bossWords.json`).
  - [ ] Update level selection UI.
  - [ ] Add tests (unlocks, word list).
  - [ ] Playtest Level 1-7.
- [ ] General Level Progression & Win Conditions:
  - [ ] Add tests to ensure Level 1-2 word generation only uses "f", "j", "g", "h".
  - [ ] Test Level 1-3 unlocks after 1-2 is marked complete in `gameState.progression`.
  - [ ] Implement robust win condition logic (e.g., defeat X enemies, survive Y waves) updating `gameState.levelStatus` to `complete`.
  - [ ] Ensure all new code is well-commented and tested, especially interactions with `gameState`.

### 2.4. Menu, World, and Level Selection UI & Logic

- [x] Implement "Back" button in level complete UI (`LevelCompleteScene`) that returns to level selection (updates `gameState.gameStatus`).
- [x] Ensure "Continue" button in `LevelCompleteScene` advances to the next level/world or menu as appropriate (updates `gameState`).
- [x] Test all navigation flows thoroughly.

## Phase 3: Polish & Expansion Prep

### 3.1. Visual & Audio Feedback

- [ ] Create finger position guidance overlays for tutorials.
- [ ] Implement letter highlighting system showing which finger should be used.
- [ ] Design unique visual effects for each world/finger group.
- [ ] Create thematic boss designs for each world.
- [ ] Implement distinctive sound effects for different finger groups.
- [ ] Create celebratory animations and sounds for level completion.

### 3.2. Documentation & Testing

- [ ] Document new levels and keys as they are completed in README.
- [ ] Document the v2 architecture, focusing on the global game state and its benefits.
- [ ] Update testing instructions for v2.
- [ ] Keep `.github/instructions/project_layout.instructions.md` updated as v2 evolves.
- [ ] Aim for high test coverage for `StateManager`, systems, and critical UI interactions.
- [ ] Write E2E tests for key user flows (starting game, completing level, pausing, etc.).

## Phase 4: Future Expansion

- [ ] Create Numbers & Symbols World with specialized levels.
- [ ] Implement Programming/Coding Mode with syntax exercises.
- [ ] Design advanced challenges combining all character types.

---

This TODO provides a focused, actionable roadmap for TypeDefense v2. Keep documentation and tests up-to-date as features are completed.

# TypeDefense v2 - TODO List

## Phase 2: Gameplay Enhancements & Content

### 2.1. Mob & Spawning System

- [ ] Playtest and balance mob spawn rates and speeds for all levels. Adjust configurations in level data files read by `MobSpawner` via `gameState`.

### 2.2. Gameplay Loop & Feedback

- [ ] Implement audio cue when a mob is defeated (triggered by state change in `gameState.mobs`).
- [ ] Add camera shake and screen flash effects for wave completion and boss defeat (triggered by relevant `gameState` changes).

### 2.3. Level & World Progression (World 1 Completion)

#### Level 1-4 (T/Y)

- [ ] Update level selection UI (`LevelSelectScene`) to include Level 1-4 and handle unlocking via `gameState.progression`.
- [ ] Add tests verifying Level 1-4 unlocks and that the correct word list is used.

#### Level 1-5 (V/M)

- [ ] Create word list JSON file (`fjghrutyvmWords.json`).
- [ ] Update level selection UI.
- [ ] Add tests (unlocks and word list).
- [ ] Playtest Level 1-5.

#### Level 1-6 (B/N)

- [ ] Create word list JSON file (`fjghrutyvmbnWords.json`).
- [ ] Update level selection UI.
- [ ] Add tests (unlocks and word list).
- [ ] Playtest Level 1-6.

#### Level 1-7 (Boss)

- [ ] Create word list JSON file (`fjghrutyvmbn_bossWords.json`).
- [ ] Update level selection UI.
- [ ] Add tests (unlocks and word list).
- [ ] Playtest Level 1-7.

#### General Level Progression & Win Conditions

- [x] Add tests to ensure Level 1-2 word generation uses only the specified letters ("f", "j", "g", "h").
- [x] Test that Level 1-3 unlocks after Level 1-2 is marked complete in `gameState.progression`.
- [x] Implement robust win condition logic (e.g., defeat X enemies, survive Y waves) updating `gameState.levelStatus`to complete.
- [x] Ensure all new code is well-commented and thoroughly tested.

## Phase 3: Polish & Expansion Prep

### 3.1. Visual & Audio Feedback

- [ ] Create finger position guidance overlays for tutorials.
- [ ] Implement a letter highlighting system to indicate the correct finger for each key.
- [ ] Design unique visual effects for each world/finger group.
- [ ] Create thematic boss designs for each world.
- [ ] Implement distinctive sound effects for different finger groups.
- [ ] Create celebratory animations and sounds for level completion.

### 3.2. Documentation & Testing

- [ ] Document new levels and keys as they are completed in the README.
- [ ] Document the v2 architecture with emphasis on global game state benefits.
- [ ] Update testing instructions for v2.
- [ ] Keep `.github/instructions/project_layout.instructions.md` updated as v2 evolves.
- [ ] Aim for high test coverage for `StateManager`, critical systems, and key UI interactions.
- [x] Write E2E tests for key user flows (e.g., starting game, completing level, pausing).
- [x] Create comprehensive game journey test with statistics tracking that simulates playing through all worlds and levels.

## Phase 4: Future Expansion

- [ ] Create Numbers & Symbols World with specialized levels.
- [ ] Implement Programming/Coding Mode with syntax exercises.
- [ ] Design advanced challenges combining all character types.

---

This TODO provides a focused, actionable roadmap for TypeDefense v2. Keep documentation and tests up-to-date as features are completed.

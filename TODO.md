# TODO

## Code Tasks

### Mob & Spawning System

- [ ] Playtest and balance spawn rate and speed for fun/challenge.

### Gameplay Loop & Feedback

- [ ] Implement action-challenge-reward loop with instant visual/audio feedback on word defeat.
  - [ ] Play an audio cue when a mob is defeated.
  - [ ] Integrate the loop into the main game update cycle.
- [ ] Add camera shake and screen flash effects for wave completion and boss defeat.
- [ ] Integrate layered audio cues for typing, combos, and wave clearances.
- [ ] Modularize game states into separate Phaser Scenes (preload, menu, waves, game over).

### Level & World Progression

- [ ] Create Level 1-4: Add T/Y (top row) with more complex patterns.
  - [x] Define Level 1-4 in the curriculum and world configuration files.
  - [x] Create a word list JSON file for Level 1-4 using only index finger letters plus T and Y, emphasizing T/Y usage.
  - [ ] Update the level selection menu to include Level 1-4 and ensure it unlocks after 1-3.
  - [ ] Add tests to verify Level 1-4 unlocks correctly and uses the correct word list.
  - [ ] Playtest Level 1-4 to ensure word patterns emphasize T/Y and gameplay is challenging but fair.
  - [ ] Update README.md to document the new level and its focus.
  - [ ] Update project layout documentation if new files are added.
  - [ ] Mark this task as complete when all subtasks are finished.
- [ ] Create Level 1-5: Add V/M (bottom row) with drills for downward reaches.
  - [x] Define Level 1-5 in the curriculum and world configuration files.
  - [ ] Create a word list JSON file for Level 1-5 using only index finger letters plus V and M, emphasizing V/M usage.
  - [ ] Update the level selection menu to include Level 1-5 and ensure it unlocks after 1-4.
  - [ ] Add tests to verify Level 1-5 unlocks correctly and uses the correct word list.
  - [ ] Playtest Level 1-5 to ensure word patterns emphasize V/M and gameplay is challenging but fair.
  - [ ] Update README.md to document the new level and its focus.
  - [ ] Update project layout documentation if new files are added.
  - [ ] Mark this task as complete when all subtasks are finished.
- [ ] Create Level 1-6: Add B/N (completing bottom row) with all index letters.
  - [x] Define Level 1-6 in the curriculum and world configuration files.
  - [ ] Create a word list JSON file for Level 1-6 using all index finger letters (F, J, G, H, R, U, T, Y, V, M, B, N).
  - [ ] Update the level selection menu to include Level 1-6 and ensure it unlocks after 1-5.
  - [ ] Add tests to verify Level 1-6 unlocks correctly and uses the correct word list.
  - [ ] Playtest Level 1-6 to ensure word patterns emphasize B/N and gameplay is challenging but fair.
  - [ ] Update README.md to document the new level and its focus.
  - [ ] Update project layout documentation if new files are added.
  - [ ] Mark this task as complete when all subtasks are finished.
- [ ] Create Level 1-7: Boss level using all index finger letters in combination.
  - [x] Define Level 1-7 in the curriculum and world configuration files.
  - [ ] Create a word list JSON file for Level 1-7 using all index finger letters (F, J, G, H, R, U, T, Y, V, M, B, N).
  - [ ] Update the level selection menu to include Level 1-7 and ensure it unlocks after 1-6.
  - [ ] Add tests to verify Level 1-7 unlocks correctly and uses the correct word list.
  - [ ] Playtest Level 1-7 to ensure word patterns are challenging and suitable for a boss level.
  - [ ] Update README.md to document the new boss level and its focus.
  - [ ] Update project layout documentation if new files are added.
  - [ ] Mark this task as complete when all subtasks are finished.
- [ ] Fix level progression so that completing a level unlocks and advances to the next.
  - [x] Review and update the logic in `LevelManager` to ensure that completing a level marks it as completed and unlocks the next level.
  - [x] Update the game flow in `GameScene` so that after a level is completed, the next level is automatically unlocked and the player is advanced to it (or returned to the menu if at the last level).
  - [x] Ensure the level selection menu (`LevelMenuScene`) reflects the unlocked status of levels immediately after completion.
    - [x] Update `LevelMenuScene` to refresh level lock/unlock status when returning from a completed level.
    - [x] Ensure the UI updates immediately, not just on scene reload.
    - [x] Test by completing a level and verifying the menu updates as expected.
  - [x] Add or update tests in `levelManager.test.ts` to verify that completing a level unlocks the next.
  - [x] Update `README.md` and `.github/instructions/project_layout.instructions.md` to document the improved progression system.
- [x] Ensure that level 1-2 uses the correct word list and includes "g" and "h" in generated words.
- [x] Verify that level 1-2 uses the correct word list (`fjghWords.json`) and that generated words include "g" and "h".
- [ ] Add or update tests to ensure that level 1-2 only uses "f", "j", "g", and "h" in generated words.
- [ ] Test and verify that after completing level 1-2, level 1-3 is unlocked and accessible in the level selection menu.
- [x] Add or update tests in `levelManager.test.ts` to confirm that completing 1-2 unlocks 1-3.
- [ ] Update `README.md` to document the correct word list usage and level unlocking behavior.
- [ ] Update `.github/instructions/project_layout.instructions.md` if any project structure changes are made.
- [ ] Implement logic in GameScene to detect when 50 enemies are defeated and trigger level completion.
- [ ] Update LevelManager to unlock and move to level 1-2 upon winning.
- [ ] Display a "Level Complete" message and transition to the next level.
- [ ] Ensure all new code is well-commented and tested.

### Menu, World, and Level Selection

- [ ] Create a world chooser scene that displays all worlds, showing locked/unlocked/completed status.
- [ ] When a player selects an unlocked world, show a level selector for that world, displaying all levels with their status.
- [ ] Allow the player to select an unlocked level to start the game.
- [ ] Integrate the new scenes into the game flow and update navigation logic.
- [ ] Refactor MenuScene to display only worlds (no levels).
- [ ] Implement LevelMenuScene to display levels for the selected world.
- [ ] Make levels clickable in LevelMenuScene to start the game at the selected level.
- [ ] Add a Back button to the level complete UI in GameScene that returns to level selection.
- [ ] Ensure the Continue button in GameScene advances to the next level, world, or menu as appropriate.
- [ ] Add keyboard shortcuts: Enter for continue, Escape for back, in the level complete UI.
- [ ] Ensure all navigation buttons and keyboard shortcuts work for progressing through the whole game.
- [x] Fix the "Continue" button in GameScene so it is always clickable after beating a level.
- [x] Update MenuScene so that worlds 2, 3, and 4 cannot be selected unless world 1-7 is completed (lock worlds until previous world is finished).
- [ ] Ensure all new code is well-commented and tested.

### Documentation & Testing

- [ ] Update README.md to mention new levels and keys.
- [ ] Update README.md to document the new win/level progression feature.
- [ ] Update README.md to document the new menu flow and particle effect.
- [ ] Update README.md with testing instructions.
- [ ] Update project layout documentation if new files/structure are added.
- [ ] Update project layout documentation if new directories are added.
- [ ] Update project layout documentation if any new files or structure are added.
- [ ] Update documentation and mark tasks as complete.

## Asset Tasks

### Visual & Audio Feedback

- [ ] Create finger position guidance overlays for tutorials.
- [ ] Implement letter highlighting system showing which finger should be used.
- [ ] Design unique visual effects for each world/finger group.
- [ ] Create thematic boss designs for each world.
- [ ] Implement distinctive sound effects for different finger groups.
- [ ] Create celebratory animations and sounds for level completion.

## Expansion Content (Future)

- [ ] Create Numbers & Symbols World with specialized levels.
- [ ] Implement Programming/Coding Mode with syntax exercises.
- [ ] Design advanced challenges combining all character types.

---
Contains AI-generated edits.

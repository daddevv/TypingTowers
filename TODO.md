# TODO

## Mob & Spawning System

- [x] Test and tune scaling parameters for smooth difficulty increase.
- [x] Add waves to the mob spawner (discrete waves, notification, and delay between waves).
- [ ] Playtest and balance spawn rate and speed for fun/challenge.

## Gameplay Loop & Feedback

- [x] Add tweened UI transitions for score pop-ups and wave notifications.
- [ ] Implement action-challenge-reward loop with instant visual/audio feedback on word defeat.
  - [ ] Play an audio cue when a mob is defeated.
  - [ ] Integrate the loop into the main game update cycle.
- [ ] Add camera shake and screen flash effects for wave completion and boss defeat.
- [ ] Integrate layered audio cues for typing, combos, and wave clearances.
- [ ] Modularize game states into separate Phaser Scenes (preload, menu, waves, game over).

## Architecture & Curriculum

- [ ] Write unit tests for the `WordGenerator` class covering all methods and edge cases.
- [ ] Integrate the `WordGenerator` into the game logic where word generation is required.
- [ ] Update documentation (`README.md` and project layout) to describe the new class and its usage.
- [ ] Implement a difficulty scaling system that adjusts spawn rates and word complexity.

## Level & World Progression

- [ ] Create Level 1-4: Add T/Y (top row) with more complex patterns.
- [ ] Create Level 1-5: Add V/M (bottom row) with drills for downward reaches.
- [ ] Create Level 1-6: Add B/N (completing bottom row) with all index letters.
- [ ] Create Level 1-7: Boss level using all index finger letters in combination.
- [ ] Update README.md to mention new levels and keys.
- [ ] Update project layout documentation if new files/structure are added.
- [ ] Ensure all new code is well-commented and tested.

## Visual & Audio Feedback

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

## Level Completion & Progression UX

- [ ] Implement logic in GameScene to detect when 50 enemies are defeated and trigger level completion.
- [ ] Update LevelManager to unlock and move to level 1-2 upon winning.
- [ ] Display a "Level Complete" message and transition to the next level.
- [ ] Update README.md to document the new win/level progression feature.
- [ ] Update README.md and project layout documentation to reflect the new menu and progression system.
- [ ] Update `README.md` with testing instructions.
- [ ] Update project layout documentation if new directories are added.
- [ ] Ensure all new code is well-commented and tested.

## Menu, World, and Level Selection

- [ ] Create a world chooser scene that displays all worlds, showing locked/unlocked/completed status.
- [ ] When a player selects an unlocked world, show a level selector for that world, displaying all levels with their status.
- [ ] Allow the player to select an unlocked level to start the game.
- [ ] Integrate the new scenes into the game flow and update navigation logic.
- [ ] Update documentation and mark tasks as complete.
- [ ] Refactor MenuScene to display only worlds (no levels).
- [ ] Implement LevelMenuScene to display levels for the selected world.
- [ ] Make levels clickable in LevelMenuScene to start the game at the selected level.
- [ ] Update README.md to document the new menu flow and particle effect.
- [ ] Update project layout documentation if needed.
- [ ] Add a Back button to the level complete UI in GameScene that returns to level selection.
- [ ] Ensure the Continue button in GameScene advances to the next level, world, or menu as appropriate.
- [ ] Add keyboard shortcuts: Enter for continue, Escape for back, in the level complete UI.
- [ ] Ensure all navigation buttons and keyboard shortcuts work for progressing through the whole game.
- [ ] Update README.md to document the new navigation and keyboard shortcuts.
- [ ] Update project layout documentation if any new files or structure are added.
- [ ] Ensure all new code is well-commented and tested.

---
Contains AI-generated edits.

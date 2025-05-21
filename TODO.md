# TODO

- [x] Set up Vite + TypeScript project structure for client
- [x] Create `client/src/main.ts` as the game entry point
- [x] Create a `GameScene` class in `client/src/scenes/GameScene.ts`
- [x] Implement a basic Phaser game instance in `main.ts` with GameScene
- [x] Set up the main game scene in `GameScene`
- [x] Set up the core game loop in `GameScene`
  - [x] Initialize and update the Player and InputHandler in GameScene
  - [x] Ensure the update method calls the necessary update functions each frame
- [x] Implement Player and InputHandler classes

## Mob & Spawning System Improvements

- [x] Refactor MobSpawner to always spawn mobs fully off-screen (right side)
- [x] Increase the number of mobs spawned per interval (support multiple spawns at once)
  - [x] Add a property to MobSpawner for mobs per interval
  - [x] Update MobSpawner's spawn logic to spawn multiple mobs per interval
  - [x] Update GameScene or config if needed to set mobs per interval
  - [x] Test that multiple mobs spawn at each interval
  - [x] Mark this task as complete
- [x] Add a base speed property to Mob and allow it to be set per spawn
- [x] Increase the default/base speed of mobs to make gameplay more challenging
- [x] Add y-position variation to mob spawner so mobs spawn at different vertical positions
- [x] Design a progression/scaling system (e.g., based on elapsed time, wave number, or score)
- [x] Update MobSpawner to adjust spawnInterval and mobBaseSpeed dynamically as the game progresses
- [x] Integrate scaling logic into GameScene or the main game loop
- [ ] Test and tune scaling parameters for smooth difficulty increase
- [ ] Update documentation and mark the task as complete
- [ ] Ensure all mobs move smoothly toward the player after spawning
- [ ] Playtest and balance spawn rate and speed for fun/challenge

## Gameplay Loop & Feedback

- [ ] Implement action-challenge-reward loop with instant visual/audio feedback on word defeat
  - [x] Add logic to match player input to mob words and trigger defeat when matched
  - [x] Implement instant visual feedback (e.g., flash, particle effect, or animation) when a word/mob is defeated
  - [ ] Play an audio cue when a mob is defeated
  - [ ] Integrate the loop into the main game update cycle
- [ ] Display real-time scores, combo multipliers, and particle bursts on each keystroke
- [ ] Add tweened UI transitions for score pop-ups and wave notifications
- [ ] Add camera shake and screen flash effects for wave completion and boss defeat
- [ ] Integrate layered audio cues for typing, combos, and wave clearances
- [ ] Modularize game states into separate Phaser Scenes (preload, menu, waves, game over)

## Architecture & Curriculum

- [ ] Write tests or usage examples for FingerGroupManager
- [ ] Implement a LevelManager to handle level transitions and progress tracking
- [ ] Design a WordGenerator class that creates appropriate words based on available letters
- [ ] Implement a difficulty scaling system that adjusts spawn rates and word complexity

## FingerGroupManager Integration

- [x] Instantiate FingerGroupManager in GameScene
- [x] On each key press, determine which finger group the key belongs to
- [x] Record the key press in FingerGroupManager with timing and correctness
- [x] Update the game loop in GameScene to call FingerGroupManager when player input occurs
- [x] (Optional) Expose progress/stats for UI or debugging

## Level & World Progression

- [x] Create Level 1-1: Basic F/J training with simple letter targets
  - [x] Define Level 1-1 configuration in curriculum (worldConfig)
  - [x] Add F/J-only word list for Level 1-1
  - [x] Implement Level 1-1 spawning logic in MobSpawner
  - [x] Integrate Level 1-1 into GameScene (load config, use word list)
  - [x] Playtest Level 1-1: verify only F/J targets spawn
  - [x] Mark Level 1-1 as completed in TODO.md
- [x] Create Level 1-2: Add G/H home row keys with simple combinations
  - [x] Define Level 1-2 configuration in curriculum (worldConfig)
  - [x] Create word list JSON for Level 1-2 (G/H/F/J simple combos)
  - [x] Update loadWordList.ts to support Level 1-2
  - [x] Mark Level 1-2 as completed in TODO.md
  - [ ] Update README.md to mention Level 1-2 and new keys
  - [ ] Update project layout documentation if new files/structure are added
  - [ ] Ensure all new code is well-commented and tested
- [x] Create Level 1-3: Add R/U (top row) with more letter combinations
- [ ] Create Level 1-4: Add T/Y (completing top row) with more complex patterns
- [ ] Create Level 1-5: Add V/M (bottom row) with drills for downward reaches
- [ ] Create Level 1-6: Add B/N (completing bottom row) with all index letters
- [ ] Create Level 1-7: Boss level using all index finger letters in combination

## Visual & Audio Feedback

- [ ] Create finger position guidance overlays for tutorials
- [ ] Implement letter highlighting system showing which finger should be used
- [ ] Design unique visual effects for each world/finger group
- [ ] Create thematic boss designs for each world
- [ ] Implement distinctive sound effects for different finger groups
- [ ] Create celebratory animations and sounds for level completion

## Expansion Content (Future)

- [ ] Create Numbers & Symbols World with specialized levels
- [ ] Implement Programming/Coding Mode with syntax exercises
- [ ] Design advanced challenges combining all character types

Contains AI-generated edits.

- [x] Refactor mob input handling to target the closest matching mob for each keypress, check others if not matched, and reset all mobs if no match (fix combo bug with multiple mobs)
- [x] Fix mob targeting system:
  - [x] If no mobs are targeted, keypresses identify a target (closest mob if multiple match)
  - [x] If a mob is targeted, next keypress is aimed at them; if correct, stay on target and advance letter; if miss, check other mobs
  - [x] Targeted mob is visually highlighted
  - [x] Matched letters are animated to inactive so player knows which letter is next
- [x] Add collision detection to mobs to prevent them from overlapping with each other
- [x] Add a win condition: defeat 50 enemies to win the level
- [ ] Implement logic in GameScene to detect when 50 enemies are defeated and trigger level completion
- [ ] Update LevelManager to unlock and move to level 1-2 upon winning
- [ ] Display a "Level Complete" message and transition to the next level
- [ ] Update README.md to document the new win/level progression feature

## World & Level Selection Menu

- [x] Design and implement a World/Level selection menu UI in the client (Phaser scene)
- [x] Implement logic to lock/unlock levels and worlds based on completion status
- [x] Store and retrieve level/world completion status in local storage
- [x] Integrate the menu into the game flow (entry point, transitions)
- [ ] Update README.md and project layout documentation to reflect the new menu and progression system

## Testing & Validation

- [x] Set up a test framework (Vitest) for the project
- [x] Configure the test environment
- [x] Add sample unit tests for a core entity (e.g., Mob or Player)
- [x] Add test scripts to `package.json`
- [ ] Update `README.md` with testing instructions
- [ ] Update project layout documentation if new directories are added
- [x] Add unit tests for FingerGroupManager (client/src/managers/fingerGroupManager.ts)
- [ ] Add unit tests for WordGenerator (client/src/utils/wordGenerator.ts)
- [ ] Add unit tests for InputHandler (client/src/entities/InputHandler.ts)
- [ ] Add unit tests for LevelManager (client/src/managers/levelManager.ts)
- [ ] Add unit tests for loadWordList utility (client/src/utils/loadWordList.ts)
- [ ] Add integration test for MobSpawner and mob spawning logic
- [ ] Ensure all new code is well-commented and tested

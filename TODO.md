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
- [ ] Increase the number of mobs spawned per interval (support multiple spawns at once)
- [ ] Add a base speed property to Mob and allow it to be set per spawn
- [ ] Increase the default/base speed of mobs to make gameplay more challenging
- [ ] Add support for scaling spawn rate and mob speed as the game progresses
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
- [ ] Create Level 1-2: Add G/H home row keys with simple combinations
- [ ] Create Level 1-3: Add R/U (top row) with more letter combinations
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

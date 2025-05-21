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
- [ ] Implement action-challenge-reward loop with instant visual/audio feedback on word defeat
  - [x] Add logic to match player input to mob words and trigger defeat when matched
  - [ ] Implement instant visual feedback (e.g., flash, particle effect, or animation) when a word/mob is defeated
  - [ ] Play an audio cue when a mob is defeated
  - [ ] Integrate the loop into the main game update cycle
- [ ] Display real-time scores, combo multipliers, and particle bursts on each keystroke
- [ ] Add tweened UI transitions for score pop-ups and wave notifications
- [ ] Add camera shake and screen flash effects for wave completion and boss defeat
- [ ] Integrate layered audio cues for typing, combos, and wave clearances
- [ ] Modularize game states into separate Phaser Scenes (preload, menu, waves, game over)
- [ ] Store wave configurations and word packs in JSON for data-driven design
- [ ] Implement escalating difficulty and unlockable word packs
- [ ] Add leaderboard and achievements integration
- [x] Implement Mob and MobSpawner classes for enemy logic
  - [x] Create Mob class in `client/src/entities/Mob.ts`
  - [x] Create MobSpawner class in `client/src/entities/MobSpawner.ts`
  - [x] Integrate Mob and MobSpawner into GameScene

## Finger-Group Curriculum Implementation

### Core Systems

- [x] Design the FingerGroupManager class interface and responsibilities
- [x] Implement the FingerGroupManager class in `client/src/managers/fingerGroupManager.ts`
- [ ] Integrate FingerGroupManager with the game loop to record finger usage and progress
- [ ] Add methods to retrieve progress and statistics for each finger group
- [ ] Write tests or usage examples for FingerGroupManager
- [ ] Implement a LevelManager to handle level transitions and progress tracking
- [ ] Design a WordGenerator class that creates appropriate words based on available letters
- [ ] Implement a difficulty scaling system that adjusts spawn rates and word complexity

### World 1: Index Fingers (F, G, R, T, V, B, J, H, Y, U, N, M)

- [ ] Create Level 1-1: Basic F/J training with simple letter targets
- [ ] Implement wave logic for Level 1-1 (spawn and manage basic mobs)
- [ ] Integrate Level 1-1 wave logic into the main game loop
- [ ] Provide basic score and combo feedback during Level 1-1
- [ ] Test the Level 1-1 loop for functionality and balance
- [ ] Create Level 1-2: Add G/H home row keys with simple combinations
- [ ] Create Level 1-3: Add R/U (top row) with more letter combinations
- [ ] Create Level 1-4: Add T/Y (completing top row) with more complex patterns
- [ ] Create Level 1-5: Add V/M (bottom row) with drills for downward reaches
- [ ] Create Level 1-6: Add B/N (completing bottom row) with all index letters
- [ ] Create Level 1-7: Boss level using all index finger letters in combination

### World 2: Middle Fingers (D, E, C, K, I, comma)

- [ ] Create Level 2-1: Basic D/K training (home row middle fingers)
- [ ] Create Level 2-2: Add E/I (top row) with mixed patterns
- [ ] Create Level 2-3: Add C/comma (bottom row) completing middle finger set
- [ ] Create Level 2-4: Practice patterns mixing index and middle finger keys
- [ ] Create Level 2-5: Boss level requiring alternating between index and middle finger letters

### World 3: Ring Fingers (S, W, X, L, O, period)

- [ ] Create Level 3-1: Basic S/L training (home row ring fingers)
- [ ] Create Level 3-2: Add W/O (top row) with mixed patterns
- [ ] Create Level 3-3: Add X/period (bottom row) completing ring finger set
- [ ] Create Level 3-4: Practice with combined index, middle, and ring finger patterns
- [ ] Create Level 3-5: Boss level focusing on ring finger letters with mixed patterns

### World 4: Pinky Fingers (A, Q, Z, semicolon, P, slash)

- [ ] Create Level 4-1: Basic A/semicolon training (home row pinky fingers)
- [ ] Create Level 4-2: Add Q/P (top row) with practice combinations
- [ ] Create Level 4-3: Add Z/slash (bottom row) completing all letter keys
- [ ] Create Level 4-4: Introduce shift key for capital letters
- [ ] Create Level 4-5: Practice sequences involving all fingers
- [ ] Create Level 4-6: Final boss using the full alphabet

### Post-Game Challenges

- [ ] Implement Mixed-Finger Challenge mode with progressively difficult sentences
- [ ] Create Endless Mode with infinite waves and random dictionary words
- [ ] Implement WPM tracking and personal best statistics
- [ ] Add Achievement system for mastering different finger groups and speed goals

### Visual & Audio Feedback

- [ ] Create finger position guidance overlays for tutorials
- [ ] Implement letter highlighting system showing which finger should be used
- [ ] Design unique visual effects for each world/finger group
- [ ] Create thematic boss designs for each world
- [ ] Implement distinctive sound effects for different finger groups
- [ ] Create celebratory animations and sounds for level completion

### Expansion Content (Future)

- [ ] Create Numbers & Symbols World with specialized levels
- [ ] Implement Programming/Coding Mode with syntax exercises
- [ ] Design advanced challenges combining all character types

Contains AI-generated edits.

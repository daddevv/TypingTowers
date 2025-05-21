# TODO

- [x] Set up Vite + TypeScript project structure for client
- [x] Create `client/src/main.ts` as the game entry point
- [x] Create a `GameScene` class in `client/src/scenes/GameScene.ts`
- [x] Implement a basic Phaser game instance in `main.ts` with GameScene
- [x] Set up the main game scene in `GameScene`
- [x] Set up the core game loop in `GameScene`
  - [x] Initialize and update the Player and InputHandler in GameScene
  - [x] Add placeholders for enemy spawning and word challenge logic in the update loop
  - [x] Ensure the update method calls the necessary update functions each frame
- [x] Implement Player and InputHandler classes
- [ ] Implement action-challenge-reward loop with instant visual/audio feedback on word defeat
- [ ] Display real-time scores, combo multipliers, and particle bursts on each keystroke
- [ ] Add tweened UI transitions for score pop-ups and wave notifications
- [ ] Add camera shake and screen flash effects for wave completion and boss defeat
- [ ] Integrate layered audio cues for typing, combos, and wave clearances
- [ ] Modularize game states into separate Phaser Scenes (preload, menu, waves, game over)
- [ ] Store wave configurations and word packs in JSON for data-driven design
- [ ] Implement escalating difficulty and unlockable word packs
- [ ] Add leaderboard and achievements integration

Contains AI-generated edits.

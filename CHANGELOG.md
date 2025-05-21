# Changelog

- Set up Vite + TypeScript project structure and entry points.
- Implemented core game loop, Player, and InputHandler.
- Refactored MobSpawner for off-screen spawning and multiple mobs per interval.
- Added base speed property to Mob and increased default speed.
- Added y-position variation and dynamic scaling to MobSpawner.
- Integrated scaling logic into GameScene.
- Added logic for matching player input to mob words and instant visual feedback on defeat.
- Instantiated and integrated FingerGroupManager.
- Created and playtested Level 1-1 and Level 1-2, including word lists and curriculum updates.
- Added collision detection to mobs and win condition (defeat 50 enemies).
- Designed and implemented World/Level selection menu with lock/unlock logic and local storage.
- Set up Vitest, added unit/integration tests for core entities and utilities.
- Added "Continue" button and Enter key handler to level complete screen.
- Implemented a real-time score system that updates on each correct keystroke.
- Added a combo multiplier that increases with consecutive correct keystrokes and resets on mistakes.
  - Combo multiplier variable added to game state.
  - Incremented combo multiplier on each correct keystroke.
  - Reset combo multiplier on incorrect keystroke.
  - Score calculation now uses combo multiplier.
  - Comments and documentation updated as needed.
  - Tests added/updated to cover combo logic.
- Displayed current score and combo multiplier in the game UI.
  - Score and combo text objects added to GameScene UI layer.
  - UI elements update in real-time as player types.
- Triggered a particle burst effect at the mob or letter location on every correct keystroke.
  - Enhanced particle system to emit at correct position on every correct keystroke.
  - Optimized effect for clarity and performance.
- Added/updated tests to verify score/combo UI and particle burst behavior.
  - Unit test: Combo/score logic updates UI as expected.
  - Integration test: Particle burst triggers on correct keystroke.
- Designed and implemented a `WordGenerator` class for generating words based on available letters.
  - Methods for random word generation, filtering valid words, and generating pseudo-words.
- Implemented a main menu scene with a "Play" button that navigates to the world chooser.
- Added a constant particle effect in the top-left corner of GameScene. (REMOVED: Feature reverted per user request; only burst effect remains)

Contains AI-generated edits.

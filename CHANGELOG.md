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
- Added or updated tests to verify score/combo UI and particle burst behavior.
  - Unit test: Combo/score logic updates UI as expected.
  - Integration test: Particle burst triggers on correct keystroke.

Contains AI-generated edits.

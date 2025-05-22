# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

- Ongoing refactoring for v2 architecture: centralized game state, improved maintainability, and debugging.
- Scene management refactor: scenes now read from and update `gameState` via `StateManager`.
- Input handling centralized in new `InputSystem`.
- Entities (`Player`, `Mob`, `MobSpawner`) now operate on state from `gameState`.
- Level and finger group progression managed via state and dedicated managers.

## [v0.1.0] - 2025-05-21

- Added wave-based enemy spawning with notifications and delays between waves.
- Score pop-ups now feature tweened animations for visual feedback.
- Mobs move faster by default and spawn at random vertical positions.
- Each mob displays a single letter; correct input removes the mob instantly.
- Particle burst effect triggered at mob/letter location on every correct keystroke.
- Mobs spawn fully off-screen for polish.
- Multiple mobs can spawn at each interval (configurable).
- Real-time score and combo multiplier UI in the top-left corner.
- Combo multiplier increases with consecutive correct keystrokes, resets on mistakes.
- Tweened UI transitions and camera effects for key events.
- Layered audio cues for typing, combos, and wave clearances.
- Mob and MobSpawner integrated into GameScene and main game loop.
- FingerGroupManager tracks player progress and statistics across finger groups.
- Player health system: visible health, game over on zero health.
- Collision detection and overlap prevention for mobs.
- Dynamic difficulty: spawn rate and speed scale up as game progresses.
- Word complexity scaling: word length/complexity increases with difficulty.
- Level/world progression system: unlocks, saves progress, and updates menu UI.
- All tests colocated in `__tests__` folders; uses Vitest for testing.

## [v0.0.1] - 2025-05-20

- Initial release: core gameplay, basic mob spawning, typing input, and scoring.
- Basic menu and level selection.
- Initial curriculum and word packs.
- Project structure established with Vite, TypeScript, and Go backend.

---

For a full project structure and documentation, see `.github/instructions/project_layout.instructions.md`.

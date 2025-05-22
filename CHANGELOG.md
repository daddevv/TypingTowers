# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

- Ongoing refactoring for v2 architecture: centralized game state, improved maintainability, and debugging.
- Scene management refactor: scenes now read from and update `gameState` via `StateManager`.
- Input handling centralized in new `InputSystem`.
- Entities (`Player`, `Mob`, `MobSpawner`) now operate on state from `gameState`.
- Level and finger group progression managed via state and dedicated managers.
- Initialized v2 branch for development.
- Reviewed and confirmed v2 goals: centralized game state, improved architecture, easier debugging, and console-inspectable state.
- Adopted new project layout per `project_layout.instructions.md`.
- Defined comprehensive `GameState` interface/type, including player, level, game, mob, spawner, UI, settings, curriculum/progression, and timing state.
- Implemented `StateManager` with:
  - Default/empty state initialization.
  - Immutable/copy getter for state.
  - Update functions for all major state parts.
  - Exposed `gameState` to `window` for debugging.
  - Event emitter for state changes.
  - Save/load to localStorage for progression.
- Refactored main game loop to:
  - Fetch and update delta time in state.
  - Call system update functions using state.
- Refactored scene management:
  - All scenes now read from `gameState.gameStatus` and relevant state.
  - Scene transitions are triggered by updating `gameState.gameStatus` via `stateManager.setGameStatus(...)`.
  - MainMenu, Menu (WorldSelect), and LevelMenu scenes render UI and handle navigation based on state, dispatching actions to StateManager.
  - All scene transitions and UI updates are reactive to state changes.
  - Added/updated tests for scene transitions and state-driven rendering.

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

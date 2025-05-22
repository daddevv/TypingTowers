---
applyTo: "**"
---

# Project Layout Documentation
This document outlines the folder structure and organization of the TypeDefense project.
Update this file as the project evolves to reflect any changes in the directory structure or organization.
It serves as a guide for developers to understand where to find specific components and how the project is structured.
It is also used to generate the project layout documentation in the GitHub repository.
It is important to keep this document up-to-date to ensure that all team members can easily navigate the project and understand its organization.

## Root Directory Structure

- `client/` - Frontend application code built with Vite and TypeScript
- `server/` - Backend Go HTTP server for serving static assets
- `.github/` - GitHub-specific files for workflows, templates, and documentation
- `.vscode/` - VS Code configuration for consistent development experience
- `dist/` - Built and bundled application output (generated)
- `node_modules/` - Frontend dependencies (generated)

## Client Directory (`client/`)

- `src/` - Source code for the game
  - `assets/` - Game assets organized by type
    - `audio/` - Sound effects and music files
    - `images/` - Sprites, backgrounds, and UI elements
  - `entities/` - Game entity classes (Mob, Player, InputHandler, etc.)
    - `Mob.ts` - Supports multi-letter words, targeting/highlighting, and matched letter animation. **Default mob speed increased for more challenge.**
    - `MobSpawner.ts` - Now supports spawning multiple mobs per interval (configurable), prevents mobs from overlapping using collision detection/repulsion, **spawns mobs at random vertical positions, and supports dynamic scaling of spawn rate and mob speed as the game progresses. Scaling logic is covered by unit tests for smooth difficulty progression.**
    - `__tests__/` - Unit and integration tests for entity classes (e.g., Mob, MobSpawner, InputHandler)
      - `ComboSystem.unit.test.ts` - Unit tests for combo multiplier and score logic
  - `scenes/` - Phaser scene classes (modular: preload, menu, waves, game over, etc.)
    - `BootScene.ts` - Initializes StateManager, loads essential assets, and sets `gameStatus` to `'mainMenu'`. Listens for `gameStatus` changes and switches scenes accordingly. **All scene transitions are now state-driven via `gameState.gameStatus` and `StateManager`.**
    - `GameScene.ts` - Main game loop. Updates global `gameState` with delta time and timestamp each frame using `StateManager`. All core systems (Player, MobSpawner, UI) are updated here and can access the latest state via the singleton. The main game loop is now fully state-driven: it updates `gameState` with the current delta time and timestamp, and all system updates are performed using the current state, enabling robust, testable, and reactive gameplay logic.
    - `MenuScene.ts` - World and level selection menu. Handles locking/unlocking of levels based on completion status, with progress stored in local storage. **UI and navigation are now reactive to state and use `StateManager` for transitions.**
    - `LevelMenuScene.ts` - Level selection for a world. **UI and navigation are now reactive to state and use `StateManager` for transitions.**
    - `MainMenuScene.ts` - Main menu. **Navigation to world select is now triggered by updating `gameState.gameStatus` via `StateManager`.**
    - `WorldSelectionScene.ts` - World selection menu. **Displays worlds based on `gameState.curriculum.worldConfig` and `gameState.progression.unlockedWorlds`. Only unlocked worlds are interactive.**
    - `__tests__/` - Integration tests for scenes
      - `GameScene.combo.integration.test.ts` - Integration tests for combo and score in the game scene
  - `levels/` - Level definitions and configurations
    - `world1/` - Index finger levels
    - `world2/` - Middle finger levels
    - `world3/` - Ring finger levels
    - `world4/` - Pinky finger levels
  - `curriculum/` - Finger group definitions and level progression structure
    - `fingerGroups.ts` - Mapping of fingers to keys for curriculum structure
    - `worldConfig.ts` - Configuration for each world and its levels
  - `wordpacks/` - Unlockable word packs and JSON data for waves (planned)
    - `fjWords.json` - Words using only F and J (Level 1-1)
    - `fjghWords.json` - Words using F, J, G, H (Level 1-2)
    - `indexFingerWords.json` - Words using only index finger letters
    - `middleFingerWords.json` - Words using middle finger letters
    - `ringFingerWords.json` - Words using ring finger letters
    - `pinkyFingerWords.json` - Words using pinky finger letters
    - `combinedWords.json` - Words using mixed finger combinations
    - `fjghrutyvmWords.json` - Words using F, J, G, H, R, U, T, Y, V, M (Levels 1-4, 1-5)
    - `fjghrutyvmbnWords.json` - Words using all index finger letters (Level 1-6)
    - `fjghrutyvmbn_bossWords.json` - Boss words using all index finger letters (Level 1-7)
  - `managers/` - (Deprecated) Old manager classes for level and finger group progression. Logic is now handled by StateManager and stored in `gameState`.
  - `utils/` - Utility functions and helper classes
    - `wordGenerator.ts` - Generates appropriate words based on available letters, **and dynamically increases word length/complexity as difficulty increases. Scaling logic is covered by unit tests.**
    - `__tests__/` - Unit tests for utility modules (e.g., wordGenerator, loadWordList)
  - `config/` - Game configuration constants (including data-driven wave/word settings)
  - `types/` - TypeScript type definitions
  - `main.ts` - Application entry point
  - `state/` - Centralized game state and state management
    - `gameState.ts` - Defines the global `GameState` interface and default state object for the entire game. All gameplay, UI, and progression data is stored here for easy debugging and testing.
    - `stateManager.ts` - Singleton StateManager class for reading/updating game state, event subscription, and save/load. Exposes `window.gameState` for debugging and is the only way to mutate state.

## Server Directory (`server/`)

- `main.go` - Go HTTP server entry point
- `static/` - Static files for development (may be symlinked to dist)

## GitHub Directory (`.github/`)

- `instructions/` - AI assistance and project guideline files
  - `ai-comments.instructions.md` - AI-generated code comment guidelines
  - `project_layout.instructions.md` - This file
- `prompts/` - Reusable Copilot prompt files
  - `*.prompt.md` - Task-specific context for Copilot

## VSCode Directory (`.vscode/`)

- `settings.json` - Editor settings for consistent formatting
- `extensions.json` - Recommended extensions
- `launch.json` - Debugging configurations

## Build Output (`dist/`)

- Contains compiled and bundled frontend assets
- Served by the Go backend server

## Future Expansion Directories

- `tests/` - Unit and integration tests
- `docs/` - Documentation beyond README
- `leaderboard/` - Online leaderboard integration (planned)
- `achievements/` - Achievement tracking and display (planned)

## Testing

- All tests are located in `__tests__` subdirectories next to the code under test.
- The project uses Vitest for running tests (`npm run test` in the `client` directory).
- Add new tests in the appropriate `__tests__` folder for each module.
- **Unit tests for `StateManager` are located in `client/src/state/__tests__/stateManager.unit.test.ts`. These tests cover state initialization, updates, and getters.**
- **System logic is tested with mocked `gameState` in system unit tests (see `client/src/systems/__tests__/InputSystem.test.ts`).**

## GameScene Updates

- The `GameScene` now includes score and combo UI elements in the top-left corner, updating in real-time as the player types. The UI is implemented using Phaser's `Text` objects and is visible throughout gameplay.
- **A particle burst effect is now triggered at the mob or letter location on every correct keystroke, providing instant visual feedback.**
- **The level completion screen features Continue (Enter) and Back (Esc) buttons for seamless navigation through levels and worlds.**

## Level & World Progression System

- Completing a level marks it as completed and unlocks the next level.
- The game automatically advances to the next unlocked level, or returns to the menu if the last level is completed.
- The level selection menu (`LevelMenuScene`) updates immediately to reflect unlocked/completed status after a level is finished.
- Level progression and unlock logic is managed by the `LevelManager` class (`client/src/managers/levelManager.ts`).
- Tests in `levelManager.test.ts` verify correct progression and unlocking.
- Player progress is saved for persistent advancement.

## Scene Management (v2 Architecture)
- Scenes are activated/deactivated based on `gameState.gameStatus`.
- `BootScene` is the entry point and listens for `gameStatus` changes to switch scenes.
- All navigation and transitions should use `stateManager.setGameStatus(...)` instead of direct `scene.start()` calls.

## Input Handling
- All input (keyboard/mouse) is handled by `client/src/systems/InputSystem.ts`, which updates `gameState` via `StateManager`. Scenes and entities do not register input listeners directly.

Contains AI-generated edits.

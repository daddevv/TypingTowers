# TypeDefense

TypeDefense is a web-based typing game designed to help players improve their typing skills through engaging gameplay. Players defend against waves of enemies by typing words or letters displayed above each enemy before they reach the player.

## Table of Contents

- [Features](#features)
- [Headless Game Engine](#headless-game-engine)
- [Screenshots](#screenshots)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Changelog](#changelog)

## Features

- Fast-paced, wave-based typing defense gameplay
- Enemies spawn in waves and move toward the player
- Defeat enemies by typing their associated words or letters
- Real-time score and combo multiplier display
- Particle effects and visual feedback for correct keystrokes
- The level completion screen features Continue (Enter) and Back (Esc) buttons for seamless navigation through levels and worlds.
  - **Keyboard shortcuts for Continue (Enter) and Back (Esc) are now handled via the centralized InputSystem, which updates gameState.**
- Dynamic difficulty scaling (spawn rate, speed, word complexity)
- Level and world progression system with unlockable content
  - **World selection menu now displays worlds based on `gameState.curriculum.worldConfig` and `gameState.progression.unlockedWorlds`.**
  - **World 1 is always unlocked, so players can always access Level 1-1.**
  - **Level 1-1 is always unlocked, regardless of player progress.**
- Player health system and game over state
- Curriculum-based finger group tracking for typing improvement
- Built with TypeScript, Vite, and Phaser
- Centralized game state managed by StateManager
- Level progression and finger group stats are now fully managed in the global game state (`gameState`) via StateManager, replacing the old LevelManager and FingerGroupManager. All progression and curriculum data is now state-driven and accessible for robust, testable gameplay and UI logic.
- The main game loop now updates delta time and timestamp in the global game state each frame, and all core systems (Player, MobSpawner, UI, etc.) are updated using the current state via StateManager. This enables robust, testable, and state-driven gameplay logic.
- Scenes are now managed via `gameState.gameStatus` and all transitions are triggered by updating state via StateManager, not by direct scene switching.
- MainMenu, Menu (WorldSelect), and LevelMenu scenes now render UI and handle navigation based on state, and listen for state changes to update or stop themselves.
- All scene transitions and UI updates are reactive to changes in `gameState.gameStatus`.
- The core game engine (`client/src/engine/HeadlessGameEngine.ts`) is fully UI-agnostic and contains no rendering or UI logic. All rendering, UI, and effects are handled in the render layer (Phaser scenes or other renderer implementations) based on the current game state.
- All input (keyboard/mouse) is now handled centrally by the InputSystem, which updates gameState via StateManager. Scenes and entities no longer register input listeners directly.
- All core entities (Player, Mob, MobSpawner) are now fully state-driven: their data and logic are managed via the global game state (`gameState`) and all updates are performed through the `StateManager`. This enables robust, testable, and reactive gameplay logic, and allows for easy debugging and inspection of all entity state at runtime.

## Headless Game Engine

The core game logic is implemented in a headless, UI-agnostic module: `client/src/engine/HeadlessGameEngine.ts`.

- **No rendering or UI dependencies:** The engine contains only gameplay logic and state management. It does not import Phaser, Three.js, or any DOM APIs.
- **Programmatic API:** The engine exposes the following interface for integration and testing:

  ```typescript
  export interface IGameEngine {
    step(delta: number, timestamp?: number): void;
    injectInput(input: string): void;
    getState(): GameState;
    on(event: string, handler: (...args: any[]) => void): void;
    off(event: string, handler: (...args: any[]) => void): void;
    reset(): void;
  }
  ```

- **Usage:**
  - Call `step(delta)` to advance the simulation.
  - Use `injectInput(input)` to simulate player keystrokes.
  - Retrieve the current state with `getState()`.
  - Listen for state changes or events with `on(event, handler)`.
  - Reset the engine with `reset()`.
- **Testing:** The engine is fully testable in Node.js with no DOM or rendering dependencies. See `client/src/engine/HeadlessGameEngine.unit.test.ts` for usage examples and integration tests.

## Screenshots

<!-- Add screenshots or animated GIFs here to showcase gameplay -->

## Installation

### Prerequisites

- Node.js (v16+ recommended)
- npm (v8+ recommended)
- Go (for backend server, optional for frontend development)

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/type-defense.git
   cd type-defense
   ```

2. Install frontend dependencies:

   ```bash
   cd client
   npm install
   ```

3. (Optional) Build the frontend:

   ```bash
   npm run build
   ```

4. (Optional) Run the Go backend server:

   ```bash
   cd ../server
   go run main.go
   ```

## Usage

- To start the development server for the frontend:

  ```bash
  cd client
  npm run dev
  ```

- Open your browser to the provided local address (usually <http://localhost:5173>).
- For production, build the frontend and serve the `dist/` directory using the Go backend or any static file server.

### Selecting the Render Backend (Phaser or Three.js)

TypeDefense supports multiple rendering backends via a runtime flag. By default, Phaser is used.  
To use the Three.js renderer (if implemented), set the `RENDER_BACKEND` flag before loading the app:

- **In development:**  
  Set the environment variable before running Vite:
  ```bash
  VITE_RENDER_BACKEND=three npm run dev
  ```
- **At runtime (in browser console):**
  ```js
  window.RENDER_BACKEND = 'three';
  // Then reload the page
  ```

- **Default:**  
  If no flag is set, Phaser is used.

The selected renderer is attached to `window.renderManager` and used by all scenes.

## Project Structure

See `.github/instructions/project_layout.instructions.md` for a detailed breakdown of the project structure and file organization.

## Testing

TypeDefense uses [Vitest](https://vitest.dev/) as the primary test runner, with [jsdom](https://github.com/jsdom/jsdom) providing a browser-like environment for running tests in Node.js.

**Browser API Polyfills:**

- The test environment automatically polyfills browser APIs using `jsdom`, `node-canvas` (for Canvas 2D), and `headless-gl` (for WebGL) in `client/setupTests.ts`.
- This allows headless testing of Phaser, PixiJS, and Three.js code in CI and local development.
- No manual setup is required; all shims are registered before tests run.

To run all tests:

```sh
npm run test
```

- All unit and integration tests are located in `__tests__` subdirectories next to the code under test.
- The test environment is configured in `client/vitest.config.ts`.
- Coverage reports are generated in text, JSON, and HTML formats.

## Contributing

Contributions are welcome! Please open issues or pull requests for bug fixes, new features, or documentation improvements.

## License

[MIT](LICENSE)

## Changelog

See [CHANGELOG.md](./CHANGELOG.md) for a detailed list of changes and feature history.

## Headless Game Engine API & Renderer Integration

The core game logic is implemented in a headless, UI-agnostic engine (`client/src/engine/HeadlessGameEngine.ts`). This engine exposes a public API for stepping the game, injecting input, retrieving state, and subscribing to events. Renderers (Phaser, Three.js, CLI, etc.) interact with the engine only via this API.

### Engine Interface

```typescript
export interface IGameEngine {
  step(delta: number, timestamp?: number): void;
  injectInput(input: string): void;
  getState(): GameState;
  on(event: string, handler: (...args: any[]) => void): void;
  off(event: string, handler: (...args: any[]) => void): void;
  reset(): void;
}
```

### Usage Example

```typescript
import HeadlessGameEngine from './client/src/engine/HeadlessGameEngine';

const engine = new HeadlessGameEngine();

// Step the game forward (e.g., in a requestAnimationFrame loop)
engine.step(16);

// Inject player input (simulate typing)
engine.injectInput('f');

// Subscribe to game events (e.g., for rendering or UI updates)
engine.on('gameStatusChanged', (status) => {
  // React to status change (e.g., switch scenes)
});

// Get the current game state for rendering
const state = engine.getState();
```

### Renderer Integration

- **Phaser/Three.js/Other Renderers:**  
  Renderers should subscribe to engine events and render based on the current state. All input should be injected via `injectInput()`.  
  Scene transitions and UI updates should be triggered by listening to state changes (e.g., `gameStatusChanged`).

- **No Direct State Mutation:**  
  All state updates must go through the engine/state manager API. Do not mutate state directly in the renderer.

- **Testing/Automation:**  
  The engine can be used in Node.js for automated play, bots, and CLI/testing scenarios. No DOM or rendering dependencies are required.

See `client/src/engine/HeadlessGameEngine.unit.test.ts` for comprehensive usage and integration tests.

### Headless Game Engine

- The core game logic (game loop, mob/player logic, win/loss, input, etc.) is now fully decoupled from Phaser and lives in the headless engine (`client/src/engine/HeadlessGameEngine.ts`).
- The engine and all core systems are tested in a pure Node.js environment with no DOM or Phaser dependencies (see `HeadlessGameEngine.unit.test.ts`).
- This engine provides a programmatic API for stepping the game, injecting input, and retrieving the full game state, enabling automated play, bots, and CLI/testing scenarios.
- All gameplay state and logic are managed via the global `gameState` and `StateManager`, ensuring robust, testable, and decoupled gameplay.
- **All rendering and UI logic is handled in the render layer (Phaser scenes or other renderer), not in the engine.**
- See `client/src/engine/HeadlessGameEngine.unit.test.ts` for usage examples and tests.
- **Comprehensive unit and integration tests now cover full game simulation, bot play, rapid input, empty/edge word lists, and extreme spawn rates/health edge cases.**

## RenderManager Abstraction

TypeDefense uses a `RenderManager` abstraction to decouple game logic from rendering. The `IRenderManager` interface defines the contract for rendering the current game state. Concrete implementations (e.g., Phaser, Three.js) implement this interface and are responsible for all rendering, UI, and effects.

- See `client/src/render/RenderManager.ts` for the interface definition.
- Renderers should implement `init`, `render`, and `destroy` methods.
- The engine and game logic interact with the renderer only via this interface, enabling easy swapping or extension of render backends.
- **Tests and mocks for the RenderManager interface are provided in `client/src/render/__tests__/RenderManager.test.ts`.** These verify that scenes and game logic call the correct rendering methods and make it easy to test new renderer implementations.

## Render Backend Selection

TypeDefense supports multiple rendering backends via the `RenderManager` abstraction. You can choose between Phaser (default) and a prototype Three.js renderer.

- **Phaser:** Full-featured, production-ready renderer.
- **Three.js:** Prototype implementation (`ThreeJsRenderManager.ts`) that renders mobs and player as colored spheres. Used to validate the render abstraction.

To use the Three.js renderer, set the backend before loading the app:

- **Build-time:** `VITE_RENDER_BACKEND=three npm run dev`
- **Runtime:** In the browser console, set `window.RENDER_BACKEND = 'three'` before loading the app.

See `client/src/render/ThreeJsRenderManager.ts` for details.

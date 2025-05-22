# TypeDefense

TypeDefense is a web-based typing game designed to help players improve their typing skills through engaging gameplay. Players defend against waves of enemies by typing words or letters displayed above each enemy before they reach the player.

## Table of Contents

- [Features](#features)
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
- Dynamic difficulty scaling (spawn rate, speed, word complexity)
- Level and world progression system with unlockable content
- Player health system and game over state
- Curriculum-based finger group tracking for typing improvement
- Built with TypeScript, Vite, and Phaser
- Centralized game state managed by StateManager
- Level progression and finger group stats are now fully managed in the global game state (`gameState`) via StateManager, replacing the old LevelManager and FingerGroupManager. All progression and curriculum data is now state-driven and accessible for robust, testable gameplay and UI logic.
- The main game loop now updates delta time and timestamp in the global game state each frame, and all core systems (Player, MobSpawner, UI, etc.) are updated using the current state via StateManager. This enables robust, testable, and state-driven gameplay logic.
- Scenes are now managed via `gameState.gameStatus` and all transitions are triggered by updating state via StateManager, not by direct scene switching.
- MainMenu, Menu (WorldSelect), and LevelMenu scenes now render UI and handle navigation based on state, and listen for state changes to update or stop themselves.
- All scene transitions and UI updates are reactive to changes in `gameState.gameStatus`.
- All input (keyboard/mouse) is now handled centrally by the InputSystem, which updates gameState via StateManager. Scenes and entities no longer register input listeners directly.
- All core entities (Player, Mob, MobSpawner) are now fully state-driven: their data and logic are managed via the global game state (`gameState`) and all updates are performed through the `StateManager`. This enables robust, testable, and reactive gameplay logic, and allows for easy debugging and inspection of all entity state at runtime.

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

## Project Structure

See `.github/instructions/project_layout.instructions.md` for a detailed breakdown of the project structure and file organization.

## Testing

- All tests are located in `__tests__` subdirectories next to the code under test.
- The project uses [Vitest](https://vitest.dev/) for running tests:

  ```bash
  cd client
  npm run test
  ```

- Add new tests in the appropriate `__tests__` folder for each module.

## Contributing

Contributions are welcome! Please open issues or pull requests for bug fixes, new features, or documentation improvements.

## License

[MIT](LICENSE)

## Changelog

See [CHANGELOG.md](./CHANGELOG.md) for a detailed list of changes and feature history.

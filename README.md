# TypeDefense

TypeDefense is a web-based typing game designed to help players improve their typing skills through engaging gameplay. Players control a static character on the left side of the screen, defending against waves of enemies that spawn from the right. To defeat enemies, players must quickly and accurately type the words or phrases displayed above each enemy before they reach the player.

## Features

- Fast-paced wave-based mob defense gameplay
- Enemies spawn from the right and move toward the player
- Defeat enemies by typing their associated words or letters
- **Each mob now displays a single letter. When the correct letter is typed, the mob disappears.**
- Designed to improve typing speed and accuracy
- Built with TypeScript and Vite for a modern web experience
- Addictive action-challenge-reward loop with instant visual and audio feedback
- Real-time scores, combo multipliers, and particle effects for engagement
- Tweened UI transitions and dramatic camera effects for key events
- Layered audio cues for typing, combos, and wave clearances
- Core game loop set up in `GameScene` with Player and InputHandler initialized and updated each frame
- **Mob and MobSpawner integrated into GameScene:**
  - `MobSpawner` handles spawning and management of enemy mobs (see `client/src/entities/MobSpawner.ts`).
  - `Mob` represents individual enemy entities (see `client/src/entities/Mob.ts`).
  - Both are now fully integrated and updated within the main game loop in `GameScene`.
- Added initial design and documentation for the `FingerGroupManager` class, which will track player progress and statistics across finger groups for the typing curriculum.
  - Implemented the `FingerGroupManager` class in `client/src/managers/fingerGroupManager.ts`.
  - Tracks player progress, key usage, accuracy, and speed for each finger group.
  - Provides methods to record key presses, retrieve stats, and determine mastery.
  - Uses curriculum-defined finger/key mappings for robust tracking.
- Integrated FingerGroupManager with the main game loop in `GameScene.ts`.
- Each key press is now recorded and mapped to its finger group using the curriculum mapping (`getKeyInfo`).
- This enables tracking of finger usage and progress for each finger group in real time.
- **Mobs now walk toward the player:** Each mob moves toward the player character's position (left side of the screen) after spawning.
- **Player health system:** The player has a visible health value above their sprite. When a mob reaches the player, the player loses health and the mob is removed. The game ends with a "Game Over" message if health reaches zero.

## Curriculum Design

TypeDefense features a unique learning approach based on finger groups rather than random letters. The game is structured into four worlds, each focusing on a specific set of fingers:

### World 1: Index Fingers

- Left Hand: F, G, R, T, V, B
- Right Hand: J, H, Y, U, N, M
- Progressive levels introduce these keys gradually, starting with home row (F/J)

### World 2: Middle Fingers

- Left Hand: D, E, C
- Right Hand: K, I, comma
- Builds on index finger skills while introducing middle finger positions

### World 3: Ring Fingers

- Left Hand: S, W, X
- Right Hand: L, O, period
- Focuses on training the typically weaker ring fingers

### World 4: Pinky Fingers

- Left Hand: A, Q, Z
- Right Hand: semicolon, P, slash
- Completes the alphabet and introduces Shift key for capitals

Each world contains multiple levels that introduce letters progressively, with boss battles that test mastery before advancing. This finger-group approach builds proper muscle memory and typing technique.

## Getting Started

### Prerequisites

- Node.js (v18 or newer recommended)
- npm (comes with Node.js)

### Setup and Run (Client)

1. Navigate to the `client` directory:

   ```bash
   cd client
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the development server:

   ```bash
   npm run dev
   ```

4. Open your browser to the local server URL (usually `http://localhost:5173`).

The main game scene is set up in `client/src/scenes/GameScene.ts` and is ready for core mechanic development.

Contains AI-generated edits.

# TypeDefense

TypeDefense is a web-based typing game designed to help players improve their typing skills through engaging gameplay. Players control a static character on the left side of the screen, defending against waves of enemies that spawn from the right. To defeat enemies, players must quickly and accurately type the words or phrases displayed above each enemy before they reach the player.

## Features

- Fast-paced wave-based mob defense gameplay
- Enemies spawn from the right and move toward the player
- Defeat enemies by typing their associated words
- Designed to improve typing speed and accuracy
- Built with TypeScript and Vite for a modern web experience
- Addictive action-challenge-reward loop with instant visual and audio feedback
- Real-time scores, combo multipliers, and particle effects for engagement
- Tweened UI transitions and dramatic camera effects for key events
- Layered audio cues for typing, combos, and wave clearances
- Modular scenes for each game state (preload, menu, waves, game over)
- Data-driven wave and word pack configuration for easy tuning
- Escalating difficulty and unlockable word packs
- Planned leaderboard and achievements integration

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

---
Contains AI-generated edits.

# TypeDefense

TypeDefense is a game built with Ebiten that lets you practice your typing skills in a fun and engaging way. The game features various levels of difficulty, different types of enemies ("mobs"), and a unique scoring system that rewards accuracy and speed.

## Gameplay

In TypeDefense, you face waves of enemies that approach you. Each enemy (mob) has a word or sequence of letters above it. You must type the letters in order to defeat the mob. The game tests your typing speed and accuracy as you progress through different levels.

## Mob System and Performance

Each mob displays a sequence of letters, and each letter can be in one of three states:

- **TARGET**: The current letter the player must type (highlighted)
- **ACTIVE**: Letters not yet typed, but not the current target
- **INACTIVE**: Letters already typed

For performance, the game uses a global cache of pre-rendered images for each letter in each state. This avoids per-frame text rendering and allows for fast switching of letter states as the player types. When all letters in a mob are INACTIVE, the mob starts a death animation and is removed from the game.

## Features

- Endless waves of enemies
- Multiple worlds with increasing letters and difficulty
- Various types of enemies with different words specific to each world
- Unique scoring system based on accuracy and speed
- Engaging gameplay that improves typing skills
- Beautiful graphics and sound effects
- Leaderboard to track your progress and compare with others
- Customizable settings for a personalized experience

## Architecture Overview

- **Mobs**: Each mob is an entity with a position, movement, and a sequence of letters. Letter states are tracked per-mob, and rendering uses the global letter image cache.
- **Letter Image Cache**: A shared cache stores images for each rune in each state (TARGET, ACTIVE, INACTIVE), generated once per font/size.
- **Game Loop**: Handles spawning, updating, and removing mobs, as well as player input and scoring.

## Rendering and Canvas

TypeDefense now uses a fixed 1920x1080 internal canvas for all gameplay, UI, and entity rendering. All coordinates and layout are specified in this space, and the engine automatically scales the entire canvas to fit the client window, maintaining aspect ratio. This approach simplifies layout and ensures pixel-perfect consistency across all devices and window sizes.

- All entity and UI positions are in 1920x1080 space.
- No per-object scaling or normalized coordinates are used.
- The engine handles scaling the final frame to the window.

## Current State and Next Steps

The current focus is on iterating the core game loop with a single mob, ensuring the gameplay is fun and robust before expanding to more content and customization. Key areas for immediate and future development include:

- Mob spawning logic and timing
- Targeting and shooting mobs (player input, feedback, and effects)
- Scoring system and feedback
- Level win/lose conditions
- UI/UX improvements, including menus and game lobby for mode selection and customization
- Game loop polish and iteration for fun and engagement

See `TODO.md` for a detailed breakdown of proposed development paths and priorities.

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
- **UI**: Uses Ebiten's text and image drawing APIs for efficient rendering.

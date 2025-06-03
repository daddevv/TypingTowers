# Project Roadmap

## Concept

A keyboard-only base defense game where players strategically place towers to defend against waves of enemies.

## Features

- **Keyboard Navigation**: All game controls are accessible via keyboard shortcuts.
- **Letter Mobs**: Enemies are represented by letters, each with unique properties.
- **Progression**: Players can unlock new mobs by spending points earned from defeating enemies.
- **Technology Tree**: Players can unlock upgrades to towers and abilities.

## Core Gameplay Loop
1. Waves spawn enemies marching toward the tower.
2. Towers auto-fire, but they must be reloaded manually by typing sequences.
3. Typing accuracy and speed affect reload time, score multiplier, and special effects.
4. Letters are unlocked over time, slowly expanding the typing difficulty.

## Main Objectives
- Survive waves as long as possible.
- Achieve a high score through efficient typing.
- Progress through the tech tree to gain tower and player enhancements.

## Typing System Design
- **Initial Set**: Only home row letters (`a`, `s`, `d`, `f`, `j`, `k`, `l`, `;`) are available.
- **Unlock Logic**: Certain upgrades (e.g., *Laser Beam*) require letters like `z` or `q` to be unlocked.
- **Typing Tasks**:
  - Reload tower: `r e l o a d`
  - Activate power: `s h o c k`
  - Repair tower: `f i x`
- **Penalty**: Mistyped letters cause delays or missed reloads.

## Tech Tree Mechanics
Branches unlock based on letter access. For example, advanced weapons become available once the player can type the required letters.

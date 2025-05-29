# TODO: TypeDefense Development Roadmap

This document outlines the proposed development paths and priorities for TypeDefense. The focus is on iterating the core gameplay loop to completion and fun, then expanding content, customization, and polish.

## 1. Core Game Loop (Immediate Priority)
- [ ] **Single Mob Iteration**
    - Refine the experience with a single mob: spawning, movement, letter targeting, and feedback.
    - Ensure the typing mechanic is responsive, rewarding, and fun.
    - Add visual and audio feedback for correct/incorrect keypresses.
    - Implement mob death animation and removal when all letters are typed.
    - Polish the update/draw cycle for smoothness and clarity.
- [ ] **Player Input and Targeting**
    - Highlight the current target letter.
    - Provide clear feedback for correct/incorrect input.
    - Allow for rapid input and error recovery.
- [ ] **Scoring System**
    - Track accuracy, speed, and combo streaks.
    - Display score and feedback in real time.
    - Prepare for leaderboard integration.
- [ ] **Level Win/Lose Conditions**
    - Define what constitutes a win or loss for a level (e.g., all mobs defeated, player survives X time, etc.).
    - Implement transitions between levels or game over states.

## 2. Mob System Expansion
- [ ] **Mob Spawning Logic**
    - Implement timed and patterned mob spawning (waves, random, etc.).
    - Support multiple mobs on screen with collision and overlap handling.
    - Vary mob speed, word length, and behavior by level/difficulty.
- [ ] **Mob Types and Behaviors**
    - Add new mob types with unique movement, words, and effects.
    - Support for special mobs (e.g., bonus, boss, shielded, etc.).

## 3. UI/UX and Menus
- [ ] **Main Menu Improvements**
    - Polish the main menu layout and navigation.
    - Add visual feedback for selection and transitions.
- [ ] **Game Lobby**
    - Implement a lobby screen for game mode selection and customization (difficulty, world, etc.).
    - Allow players to configure settings before starting a game.
    - Support multiplayer or co-op options in the future.
- [ ] **In-Game UI**
    - Display score, combo, and other stats during gameplay.
    - Add pause, resume, and quit options.
    - Show level progress and win/lose feedback.

## 4. Content and Customization
- [ ] **Worlds and Levels**
    - Expand with new worlds, backgrounds, and word lists.
    - Add level progression and unlocks.
- [ ] **Customization**
    - Allow players to select fonts, colors, and accessibility options.
    - Support for custom word lists and challenges.

## 5. Technical and Architecture
- [ ] **Performance Optimization**
    - Profile and optimize rendering and update loops.
    - Ensure smooth performance with many mobs and effects.
- [ ] **Codebase Cleanup**
    - Refactor for clarity, maintainability, and extensibility.
    - Improve documentation and inline comments.
- [ ] **Testing and QA**
    - Add unit and integration tests for core systems.
    - Implement automated build and test pipelines.

## 6. Future Features (Post-MVP)
- [ ] **Leaderboards and Progression**
    - Online and local leaderboards.
    - Player profiles and stats tracking.
- [ ] **Multiplayer/Co-op Modes**
    - Real-time or turn-based multiplayer typing challenges.
- [ ] **Modding and Community Content**
    - Support for user-generated word lists, levels, and mods.

---

**Current Focus:**
- Iterate and polish the single-mob game loop until it is complete and fun.
- Defer additional content and features until the core gameplay is robust and engaging.

This roadmap is a living document and should be updated as priorities shift and milestones are achieved.

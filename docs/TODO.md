# TODO List

## Completed ✅

- [x] ~~Create wave system that spawns enemy units marching toward the tower~~
- [x] ~~Allow towers to auto-fire and reload whenever ammo is below full using letter prompts~~
- [x] ~~Implement basic projectile intercept calculations for moving targets~~
- [x] ~~Add jamming system for incorrect reload inputs~~
- [x] ~~Create base health system and game over conditions~~
- [x] ~~Implement bouncing projectile mechanics~~
- [x] ~~Add real-time configuration reloading system~~
- [x] ~~Implement functional upgrade purchasing with gold~~
- [x] ~~Add tower damage, range, and fire rate upgrades~~
- [x] ~~Implement two-queue ammo system (firing queue + reload queue)~~
- [x] ~~Add ammo capacity upgrades to shop~~
- [x] ~~Show reload queue preview in HUD~~

## Core Gameplay Loop (Next Sprint)

### 1. Multiple Tower Placement & Management (Vim/Qutebrowser Navigation)

- [ ] **Backend:** Implement logic to place new towers on the grid using keyboard input (no mouse; Vim-style navigation)
  - [ ] Define valid placement areas (e.g., specific tile types).
  - [ ] Add a cost for building new towers (`TowerConstructionCost`).
  - [ ] Deduct gold when a tower is placed.
- [ ] **UI:** Visual indicator for valid/invalid tower placement locations (keyboard-driven, e.g., highlight with Vim-style cursor).
- [ ] **UI:** Tower selection mechanism using Vim/Qutebrowser keys (`h/j/k/l`, `gg/G`, `/`, etc.) to move/select towers.
- [ ] **Config:** Add `TowerConstructionCost` to `config.json`.
- [ ] **UI:** Implement modal navigation (normal/insert/command mode) for all in-game menus and overlays.
- [ ] **UI:** Display keyboard hints/overlays for available actions (e.g., `[h/j/k/l] move`, `[d] delete`, `[u] upgrade`, `[q] quit`).

### 2. Enhanced Shop Functionality (Vim/Qutebrowser Navigation)

- [ ] **Backend:** Modify upgrade logic to apply to the currently *selected* tower (selected via Vim-style navigation).
- [ ] **UI:** Update HUD to clearly show which tower is selected (keyboard highlight/cursor).
- [ ] **UI:** Shop upgrade options should be navigable and activated via Vim/Qutebrowser keys (e.g., `j/k` to move, `Enter` to select, `1-5` for upgrades).
- [ ] **Feature:** "Foresight" upgrade:
  - [ ] **Backend:** Increase the number of reload queue letters shown.
  - [ ] **UI:** Update HUD to display more reload queue letters.
  - [ ] **Shop:** Add "Foresight" as a purchasable upgrade (e.g., key '5').
- [ ] **UI:** Add visual feedback for purchased upgrades (updated stats in HUD/Shop UI).
- [ ] **UI:** All shop and upgrade navigation must be fully keyboard-driven, with no mouse dependency.

## Advanced Mechanics & Systems

### Word-Based Reloading

- [ ] Design: Define short words for reloading (e.g., "ammo", "load", "zap").
- [ ] Implement: Replace single-letter reload prompts with word prompts.
- [ ] Implement: Logic for typing words correctly for reload.
- [ ] UI: Update HUD to display word prompts.

### Typing Performance Modifiers

- [ ] Track typing accuracy for reloads/targeting.
- [ ] Track typing speed (WPM or similar metric).
- [ ] Implement: Reload speed bonuses for fast/accurate typing.
- [ ] Implement: Penalties for slow/inaccurate typing (e.g., increased reload time, temporary jam).
- [ ] UI: Display basic typing stats (accuracy, WPM) in HUD.

### Letter Pool Progression

- [x] ~~Design: System for unlocking/expanding available letters for reloading (e.g., start with f and j, proceed to other index fingers, then middle then ring then pinky, unlock more via shop or wave progression).~~
- [x] ~~Implement: Mechanism to expand the letter pool.~~

## Medium Priority

### Additional Tower Types

- [ ] Design sniper tower (high damage, slow fire, precise typing)
- [ ] Design rapid-fire tower (low damage, fast fire, simple inputs)
- [ ] Implement tower type selection in shop/build menu (keyboard-driven)
- [ ] Balance different tower types for strategic variety

### Enhanced Enemy Variety

- [ ] Create different enemy types with varying health/speed
- [ ] Add special enemies that require specific letter combinations or words
- [ ] Implement enemy abilities (e.g., armor, shields, speed bursts)
- [ ] Design boss enemies for milestone waves

### Typing Performance Metrics

- [ ] Implement performance-based score multipliers for gold/score.
- [ ] Add detailed typing statistics page/summary post-game.
- [ ] Create performance history tracking (e.g., best WPM, accuracy trends).

## Low Priority

### Progressive Letter Unlocks (Tech Tree)

- [ ] Design tech tree for unlocking new letters, words, or typing abilities.
- [ ] Implement letter/word unlock conditions and rewards.
- [ ] Create visual tech tree interface (keyboard navigable).
- [ ] Add achievement system for unlock milestones.

### Advanced Reload Mechanics

- [ ] Implement variable reload sequences based on tower type or upgrades.
- [ ] Add combo system for consecutive accurate inputs during reload/targeting.
- [ ] Create special reload challenges for bonus effects.

### Quality of Life Improvements

- [ ] Add pause menu with options (resume, restart, quit, settings) navigable via Vim/Qutebrowser keys.
- [ ] Implement save/load game state.
- [ ] Create settings menu for key bindings, audio, and display options (keyboard-driven).
- [ ] Add sound effects and background music.

## Technical Debt

- [ ] Refactor tower creation to support multiple tower types efficiently.
- [ ] Abstract enemy creation for different enemy types and behaviors.
- [ ] Improve configuration system for more complex parameters (e.g., per-tower type configs, enemy-specific letter pools).
- [ ] Add proper error handling and logging for edge cases.
- [ ] Implement a more robust asset management system.

## Testing

- [ ] Write comprehensive unit tests for new systems:
  - [ ] Letter-based enemy targeting logic.
  - [ ] Multiple tower placement, selection, and gold deduction.
  - [ ] Shop upgrades for selected towers.
  - [ ] Word-based reloading.
  - [ ] Vim/Qutebrowser-style navigation for all UI/UX.
- [ ] Add integration tests for complex interactions (e.g., multiple towers targeting, shop affecting game state).
- [ ] Performance optimization for large numbers of entities.

## Documentation Updates

- [ ] Create gameplay tutorial/guide for new mechanics (targeting, tower placement, Vim/Qutebrowser navigation).
- [ ] Document new configuration file parameters (`TowerConstructionCost`, letter pools, etc.).
- [ ] Update developer setup instructions if new dependencies or build steps arise.
- [ ] Update contribution guidelines for new features and testing requirements.
- [ ] Code cleanup and documentation updates within the codebase.

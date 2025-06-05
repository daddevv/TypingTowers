# TODO List

## Core Gameplay Loop

### Multiple Tower Placement & Management (Vim/Qutebrowser Navigation)

- [x] Implement additional tower types (sniper, rapid-fire) with unique mechanics and balance for strategic variety.
- [x] Ensure sniper tower has longer range and rapid-fire tower has faster fire rate than basic tower. Add/verify unit tests for these properties.
- [ ] Implement tower type selection in shop/build menu (keyboard-driven).
- [x] Refactor tower creation to efficiently support multiple tower types.
- [x] Fix and expand unit tests to verify correct tower type is built and gold is deducted appropriately.

### Enhanced Enemy Variety

- [x] Create enemy types with varying health, speed, and abilities (armor, shields, speed bursts).
- [x] Add special enemies requiring specific letter combinations or words.
- [x] Design and implement boss enemies for milestone waves.
- [x] Abstract enemy creation for extensibility.
- [x] Ensure new mob types (boss, armored, fast) are properly implemented and tested for unique behaviors and stats.
- [x] Add/verify unit tests for mob type instantiation and special abilities.

### Typing Performance & Metrics

- [ ] Implement performance-based score multipliers for gold/score.
- [ ] Add detailed typing statistics page/summary post-game.
- [ ] Track performance history (best WPM, accuracy trends).
- [ ] Implement performance-based achievements and rewards.

## Technical Debt

- [ ] Improve configuration system for complex parameters (per-tower type configs, enemy-specific letter pools).
- [ ] Add proper error handling and logging for edge cases.
- [ ] Implement robust asset management system.

## Testing

- [ ] Write comprehensive unit tests for:
  - Multiple tower placement, selection, and gold deduction.
  - Shop upgrades for selected towers.
  - Word-based reloading.
  - Vim/Qutebrowser-style navigation for all UI/UX.
  - Tower type properties (range, fire rate, etc.) and correct instantiation.
  - Mob type properties (health, speed, abilities) and correct instantiation.
- [ ] Add integration tests for complex interactions (e.g., multiple towers targeting, shop affecting game state).
- [ ] Optimize performance for large numbers of entities.

## Documentation Updates

- [ ] Create gameplay tutorial/guide for new mechanics (targeting, tower placement, Vim/Qutebrowser navigation).
- [ ] Document new configuration file parameters (`TowerConstructionCost`, letter pools, etc.).
- [ ] Update developer setup instructions for new dependencies or build steps.
- [ ] Update contribution guidelines for new features and testing requirements.
- [ ] Code cleanup and documentation updates within the codebase.

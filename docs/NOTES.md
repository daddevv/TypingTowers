# Game Notes

## Vim/Qutebrowser-Style Navigation

All user interface and gameplay controls are designed for Vim/Qutebrowser-style keyboard navigation. This means:

- **No mouse required**: All actions are performed via keyboard.
- **Modal navigation**: The game uses modes (normal/insert/command) for different contexts.
- **Navigation keys**: Use `h/j/k/l` to move selection, `gg/G` to jump, `/` to search, and other Vim/Qutebrowser conventions.
- **Action hints**: UI overlays display available keyboard actions and current mode.
- **All new features must be accessible and testable via keyboard navigation only.**

## Current Implementation

### Controls

- **h/j/k/l**: Move selection cursor (towers, shop, menus)
- **Space**: Pause/unpause game
- **F5**: Reload configuration file  
- **Escape**: Quit game or exit to previous mode
- **Enter**: Continue to next wave (shop mode) or confirm selection
- **F** or **J**: Reload ammunition when prompted
- **Backspace**: Clear jammed tower
- **1-5**: Purchase upgrades in shop (when available)
- **: (colon)**: Enter command mode (for advanced actions)
- **/ (slash)**: Search/select towers or upgrades

### Game Flow

1. **Wave Phase**: Enemies spawn from the right and move toward the base
2. **Auto-Combat**: Tower automatically targets and fires at closest enemy in range
3. **Reload Phase**: When ammo is low, player must type prompted letters to reload
4. **Shop Phase**: Between waves, player can proceed to next wave or purchase upgrades using keyboard navigation

### Tower Mechanics

- **Auto-Targeting**: Selects closest enemy within range
- **Limited Ammo**: Tower has finite ammunition that must be manually reloaded
- **Reload Prompts**: Shows 'f' or 'j' letters that must be typed to add ammunition
- **Jamming**: Incorrect keystrokes jam the tower until backspace is pressed
- **Projectile Intercept**: Projectiles calculate intercept paths for moving targets
- **Bouncing**: Projectiles can bounce between multiple targets

### Enemy Mechanics

- **Wave Spawning**: Enemies spawn at intervals from the right side
- **Health Scaling**: Enemy health increases with each wave
- **Movement**: Enemies move directly toward the player's base
- **Collision**: Enemies damage the base when they reach it

### Base Defense

- **Health System**: Base has limited health that decreases when enemies reach it
- **Game Over**: Game ends when base health reaches zero

## Configuration System

The game supports real-time configuration changes through:

- Tower fire rate, damage, range, and ammunition capacity
- Enemy spawn rates, health, and movement speed  
- Projectile speed and bounce mechanics
- Wave difficulty scaling parameters

## Future Gameplay Ideas

- All future UI/UX features must be designed for Vim/Qutebrowser-style keyboard navigation.

# Game Notes

## Current Implementation

### Controls

- **Space**: Pause/unpause game
- **F5**: Reload configuration file  
- **Escape**: Quit game
- **Enter**: Continue to next wave (shop mode)
- **F** or **J**: Reload ammunition when prompted
- **Backspace**: Clear jammed tower

### Game Flow

1. **Wave Phase**: Enemies spawn from the right and move toward the base
2. **Auto-Combat**: Tower automatically targets and fires at closest enemy in range
3. **Reload Phase**: When ammo is low, player must type prompted letters to reload
4. **Shop Phase**: Between waves, player can proceed to next wave

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

### Planned Features

- **Letter-Based Enemies**: Enemies display letters that must be typed to target them
- **Typing Challenges**: Complex reload sequences requiring full words
- **Multiple Tower Types**: Different towers with unique ammunition and reload mechanics
- **Tech Tree**: Unlock new towers and abilities through progression
- **Accuracy Scoring**: Track typing accuracy and speed for performance bonuses

### Shop System Framework

- Basic shop phase exists between waves
- Framework for purchasing upgrades with collected gold
- Placeholder for tower and ability upgrades

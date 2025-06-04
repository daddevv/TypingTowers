# TypingTowers

TypingTowers is a keyboard-controlled tower defense game built with [Ebiten](https://ebiten.org/). Players defend their base against waves of enemies using towers that require manual reloading through typing mechanics.

## Gameplay

- **Wave Defense**: Enemies spawn from the right side of the screen and move toward your base on the left
- **Auto-Targeting**: Towers automatically target and fire at the closest enemy within range
- **Manual Reloading**: When a tower's ammo runs low, players must type the prompted letters ('f' or 'j') to reload
- **Jamming System**: Typing incorrect letters jams the tower, requiring backspace to clear
- **Progressive Difficulty**: Each wave spawns more enemies with increased health

## Controls

- **Space**: Pause/unpause the game
- **F5**: Reload configuration file
- **Escape**: Quit game
- **Enter**: Proceed to next wave (when in shop mode)
- **F/J**: Reload ammunition when prompted
- **Backspace**: Clear jammed tower

## Structure

- `v1/cmd/game`: Main entrypoint for running the game
- `v1/internal/game`: Core game logic, including entities, input handling, mobs, and towers
- `assets`: Graphic and audio resources
- `docs/`: Documentation including design notes and roadmap

## Running

Ensure you have Go installed. From the `v1` directory run:

```bash
cd v1
go run ./cmd/game
```

## Testing

Run Go tests from the `v1` directory:

```bash
cd v1
go test ./...
```

## Configuration

The game loads settings from a configuration file that can be reloaded during gameplay with F5. This allows for real-time tuning of game parameters like tower damage, mob health, and spawn rates.

## Contributing

Code should be formatted with `gofmt` and accompanied by unit tests when possible. See `.github/instructions` for detailed development guidelines.

## Current Features

- Tower auto-targeting and firing system
- Manual reload mechanics with typing prompts
- Wave-based enemy spawning with scaling difficulty
- Base health system and game over conditions
- Real-time configuration reloading
- Projectile intercept calculations for moving targets
- Bouncing projectile mechanics

## Automation

Custom prompts in `.github/prompts` and guidelines in `.github/instructions` are used with GitHub Copilot to streamline development and reduce bugs.

# TypingTowers

TypingTowers is a keyboard-controlled tower defense game built with [Ebiten](https://ebiten.org/). Players defend their base against waves of enemies using towers that require manual reloading through typing mechanics.

## Keyboard-First Navigation (Vim/Qutebrowser Style)

All UI and gameplay interactions are designed for Vim/Qutebrowser-style keyboard navigation. This means:

- **No mouse required**: All actions (tower placement, selection, upgrades, shop navigation) are performed via keyboard.
- **Modal navigation**: The game uses modes (normal/insert/command) similar to Vim for different contexts (e.g., moving selection, typing, issuing commands).
- **Navigation keys**: Use `h/j/k/l` to move selection, `gg/G` to jump, `/` to search, and other Vim/Qutebrowser conventions.
- **Action hints**: UI overlays display available keyboard actions and current mode.

## Gameplay

- **Wave Defense**: Enemies spawn from the right side of the screen and move toward your base on the left
- **Auto-Targeting**: Towers automatically target and fire at the closest enemy within range
- **Manual Reloading**: When a tower's ammo runs low, players must type the prompted letters to reload
- **Letter Unlocking**: Available reload letters expand over time, starting with 'f' and 'j' and adding more keys each wave
- **Jamming System**: Typing incorrect letters jams the tower, requiring backspace to clear
- **Progressive Difficulty**: Each wave spawns more enemies with increased health
- **Upgrade Purchasing**: Spend gold earned from defeating enemies to upgrade tower damage, range, and fire rate between waves
  - Press [1] to upgrade damage (+1)
  - Press [2] to upgrade range (+50)
  - Press [3] to upgrade fire rate (faster)
  - Each upgrade costs 5 gold and applies to the selected tower
  - Upgrades are available in the shop between waves and are covered by automated unit tests
- **Enemy Variety**: Multiple enemy types are implemented and tested, including:
  - **Armored mobs**: Take reduced damage from attacks
  - **Fast mobs**: Periodically move at burst speed
  - **Boss mobs**: Appear on milestone waves with high health and unique stats
  - All mob types are covered by unit tests for instantiation and special abilities

## Controls

- **h/j/k/l**: Move selection cursor (towers, shop, menus)
- **Enter**: Confirm selection or action
- **Space**: Pause/unpause the game
- **F5**: Reload configuration file
- **Escape**: Quit game or exit to previous mode
- **F/J**: Reload ammunition when prompted
- **Backspace**: Clear jammed tower
- **1-5**: Purchase upgrades in shop (when available)
- **: (colon)**: Enter command mode (for advanced actions, e.g., `:quit`, `:save`)
- **/ (slash)**: Search/select towers or upgrades

## Structure

- `v1/cmd/game`: Main entrypoint for running the game
- `v1/internal/game`: Core game logic, including entities, input handling, mobs, and towers
- `assets`: Graphic and audio resources
- `docs/`: Documentation including design notes, roadmap and balancing approach

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

Code should be formatted with `gofmt` and accompanied by unit tests when possible. See `.github/instructions` for detailed development guidelines. All new UI/UX features must support Vim/Qutebrowser-style keyboard navigation.

## Current Features

- Tower auto-targeting and firing system
- Manual reload mechanics with typing prompts
- Wave-based enemy spawning with scaling difficulty
- Varied enemy types including armored, fast and boss enemies
- Base health system and game over conditions
- Real-time configuration reloading
- Projectile intercept calculations for moving targets
- Bouncing projectile mechanics
- Upgrade purchasing system using gold between waves (damage, range, fire rate upgrades implemented and tested)
- Technology tree loaded from `tech_tree.yaml` with keyboard-driven purchase menu (`/` to search, `Enter` to buy)
- **Keyboard-driven navigation for all menus and gameplay**
- Typing accuracy and WPM tracking with bonuses and penalties

## Automation

Custom prompts in `.github/prompts` and guidelines in `.github/instructions` are used with GitHub Copilot to streamline development and reduce bugs.

## Balance Editor

A Python helper script located at `tools/balance_editor.py` can visualize expected game progression and tweak `v1/config.json`. The script simulates up to 100 waves using the formulas described in `docs/BALANCING.md` and plots mob health, time to kill, and survival time.

Run it with optional overrides:

```bash
python3 tools/balance_editor.py --set tower_damage=15 --save
```

`matplotlib` is required for plotting.

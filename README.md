# Type Defense

Type Defense is a keyboard-controlled tower defense game built with [Ebiten](https://ebiten.org/). Players type letters to shoot enemy mobs and manage upgrades through keyboard shortcuts.

## Structure
- `v1/cmd/game`: Main entrypoint for running the game.
- `v1/internal/game`: Core game logic, including entities, input handling, mobs, and towers.
- `assets`: Graphic and audio resources.
- `NOTES.md` and `ROADMAP.md`: Design notes and planned features.

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

## Contributing
Code should be formatted with `gofmt` and accompanied by unit tests when possible. See `.github/instructions` for detailed development guidelines.


## Automation
Custom prompts in `.github/prompts` and guidelines in `.github/instructions` are used with GitHub Copilot to streamline development and reduce bugs.


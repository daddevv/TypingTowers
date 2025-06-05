# Project Layout

```
/
├── README.md             - project overview and setup
├── NOTES.md              - gameplay notes
├── ROADMAP.md            - future features
├── v1/                   - main game module
│   ├── cmd/game          - entry point for the Ebiten application
│   ├── internal/game     - game logic packages (tower upgrades, shop logic, etc.)
│   └── assets            - sprites and other resources
```

All Go commands should be run from the `v1` directory which contains the module file `go.mod`.


# Project Layout

```
/
├── README.md             - project overview and setup
├── NOTES.md              - gameplay notes
├── ROADMAP.md            - future features
├── v1/                   - main game module
│   ├── cmd/game          - entry point for the Ebiten application
│   ├── internal/game     - main coordinator and shared logic
│   ├── internal/content  - asset loading and management
│   ├── internal/entity   - units and base entities
│   ├── internal/event    - Core event types and EventBus
│   ├── internal/phase    - turn/phase manager
│   ├── internal/sprite   - sprite definitions and animations
│   ├── internal/tech     - tech tree logic
│   ├── internal/tower    - tower stats and behavior
│   ├── internal/ui       - HUD and UI overlays
│   └── assets            - sprites and other resources
```

All Go commands should be run from the `v1` directory which contains the module file `go.mod`.


# Tech Tree YAML Schema

Each building's tech tree is stored as a YAML file in `data/trees/`. A file contains a `nodes` list where each node defines:

- `id` — unique string identifier
- `name` — display name
- `type` — node category such as `UnlockLetter`, `StatBoost`, `NewUnit`, `CDReduction`, `GlobalPassive`
- `cost` — King's Points or resource cost to purchase
- `effects` — key/value map describing the modifiers granted on unlock
- `prereqs` — list of node IDs that must be unlocked first

See `data/trees/letters_basic.yaml` for an example.

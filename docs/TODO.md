# Modularization & Refactor Backlog

- [x] **MIG-002** Rename `mob` package to `enemy`
  - [x] Move all code from `internal/mob` to `internal/enemy`
  - [x] Update all import paths and references from `mob` to `enemy`
  - [ ] Update tests and documentation to use new package name
- [x] **MIG-003** Rename `structure` package to `building`
  - [x] Move all code from `internal/structure` to `internal/building`
  - [x] Update all import paths and references from `structure` to `building`
  - [ ] Update tests and documentation to use new package name
- [x] **MIG-004** Move worker types under `building/worker`
  - [x] Create `internal/building/worker` subpackage
  - [x] Move Farmer, Miner, Lumberjack, etc. from `worker` to `building/worker`
  - [x] Update all references and imports accordingly
  - [ ] Remove `worker` package after migration
- [ ] **MIG-005** Update all imports and tests after renames
  - [ ] Search for any lingering references to old package names
  - [ ] Refactor integration and unit tests to use new structure
- [ ] **MIG-006** Remove deprecated package references
  - [ ] Delete any obsolete files or directories left from the monolithic `game` package
  - [ ] Clean up documentation and comments referencing old structure

---

### Additional Modularization Tasks

- [ ] Audit all handler dependencies to ensure no circular imports remain
- [ ] Ensure all modules expose a `Handler` with `Update(dt)` and are registered in `game.Game`
- [ ] Refactor any remaining monolithic logic in `game` to delegate to handlers
- [ ] Update event bus usage: ensure all inter-module communication uses events, not direct calls
- [ ] Review and update the import map in `docs/INTERNAL_RESTRUCTURE.md` after all moves
- [ ] Update architecture documentation to reflect final modular structure

---

### Game Loop & Handler Refactor

- [ ] Refactor any remaining direct field access in `Game` to use handler interfaces where possible (e.g., mobs, towers, projectiles)
- [ ] Move all per-frame update logic for entities, towers, projectiles, and buildings into their respective handler modules
- [ ] Decouple direct calls to `g.base`, `g.farmer`, `g.lumberjack`, `g.miner`, and `g.barracks` by routing updates and events through handlers
- [ ] Remove any legacy logic that bypasses the event bus for inter-module communication
- [ ] Ensure all game state transitions (e.g., phase changes, game over, pause) are handled via a centralized state manager or event system
- [ ] Move shop, build, upgrade, and skill menu logic into dedicated UI handler modules
- [ ] Refactor input handling to delegate to handlers/UI modules instead of being processed directly in the main loop
- [ ] Remove any remaining references to the old monolithic `mob` or `worker` packages
- [ ] Implement event-driven mob spawning and wave progression, removing direct counter logic from the main loop
- [ ] Ensure all save/load and config reload logic is handled by a persistence or config handler, not directly in `Game`
- [ ] Audit for any remaining circular dependencies or tight coupling between `Game` and submodules
- [ ] Add tests for the refactored handler-based update and event systems

---

### Finalization

- [ ] Run all tests and CI to verify refactor stability
- [ ] Archive completed refactor tasks to `TODO_ARCHIVE.md`
- [ ] Announce completion and update contributor onboarding docs

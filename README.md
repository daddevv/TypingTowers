# Kingdoms of Keystoria (working title)

*Pixel-art city-builder / tower-defense where **every action is triggered by typing.***  
Build a medieval settlement, train armies of knights, and repel orc hordes by mastering the keyboard.  
No mouse required – the entire game (UI **and** gameplay) is navigated with Vim / Qutebrowser-style keys.

---

## Introduction and Core Vision

**TypingTowers** is a keyboard-controlled tower defense game where players defend their base against waves of enemies using towers that must be reloaded by typing.  
This expanded design envisions TypingTowers as a deep, replayable indie title that adds rich optional layers of gameplay while maintaining keyboard-only controls and a focus on improving typing speed and accuracy.  
The goal is to preserve the core educational typing benefits and Vim/Qutebrowser-style navigation, adding progression, complexity, and lore to engage players for 100+ hours.  
All new features remain optional enhancements on top of the base game, ensuring the game is welcoming to casual typists and challenging for hardcore players alike.

---

## Key Pillars

| Key pillars | Short version |
|-------------|---------------|
| **Keyboard-First** | `h/j/k/l`, `/` search, `:` command mode – every screen is accessible by keys alone. |
| **Letter Streams** | Each building owns a cooldown; when it expires it queues a random word from its letter-pool. Typing that word finishes a construction step, reloads a tower, or trains a soldier. |
| **Per-Building Tech Trees & Deep Progression** | Unlock letters one family at a time (Farmer → `f j`, Barracks → `f j d k`, …) to shorten cooldowns, add units, and widen the global letter pool. A massive skill tree with 100+ nodes enables long-term progression, branching upgrades, and new features. |
| **Autonomous Minions & Heroes** | Summon and command minions or heroes by typing keywords. Minions have unique roles and can be upgraded or managed via typed commands, complementing towers and enriching strategy. |
| **Incremental & Idle Mechanics** | Optional idle progression: auto-collection, offline progress, auto-reloading towers, and upgradable generators allow for meaningful progress with minimal input. Prestige/reset mechanics extend replayability. |
| **Typing Minigames & Challenges** | Speed trials, accuracy challenges, word puzzles, and boss practice modes provide fun breaks, reinforce typing skills, and grant in-game rewards. |
| **Multiple Playstyle Support** | Grind, optimize, idle, or embrace chaos—each playstyle is valid and rewarding, with systems and skill tree branches to support them. |
| **Global Skill Tree** | Sample in-memory tree defines offense, defense, typing, automation and utility nodes. |
| **Pixel-art Charm** | SNES-style sprites & slap-stick combat (exploding cabbages, bouncing arrows, comic “BONK!” bubbles). Violence stays E10+ but feels impactful. |

---

## Quick-Start (dev build)

```bash
git clone …
cd keystoria
go run ./cmd/game          # Ebiten entry
```

## Dependencies

- Go 1.22+, Ebiten, no GPU shaders beyond WebGL

## Run tests

```bash
go test -tags test ./...
```

The `resources_integration_test.go` integration test runs the game headlessly
for three minutes and verifies that all resources accumulate from zero.

-## Current prototype

- Shared FIFO queue manager implemented. Buildings enqueue words that are processed **letter by letter**. Completing a Barracks word spawns a Footman.
- Global queue is displayed on the HUD at `(400,900)` with a conveyor belt animation. Mistypes jam the queue until Backspace is pressed.
- Mistypes now trigger a brief red flash and a "clank" sound effect.
- Basic orc grunt waves scale every 45 s.
- Footmen automatically attack nearby orc grunts each tick.
- Back-pressure damage: if the queue grows past 20 letters, the base loses 1 HP each second.
- Typing speed/accuracy multiplier working.
- Vim navigation for pause/menu/shop implemented.
- Title screen with animated background and keyboard-only menu.
- Pre-game setup lets you choose a character and difficulty, shows a quick
  tutorial and typing test, then prompts for mode selection.
- Letter unlock order and costs documented (see `docs/LETTER_UNLOCKS.md`).
- Letters can now be unlocked in-game using King's Points, expanding each building's word pool.
- Tech trees are defined in YAML under `data/trees/` (see `letters_basic.yaml`). They are loaded at runtime via a Go parser that builds an in-memory graph and verifies all prerequisites.
- Skill tree nodes can be purchased with King's Points once prerequisites are met.

## Tech Tree YAML

Tech tree files live in `data/trees/` and describe upgrade nodes. The example
`letters_basic.yaml` shows the schema with fields for `type`, `cost`, `effects`
and `prereqs`. The `LoadTechTree` function parses these YAML files and returns a
validated in-memory graph.

---

## Prototype Sprint

The current focus is on prototyping the Gathering (Farmer) and Military (Barracks) families, implementing a shared queue manager, per-building cooldowns, and playtesting word density.  
See the [ROADMAP.md](./ROADMAP.md) for detailed tasks.

---

-## Farmer (Gathering Building)

- The Farmer building is now implemented as a Gathering structure that generates a word from its letter pool every cooldown cycle (default 1.5s).
- Typing the generated word completes the cycle and produces Food and Gold, which are added to the global resource pool.
- Each building's cooldown timer pauses once a word is queued and only resets after that word is typed.
- The Farmer's cooldown, letter pool, and word length can be configured and extended for progression.
- See `v1/internal/game/farmer.go` for implementation and `farmer_test.go` for tests.

## Lumberjack (Gathering Building)

- The Lumberjack produces Wood and Gold when its queued word is typed.
- Shares similar mechanics to the Farmer with a 1.5s base cooldown.
- Implementation lives in `v1/internal/game/lumberjack.go`.

## Miner (Gathering Building)

- The Miner yields Stone and Iron plus a bit of Gold on completion.
- Follows the same word-based workflow and cooldown system.
- See `v1/internal/game/miner.go` for code and tests.

-## Barracks (Military Building)

- The Barracks generates a word from its letter pool every few seconds.
- Each letter must be typed in order. When the final letter is entered a Footman unit spawns.
- Footmen start with **10 HP**, deal **1 damage** and move at a speed of **50 px/s**.
- Spawned units are tracked by the `Military` system.
- Word generation logic, cooldown behaviour, and letter-by-letter queue processing are tested in `barracks_test.go`.
- The HUD highlights progress on the first queued word so you can see partially typed letters.
- Cooldown progress bars for the Farmer and Barracks are shown along with resource icons.

- Press `/` to enter tower selection mode. The screen dims and towers are labeled with letters; press a letter to open that tower's upgrade menu.
- Press `/` again to open the tech menu. Type to search unlocked technologies and press `Enter` to purchase the highlighted upgrade.
- Press `Tab` to open the global skill tree. Use arrow keys to change categories and highlight skills. Locked and unlocked skills are indicated.
- Unlocking a skill immediately applies its bonuses and is saved with your game.
- Press `:` to enter command mode for quick text commands like `pause` or `quit`.

See docs/REQUIREMENTS.md for the full feature scaffold, ROADMAP.md for planned phases, and TODO.md for sprint tasks.

Happy typing & may your catapults stay jam-free!

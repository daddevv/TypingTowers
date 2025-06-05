# Kingdoms of Keystoria (working title)

*Pixel-art city-builder / tower-defense where **every action is triggered by typing.***  
Build a medieval settlement, train armies of knights, and repel orc hordes by mastering the keyboard.  
No mouse required – the entire game (UI **and** gameplay) is navigated with Vim / Qutebrowser-style keys.

| Key pillars | Short version |
|-------------|---------------|
| **Keyboard-First** | `h/j/k/l`, `/` search, `:` command mode – every screen is accessible by keys alone. |
| **Letter Streams** | Each building owns a cooldown; when it expires it queues a random word from its letter-pool. Typing that word finishes a construction step, reloads a tower, or trains a soldier. |
| **Per-Building Tech Trees** | Unlock letters one family at a time (Farmer → `f j`, Barracks → `f j d k`, …) to shorten cooldowns, add units, and widen the global letter pool. |
| **Endless Waves** | Orcs scale forever. Your throughput scales via new letters, faster cooldowns, smarter tech and player skill (accuracy + WPM). |
| **Pixel-art Charm** | SNES-style sprites & slap-stick combat (exploding cabbages, bouncing arrows, comic “BONK!” bubbles). Violence stays E10+ but feels impactful. |

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

go test ./...

```

## Current prototype

- Global FIFO queue, single Farmer + Barracks building (f j words).
- Basic orc grunt waves scale every 45 s.
- Typing speed/accuracy multiplier working.
- Vim navigation for pause/menu/shop implemented.

---

See docs/REQUIREMENTS.md for the full feature scaffold, ROADMAP.md for planned phases, and TODO.md for sprint tasks.

Happy typing & may your catapults stay jam-free!

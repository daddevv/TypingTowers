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
go test ./...
```

-## Current prototype

 - Shared FIFO queue manager implemented. Farmer and Barracks enqueue words (f j pool) that must be typed in order. Completing a Barracks word spawns a Footman.
- Basic orc grunt waves scale every 45 s.
- Typing speed/accuracy multiplier working.
- Vim navigation for pause/menu/shop implemented.

---

## Prototype Sprint

The current focus is on prototyping the Gathering (Farmer) and Military (Barracks) families, implementing a shared queue manager, per-building cooldowns, and playtesting word density.  
See the [ROADMAP.md](./ROADMAP.md) for detailed tasks.

---

## Farmer (Gathering Building)

- The Farmer building is now implemented as a Gathering structure that generates a word from its letter pool every cooldown cycle (default 3s).
- Typing the generated word completes the cycle and produces Food resources.
- Each building's cooldown timer pauses once a word is queued and only resets after that word is typed.
- The Farmer's cooldown, letter pool, and word length can be configured and extended for progression.
- See `v1/internal/game/farmer.go` for implementation and `farmer_test.go` for tests.

## Barracks (Military Building)

- The Barracks building generates a word from its letter pool every 5 seconds.
- Typing the generated word spawns a Footman unit.
- Word generation logic and cooldown behavior are tested in `barracks_test.go`.

See docs/REQUIREMENTS.md for the full feature scaffold, ROADMAP.md for planned phases, and TODO.md for sprint tasks.

Happy typing & may your catapults stay jam-free!

# Functional & Technical Requirements – Kingdoms of Keystoria

> **Goal:** Scaffold every planned mechanic so design & code can evolve without surprise scope creep.  
> Use this as a living contract; refine granularity during implementation.

---

## Expanded Project Vision

TypingTowers is a keyboard-controlled tower defense game with a focus on typing skill, Vim-style navigation, and city-building.  
The expanded vision includes:

- **Deep progression:** A massive skill tree with 100+ nodes, unlocking new towers, upgrades, and features.
- **Autonomous minions & heroes:** Summonable units managed by typing, with unique roles and upgrades.
- **Incremental & idle mechanics:** Optional auto-collection, offline progress, and prestige/reset systems.
- **Typing minigames:** Speed trials, accuracy challenges, word puzzles, and boss practice for rewards.
- **Multiple playstyle support:** Systems and upgrades for grinding, optimization, idle, and chaos-maximizing play.

All new features are optional enhancements, preserving the educational and accessible core.

---

## 1 Core Gameplay Systems

| ID | Requirement |
|----|-------------|
| **Q-QUEUE-1** | The game maintains a **single global word queue**. Buildings push words in FIFO order when their cooldowns hit 0. |
| **Q-QUEUE-2** | The next word must be typed **exactly once** – accuracy & WPM captured. Mismatched key → jam state until `Backspace` clears. |
| **Q-QUEUE-3** | If ≥ 5 words backlog → base takes chip damage each second (prevent AFK). |
| **Q-QUEUE-4** | Queue is rendered top center with colour coding by building family and word length caps (Basic 2-3, Power 4-6, Epic 6-8). |

### 1.1 Letter Pools

- **LP-1** Start pool = `f j`.  
- **LP-2** Letter unlock order is ergonomic (see `docs/LETTER_UNLOCKS.md`).  
- **LP-3** Each building stores its **own unlock pointer** so two buildings may request different letters at the same tech tier.  
- **LP-4** Future option: purchase letters directly with Gold (tech-tree node).  

---

## 2 Buildings & Families

| Family | Buildings | Primary Resource / Effect | Base CD |
|--------|-----------|---------------------------|---------|
| **Gathering** | Farmer (Food), Lumberjack (Wood), Miner (Stone+Iron) | +Resources | 3 s |
| **Military** | Barracks (melee), Archery Range (ranged), Stable (cavalry), Siege Workshop (siege) | Spawns units | 4–6 s |
| **Craft & Defense** | Blacksmith, Armorer, Library | Upgrades towers & soldiers | 5 s |
| **Economy** | Market | Multiplies Gold drop + passive trade | 4 s |
| **Spiritual** | Sanctum of Light | Generates Mana & spell words | 6 s |
| **Housing** | House ↔ Carpenter pair | Raises pop cap & global cooldown Δ | variable |

### Per-Building Specs

- **B-GEN-1** Each building has: `letter_pool[]`, `word_len_min/max`, `cooldown`, `base_power`, `tech_tree`.
- **B-GEN-2** Level ups may grant:  
  - Shorter CD, longer word (more output), additional effect (crit, AoE, heal).
- **B-GEN-3** Buildings can be paused (`p` key) – stops pushing words.
- **B-MIL-1** Barracks generates a word every cooldown and spawns a Footman unit when that word is completed.

---

## 3 Combat & Units

- **MOB-1** Wave spawner increments `wave_id`, HP ×1.25, DPS ×1.15 each wave.  
- **MOB-2** Families: Grunt, Armored, Fast, Shaman, Boss (10th wave).  
- **UNIT-1** Each trained unit consumes `Food` and spawns at left, moves right.  
- **DAMAGE** `Effective = base × accuracy × (1+speed_bonus)` (speed bonus tiers 0 / +25 / +50 %).  
- **CRIT** 100 % accuracy **and** top speed bonus → +50 % power.  
- **TOWER RELOADS** Work exactly like buildings: arrow tower pushes its own queue words.

*Planned: Autonomous minions and hero units can be summoned and managed via typing. Minions have unique roles and can be upgraded through the skill tree.*

---

## 4 Economy

| Resource | Earned by | Spent on |
|----------|-----------|----------|
| **Gold** | Killing mobs, Market multiplier | Building construction, upgrades |
| **Food** | Farmer cycles | Military unit upkeep |
| **Wood / Stone / Iron** | Lumberjack / Miner | Advanced buildings, towers |
| **Mana** | Sanctum cycles | Spells (AoE heal, smite) |
| **King’s Points** | Wave completion | Letter unlocks & global passives |

*Planned: Idle resource generators and auto-collection upgrades available via skill tree. Offline progress and prestige/reset mechanics as optional late-game systems.*

---

## 5 Tech Trees

- **TREE-1** Each building holds a **directed-graph tech tree** of nodes (YAML in `data/trees/`).  
- **TREE-2** Node types: `UnlockLetter`, `StatBoost`, `NewUnit`, `CDReduction`, `GlobalPassive`.  
- **TREE-3** Infinite spiral: on outer-ring completion, costs×2, effects×1.1.  

*Planned: Skill tree expands to 100+ nodes, with branches for offense, defense, typing proficiency, automation, and utility. Certain nodes require WPM/accuracy milestones.*

---

## 6 UI / Input

| ID | Requirement |
|----|-------------|
| **UI-NAV-1** | All menus navigable via `h/j/k/l`, arrows in Normie mode, plus `Enter`, `:`, `/`. |
| **UI-NAV-2** | Mode indicator visible (normal/insert/command). |
| **UI-NAV-3** | Global queue must be readable at 720p and 4k; colour-blind palette support. |
| **UI-NAV-4** | Hot-reload config `F5`. |
| **UI-NAV-5** | Key-rebinder supports QWERTY, Colemak, Dvorak. |

*Planned: Typed commands for minion management, skill tree navigation, and minigames. All new features remain keyboard-only.*

---

## 7 Sound / FX

- **SND-1** 8-bit SFX for key hits, crits, jams.  
- **SND-2** Optional voice-over reads next word for early readers (accessibility toggle).  

---

## 8 Persistence

- **SAVE-1** Auto-save every wave: letter unlock state, tech trees, resources.  
- **SAVE-2** Three local save-slots; JSON format, versioned.  

*Planned: Save skill tree, minion unlocks, idle progress, and minigame achievements.*

---

## 9 Testing & Tooling

- **TEST-1** Unit tests cover queue manager, letter generator, tech unlock gating, damage math, Vim navigation.  
- **TEST-2** Integration tests simulate 100 waves at 40 WPM 95 % accuracy – must not crash; TTK < Tsurvive.  
- **BENCH** Benchmark reload throughput vs 60 Hz update loop (Ebiten frame).  

---

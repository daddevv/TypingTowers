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
| **Q-QUEUE-5** | Queue items are processed letter-by-letter; words remain in the queue until all letters are typed correctly. |
| **Q-QUEUE-6** | Per-word accuracy and completion time are recorded and stored in a history log. |

### 1.1 Letter Pools

- **LP-1** Start pool = `f j`.  
- **LP-2** Letter unlock order is ergonomic (see `docs/LETTER_UNLOCKS.md`).  
- **LP-3** Each building stores its **own unlock pointer** so two buildings may request different letters at the same tech tier.  
- **LP-4** Future option: purchase letters directly with Gold (tech-tree node).
- **LP-5** Unlocking each letter stage costs King's Points and scales per stage (see `docs/LETTER_UNLOCKS.md`).

---

## 2 Buildings & Families

| Family | Buildings | Primary Resource / Effect | Base CD |
|--------|-----------|---------------------------|---------|
| **Gathering** | Farmer (Food), Lumberjack (Wood), Miner (Stone+Iron) | +Resources | 1.5 s |
| **Military** | Barracks (melee), Archery Range (ranged), Stable (cavalry), Siege Workshop (siege) | Spawns units | 2 s |
| **Craft & Defense** | Blacksmith, Armorer, Library | Upgrades towers & soldiers | 5 s |
| **Economy** | Market | Multiplies Gold drop + passive trade | 4 s |
| **Spiritual** | Sanctum of Light | Generates Mana & spell words | 6 s |
| **Housing** | House ↔ Carpenter pair | Raises pop cap & global cooldown Δ | variable |

### Per-Building Specs

- **B-GEN-1** Each building has: `letter_pool[]`, `word_len_min/max`, `cooldown`, `base_power`, `tech_tree`.
- **B-GEN-2** Level ups may grant:  
  - Shorter CD, longer word (more output), additional effect (crit, AoE, heal).
- **B-GEN-3** Buildings can be paused (`p` key) – stops pushing words.
- **B-GEN-4** Cooldown only resets after the building's pending word is completed.
- **B-GEN-5** Default Farmer and Barracks generate roughly 1–1.5 words per second combined.
- **B-MIL-1** Barracks generates a word every cooldown and spawns a Footman unit when that word is completed.

---

## 3 Combat & Units

- **MOB-1** Wave spawner increments `wave_id`, HP ×1.25, DPS ×1.15 each wave.  
- **MOB-2** Families: Grunt, Armored, Fast, Shaman, Boss (10th wave).  
- **UNIT-1** Each trained unit consumes `Food` and spawns at left, moves right.
- **UNIT-2** The Military system tracks spawned units and updates them each frame.
- **UNIT-3** Footmen have 10 HP, deal 1 damage, and move at speed 50 px/s.
- **UNIT-4** Orc Grunts have 5 HP, deal 1 damage, and march toward the base.
- **COMBAT-1** When a Footman and Orc Grunt overlap, each deals damage equal to its stat once per tick.
- **COMBAT-2** With perfect typing, a Footman must defeat an Orc Grunt in under eight seconds.
- **COMBAT-3** Units with 0 HP are removed immediately and cannot attack further.
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

### Economy Implementation Notes

- **ECO-1** The game maintains a `ResourcePool` struct aggregating Gold, Food, Wood, Stone and Iron.
- **ECO-2** Buildings like the Farmer deposit resources into this pool when their words are completed.
- **ECO-3** King's Points are tracked in the `ResourcePool` and spent on letter unlocks.

---

## 5 Tech Trees

- **TREE-1** Each building holds a **directed-graph tech tree** of nodes (YAML in `data/trees/`).  
- **TREE-2** Node types: `UnlockLetter`, `StatBoost`, `NewUnit`, `CDReduction`, `GlobalPassive`.  
- **TREE-3** Infinite spiral: on outer-ring completion, costs×2, effects×1.1.
- **TREE-4** Each YAML node defines `type`, `cost`, `effects`, and `prereqs` fields.
- **TREE-5** Tech trees are loaded at runtime from YAML via `LoadTechTree`.
- **TREE-6** The loader validates graphs, rejecting cycles or missing prerequisites.

*Planned: Skill tree expands to 100+ nodes, with branches for offense, defense, typing proficiency, automation, and utility. Certain nodes require WPM/accuracy milestones.*

- **SKILL-1** The global skill tree organizes nodes into Offense, Defense, Typing, Automation, and Utility categories.
- **SKILL-2** Go structs `SkillCategory`, `SkillNode`, and `SkillTree` define the skill tree in memory.
- **SKILL-3** `SampleSkillTree` provides a validated in-memory tree with example nodes for each category.
- **SKILL-4** Skills can be unlocked by spending King's Points once prerequisites are met.
- **SKILL-5** `Tab` opens the skill tree menu. Arrow keys switch categories and navigate nodes, highlighting the current selection and showing locked/unlocked status.
- **SKILL-6** Unlocking a skill applies its effects to the game (tower stats, base HP, etc.).

---

## 6 UI / Input

| ID | Requirement |
|----|-------------|
| **UI-NAV-1** | All menus navigable via `h/j/k/l`, arrows in Normie mode, plus `Enter`, `:`, `/`. |
| **UI-NAV-2** | Mode indicator visible (normal/insert/command). |
| **UI-NAV-3** | Global queue must be readable at 720p and 4k; colour-blind palette support. |
| **UI-NAV-4** | Hot-reload config `F5`. |
| **UI-NAV-5** | Key-rebinder supports QWERTY, Colemak, Dvorak. |
| **UI-BUILD-1** | HUD displays cooldown progress for Farmer and Barracks. |
| **UI-RES-1** | HUD shows resource icons for Gold, Wood, Stone, Iron and Mana. |
| **UI-QUEUE-1** | HUD displays the word queue with a conveyor belt animation. |
| **UI-QUEUE-2** | The first queued word shows typed letters in gray to indicate progress. |
| **UI-TITLE-1** | Game starts at a title screen with Start, Settings and Quit options. |
| **UI-TITLE-2** | Title screen has a simple animated background and is keyboard navigable. |
| **UI-PREGAME-1** | Pre-game setup screen handles character and difficulty selection, tutorial, typing test and mode selection. |
| **UI-STATE-1** | Game phases include MainMenu, PreGame, Playing, Paused, Settings and GameOver with keyboard navigation between them. |
| **UI-TOWER-1** | `/` enters tower selection mode, labels towers with letters and opens an upgrade menu for the chosen tower. |
| **UI-TOWER-2** | While selecting, the HUD overlays each tower with a yellow box and its letter label. |
| **UI-TOWER-2** | Tower selection mode dims the background and displays letter labels above towers. |
| **UI-TECH-1** | `/` toggles the tech menu. Typing filters nodes and `Enter` purchases the highlighted technology. |
| **UI-SKILL-1** | `F4` opens the global skill tree menu and arrow keys navigate categories and nodes. |
| **UI-CMD-1** | `:` enters command mode allowing text commands like `pause` or `quit`. |

*Planned: Typed commands for minion management, skill tree navigation, and minigames. All new features remain keyboard-only.*

---

## 7 Sound / FX

- **SND-1** 8-bit SFX for key hits, crits, jams.
- **FX-1** Mistyped letters briefly flash the screen red and play a "clank" sound.
- **SND-2** Optional voice-over reads next word for early readers (accessibility toggle).
- **SND-3** Background music loops on the title screen.

---

## 8 Persistence

- **SAVE-1** Auto-save every wave: letter unlock state, tech trees, resources.  
- **SAVE-2** Three local save-slots; JSON format, versioned.
- **SAVE-3** Skill tree unlocks persist in save files and are restored on load.
- **SAVE-4** Save file structure supports multiple slots and versioning for future compatibility.
- **SAVE-5** Save/load menu lets the player choose one of three slots.
- **SAVE-6** Version mismatches are detected and reported gracefully.

*Planned: Save skill tree, minion unlocks, idle progress, and minigame achievements.*

---

## 9 Testing & Tooling

- **TEST-1** Unit tests cover queue manager, letter generator, tech unlock gating, damage math, Vim navigation.  
- **TEST-2** Integration tests simulate 100 waves at 40 WPM 95 % accuracy – must not crash; TTK < Tsurvive.
- **TEST-3** A headless `Step` function (build tag `test`) allows end-to-end simulation of the core gameplay loop.
- **BENCH** Benchmark reload throughput vs 60 Hz update loop (Ebiten frame).

---

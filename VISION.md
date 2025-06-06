# **Kingdoms of Keystoria**

## *Pixel-art Knights vs Orcs — Keyboard-Powered City-Builder / Defense*

---

## 0 ▪ Expanded Project Vision

**TypingTowers** is a keyboard-controlled tower defense game where players defend their base against waves of enemies using towers that must be reloaded by typing.  
The expanded vision is to create a deep, replayable indie title with rich optional layers of gameplay, while maintaining keyboard-only controls and a focus on improving typing speed and accuracy.  
The game preserves its educational core and Vim/Qutebrowser-style navigation, adding progression, complexity, and lore to engage players for 100+ hours.  
All new features are optional enhancements, ensuring TypingTowers is welcoming to casual typists and challenging for hardcore players alike.

### Deep Progression and Complexity

- **Expansive Skill Tree:** A massive tech tree with 100+ nodes, grouped into offense, defense, typing proficiency, automation, and utility. Unlocks new towers, minions, upgrades, and game modes. Progression is tied to demonstrated typing skill (WPM/accuracy gates).
- **Autonomous Minions & Heroes:** Summon and command minions or heroes by typing keywords. Each minion has unique roles and can be upgraded or managed via typed commands, complementing towers and enriching strategy.
- **Incremental & Idle Mechanics:** Optional idle progression—auto-collection, offline progress, auto-reloading towers, and upgradable generators. Prestige/reset mechanics extend replayability.
- **Typing Minigames & Challenges:** Speed trials, accuracy challenges, word puzzles, and boss practice modes provide fun breaks, reinforce typing skills, and grant in-game rewards.
- **Multiple Playstyle Support:** Grind, optimize, idle, or embrace chaos—each playstyle is valid and rewarding, with systems and skill tree branches to support them.

---

## 1 ▪ Core Loop (Re-mapped to Buildings)

| Phase                         | What Fires the Letter-Stream                                                                             | What You Type                                                                    | What It Achieves                                                                                                     |
| ----------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| **Construction & Production** | Every **building** has an internal cooldown.<br>When it reaches 0 it *dumps a word* to the global queue. | Type the word once (accuracy & speed scale the result).                          | • Finish a wall segment / craft weapons / train a soldier.<br>• Cooldown resets, generating another word next cycle. |
| **Combat**                    | Enemy waves march while your towers & troops act.                                                        | Same queue — towers reload words; garrison skills may add extra words in battle. | Correct typing ⟶ arrows fired, catapults re-armed, etc.                                                              |
| **Upgrade & Tech**            | Spend **Gold, Wood, Stone, Iron, Mana** earned that wave.                                                | Navigate tech trees (keyboard-only menus).                                       | • Unlock new letters for each building track.<br>• Shorten cooldowns, lengthen words’ damage, add new abilities.     |
| **Next Wave**                 | Difficulty ↑; building timers continue.                                                                  | Keep typing: the queue never sleeps.                                             | Ever-faster letter stream as the city grows.                                                                         |

The **queue is unified** across all structures, so the player is constantly juggling production, construction, and defense by typing.

---

## 2 ▪ Building Families & Their Letter Tracks

| Family                | Buildings                                           | Base Cool-down | Starting Pool | Resource / Effect                               | Tech-Tree Focus                                                                                                                   |
| --------------------- | --------------------------------------------------- | -------------- | ------------- | ----------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| **Gathering**         | **Farmer**, Lumberjack, Miner                       | 3 s            | `f j`         | +Food / Wood / Stone each completion            | • Unlock *g h* → **Silo** auto-stores surplus.<br>• Cooldown ↓.                                                                   |
| **Military**          | **Barracks**, Archery Range, Stable, Siege Workshop | 4–6 s          | `f j`         | Trains footman / archer / rider / catapult      | • Unlock *d k s l* → elite units.<br>• “Battle Cry” nodes: extra troops per word if 100 % accuracy.                               |
| **Defense / Craft**   | **Blacksmith**, Armorer, Library                    | 5 s            | `f j`         | +Weapon/Armor tiers or tower upgrades           | • Unlock *a ; g h* → magical alloys.<br>• “Masterwork” crit chance on perfect speed.                                              |
| **Economy**           | **Market**                                          | 4 s            | `f j`         | Multiplies Gold drops                           | • Unlock *e i o* → trade caravans (Gold over time).                                                                               |
| **Spiritual**         | **Sanctum of Light** (Place of Worship)             | 6 s            | `f j`         | Generates **Mana** used by towers & hero spells | • Unlock *r u y t* → miracles (AoE heal, smite).                                                                                  |
| **Housing & Support** | **House** ↔ **Carpenter** (paired)                  | House: passive | —             | Raises Pop. cap → more workers                  | • Upgrading a House spawns a *Carpenter* word.<br>• Carpenter’s letter track upgrades **all other buildings’ cooldowns by –1 %**. |

> **Rule of Thumb**
> *Every building’s letter pool contains **only letters your city has globally unlocked**, but*\* each building has its own **order** of adding them.\*
> *F-and-J first for *every* building ➜ intuitive early game. After that, you choose which building’s next letter to research, creating strategic variety.*

---

## 3 ▪ Sample Letter Progression per Building

| Tier  | Farmer Pool                        | Barracks Pool                  | Blacksmith Pool                    |
| ----- | ---------------------------------- | ------------------------------ | ---------------------------------- |
| **0** | `f j`                              | `f j`                          | `f j`                              |
| **1** | `f j d`                            | `f j d`                        | `f j k`                            |
| **2** | `f j d k`                          | `f j d k s`                    | `f j k s`                          |
| **3** | `f j d k s l`                      | `f j d k s l`                  | `f j k s l g`                      |
| **4** | full home row unlocked → add *a ;* | add *g h* (unlock **Pikemen**) | add *a ;* (unlock **Runic Steel**) |

*The Barracks path gives *combat* pay-offs, so its letters may cost more Gator—err, **King’s Points** —to gate balance.*

---

## 4 ▪ Tech-Tree Node Examples (per Building)

**Barracks Tree**

```
 • Letter 'd' Unlock  (100 Gold, 50 Food)
 • Basic Footman     (auto after 'd')
 ├─ Quick-March I    (Cooldown –4 %)          (150 Gold)
 |   ├─ Letter 'k' Unlock (Wood + Iron)
 |   |   └─ Archer Company (requires Archery Range)
 |   └─ Battle-Hymn: +5 % dmg on 100 % accuracy
 └─ Letter 's' Unlock (300 Iron)
     ├─ Elite Footman
     └─ Rally: add extra word on crit
```

**Blacksmith Tree**

```
 • Letter 'k' Unlock  (Stone + Iron)
 ├─ Tempered Blades (+1 troop dmg)
 |   └─ Letter 's' Unlock
 |       └─ Masterwork Edge (+crit dmg)
 └─ Alloy Research (global tower armor +1)
```

*Each node can continue spiraling outward, doubling cost and effect size forever (infinite growth).*

---

## 5 ▪ Early-Game Timeline (First 8 Minutes)

| Time | Event                                        | New Typing Load                                                 |
| ---- | -------------------------------------------- | --------------------------------------------------------------- |
| 0:00 | Start with **Farmer** + **Barracks** (`f j`) | 1 word every 3–4 s                                              |
| 1:30 | Spend 100 Gold → Unlock `d` for Farmer       | Occasionally 3-letter crops                                     |
| 3:00 | Build **Lumberjack** (shares `f j`)          | Two words often overlap                                         |
| 4:30 | Unlock `d` for Barracks → trains Footman     | Three-letter barrage appears                                    |
| 6:00 | House → Carpenter chain created (`fj`)       | Carpenter buffers cooldown                                      |
| 7:30 | Build **Blacksmith** (`fj`)                  | Four-building stream; player enters near-continuous typing zone |

---

## 6 ▪ Combat & Damage Translation

- **Towers reload** via their own cooldown words (e.g., `fj`, `dfj`, …).
- **Troops** spawn automatically from Barracks words; they path to fight orcs in side-scroll style.
- **Orc Waves** escalate every 45 s. If the typing backlog > 5 words, the city gate starts taking chip damage (so the player can’t just let words pile).
- **Spell Triggers** from Sanctum appear as *blue-border* words; typing them casts heals or AoE smites.

> **Speed & Accuracy Bonus** — identical to earlier drafts (perfect & fast = +50 % power + crit).
> **Mistype Jam** — only the current word; other buildings keep feeding the queue, raising the pressure.

---

## 7 ▪ Pixel-Art & UI Sketch

```
[ Global Queue HUD ]
[ fj ]  [ dfj ]  [ fj ]  [ fj ]  [ sjlk ]

Castle Gate ███████░   Wave 5
Gold 143 | Wood 85 | Stone 60 | Iron 20 | Mana 5

FARMS   BARRACKS   LUMBER   MINES   MARKET
  1/1     1/1        0/1      0/1     0/1
(hjkl / arrows to select, Enter to open tech)
```

*Visual*: 16-bit SNES palette — rolling green hills, pixel knights, squat orcs, chunky word bubbles color-coded by building family.

---

## 8 ▪ Why This Fits Your Vision

- **Same joyful typing torrent** — now mapped to a full city-builder fantasy.
- **Per-building letter trees** give *dozens* of independent progression tracks; players choose order (strategy).
- **Resource fantasy** resonates with classic RTS (Wood, Stone, Iron) yet 100 % keyboard.
- **Pixel art knights-vs-orcs** is familiar, kid-friendly, and nostalgia-rich without meme baggage.
- **Deep progression and replayability** — skill trees, minions, idle mechanics, and minigames ensure TypingTowers can engage players for 100+ hours, supporting a wide range of playstyles and skill levels.

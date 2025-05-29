# TypeDefense Content Configuration Guide

This document explains how to add or edit game content (levels, mobs, worlds) using JSON configuration files. No code changes or recompilation are required—just edit the JSON files and restart the game.

## Content Directory Structure

All content configuration files are located in:

```
desktop/content/
    levels.json   # Level definitions
    mobs.json     # Mob (enemy) definitions
    worlds.json   # World/biome definitions
```

## Editing Content

### 1. Editing or Adding Mobs

- Open `desktop/content/mobs.json`.
- Each entry defines a mob type. Example:

```json
{
  "type": "BeachballMob",
  "spriteSheet": "assets/images/mob/mob_beachball_sheet.png",
  "frameRows": 1,
  "frameCols": 7,
  "frameWidth": 48,
  "frameHeight": 48,
  "frameDuration": 6,
  "defaultSpeed": 2.0,
  "letterFont": "Mob",
  "letterFontSize": 48,
  "minLetters": 3,
  "maxLetters": 7
}
```

- To add a new mob, copy an entry and adjust the fields. Add new sprite sheets to `assets/images/mob/`.
- `minLetters` and `maxLetters` control the range of letters per mob instance.

### 2. Editing or Adding Levels

- Open `desktop/content/levels.json`.
- Each entry defines a level. Example:

```json
{
  "name": "World 1",
  "difficulty": "Easy",
  "world": "BEACH",
  "startingLetters": ["a", "e", "i", "o", "u"],
  "possibleLetters": ["a", "e", "i", "o", "u", "f", "g", "h", "j"],
  "waves": [
    {
      "scoreThreshold": 5,
      "possibleLetters": ["a", "e", "i", "o", "u"]
    },
    {
      "scoreThreshold": 10,
      "possibleLetters": ["a", "e", "i", "o", "u", "f", "g"]
    },
    {
      "scoreThreshold": 15,
      "possibleLetters": ["a", "e", "i", "o", "u", "f", "g", "h", "j"]
    }
  ]
}
```

- `world`: Reference to a world/biome defined in `worlds.json`.
- `startingLetters`: Letters available at the start of the level.
- `possibleLetters`: All letters that can be unlocked in this level.
- `waves`: Each wave has a `scoreThreshold` (score required to complete the wave) and can optionally override `possibleLetters` for that wave.

### 3. Editing or Adding Worlds

- Open `desktop/content/worlds.json`.
- Each entry defines a world/biome. Example:

```json
{
  "name": "BEACH",
  "background": "assets/images/background/beach.png"
}
```

- Add new worlds/biomes and reference new backgrounds in `assets/images/background/`.

## Reloading Content

- After editing JSON files, restart the game to load the new or updated content.
- No recompilation is needed unless you add new code features.

## Tips

- Always validate your JSON for syntax errors.
- Keep asset paths correct and relative to the `desktop/` directory.
- For new mobs, ensure the sprite sheet and animation parameters match your image.

---

For more details on content structure, see the main `README.md` and `CONTRIBUTING.md`.

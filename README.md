# TypeDefense

TypeDefense is a web-based typing game designed to help players improve their typing skills through engaging gameplay. Players control a static character on the left side of the screen, defending against waves of enemies that spawn from the right. To defeat enemies, players must quickly and accurately type the words or phrases displayed above each enemy before they reach the player.

## Features

- Fast-paced wave-based mob defense gameplay
- Enemies spawn from the right and move toward the player
- Defeat enemies by typing their associated words or letters
- **Mobs now move faster by default for a more challenging experience.**
- **Mobs spawn at random vertical positions, adding variety to each wave.**
- **Each mob now displays a single letter. When the correct letter is typed, the mob disappears.**
- **Mobs now respond instantly to correct key presses (bug fix):** Mobs are defeated when the player types the correct letter or word, and input is cleared after a mob is defeated.
- **When a mob is defeated, an instant visual feedback effect (flash and particle burst) is triggered.**
- **A particle burst effect is now triggered at the mob or letter location on every correct keystroke, providing instant visual feedback.**
- **Mobs now always spawn fully off-screen on the right for a more polished experience.**
- **Planned: Multiple mobs can spawn at once, and their base speed will increase for greater challenge.**
- **Multiple mobs can now spawn at each interval:** The MobSpawner supports spawning more than one mob at a time for increased challenge. This is configurable per level or via code.
- Designed to improve typing speed and accuracy
- Built with TypeScript and Vite for a modern web experience
- Addictive action-challenge-reward loop with instant visual and audio feedback
- Real-time scores and combo multipliers are now displayed in the top-left corner of the game UI. The score updates on every correct keystroke, and the combo multiplier appears when greater than 1, rewarding consecutive correct inputs.
- Particle effects are now displayed and updated on each keystroke. The score and combo UI appear in the top-left, and a particle burst is triggered at the mob location for every correct letter typed. The combo multiplier increases with consecutive correct keystrokes and resets on mistakes, rewarding accuracy and speed.
- Tweened UI transitions and dramatic camera effects for key events
- Layered audio cues for typing, combos, and wave clearances
- Core game loop set up in `GameScene` with Player and InputHandler initialized and updated each frame
- **Mob and MobSpawner integrated into GameScene:**
  - `MobSpawner` handles spawning and management of enemy mobs (see `client/src/entities/MobSpawner.ts`).
  - `Mob` represents individual enemy entities (see `client/src/entities/Mob.ts`).
  - Both are now fully integrated and updated within the main game loop in `GameScene`.
- Added initial design and documentation for the `FingerGroupManager` class, which will track player progress and statistics across finger groups for the typing curriculum.
  - Implemented the `FingerGroupManager` class in `client/src/managers/fingerGroupManager.ts`.
  - Tracks player progress, key usage, accuracy, and speed for each finger group.
  - Provides methods to record key presses, retrieve stats, and determine mastery.
  - Uses curriculum-defined finger/key mappings for robust tracking.
- Integrated FingerGroupManager with the main game loop in `GameScene.ts`.
- Each key press is now recorded and mapped to its finger group using the curriculum mapping (`getKeyInfo`).
- This enables tracking of finger usage and progress for each finger group in real time.
- **Mobs now walk toward the player:** Each mob moves toward the player character's position (left side of the screen) after spawning.
- **Player health system:** The player has a visible health value above their sprite. When a mob reaches the player, the player loses health and the mob is removed. The game ends with a "Game Over" message if health reaches zero.
- **Improved mob input targeting and combo logic:** When multiple mobs are on screen, the game now targets the closest mob that matches the player's keypress. If the keypress doesn't match any mob, all mobs' progress is reset, fixing the combo bug with multiple mobs.
- Added collision detection and overlap prevention to mobs. Mobs now repel each other if they get too close, ensuring they do not overlap as they move toward the player. This improves gameplay clarity and visual polish.
- **Dynamic difficulty:** As the game progresses, the spawn rate of enemies increases and their movement speed scales up, providing a smooth and challenging difficulty curve.

## Curriculum Design

TypeDefense features a unique learning approach based on finger groups rather than random letters. The game is structured into four worlds, each focusing on a specific set of fingers:

### World 1: Index Fingers

- Left Hand: F, G, R, T, V, B
- Right Hand: J, H, Y, U, N, M
- Progressive levels introduce these keys gradually, starting with home row (F/J)

### World 2: Middle Fingers

- Left Hand: D, E, C
- Right Hand: K, I, comma
- Builds on index finger skills while introducing middle finger positions

### World 3: Ring Fingers

- Left Hand: S, W, X
- Right Hand: L, O, period
- Focuses on training the typically weaker ring fingers

### World 4: Pinky Fingers

- Left Hand: A, Q, Z
- Right Hand: semicolon, P, slash
- Completes the alphabet and introduces Shift key for capitals

Each world contains multiple levels that introduce letters progressively, with boss battles that test mastery before advancing. This finger-group approach builds proper muscle memory and typing technique.

## Recent Changes

- Added a `baseSpeed` property to `Mob` and updated the spawning system to allow setting mob speed per spawn. This enables more flexible and challenging gameplay tuning.
- Improved mob input handling: The game now targets the closest matching mob for each keypress, checks others if not matched, and resets all mobs if no match. This fixes the combo bug with multiple mobs on screen.
- **New mob targeting system:** If no mob is targeted, keypresses identify the closest matching mob. If a mob is targeted, the next keypress is aimed at them; if correct, the target advances, otherwise the system checks for other matches. The targeted mob is visually highlighted, and matched letters are animated to inactive so the player knows which letter is next.
- **World & Level Selection Menu:**
  - Added a new menu scene (`MenuScene.ts`) that allows players to select worlds and levels.
  - Levels are locked/unlocked based on completion status, which is tracked and persisted in local storage.
  - The menu UI displays locked, unlocked, and completed levels with appropriate visual cues.
  - Progression is saved and loaded automatically, so players can continue where they left off.
  - Selecting an unlocked level starts the game at that level.
- **Level 1-3 ("Reaching Up"):** Adds R and U (top row) to the available keys, with a new word pack (`fjghruWords.json`) and more letter combinations. The curriculum and word loader have been updated to support this level.
- Add a "Continue" button and Enter key handler to the level complete screen. When a level is completed, players can click the button or press Enter to advance to the next level (if available) or return to the level select menu.

### WordGenerator Utility

A new `WordGenerator` class is available in `client/src/utils/wordGenerator.ts`.

- Generates words using only the available letters (for curriculum-based levels).
- Supports random letter words, filtering valid words, and generating pronounceable pseudo-words (CVC pattern).
- Methods:
  - `generateWord(length)`: Generates a random word using available letters.
  - `filterValidWords(words)`: Filters a list of words to only those that can be made from available letters.
  - `generatePseudoWord(length)`: Generates a pronounceable pseudo-word (CVC pattern) using available letters.
  - `getWord(length)`: Returns a pseudo-word if enabled, otherwise a random word.
- Used for creating level-appropriate word challenges and drills.
- See `client/src/utils/__tests__/wordGenerator.test.ts` for usage examples and tests.

**Usage Example:**

```ts
import WordGenerator from './utils/wordGenerator';
const gen = new WordGenerator(['f', 'j', 'g', 'h']);
const word = gen.getWord(4); // e.g., 'fjgh', 'ghjf', etc.
const pseudo = gen.generatePseudoWord(5); // e.g., 'gafih', 'hajig', etc.
const valid = gen.filterValidWords(['fish', 'jag', 'hug']); // Only words using available letters
```

## Testing

To run unit tests (using Vitest):

```bash
cd client
npm run test
```

To run only the WordGenerator tests:

```bash
npm run test -- src/utils/__tests__/wordGenerator.test.ts
```

- The project uses [Vitest](https://vitest.dev/) for unit and integration testing.
- Test files are located in `client/src/**/__tests__/` and follow the `.test.ts` naming convention.
- To run all tests:

  ```bash
  cd client
  npm run test
  ```

- To run tests in watch/UI mode:

  ```bash
  cd client
  npm run test:ui
  ```

- Sample unit tests are provided for core modules:
  - `Mob`, `MobSpawner`, `InputHandler`, `FingerGroupManager`, `WordGenerator`, `LevelManager`, and `loadWordList` utility.
  - **New:** Integration tests for `MobSpawner` and mob spawning logic are included in `client/src/entities/__tests__/MobSpawner.test.ts` and cover multi-interval spawning, overlap prevention, scaling, and mob removal.
  - **New:** Comprehensive unit tests for `InputHandler` are located in `client/src/entities/__tests__/InputHandler.test.ts` and cover input accumulation, clearing, event registration, and cleanup.
  - **New:** Unit and integration tests for the score/combo UI and particle burst effect are included:
    - `client/src/entities/__tests__/ComboSystem.unit.test.ts` (combo/score logic and UI updates)
    - `client/src/scenes/__tests__/GameScene.combo.integration.test.ts` (particle burst triggers on correct keystroke)
- Unit tests for `FingerGroupManager` are located in `client/src/managers/__tests__/fingerGroupManager.test.ts` and cover:
  - Initialization of stats for all finger types
  - Recording key presses and updating stats
  - Calculating average speed for a finger
  - Retrieving correct keys for a finger
  - Determining if a key is mastered (accuracy and speed criteria)
- Unit tests for `LevelManager` are located in `client/src/managers/__tests__/levelManager.test.ts` and cover initialization, progress tracking, unlocking levels, and persistence via localStorage.
- Added unit and integration tests for the combo multiplier and score system. See `client/src/entities/__tests__/ComboSystem.unit.test.ts` and `client/src/scenes/__tests__/GameScene.combo.integration.test.ts`.

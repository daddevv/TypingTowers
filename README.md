# TypeDefense

TypeDefense is a typing-based tower defense game built with Ebiten. Players defend against waves of enemies by typing the letters displayed above each mob. The game features dynamic difficulty scaling, projectile-based combat feedback, and an endless mode that progressively introduces new letters and increases challenge.

## Gameplay

In TypeDefense, you face waves of enemies (mobs) that move across the screen from right to left. Each mob displays a sequence of letters above it. You must type the letters in the correct order to defeat the mob:

- **Red letters** indicate the current target letter you must type
- **White letters** are upcoming letters in the sequence  
- **Gray letters** have already been typed

When you type a letter correctly, a projectile fires from the player position to the targeted mob, providing visual feedback. Once all letters in a mob's sequence are typed and all projectiles hit, the mob is defeated and your score increases.

## Current Features

### Core Mechanics

- **Projectile Combat System**: Visual projectiles fire when letters are typed correctly
- **Immediate Input Response**: Letter states update instantly for responsive typing feel
- **Collision Detection**: Projectiles track and hit their intended targets
- **Rapid Typing Support**: Can type multiple letters quickly without missing inputs

### Game Modes

- **Endless Mode**: Continuous waves with increasing difficulty and expanding letter sets
- **Dynamic Letter Pool**: Starts with vowels (a,e,i,o,u) and unlocks consonants based on score
- **Progressive Difficulty**: Spawn rates increase as score grows (faster mob spawning)

### Mob System

- **BeachballMob**: Animated beach ball enemies with customizable letter sequences
- **Smart Spawning**: MobSpawner handles timing, letter generation, and difficulty scaling
- **Visual States**: Each letter has distinct visual states (target/active/inactive)
- **Death Animations**: Mobs play death animations when defeated

### Technical Features

- **Performance Optimization**: Global letter image cache prevents per-frame text rendering
- **Fixed Canvas**: 1920x1080 internal resolution with automatic scaling to window size
- **Parallel Processing**: Mob updates run in parallel for better performance
- **State Management**: Clean separation between menu, game, and pause states

## Architecture Overview

### Rendering System

TypeDefense uses a fixed 1920x1080 internal canvas for all gameplay, UI, and entity rendering. All coordinates are specified in this space, and the engine automatically scales the canvas to fit any window size while maintaining aspect ratio. This ensures pixel-perfect consistency across devices.

### Entity System

- **Entity Interface**: Common interface for all game objects (Player, Mobs, Projectiles)
- **Mob Interface**: Specialized interface for enemy types with letter management
- **Component-Based**: Entities have position, sprites, animations, and behavior components

### Letter Management

- **Letter States**: TARGET (red), ACTIVE (white), INACTIVE (gray)
- **Image Caching**: Pre-rendered letter images cached globally for performance
- **State Transitions**: Immediate letter state updates with visual projectile feedback

### Input System

- **InputHandler**: Processes keyboard input and manages projectile creation
- **Target Prioritization**: Automatically targets closest mob with matching letter
- **Rapid Typing**: Supports fast consecutive inputs without dropping keypresses

## Controls

- **A-Z Keys**: Type letters to target and defeat mobs
- **Space**: Force spawn a new mob (for testing)
- **Escape**: Pause game / Return to main menu
- **F**: Toggle fullscreen
- **Arrow Keys**: Navigate menus

## Game Progression

The game uses a score-based progression system:

- **Score**: Increases by 1 for each mob defeated
- **Letter Unlocks**: New letters unlock at specific score thresholds (every 10 points)
- **Spawn Rate**: Mob spawn intervals decrease as score increases
- **Difficulty Scaling**: Maintains challenge while gradually introducing complexity

## Current Development Status

TypeDefense has a fully functional core game loop with:

- ✅ Complete typing mechanics with visual feedback
- ✅ Projectile system with collision detection  
- ✅ Dynamic mob spawning and difficulty scaling
- ✅ Score tracking and letter pool expansion
- ✅ Performance-optimized rendering
- ✅ Responsive input handling for rapid typing

### Immediate Development Areas

- Enhanced visual effects and animations
- Additional mob types and behaviors
- Level progression and win/lose conditions
- Improved UI and menu systems
- Audio feedback and sound effects

See `TODO.md` for detailed development roadmap and priorities.

## Contributing

We welcome contributions! Whether you want to add new mob types, create worlds and biomes, implement game modes, or help with features, there are many ways to contribute.

**📖 [Read the Developer Contribution Guide](CONTRIBUTING.md)** - Comprehensive guide covering:

- How to add new mob types with custom behaviors
- Creating new worlds, biomes, and levels
- Implementing new game modes and features  
- Extending the letter and typing systems
- UI development and menu creation
- Testing guidelines and best practices

The codebase is designed to be modular and extensible, making it easy to add content without modifying core systems. Perfect for developers who want to contribute while learning game development!

## Building and Running

```bash
# Run the main game
go run ./cmd/game/main.go

# Preview sprite animations (development tool)
go run ./cmd/preview/main.go <image_path> <rows> <cols> <height> <width>
```

## Project Structure

```text
desktop/
├── cmd/
│   ├── game/main.go          # Main game entry point
│   └── preview/main.go       # Sprite animation preview tool
├── internal/
│   ├── engine/               # Game engine and state management
│   ├── game/                 # Core game logic and loop
│   ├── entity/               # Game entities (Player, Mobs, Projectiles)
│   ├── ui/                   # User interface components
│   ├── world/                # Level and world definitions
│   └── utils/                # Utility functions
└── assets/                   # Game assets (images, fonts, sounds)
```

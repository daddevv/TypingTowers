# TypeDefense Developer Contribution Guide

Welcome to TypeDefense! This guide will help new developers understand how to contribute content and features to the game. The codebase is designed to be modular and extensible, making it easy to add new mobs, worlds, game modes, and other content.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Project Architecture](#project-architecture)
3. [Adding New Mob Types](#adding-new-mob-types)
4. [Creating New Worlds and Biomes](#creating-new-worlds-and-biomes)
5. [Adding Game Modes](#adding-game-modes)
6. [Extending the Letter System](#extending-the-letter-system)
7. [UI and Menu Development](#ui-and-menu-development)
8. [Testing and Quality Assurance](#testing-and-quality-assurance)
9. [Best Practices](#best-practices)
10. [Common Patterns](#common-patterns)

## Getting Started

### Prerequisites
- Go 1.21 or later
- Basic understanding of the Ebiten game engine
- Familiarity with Go interfaces and struct embedding

### Development Setup
```bash
# Clone the repository
git clone <repository-url>
cd type-defense

# Run the game
go run ./cmd/game/main.go

# Preview sprite animations (useful for mob development)
go run ./cmd/preview/main.go <image_path> <rows> <cols> <height> <width>
```

### Code Style
- Follow standard Go conventions (gofmt, golint)
- Use meaningful variable and function names
- Add comments for public interfaces and complex logic
- Keep functions focused and modular

## Project Architecture

### Directory Structure
```
desktop/
├── cmd/
│   ├── game/           # Main game entry point
│   └── preview/        # Development tools
├── internal/
│   ├── engine/         # Core game engine and state management
│   ├── game/           # Game loop, input handling, main game logic
│   ├── entity/         # All game entities (Player, Mobs, Projectiles, etc.)
│   ├── ui/             # UI components, fonts, animations
│   ├── world/          # World definitions, levels, biomes
│   └── utils/          # Utility functions
└── assets/             # Game assets (images, fonts, sounds)
    ├── images/
    │   ├── mob/        # Mob sprite sheets
    │   ├── world/      # Background images
    │   └── ui/         # UI elements
    ├── fonts/          # Font files
    └── sounds/         # Audio files (future)
```

### Key Interfaces

**Entity Interface** - Base interface for all game objects:
```go
type Entity interface {
    Draw(screen *ebiten.Image)
    Update() error
    SetPosition(x, y float64)
    GetPosition() ui.Location
}
```

**Mob Interface** - Specialized interface for enemies:
```go
type Mob interface {
    Entity
    GetLetters() []Letter
    IsDead() bool
    IsPendingDeath() bool
    StartDeath()
    IncrementPendingProjectiles()
    GetPendingProjectiles() int
}
```

## Adding New Mob Types

Mobs are the core enemies that players type to defeat. Here's how to create a new mob type:

### 1. Create the Mob Struct

Create a new file in `internal/entity/` (e.g., `mob_spider.go`):

```go
package entity

import (
    "math/rand"
    "td/internal/ui"
    "github.com/hajimehoshi/ebiten/v2"
)

// SpiderMob represents a fast-moving spider enemy
type SpiderMob struct {
    MobBase  // Embed the common mob functionality
    // Add spider-specific fields here
    WebShots int // Number of web attacks remaining
}

// Ensure SpiderMob implements required interfaces
var _ Mob = (*SpiderMob)(nil)
var _ MobLetterController = (*SpiderMob)(nil)
```

### 2. Implement Required Methods

All mobs must implement the `Mob` interface. Most functionality is provided by `MobBase`, but you need to implement specific behaviors:

```go
// Constructor function
func NewSpiderMobWithLetters(letters []string) *SpiderMob {
    // Load sprite animation
    moveAnimation, err := ui.NewAnimation("assets/images/mob/spider_sheet.png", 1, 6, 32, 32, 8)
    if err != nil {
        return nil
    }
    
    // Create the mob with SpiderMob-specific properties
    m := &SpiderMob{
        MobBase: MobBase{
            Position:       ui.Location{X: 1920, Y: rand.Float64()*300 + 600},
            Speed:          3.0, // Spiders are faster than beach balls
            MoveAnimation:  moveAnimation,
            MoveTarget:     ui.Location{X: 100, Y: 900},
            Letters:        make([]Letter, len(letters)),
            // ... other standard fields
        },
        WebShots: 3, // Spider-specific field
    }
    
    // Initialize letters (standard pattern)
    font := ui.Font("Mob", 32)
    for i := range m.MobBase.Letters {
        char := []rune(letters[i])[0]
        if i == 0 {
            m.MobBase.Letters[i] = NewLetter(GetLetterImage(char, LetterTarget, font), LetterTarget, char)
        } else {
            m.MobBase.Letters[i] = NewLetter(GetLetterImage(char, LetterActive, font), LetterActive, char)
        }
    }
    
    // Calculate word width (standard pattern)
    for _, letter := range m.MobBase.Letters {
        m.WordWidth += float64(letter.Sprite.Bounds().Dx())
    }
    m.WordWidth += float64(len(m.MobBase.Letters)-1) * 24
    
    m.Sprite = m.MoveAnimation.Update()
    return m
}

// Override Update for custom behavior
func (m *SpiderMob) Update() error {
    // Call base update first
    if m.Dead {
        m.DeathTimer--
        if m.DeathTimer <= 0 {
            m.Position.X = -9999
        }
        return nil
    }
    
    // Spider-specific movement (zigzag pattern)
    m.Position.X -= m.Speed
    m.Position.Y += 30 * math.Sin(m.Position.X * 0.01) // Zigzag movement
    
    m.Sprite = m.MoveAnimation.Update()
    
    // Standard letter completion check
    allInactive := true
    for _, letter := range m.Letters {
        if letter.State != LetterInactive {
            allInactive = false
            break
        }
    }
    if allInactive && !m.Dead && !m.PendingDeath {
        m.PendingDeath = true
    }
    
    return nil
}

// Use standard implementations for most methods
func (m *SpiderMob) GetLetters() []Letter { return m.Letters }
func (m *SpiderMob) IsDead() bool { return m.Dead }
// ... other standard methods

// Custom drawing if needed
func (m *SpiderMob) Draw(screen *ebiten.Image) {
    // Draw spider with custom scaling or effects
    opts := ebiten.DrawImageOptions{}
    opts.GeoM.Scale(2, 2) // Smaller than beach balls
    opts.GeoM.Translate(m.Position.X, m.Position.Y)
    screen.DrawImage(m.Sprite, &opts)
    
    // Draw letters (standard pattern, or customize positioning)
    letterSpacing := 20.0 // Tighter spacing for smaller mobs
    baseX := m.Position.X + float64(m.Sprite.Bounds().Dx())*1.0 - m.WordWidth/2.0
    baseY := m.Position.Y - 30.0
    
    for i := 0; i < len(m.Letters); i++ {
        img := m.Letters[i].Sprite
        letterX := baseX + float64(i)*letterSpacing + float64(i)*float64(img.Bounds().Dx())
        letterY := baseY
        imgOpts := &ebiten.DrawImageOptions{}
        imgOpts.GeoM.Translate(letterX, letterY)
        screen.DrawImage(img, imgOpts)
    }
}

// Standard letter advancement (or customize if needed)
func (m *SpiderMob) AdvanceLetterState(char rune) {
    // Standard implementation - copy from BeachballMob
    // Most mobs will use the same letter advancement logic
}
```

### 3. Register with MobSpawner

In `internal/entity/mob_spawner.go`, add your mob to the spawning system:

```go
// In NewMobSpawner or wherever appropriate
spawner.RegisterMobFactory(func(letters []string) Mob { 
    return NewSpiderMobWithLetters(letters) 
})
```

### 4. Add Assets

Create sprite sheets in `assets/images/mob/`:
- Follow naming convention: `mob_[name]_sheet.png`
- Use consistent frame sizes
- Include animation frames for movement
- Consider death animation frames

### Mob Development Tips

- **Sprite Requirements**: Use sprite sheets with consistent frame sizes
- **Animation**: Most mobs need at least movement animation; death animations are optional
- **Scaling**: Use consistent scaling relative to the 1920x1080 canvas
- **Letter Positioning**: Adjust letter positioning based on mob size and sprite
- **Performance**: Keep Update() methods efficient as they run every frame
- **Unique Behaviors**: Add special movement patterns, abilities, or visual effects

## Creating New Worlds and Biomes

Worlds provide different visual environments and can have unique mob types, backgrounds, and themes.

### 1. Understanding the World System

The world system is defined in `internal/world/` and consists of:
- **Levels**: Individual game sessions with specific settings
- **Biomes**: Visual themes with backgrounds and environmental effects
- **World Definitions**: Collections of levels and progression

### 2. Adding a New Biome

Create a new biome in `internal/world/biome.go` (or create a new file):

```go
// Add to the Biome enum
const (
    BiomeForest Biome = iota
    BiomeDesert
    BiomeMountain
    BiomeOcean     // New biome
    BiomeSpace     // Another new biome
)

// Add to biomeName map
var biomeName = map[Biome]string{
    BiomeForest:   "Forest",
    BiomeDesert:   "Desert", 
    BiomeMountain: "Mountain",
    BiomeOcean:    "Ocean",    // New
    BiomeSpace:    "Space",    // New
}
```

### 3. Create Background Assets

Add background images to `assets/images/world/`:
- `ocean_background.png` - Main background image
- `ocean_layer1.png` - Optional parallax layer
- `ocean_layer2.png` - Optional parallax layer

### 4. Implement Biome Rendering

In the Level's `DrawBackground` method, add rendering for your biome:

```go
func (l Level) DrawBackground(screen *ebiten.Image) {
    switch l.Biome {
    case BiomeOcean:
        // Load and draw ocean background
        bg, err := ui.LoadImage("assets/images/world/ocean_background.png")
        if err == nil {
            opts := &ebiten.DrawImageOptions{}
            // Scale to fit 1920x1080 canvas
            screen.DrawImage(bg, opts)
        }
        
        // Add animated water effects, moving clouds, etc.
        
    case BiomeSpace:
        // Draw space background with moving stars
        // ...
    }
}
```

### 5. Create Themed Levels

Define levels that use your new biome:

```go
func NewOceanLevel() Level {
    return Level{
        Name:           "Ocean Depths",
        Biome:          BiomeOcean,
        Difficulty:     2,
        SpecialRules:   []string{"water_physics", "tide_effects"},
        MobTypes:       []string{"jellyfish", "shark", "seahorse"},
        BackgroundMusic: "ocean_ambient.ogg",
    }
}
```

### World Development Tips

- **Visual Consistency**: Ensure backgrounds work well with mob visibility
- **Performance**: Optimize background rendering for smooth gameplay
- **Parallax**: Use multiple layers for depth (optional but adds polish)
- **Color Schemes**: Choose colors that contrast well with letter visibility
- **Theme Integration**: Consider biome-specific mobs and sound effects

## Adding Game Modes

Game modes provide different gameplay experiences while using the core typing mechanics.

### 1. Understanding Game States

The game uses a state machine in `internal/engine/` with states like:
- `StateMenu` - Main menu
- `StateGame` - Core gameplay
- `StatePause` - Paused game
- `StateGameOver` - End of game

### 2. Creating a New Game Mode

To add a training mode, create `internal/game/training_mode.go`:

```go
package game

import (
    "td/internal/entity"
    "td/internal/ui"
    "github.com/hajimehoshi/ebiten/v2"
)

type TrainingMode struct {
    // Embed or compose with Game for shared functionality
    *Game
    
    // Training-specific fields
    TargetLetters    []string
    AccuracyTarget   float64
    TimeLimit        float64
    ElapsedTime     float64
    CorrectTyped    int
    TotalTyped      int
}

func NewTrainingMode(targetLetters []string, accuracyTarget float64, timeLimit float64) *TrainingMode {
    baseGame := NewGame(GameOptions{
        // Configure for training
    })
    
    return &TrainingMode{
        Game:           baseGame,
        TargetLetters:  targetLetters,
        AccuracyTarget: accuracyTarget,
        TimeLimit:      timeLimit,
        ElapsedTime:    0,
        CorrectTyped:   0,
        TotalTyped:     0,
    }
}

func (tm *TrainingMode) Update() error {
    // Update base game
    if err := tm.Game.Update(); err != nil {
        return err
    }
    
    // Training-specific logic
    tm.ElapsedTime += 1.0/60.0 // Assuming 60 FPS
    
    // Check completion conditions
    accuracy := float64(tm.CorrectTyped) / float64(tm.TotalTyped)
    if tm.ElapsedTime >= tm.TimeLimit || accuracy >= tm.AccuracyTarget {
        // Training complete - transition to results
        return fmt.Errorf("training_complete")
    }
    
    return nil
}

func (tm *TrainingMode) Draw(screen *ebiten.Image) {
    // Draw base game
    tm.Game.Draw(screen)
    
    // Add training-specific UI
    tm.drawTrainingHUD(screen)
}

func (tm *TrainingMode) drawTrainingHUD(screen *ebiten.Image) {
    // Display time remaining, accuracy, target letters, etc.
    accuracy := float64(tm.CorrectTyped) / float64(tm.TotalTyped) * 100
    timeRemaining := tm.TimeLimit - tm.ElapsedTime
    
    // Draw HUD elements using ui.DrawText or similar
}
```

### 3. Integrate with State Machine

In your main game state handler, add logic to handle the new mode:

```go
// In engine state management
case "training_mode":
    trainingMode := NewTrainingMode(selectedLetters, targetAccuracy, timeLimit)
    // Run training mode loop
```

### Game Mode Development Tips

- **Reuse Core Systems**: Build on top of existing Game struct when possible
- **Clear Objectives**: Make goals and progress visible to players
- **State Transitions**: Handle transitions to/from your mode cleanly
- **Save Progress**: Consider persistence for training progress
- **Customization**: Allow players to configure mode parameters

## Extending the Letter System

The letter system manages which letters are available for mobs and how they're unlocked.

### 1. Understanding LetterPool

The `LetterPool` interface in `internal/entity/entity.go` manages available letters:

```go
type LetterPool interface {
    GetPossibleLetters() []string
    Update(score int)
    AddLetter(letter string)
    AddLetters(letters []string)
}
```

### 2. Creating Custom Letter Pools

Create themed letter pools for specific game modes or worlds:

```go
// In internal/entity/letter_pool_themed.go
type ThemeLetterPool struct {
    theme        string
    baseLetters  []string
    themeWords   []string
    currentPool  []string
}

func NewThemeLetterPool(theme string) *ThemeLetterPool {
    var baseLetters []string
    var themeWords []string
    
    switch theme {
    case "ocean":
        baseLetters = []string{"a", "e", "i", "o", "u", "w", "v", "s", "h"}
        themeWords = []string{"wave", "fish", "sea", "dive", "swim"}
    case "space":
        baseLetters = []string{"a", "e", "i", "o", "u", "s", "t", "r", "n"}
        themeWords = []string{"star", "moon", "mars", "orbit", "comet"}
    }
    
    return &ThemeLetterPool{
        theme:       theme,
        baseLetters: baseLetters,
        themeWords:  themeWords,
        currentPool: baseLetters,
    }
}

func (tlp *ThemeLetterPool) GetPossibleLetters() []string {
    return tlp.currentPool
}

func (tlp *ThemeLetterPool) Update(score int) {
    // Custom unlocking logic for themed content
    if score >= 50 && len(tlp.currentPool) == len(tlp.baseLetters) {
        // Unlock all letters needed for theme words
        for _, word := range tlp.themeWords {
            for _, char := range word {
                letter := string(char)
                if !contains(tlp.currentPool, letter) {
                    tlp.currentPool = append(tlp.currentPool, letter)
                }
            }
        }
    }
}
```

### 3. Word-Based Content

Create word lists for specific themes or difficulty levels:

```go
// In internal/utils/word_lists.go
var WordLists = map[string][]string{
    "beginner": {"cat", "dog", "run", "fun", "sun"},
    "ocean":    {"wave", "fish", "coral", "deep", "blue"},
    "space":    {"star", "moon", "planet", "galaxy", "nebula"},
    "advanced": {"complex", "challenging", "sophisticated"},
}

func GetWordList(category string) []string {
    if words, exists := WordLists[category]; exists {
        return words
    }
    return WordLists["beginner"]
}
```

## UI and Menu Development

### 1. UI Components

UI components are in `internal/ui/`. Key patterns:

```go
// Button component example
type Button struct {
    Text     string
    Position ui.Location
    Size     ui.Size
    OnClick  func()
    Hover    bool
}

func (b *Button) Update(mouseX, mouseY float64, mousePressed bool) {
    // Check if mouse is over button
    b.Hover = mouseX >= b.Position.X && mouseX <= b.Position.X + b.Size.Width &&
              mouseY >= b.Position.Y && mouseY <= b.Position.Y + b.Size.Height
    
    // Handle click
    if b.Hover && mousePressed && b.OnClick != nil {
        b.OnClick()
    }
}

func (b *Button) Draw(screen *ebiten.Image) {
    // Draw button background, text, hover effects
}
```

### 2. Menu Systems

Create new menus by implementing the state pattern:

```go
// Training mode selection menu
type TrainingMenuState struct {
    buttons []Button
    selectedMode string
}

func (tms *TrainingMenuState) Update() error {
    // Update buttons, handle input
    return nil
}

func (tms *TrainingMenuState) Draw(screen *ebiten.Image) {
    // Draw menu background, buttons, etc.
}
```

## Testing and Quality Assurance

### 1. Manual Testing

- **Gameplay Testing**: Play your new content extensively
- **Edge Cases**: Test unusual scenarios (empty letter lists, extreme scores)
- **Performance**: Monitor frame rate with many entities
- **Visual Testing**: Check appearance on different screen sizes

### 2. Code Testing

```go
// Example unit test for new mob
func TestSpiderMobCreation(t *testing.T) {
    letters := []string{"s", "p", "i", "d", "e", "r"}
    spider := NewSpiderMobWithLetters(letters)
    
    if spider == nil {
        t.Fatal("Spider mob creation failed")
    }
    
    if len(spider.GetLetters()) != len(letters) {
        t.Errorf("Expected %d letters, got %d", len(letters), len(spider.GetLetters()))
    }
    
    // Test that first letter is target
    if spider.GetLetters()[0].State != LetterTarget {
        t.Error("First letter should be target state")
    }
}
```

### 3. Integration Testing

- Test new mobs in actual gameplay
- Verify new worlds render correctly
- Check that new game modes integrate with save/load systems

## Best Practices

### Code Organization

1. **Single Responsibility**: Each file/struct should have one clear purpose
2. **Interface Usage**: Use interfaces for extensibility (Mob, Entity, LetterPool)
3. **Composition**: Embed MobBase for common mob functionality
4. **Error Handling**: Handle errors gracefully, especially for asset loading

### Performance

1. **Asset Caching**: Cache images and fonts (see letter_state.go for example)
2. **Efficient Updates**: Keep Update() methods fast - they run every frame
3. **Memory Management**: Reuse objects when possible, avoid allocations in hot paths
4. **Profiling**: Use Go's built-in profiler to identify bottlenecks

### Asset Guidelines

1. **Naming Convention**: Use consistent naming (mob_[name]_sheet.png)
2. **Sprite Sheets**: Use sprite sheets instead of individual images
3. **Resolution**: Target the 1920x1080 internal canvas resolution
4. **Optimization**: Optimize images for size without losing quality

### Documentation

1. **Comment Interfaces**: Document all public interfaces clearly
2. **Example Usage**: Include examples in complex systems
3. **Change Logs**: Document breaking changes
4. **Asset Attribution**: Credit asset sources appropriately

## Common Patterns

### Entity Creation Pattern

```go
// 1. Define struct with embedded base
type NewMob struct {
    MobBase
    // Custom fields
}

// 2. Implement required interfaces
var _ Mob = (*NewMob)(nil)

// 3. Constructor with asset loading
func NewNewMobWithLetters(letters []string) *NewMob {
    // Load assets, initialize struct
}

// 4. Custom Update if needed
func (m *NewMob) Update() error {
    // Custom logic + base functionality
}
```

### State Management Pattern

```go
// 1. Define state struct
type NewGameState struct {
    // State-specific fields
}

// 2. Implement state interface
func (ngs *NewGameState) Update() error { /* ... */ }
func (ngs *NewGameState) Draw(screen *ebiten.Image) { /* ... */ }

// 3. Handle transitions
func (ngs *NewGameState) HandleTransition() (string, error) {
    // Return next state name or error
}
```

### Asset Loading Pattern

```go
// 1. Define asset paths as constants
const (
    AssetNewMobSprite = "assets/images/mob/new_mob_sheet.png"
    AssetNewBG        = "assets/images/world/new_background.png"
)

// 2. Load with error handling
func loadAssets() error {
    sprite, err := ui.LoadImage(AssetNewMobSprite)
    if err != nil {
        return fmt.Errorf("failed to load sprite: %w", err)
    }
    // Use sprite...
}
```

---

## Getting Help

- **Code Review**: Submit PRs for code review before merging
- **Questions**: Use GitHub issues for questions about implementation
- **Testing**: Test thoroughly on different systems before submitting
- **Documentation**: Update this guide when adding new patterns or systems

This guide should give you everything you need to start contributing content to TypeDefense. The codebase is designed to be extensible, so most new content can be added without modifying core systems. Happy coding!

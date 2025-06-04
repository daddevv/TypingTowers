# Project Roadmap

## Vision

A keyboard-focused tower defense game that combines strategic placement with typing mechanics, where accuracy and speed directly impact defensive capabilities.

## Current State (Implemented)

### Core Systems ✅

- **Wave-based enemy spawning** with configurable intervals and scaling
- **Auto-targeting tower system** that fires at closest enemies
- **Manual reload mechanics** requiring 'f' or 'j' key presses
- **Jamming system** that penalizes incorrect keystrokes
- **Projectile intercept calculations** for accurate targeting of moving enemies
- **Base health and game over conditions**
- **Real-time configuration reloading** for gameplay tuning

### Technical Foundation ✅

- Entity system with BaseEntity for shared behavior
- Input handling abstraction with InputHandler interface
- Modular game state management
- Unit test framework with example tests
- Asset loading and rendering pipeline

## Phase 1: Enhanced Typing Mechanics 🎯

### Letter-Based Targeting

- **Enemy Letters**: Each enemy displays a letter that must be typed to target it
- **Selective Targeting**: Tower fires only when the correct enemy letter is typed
- **Letter Pools**: Start with home row keys, expand as players progress

### Advanced Reload System

- **Word-Based Reloading**: Replace single letters with short words (e.g., "reload", "fire")
- **Typing Speed Bonuses**: Faster typing reduces reload time
- **Accuracy Penalties**: Mistakes increase reload time or reduce effectiveness

## Phase 2: Strategic Depth 📈

### Multiple Tower Types

- **Sniper Tower**: High damage, slow fire rate, requires precise typing
- **Rapid Fire Tower**: Low damage, fast fire rate, simple key sequences
- **Splash Tower**: Area damage, requires complex typing patterns

### Resource Management

- **Gold Economy**: Expand current gold system for meaningful upgrades
- **Tower Placement**: Allow multiple towers with strategic positioning
- **Upgrade Paths**: Damage, range, ammunition capacity, reload speed

## Phase 3: Progression Systems 🏆

### Technology Tree

- **Letter Unlocks**: New letters unlock new tower types and abilities
- **Typing Challenges**: Specific key combinations unlock advanced features
- **Skill Gates**: Advanced content requires demonstrated typing proficiency

### Scoring and Metrics

- **Typing Statistics**: Words per minute, accuracy percentage, error tracking
- **Performance Bonuses**: High accuracy multiplies score and gold rewards
- **Leaderboards**: Track best wave survival and typing performance

## Phase 4: Content Expansion 🌟

### Game Modes

- **Endless Mode**: Infinite waves with exponential difficulty scaling
- **Challenge Modes**: Specific typing constraints or enemy patterns
- **Tutorial Mode**: Guided introduction to mechanics and typing skills

### Quality of Life

- **Custom Key Bindings**: Support for different keyboard layouts
- **Accessibility Options**: Colorblind support, font size options
- **Save System**: Persistent progress and unlocks

## Technical Priorities

### Immediate (Next 2-4 weeks)

- Implement letter-based enemy targeting
- Add multiple tower placement system
- Expand shop functionality with real upgrades
- Create comprehensive unit test coverage

### Medium Term (1-3 months)

- Design and implement technology tree
- Add typing performance metrics
- Create multiple enemy and tower types
- Implement save/load system

### Long Term (3+ months)

- Add multiplayer/competitive modes
- Create level editor for custom scenarios
- Implement mod support for community content
- Port to additional platforms

## Success Metrics

- **Engagement**: Players complete multiple waves consistently
- **Skill Development**: Observable improvement in typing speed/accuracy
- **Replayability**: Multiple viable strategies and tower combinations
- **Community**: Active sharing of strategies and custom content

## Design Philosophy

- **Keyboard First**: All interactions should feel natural on keyboard
- **Progressive Difficulty**: Gentle introduction with meaningful skill gates
- **Immediate Feedback**: Clear visual and audio cues for all actions
- **Meaningful Choice**: Every upgrade and strategy should feel impactful

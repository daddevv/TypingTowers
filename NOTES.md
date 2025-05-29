# Drawing and Animation

`Draw(screen *ebiten.Image)` methods are used to call render functions that draw the game objects on the screen.

`Render(x, y float64, screen *ebiten.Image)` methods are used to draw the game objects on the screen.
x and y coordinates are 0-1 normalized values where 0,0 is the top left corner of the screen and 1,1 is the bottom right corner of the screen.

# Mob Letter State System (2025)

- Each mob displays a sequence of letters above it. Each letter can be in one of three states:
  - **TARGET**: The current letter to type (highlighted)
  - **ACTIVE**: Not yet typed, not the current target
  - **INACTIVE**: Already typed
- For performance, the game uses a global cache of pre-rendered images for each letter in each state. This avoids per-frame text rendering and allows for fast switching of letter states.
- When all letters in a mob are INACTIVE, the mob starts a death animation and is removed from the game.

# Architecture Notes

- Mobs track their own letter states and update them as the player types.
- The letter image cache is shared and keyed by rune and state.
- The game loop manages mob spawning, updating, and removal.

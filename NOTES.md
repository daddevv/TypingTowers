# Drawing and Animation

`Draw(screen *ebiten.Image)` methods are used to call render functions that draw the game objects on the screen.

`Render(x, y float64, screen *ebiten.Image)` methods are used to draw the game objects on the screen.
x and y coordinates are 0-1 normalized values where 0,0 is the top left corner of the screen and 1,1 is the bottom right corner of the screen.

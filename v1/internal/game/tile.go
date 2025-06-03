package game

// tileAtPosition returns the tile coordinates at the given screen position.
func tileAtPosition(x, y int) (int, int) {
	tileX := x / 32
	tileY := (y - 28) / 32
	return tileX, tileY
}

// tilePosition returns the screen position to draw a tile at the given tile coordinates.
func tilePosition(tileX, tileY int) (int, int) {
	x := tileX * 32    // Center of the tile
	y := tileY*32 + 28 // Center of the tile, adjusted for the top margin
	return x, y
}

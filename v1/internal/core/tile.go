package core

var (
	TileSize  = 32
	TopMargin = 28 // Top margin for the grid
)

// tileAtPosition returns the tile coordinates at the given screen position.
func TileAtPosition(x, y int) (int, int) {
	tileX := x / TileSize
	tileY := (y - TopMargin) / TileSize
	return tileX, tileY
}

// tilePosition returns the screen position to draw a tile at the given tile coordinates.
func TilePosition(tileX, tileY int) (int, int) {
	x := tileX * TileSize           // Center of the tile
	y := tileY*TileSize + TopMargin // Center of the tile, adjusted for the top margin
	return x, y
}

package tiles

import (
	"typingtowers/internal/content"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileMap struct {
	TileSize   int       // Size of each tile in pixels (assuming square tiles)
	Width      int       // Number of tiles in the horizontal direction
	Height     int       // Number of tiles in the vertical direction
	TopMargin  int       // Margin at the top of the tile map
	LeftMargin int       // Margin at the left of the tile map
	Tiles      [][]*Tile // 2D slice of tiles, where each tile is a pointer to a Tile struct
}

// NewTileMap creates a new TileMap with the specified dimensions and margins.
func NewTileMap(width, height, topMargin, leftMargin int, tileSize int) *TileMap {
	tileMap := &TileMap{
		TileSize:   tileSize,
		Width:      width,
		Height:     height,
		TopMargin:  topMargin,
		LeftMargin: leftMargin,
		Tiles:      make([][]*Tile, height),
	}

	for i := range tileMap.Tiles {
		tileMap.Tiles[i] = make([]*Tile, width)
	}

	return tileMap
}

// InitializeTiles initializes the tile map with default tiles.
func (tm *TileMap) InitializeTiles() {
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			// Create a default tile (you can customize this as needed)
			tm.Tiles[y][x] = &Tile{
				ID:       y*tm.Width + x,                       // Unique ID based on position
				Name:     "Tile_" + string(rune(y*tm.Width+x)), // Name based on position
				Image:    content.DefaultTile,                  // Placeholder for tile image, can be set later
				Walkable: true,                                 // Default to walkable
			}
		}
	}
}

// Draw draws the tile map on the provided image.
func (tm *TileMap) Draw(screen *ebiten.Image) {
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			tile := tm.Tiles[y][x]
			if tile != nil && tile.Image != nil {
				// Calculate the pixel position considering the margins
				px := x*tm.TileSize + tm.LeftMargin
				py := y*tm.TileSize + tm.TopMargin
				// Draw the tile image at the calculated position
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(px), float64(py))
				screen.DrawImage(tile.Image, op)
			}
		}
	}
}

// SetTile sets a tile at the specified coordinates in the tile map.
func (tm *TileMap) SetTile(x, y int, tile *Tile) {
	if x < 0 || x >= tm.Width || y < 0 || y >= tm.Height {
		return // Out of bounds
	}
	tm.Tiles[y][x] = tile
}

// GetTile retrieves the tile at the specified coordinates in the tile map.
func (tm *TileMap) GetTile(x, y int) *Tile {
	if x < 0 || x >= tm.Width || y < 0 || y >= tm.Height {
		return nil // Out of bounds
	}
	return tm.Tiles[y][x]
}

// Clear clears the tile map by setting all tiles to nil.
func (tm *TileMap) Clear() {
	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			tm.Tiles[y][x] = nil
		}
	}
}

// Dimensions returns the dimensions of the tile map in pixels, considering the margins.
func (tm *TileMap) Dimensions(tileWidth, tileHeight int) (int, int) {
	width := tm.Width*tileWidth + tm.LeftMargin*2
	height := tm.Height*tileHeight + tm.TopMargin*2
	return width, height
}

// TileAt returns the tile at the specified pixel coordinates, considering the margins.
func (tm *TileMap) TileAt(px, py, tileWidth, tileHeight int) (*Tile, int, int) {
	x := (px - tm.LeftMargin) / tileWidth
	y := (py - tm.TopMargin) / tileHeight

	if x < 0 || x >= tm.Width || y < 0 || y >= tm.Height {
		return nil, -1, -1 // Out of bounds
	}

	tile := tm.GetTile(x, y)
	return tile, x, y
}

// PixelToTile converts pixel coordinates to tile coordinates, considering the margins.
func (tm *TileMap) PixelToTile(px, py, tileWidth, tileHeight int) (int, int) {
	x := (px - tm.LeftMargin) / tileWidth
	y := (py - tm.TopMargin) / tileHeight

	if x < 0 || x >= tm.Width || y < 0 || y >= tm.Height {
		return -1, -1 // Out of bounds
	}

	return x, y
}

// TileToPixel converts tile coordinates to pixel coordinates, considering the margins.
func (tm *TileMap) TileToPixel(x, y, tileWidth, tileHeight int) (int, int) {
	if x < 0 || x >= tm.Width || y < 0 || y >= tm.Height {
		return -1, -1 // Out of bounds
	}

	px := x*tileWidth + tm.LeftMargin
	py := y*tileHeight + tm.TopMargin

	return px, py
}

// Size returns the size of the tile map in tiles.
func (tm *TileMap) Size() (int, int) {
	return tm.Width, tm.Height
}

// SetSize sets the size of the tile map in tiles.
func (tm *TileMap) SetSize(width, height int) {
	if width < 0 || height < 0 {
		return // Invalid size
	}

	tm.Width = width
	tm.Height = height

	// Resize the tile map
	tm.Tiles = make([][]*Tile, height)
	for i := range tm.Tiles {
		tm.Tiles[i] = make([]*Tile, width)
	}
}

// GetTileMap returns the tile map.
func (tm *TileMap) GetTileMap() [][]*Tile {
	return tm.Tiles
}

// SetTileMap sets the tile map to the provided 2D slice of tiles.
func (tm *TileMap) SetTileMap(tiles [][]*Tile) {
	if len(tiles) != tm.Height || len(tiles[0]) != tm.Width {
		return // Invalid tile map size
	}

	tm.Tiles = tiles
}

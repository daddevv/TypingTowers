package tilemap

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Cliff is a globally available tilemap, autoloaded on package import.
var (
	Cliff = loadCliffTiles() // Global variable to hold the cliff tiles
)

func loadCliffTiles() map[string]*ebiten.Image {
	cliffTileMap := make(map[string]*ebiten.Image)
	scale := 3 // Scale factor for the tiles

	// Load the cliff tiles from the image file
	cliffImage, _, err := ebitenutil.NewImageFromFile("assets/sprites/background/cliff.png")
	if err != nil {
		panic(err)
	}

	// Define the positions of the tiles in the cliff image
	tileWidth := 16  // 16 pixels wide
	tileHeight := 16 // 16 pixels tall
	positions := []struct {
		name string
		x, y int
	}{
		{"top_left", 0, 0},
		{"top", tileWidth, 0},
		{"top_right", 2 * tileWidth, 0},
		{"left", 0, tileHeight},
		{"center", tileWidth, tileHeight},
		{"right", 2 * tileWidth, tileHeight},
		{"bottom_left", 0, 2 * tileHeight},
		{"bottom", tileWidth, 2 * tileHeight},
		{"bottom_right", 2 * tileWidth, 2 * tileHeight},
	}
	// Extract each tile from the cliff image and store it in the map
	for _, pos := range positions {
		subImg := cliffImage.SubImage(image.Rect(pos.x, pos.y, pos.x+tileWidth, pos.y+tileHeight)).(*ebiten.Image)
		// Scale the tile image
		scaledTile := ebiten.NewImage(tileWidth*scale, tileHeight*scale)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Scale(float64(scale), float64(scale)) // Scale the image
		scaledTile.DrawImage(subImg, opts)
		cliffTileMap[pos.name] = scaledTile
	}

	return cliffTileMap
}

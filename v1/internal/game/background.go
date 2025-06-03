package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	BACKGROUND_GRID = generateBackground(1920, 1080, 32, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	})
)

// generateBackground creates a background image
func generateBackground(width, height, gridSize int, color color.RGBA) *ebiten.Image {
	bg := ebiten.NewImage(width, height)
	tile, _, err := ebitenutil.NewImageFromFile("assets/basic_tile_32.png")
	if err != nil {
		panic(err)
	}
	for x := range 60 {
		for y := range 32 {
			tileX, tileY := tilePosition(x, y)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tileX), float64(tileY))
			bg.DrawImage(tile, op)
		}
	}

	return bg
}

package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ImgBackgroundBasicTiles = generateBackground()
	ImgBackgroundTile       = loadImage("assets/basic_tile_32.png")
	ImgHighlightTile        = loadImage("assets/basic_tile_highlight_32.png")
	ImgHouseTile            = loadImage("assets/blue_house_32.png")
)

// loadImage is the utility function to load an image from a file path.
func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}
	imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()
	log.Println("Loaded image:", path, "(", imgWidth, "x", imgHeight, ")")
	return img
}

// generateBackground creates a background image
func generateBackground() *ebiten.Image {
	bg := ebiten.NewImage(1920, 1080)
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

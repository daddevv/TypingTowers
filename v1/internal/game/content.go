package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ImgBackgroundTile = loadImage("assets/basic_tile_32.png")
	ImgHighlightTile  = loadImage("assets/basic_tile_highlight_32.png")
	ImgHouseTile      = loadImage("assets/blue_house_32.png")
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

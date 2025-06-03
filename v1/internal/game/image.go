package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ImgBackgroundTile = loadImage("assets/basic_tile_32.png")
	ImgHighlightTile  = loadImage("assets/basic_tile_highlight_32.png")
)

func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic(err)
	}
	return img
}

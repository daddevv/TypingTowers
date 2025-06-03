package game

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ImgBackgroundBasicTiles = generateBackground()
	ImgBackgroundTile       = loadImage("assets/basic_tile_32.png")
	ImgHighlightTile        = loadImage("assets/basic_tile_highlight_32.png")
	ImgHouseTile            = loadImage("assets/blue_house_32.png")
	ImgTower                = generateTowerImage()
	ImgMobA                 = generateMobImage(color.RGBA{255, 0, 0, 255})
	ImgMobB                 = generateMobImage(color.RGBA{255, 128, 0, 255})
	ImgProjectile           = generateProjectileImage()
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

func generateTowerImage() *ebiten.Image {
	img := ebiten.NewImage(32, 32)
	img.Fill(color.RGBA{0, 0, 200, 255})
	return img
}

func generateMobImage(c color.Color) *ebiten.Image {
	img := ebiten.NewImage(32, 32)
	img.Fill(c)
	return img
}

func generateProjectileImage() *ebiten.Image {
	img := ebiten.NewImage(8, 8)
	clr := color.RGBA{255, 255, 0, 255}
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			dx := x - 4
			dy := y - 4
			if dx*dx+dy*dy <= 16 {
				img.Set(x, y, clr)
			}
		}
	}
	return img
}

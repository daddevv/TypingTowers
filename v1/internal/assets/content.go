package assets

import (
	"image/color"
	"log"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	// AssetPrefix allows tests to override asset path prefix.
	AssetPrefix = ""

	ImgBackgroundBasicTiles *ebiten.Image
	ImgBackgroundTile       *ebiten.Image
	ImgHighlightTile        *ebiten.Image
	ImgHouseTile            *ebiten.Image
	ImgBase                 *ebiten.Image
	ImgTower                *ebiten.Image
	ImgMobA                 *ebiten.Image
	ImgMobB                 *ebiten.Image
	ImgFootman              *ebiten.Image
	ImgProjectile           *ebiten.Image
)

// InitImages loads all image assets. Must be called after setting assetPrefix.
func InitImages() {
	ImgBackgroundTile = loadImage("assets/basic_tile_32.png")
	ImgHighlightTile = loadImage("assets/basic_tile_highlight_32.png")
	ImgHouseTile = loadImage("assets/blue_house_32.png")
	ImgBase = generateBaseImage()
	ImgTower = generateTowerImage()
	ImgMobA = generateMobImage(color.RGBA{255, 0, 0, 255})
	ImgMobB = generateMobImage(color.RGBA{255, 128, 0, 255})
	ImgFootman = generateMobImage(color.RGBA{0, 0, 255, 255})
	ImgProjectile = generateProjectileImage()
	ImgBackgroundBasicTiles = ebiten.NewImage(1920, 1080)
}

// loadImage is the utility function to load an image from a file path.
func loadImage(path string) *ebiten.Image {
	path = filepath.FromSlash(AssetPrefix + path) // Use exported AssetPrefix
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		// try an extra level up in case of test environment
		fp := filepath.Base(path)
		path = "/home/bobbitt/projects/public/TypingTowers/v1/assets/images/" + fp
		img, _, err = ebitenutil.NewImageFromFile(path)
		if err != nil {
			log.Fatal("failed to load image:", path, ":", err)
		}
	}
	// imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()
	// log.Println("Loaded image:", path, "(", imgWidth, "x", imgHeight, ")")
	return img
}

func generateBaseImage() *ebiten.Image {
	w, h := 96, 64
	img := ebiten.NewImage(w, h)
	clr := color.RGBA{0, 128, 128, 255}
	topWidth := 40
	bottomWidth := 80
	for y := 0; y < h; y++ {
		t := float64(y) / float64(h-1)
		rowW := int(float64(topWidth) + (float64(bottomWidth-topWidth) * t))
		startX := w/2 - rowW/2
		for x := startX; x < startX+rowW; x++ {
			img.Set(x, y, clr)
		}
	}
	return img
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

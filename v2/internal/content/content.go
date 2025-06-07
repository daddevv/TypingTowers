package content

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	// This is where images, sounds, and other game assets will be stored.
	// For now, this is a placeholder.
	RedSquare = generateSquare(255, 0, 0, 255, 50, 50)

	DefaultTile = loadImage("assets/image/basic_tile_32.png")
)

func InitializeContent() {
	// Initialize game content, such as loading images, sounds, etc.
	// This function can be expanded to load various game assets.
	// For now, it is a placeholder.
	// Example: LoadImages(), LoadSounds(), etc.
}

func loadImage(path string) *ebiten.Image {
	// This function loads an image from the specified path.
	// It is a placeholder for the actual image loading logic.
	// In a real implementation, you would use an image library to load the image.
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		panic("failed to load image: " + path + ": " + err.Error())
	}
	return img
}

func generateSquare(r, g, b, a, width, height int) *ebiten.Image {
	// This function generates a square image with the specified color and dimensions.
	// It is a placeholder for the actual image generation logic.
	// In a real implementation, you would use an image library to create an image.
	img := ebiten.NewImage(width, height)
	img.Fill(color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
	return img
}

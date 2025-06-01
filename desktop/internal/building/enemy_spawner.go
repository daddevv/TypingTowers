package building

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	MobSpawnerLevel1 = loadMobSpawnerLevel1() // Load the enemy spawner sprite
)

func loadMobSpawnerLevel1() *ebiten.Image {
	// Load the enemy spawner sprite from a file or resource
	// For now, we will return a placeholder image
	sprite, _, err := ebitenutil.NewImageFromFile("assets/miniworld/building_enemy_spawner_level1.png")
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}
	scale := 2.0 // Scale factor for the sprite
	scaledSprite := ebiten.NewImage(int(float64(sprite.Bounds().Dx())*scale), int(float64(sprite.Bounds().Dy())*scale))
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scale, scale) // Scale the image by the scale factor
	scaledSprite.DrawImage(sprite, opts)
	// Return the scaled sprite
	sprite = scaledSprite
	return sprite
}

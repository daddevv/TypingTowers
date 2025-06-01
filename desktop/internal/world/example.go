package world

import (
	"td/internal/tilemap"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	Background *ebiten.Image // Background image for the world
}

var (
	Example = generateExampleWorld()
)

func generateExampleWorld() *World {
	bg := ebiten.NewImage(1920, 1080)

	islandWidth := 40     // Number of tiles horizontally (max 39)
	islandHeight := 18    // Number of tiles vertically (max 22)
	verticalOffset := 2   // Number of tiles to offset vertically
	horizontalOffset := 2 // Number of tiles to offset horizontally

	for y := verticalOffset * 48; y < 1080; y += 48 {
		for x := horizontalOffset * 48; x < 1920; x += 48 {
			tileX := x / 48
			tileY := y / 48
			var tileImg *ebiten.Image
			if tileX == horizontalOffset && tileY == verticalOffset {
				tileImg = tilemap.Cliff["top_left"]
			} else if tileX == horizontalOffset && tileY == verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["bottom_left"]
			} else if tileX == horizontalOffset+islandWidth-1 && tileY == verticalOffset {
				tileImg = tilemap.Cliff["top_right"]
			} else if tileX == horizontalOffset+islandWidth-1 && tileY == verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["bottom_right"]
			} else if tileY == verticalOffset && tileX > horizontalOffset && tileX < horizontalOffset+islandWidth-1 {
				tileImg = tilemap.Cliff["top"]
			} else if tileY == verticalOffset+islandHeight-1 && tileX > horizontalOffset && tileX < horizontalOffset+islandWidth-1 {
				tileImg = tilemap.Cliff["bottom"]
			} else if tileX == horizontalOffset && tileY > verticalOffset && tileY < verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["left"]
			} else if tileX == horizontalOffset+islandWidth-1 && tileY > verticalOffset && tileY < verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["right"]
			} else if tileX > horizontalOffset && tileX < horizontalOffset+islandWidth-1 && tileY > verticalOffset && tileY < verticalOffset+islandHeight-1 {
				tileImg = tilemap.Cliff["center"] // Default to center tile
			}

			if tileImg != nil {
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Translate(float64(x), float64(y)) // Position the tile
				bg.DrawImage(tileImg, opts)
			}
		}
	}

	return &World{
		Background: bg,
	}
}

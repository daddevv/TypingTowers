package main

import (
	"log"
	"os"

	"td/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// VERSION is the current version of the application.
	VERSION = "0.1.0"
)

func init() {
	ebiten.SetTPS(100) // Set the target frames per second
}

func main() {
	canvasWidth, canvasHeight := 1920, 1080
	// Ensure plugins directory exists
	os.MkdirAll("plugins", 0755)
	engine := engine.NewEngine(canvasWidth, canvasHeight, VERSION)

	ebiten.SetWindowSize(canvasWidth, canvasHeight)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Type Defense")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetWindowDecorated(false)

	if err := ebiten.RunGame(engine); err != nil {
		log.Fatal(err)
	}
}

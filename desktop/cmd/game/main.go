package main

import (
	"log"

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
	engine := engine.NewEngine()

	ebiten.SetWindowSize(1920/2,1080/2) 
	ebiten.SetWindowTitle("Type Defense")
	// ebiten.SetFullscreen(true)
	// ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetWindowDecorated(false)

	if err := ebiten.RunGame(engine); err != nil {
		log.Fatal(err)
	}
}

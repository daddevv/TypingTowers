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
	// Ensure plugins directory exists
	os.MkdirAll("plugins", 0755)
	engine := engine.NewEngine(VERSION)

	ebiten.SetWindowSize(1920/2,1080/2) // Set the window size to half of 1920x1080
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Type Defense")
	// ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	// ebiten.SetWindowDecorated(false)

	if err := ebiten.RunGame(engine); err != nil {
		log.Fatal(err)
	}
}

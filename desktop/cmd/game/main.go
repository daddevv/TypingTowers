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
	screenWidth, screenHeight := ebiten.Monitor().Size()
	engine := engine.NewGame(screenWidth, screenHeight, VERSION)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Type Defense")

	if err := ebiten.RunGame(engine); err != nil {
		log.Fatal(err)
	}
}

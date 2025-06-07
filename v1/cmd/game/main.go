package main

import (
	"log"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	ebiten.SetTPS(100)
}

func main() {
	assets.InitImages()
	g := game.NewDebugGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

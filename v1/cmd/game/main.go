package main

import (
	"log"

	"github.com/daddevv/type-defense/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	ebiten.SetTPS(100)
}

func main() {
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

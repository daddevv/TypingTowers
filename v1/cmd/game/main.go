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
	game.InitImages()
	cfg, err := game.LoadConfig(game.ConfigFile)
	if err != nil {
		log.Println("using default config:", err)
	}
	g := game.NewGameWithConfig(cfg)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

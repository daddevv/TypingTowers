package main

import (
	"log"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/config"
	"github.com/daddevv/type-defense/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	ebiten.SetTPS(100)
}

func main() {
	assets.InitImages()
	cfg, err := config.LoadConfig(config.ConfigFile)
	if err != nil {
		log.Println("using default config:", err)
	}
	g := game.NewGameWithConfig(cfg)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

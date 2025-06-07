package main

import (
	"log"
	"typingtowers/internal/content"
	"typingtowers/internal/core"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	ebiten.SetTPS(100)
}

// main initializes the game and starts the ebiten game loop.
func main() {
	content.InitializeContent()
	// cfg, err := config.LoadConfig(config.ConfigFile)
	// if err != nil {
	// 	log.Println("using default config:", err)
	// }
	g := core.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

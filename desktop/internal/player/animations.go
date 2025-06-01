package player

import "td/internal/ui"

var (
	IdleAnimation = loadPlayerIdle()
)

func loadPlayerIdle() *ui.Animation {
	animation, err := ui.NewAnimation("assets/sprites/player/spr_idle_strip9.png", 1, 9, 96, 64, 10, 4.0)
	if err != nil {
		panic("Failed to load player idle animation: " + err.Error())
	}
	return animation
}

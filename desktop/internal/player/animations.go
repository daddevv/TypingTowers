package player

import "td/internal/ui"

var (
	IdleAnimation = loadPlayerIdle()
	WalkingAnimation = loadPlayerWalking()
	// Add more animations as needed
)

func loadPlayerIdle() *ui.Animation {
	animation, err := ui.NewAnimation("assets/sprites/player/spr_idle_strip9.png", 1, 9, 96, 64, 10, 2.0)
	if err != nil {
		panic("Failed to load player idle animation: " + err.Error())
	}
	return animation
}

func loadPlayerWalking() *ui.Animation {
	animation, err := ui.NewAnimation("assets/sprites/player/spr_walking_strip8.png", 1, 8, 96, 64, 10, 2.0)
	if err != nil {
		panic("Failed to load player walking animation: " + err.Error())
	}
	return animation
}

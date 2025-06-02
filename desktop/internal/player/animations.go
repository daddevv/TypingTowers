package player

import (
	"td/internal/sprite"
)

var (
	IdleAnimation = loadPlayerIdle()
	WalkingAnimation = loadPlayerWalking()
	// Add more animations as needed
)

func loadPlayerIdle() *sprite.AnimatedSprite {
	// animation, err := ui.NewAnimation("assets/sprites/player/spr_idle_strip9.png", 1, 9, 96, 64, 10, 2.0)
	// if err != nil {
	// 	panic("Failed to load player idle animation: " + err.Error())
	// }
	frames := sprite.SliceAnimationFrames("assets/sprites/player/spr_idle_strip9.png", 1, 9, 96, 64, 10, 2.0)
	if len(frames) == 0 {
		panic("Failed to load player idle animation frames")
	}
	return sprite.NewAnimatedSprite(frames, 10)
}

func loadPlayerWalking() *sprite.AnimatedSprite {
	// animation, err := ui.NewAnimation("assets/sprites/player/spr_walking_strip8.png", 1, 8, 96, 64, 10, 2.0)
	// if err != nil {
	// 	panic("Failed to load player walking animation: " + err.Error())
	// }
	frames := sprite.SliceAnimationFrames("assets/sprites/player/spr_walking_strip8.png", 1, 8, 96, 64, 10, 2.0)
	if len(frames) == 0 {
		panic("Failed to load player walking animation frames")
	}
	return sprite.NewAnimatedSprite(frames, 10)
}

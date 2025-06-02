package sprite

import (
	"image"
	"td/internal/math"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimatedSprite struct {
	Sprite
	Frames       []*ebiten.Image // List of frames for the animation
	CurrentFrame int             // Index of the current frame
	Tick         int             // Current ms count for the animation
	FrameDelay   int             // Delay between frames in ms
}

func NewAnimatedSprite(frames []*ebiten.Image, frameDelay int) *AnimatedSprite {
	return &AnimatedSprite{
		Frames:       frames,
		CurrentFrame: 0,
		Tick:         0,
		FrameDelay:   frameDelay,
	}
}

func (a *AnimatedSprite) Frame() *ebiten.Image {
	if len(a.Frames) == 0 {
		return nil // Return nil if there are no frames
	}
	return a.Frames[a.CurrentFrame]
}

func (a *AnimatedSprite) GetFrame() *ebiten.Image {
	if len(a.Frames) == 0 {
		return nil // Return nil if there are no frames
	}
	return a.Frames[a.CurrentFrame]
}

func (a *AnimatedSprite) GetFrameRect() image.Rectangle {
	if len(a.Frames) == 0 {
		return image.Rectangle{} // Return empty rectangle if there are no frames
	}
	return a.Frames[a.CurrentFrame].Bounds()
}

func (a *AnimatedSprite) GetPosition() math.Vec2 {
	return a.Sprite.GetPosition()
}

func (a *AnimatedSprite) SetPosition(pos math.Vec2) {
	a.Sprite.SetPosition(pos)
}

func (a *AnimatedSprite) Update(deltaTime float64) {
	a.Tick += int(deltaTime * 1000) // Convert deltaTime to milliseconds
	if a.Tick >= a.FrameDelay {
		framesToMove := a.Tick / a.FrameDelay                            // Calculate how many frames to move based on the tick count
		a.Tick = a.Tick % a.FrameDelay                                   // Reset tick to the remainder after moving frames
		a.CurrentFrame = (a.CurrentFrame + framesToMove) % len(a.Frames) // Loop through frames
		a.Tick = 0                                                       // Reset tick after updating frame
	}
}

func (a *AnimatedSprite) Draw(canvas *ebiten.Image, opts *ebiten.DrawImageOptions) {
	if len(a.Frames) == 0 {
		return // Do not draw if there are no frames
	}
	canvas.DrawImage(a.Frames[a.CurrentFrame], opts) // Draw the current frame with the provided options
}

func (a *AnimatedSprite) SetAnimation(other *AnimatedSprite) {
	a.Frames = other.Frames
	a.CurrentFrame = other.CurrentFrame
	a.Tick = other.Tick
	a.FrameDelay = other.FrameDelay
}

func (a *AnimatedSprite) Reset() {
	a.CurrentFrame = 0
	a.Tick = 0
}

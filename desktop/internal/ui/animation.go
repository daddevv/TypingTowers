package ui

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Animation struct {
	SpriteSheet  *ebiten.Image
	Frames       []*ebiten.Image
	Rows         int
	Cols         int
	FrameWidth   int
	FrameHeight  int
	CurrentFrame int
	Tick         int
	FrameDelay   int
	ImagePath    string
}

func NewAnimation(imagePath string, rows, cols, srcFrameWidth, srcFrameHeight, delay int, scale float64) (*Animation, error) {
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		return nil, err
	}

	frames := make([]*ebiten.Image, 0, rows*cols)
	for r := range rows {
		for c := range cols {
			x0 := c * srcFrameWidth
			y0 := r * srcFrameHeight
			rect := image.Rect(x0, y0, x0+srcFrameWidth, y0+srcFrameHeight)
			frame := img.SubImage(rect).(*ebiten.Image)
			// Scale the frame if a scale factor is provided
			if scale != 1.0 {
				scaledFrame := ebiten.NewImage(int(float64(srcFrameWidth)*scale), int(float64(srcFrameHeight)*scale))
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Scale(scale, scale) // Scale the image
				scaledFrame.DrawImage(frame, opts)
				frame = scaledFrame
			}
			// Add the frame to the list of frames
			frames = append(frames, frame)
		}
	}

	return &Animation{
		SpriteSheet:  img,
		Frames:       frames,
		Rows:         rows,
		Cols:         cols,
		FrameWidth:   srcFrameWidth,
		FrameHeight:  srcFrameHeight,
		CurrentFrame: 0,
		Tick:         0,
		FrameDelay:   delay,
		ImagePath:    imagePath,
	}, nil
}

func (a *Animation) Update() {
	a.Tick++
	if a.Tick >= a.FrameDelay {
		a.CurrentFrame = (a.CurrentFrame + 1) % len(a.Frames)
		a.Tick = 0
	}
}

func (a *Animation) Frame() *ebiten.Image {
	return a.Frames[a.CurrentFrame]
}

// Reset resets the animation to the first frame and tick count.
func (a *Animation) Reset() {
	a.CurrentFrame = 0
	a.Tick = 0
}

// SetFrame sets the current frame of the animation to a specific index.
func (a *Animation) SetFrame(frameIndex int) {
	if frameIndex < 0 || frameIndex >= len(a.Frames) {
		return // Ignore invalid frame index
	}
	a.CurrentFrame = frameIndex
	a.Tick = 0 // Reset tick count when setting a specific frame
}

// SetAnimation updates the animation's sprite sheet and frame properties.
func (a *Animation) SetAnimation(other *Animation) {
	a.SpriteSheet = other.SpriteSheet
	a.Frames = other.Frames
	a.Rows = other.Rows
	a.Cols = other.Cols
	a.FrameWidth = other.FrameWidth
	a.FrameHeight = other.FrameHeight
	a.CurrentFrame = 0
	a.Tick = 0
}
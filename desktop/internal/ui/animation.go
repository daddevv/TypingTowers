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

func NewAnimation(imagePath string, rows, cols, frameWidth, frameHeight, delay int) (*Animation, error) {
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		return nil, err
	}

	frames := make([]*ebiten.Image, 0, rows*cols)
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			x0 := c * frameWidth
			y0 := r * frameHeight
			rect := image.Rect(x0, y0, x0+frameWidth, y0+frameHeight)
			frame := img.SubImage(rect).(*ebiten.Image)
			frames = append(frames, frame)
		}
	}

	return &Animation{
		SpriteSheet:  img,
		Frames:       frames,
		Rows:         rows,
		Cols:         cols,
		FrameWidth:   frameWidth,
		FrameHeight:  frameHeight,
		CurrentFrame: 0,
		Tick:         0,
		FrameDelay:   delay,
		ImagePath:    imagePath,
	}, nil
}

func (a *Animation) Update() *ebiten.Image {
	a.Tick++
	if a.Tick >= a.FrameDelay {
		a.CurrentFrame = (a.CurrentFrame + 1) % len(a.Frames)
		a.Tick = 0
	}
	return a.Frames[a.CurrentFrame]
}

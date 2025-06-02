package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func SliceAnimationFrames(
	imagePath string,
	sheetRows, sheetCols int,
	frameWidth, frameHeight int,
	frameDuration float64,
	scale float64,
) []*ebiten.Image {
	frames := make([]*ebiten.Image, 0, sheetRows*sheetCols)
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}
	for r := range sheetRows {
		for c := range sheetCols {
			x0 := c * frameWidth
			y0 := r * frameHeight
			rect := image.Rect(x0, y0, x0+frameWidth, y0+frameHeight)
			frame := img.SubImage(rect).(*ebiten.Image)

			if scale != 1.0 {
				scaledFrame := ebiten.NewImage(int(float64(frameWidth)*scale), int(float64(frameHeight)*scale))
				opts := &ebiten.DrawImageOptions{}
				opts.GeoM.Scale(scale, scale) // Scale the image
				scaledFrame.DrawImage(frame, opts)
				frame = scaledFrame
			}

			frames = append(frames, frame)
		}
	}
	return frames
}

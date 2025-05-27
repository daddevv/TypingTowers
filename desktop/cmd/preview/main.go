package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Previewer struct {
	SpriteSheet   *ebiten.Image
	Frames        []*ebiten.Image
	Rows, Cols    int
	FrameWidth    int
	FrameHeight   int
	CurrentFrame  int
	Tick          int
	FrameDelay    int // ticks per frame
	AntiSpamTick  int // Anti-spam tick for frame updates
	AntiSpamDelay int // Default frame delay
	ImagePath     string
}

func NewPreviewer(imagePath string, rows, cols, height, width int) *Previewer {
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		log.Fatalf("failed to load image: %v", err)
	}
	frames := []*ebiten.Image{}
	for r := range rows {
		for c := range cols {
			x0 := c * width
			y0 := r * height
			rect := image.Rect(x0, y0, x0+width, y0+height)
			frame := img.SubImage(rect).(*ebiten.Image)
			frames = append(frames, frame)
		}
	}
	return &Previewer{
		SpriteSheet:   img,
		Frames:        frames,
		Rows:          rows,
		Cols:          cols,
		FrameWidth:    width,
		FrameHeight:   height,
		CurrentFrame:  0,
		Tick:          0,
		FrameDelay:    10,
		AntiSpamTick:  0,  // Anti-spam tick for frame updates
		AntiSpamDelay: 10, // Default frame delay
		ImagePath:     imagePath,
	}
}

func (p *Previewer) Update() error {
	if p.AntiSpamTick > 0 {
		p.AntiSpamTick--
	}

	// Controls: Up/Down to increase/decrease frame rate
	if ebiten.IsKeyPressed(ebiten.KeyUp) && p.AntiSpamTick <= 0 {
		p.AntiSpamTick = p.AntiSpamDelay
		if p.FrameDelay > 1 {
			p.FrameDelay--
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && p.AntiSpamTick <= 0 {
		p.AntiSpamTick = p.AntiSpamDelay
		p.FrameDelay++
	}

	p.Tick++
	if p.Tick >= p.FrameDelay {
		p.CurrentFrame = (p.CurrentFrame + 1) % len(p.Frames)
		p.Tick = 0
	}

	return nil
}

func (p *Previewer) Draw(screen *ebiten.Image) {
	screen.Fill(image.White)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(100, 100)
	screen.DrawImage(p.Frames[p.CurrentFrame], opts)

	// Draw config values
	info := fmt.Sprintf(
		"Image: %s\nRows: %d  Cols: %d\nFrame Size: %dx%d\nFrameDelay: %d (Up/Down to change)\nCurrentFrame: %d",
		p.ImagePath, p.Rows, p.Cols, p.FrameWidth, p.FrameHeight, p.FrameDelay, p.CurrentFrame,
	)
	ebitenutil.DebugPrintAt(screen, info, 10, 10)
}

func (p *Previewer) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	if len(os.Args) < 6 {
		fmt.Println("Usage: go run ./cmd/preview/main.go <image_path> <rows> <cols> <height> <width>")
		os.Exit(1)
	}
	imagePath := os.Args[1]
	rows, _ := strconv.Atoi(os.Args[2])
	cols, _ := strconv.Atoi(os.Args[3])
	height, _ := strconv.Atoi(os.Args[4])
	width, _ := strconv.Atoi(os.Args[5])

	previewer := NewPreviewer(imagePath, rows, cols, height, width)
	ebiten.SetWindowTitle("Sprite Previewer")
	ebiten.SetWindowSize(800, 600)
	if err := ebiten.RunGame(previewer); err != nil {
		log.Fatal(err)
	}
}

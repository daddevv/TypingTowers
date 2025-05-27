package entity

import (
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	PLAYER_Y = 0.8
)

type Player struct {
	Pos   ui.Location
	Image *ebiten.Image
}

func NewPlayer() *Player {
	image, _, err := ebitenutil.NewImageFromFile("assets/images/player/idle.png")
	if err != nil {
		panic(err)
	}
	return &Player{
		Pos:   ui.Location{X: 0.1, Y: PLAYER_Y},
		Image: image,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Draw the player on the screen
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4*float64(screen.Bounds().Dx())/1920, 4*float64(screen.Bounds().Dy())/1080) // Scale the player image
	opts.GeoM.Translate(
		p.Pos.X * float64(screen.Bounds().Dx()),
		p.Pos.Y * float64(screen.Bounds().Dy()),
	) // Position based on screen size
	screen.DrawImage(p.Image, opts)
}

func (p *Player) Update() error {
	// Satisfy the Entity interface, but no specific update logic for Player
	return nil
}

func (p *Player) GetPosition() ui.Location {
	return p.Pos
}

func (p *Player) SetPosition(x, y float64) {
	p.Pos.X = x
	p.Pos.Y = y
}

func (p *Player) StartDeath() {
	// TODO: Implement player death logic (animation, state, etc)
}

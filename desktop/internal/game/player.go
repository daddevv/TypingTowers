package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	Pos   Location
	Image *ebiten.Image
}

func NewPlayer() *Player {
	image, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		panic(err)
	}
	return &Player{
		Pos:   Location{X: 100, Y: 860},
		Image: image,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Draw the player on the screen
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4) // Scale the player image
	opts.GeoM.Translate(p.Pos.X, p.Pos.Y)
	screen.DrawImage(p.Image, opts)
}

func (p *Player) Update() error {
	// Satisfy the Entity interface, but no specific update logic for Player
	return nil
}

func (p *Player) GetPosition() Location {
	return p.Pos
}

func (p *Player) SetPosition(x, y float64) {
	p.Pos.X = x
	p.Pos.Y = y
}

func (p *Player) StartDeath() {
	// TODO: Implement player death logic (animation, state, etc)
}

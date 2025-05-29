package entity

import (
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	PLAYER_Y = 850.0 // px
)

type Player struct {
	Pos    ui.Location
	Image  *ebiten.Image
	Health int // Player health
}

func NewPlayer() *Player {
	image, _, err := ebitenutil.NewImageFromFile("assets/images/player/idle.png")
	if err != nil {
		panic(err)
	}
	return &Player{
		Pos:    ui.Location{X: 100, Y: PLAYER_Y}, // px
		Image:  image,
		Health: 5, // Default health now 5
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(4, 4)
	opts.GeoM.Translate(p.Pos.X, p.Pos.Y)
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

func (p *Player) DecrementHealth() {
	if p.Health > 0 {
		p.Health--
	}
}

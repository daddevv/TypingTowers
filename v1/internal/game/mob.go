package game

// Mob represents a basic enemy moving left.
type Mob struct {
	BaseEntity
	speed      float64
	animTicker int
	alive      bool
}

// NewMob returns a new mob at the given position.
func NewMob(x, y float64) *Mob {
	w, h := ImgMobA.Bounds().Dx(), ImgMobA.Bounds().Dy()
	return &Mob{
		BaseEntity: BaseEntity{
			pos:          Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgMobA,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
		},
		speed: 1,
		alive: true,
	}
}

// Update moves the mob and handles animation.
func (m *Mob) Update() {
	m.pos.X -= m.speed
	m.animTicker++
	if m.animTicker%30 < 15 {
		m.frame = ImgMobA
	} else {
		m.frame = ImgMobB
	}
	if m.pos.X < -float64(m.width) {
		m.alive = false
	}
}

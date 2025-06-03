package game

// Base represents the player's base that mobs try to destroy.
const BaseStartingHealth = 10

// Base represents the player's base that mobs try to destroy.
type Base struct {
	BaseEntity
	health int
}

// NewBase creates a new base at the given position.
func NewBase(x, y float64) *Base {
	w, h := ImgBase.Bounds().Dx(), ImgBase.Bounds().Dy()
	return &Base{
		BaseEntity: BaseEntity{
			pos:          Point{x, y},
			width:        w,
			height:       h,
			frame:        ImgBase,
			frameAnchorX: float64(w) / 2,
			frameAnchorY: float64(h) / 2,
			static:       true,
		},
		health: BaseStartingHealth,
	}
}

// Update updates the base state. Currently does nothing.
func (b *Base) Update() {}

// Damage reduces the base's health by the given amount.
func (b *Base) Damage(amount int) {
	b.health -= amount
}

// Health returns the current health of the base.
func (b *Base) Health() int {
	return b.health
}

// Alive reports whether the base still has health remaining.
func (b *Base) Alive() bool {
	return b.health > 0
}

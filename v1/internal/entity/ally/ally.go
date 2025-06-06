package ally

import "github.com/daddevv/type-defense/internal/entity"

// MobType defines ally categories.
type MobType int

const (
	MobBasic MobType = iota
	MobArmored
	MobShielded
	MobFast
	MobWord
	MobBoss
)

// Ally describes common ally behavior.
type Ally interface {
	entity.Entity
	Velocity() (float64, float64)
	Alive() bool
	Damage(amount int)
	Type() MobType
}

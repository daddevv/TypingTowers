package game

// MobType defines enemy categories.
type MobType int

const (
	MobBasic MobType = iota
	MobArmored
	MobShielded
	MobFast
	MobWord
	MobBoss
)

// Enemy describes common enemy behavior.
type Enemy interface {
	Entity
	Velocity() (float64, float64)
	Alive() bool
	Damage(amount int)
	Type() MobType
}

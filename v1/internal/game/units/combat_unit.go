package game

// CombatUnit is a base struct for mobile combatants.
type CombatUnit struct {
	BaseEntity
	hp     int
	damage int
	speed  float64
	alive  bool
}

// Alive reports whether the unit is still active.
func (u *CombatUnit) Alive() bool { return u.alive }

// Damage reduces the unit's HP.
func (u *CombatUnit) Damage(amount int) {
	if !u.alive {
		return
	}
	u.hp -= amount
	if u.hp <= 0 {
		u.alive = false
	}
}

// Health returns the current HP.
func (u *CombatUnit) Health() int { return u.hp }

// AttackDamage returns the damage dealt in melee.
func (u *CombatUnit) AttackDamage() int { return u.damage }

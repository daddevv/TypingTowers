package game

import (
	"github.com/daddevv/type-defense/internal/entity"
	"testing"
)

// TestOrcGruntDefaults verifies default stats.
func TestOrcGruntDefaults(t *testing.T) {
	o := entity.NewOrcGrunt(0, 0)
	if o.Health() != 5 {
		t.Errorf("expected default HP 5 got %d", o.Health())
	}
	if o.AttackDamage() != 1 {
		t.Errorf("expected default damage 1 got %d", o.AttackDamage())
	}
}

// TestCombatFootmanKillsGrunt ensures a footman defeats a grunt in melee.
func TestCombatFootmanKillsGrunt(t *testing.T) {
	f := entity.NewFootman(0, 0)
	f.speed = 0
	o := entity.NewOrcGrunt(0, 0)
	o.speed = 0

	m := NewMilitary()
	m.AddUnit(f)

	for i := 0; i < 6; i++ {
		o.Update(0.1)
		m.Update(0.1, []*entity.OrcGrunt{o})
	}

	if o.Alive() {
		t.Errorf("expected orc grunt to be defeated")
	}
	if !f.Alive() {
		t.Errorf("footman should survive with remaining HP")
	}
}

// TestCombatFootmanDies verifies footman removal when killed.
func TestCombatFootmanDies(t *testing.T) {
	f := entity.NewFootman(0, 0)
	f.speed = 0
	f.hp = 2
	o := entity.NewOrcGrunt(0, 0)
	o.speed = 0
	o.damage = 5

	m := NewMilitary()
	m.AddUnit(f)

	m.Update(0.1, []*entity.OrcGrunt{o})

	if m.Count() != 0 {
		t.Errorf("expected footman removed after death")
	}
}

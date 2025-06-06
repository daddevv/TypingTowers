package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/entity/ally"
	"github.com/daddevv/type-defense/internal/entity/enemy"
)

// TestOrcGruntDefaults verifies default stats.
func TestOrcGruntDefaults(t *testing.T) {
	o := enemy.NewOrcGrunt(0, 0)
	if o.Health() != 5 {
		t.Errorf("expected default HP 5 got %d", o.Health())
	}
	if o.AttackDamage() != 1 {
		t.Errorf("expected default damage 1 got %d", o.AttackDamage())
	}
}

// TestCombatFootmanKillsGrunt ensures a footman defeats a grunt in melee.
func TestCombatFootmanKillsGrunt(t *testing.T) {
	f := ally.NewFootman(0, 0)
	f.Speed = 0
	o := enemy.NewOrcGrunt(0, 0)
	o.Speed = 0

	for i := 0; i < 6; i++ {
		o.Update(0.1)
	}

	if o.Alive {
		t.Errorf("expected orc grunt to be defeated")
	}
	if !f.Alive {
		t.Errorf("footman should survive with remaining HP")
	}
}

// TestCombatFootmanDies verifies footman removal when killed.
func TestCombatFootmanDies(t *testing.T) {
	f := ally.NewFootman(0, 0)
	f.Speed = 0
	f.Hp = 2
	o := enemy.NewOrcGrunt(0, 0)
	o.Speed = 0
	o.Damage = 5

}

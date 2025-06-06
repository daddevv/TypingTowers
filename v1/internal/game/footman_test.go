package game

import (
	"testing"

	"github.com/daddevv/type-defense/internal/entity/ally"
)

func TestFootmanMovement(t *testing.T) {
	f := ally.NewFootman(0, 0)
	f.Speed = 10
	f.Update(1.0)
	x := f.Position.X
	if x <= 0 {
		t.Errorf("expected footman to move right, got %f", x)
	}
}

func TestFootmanDefaults(t *testing.T) {
	f := ally.NewFootman(0, 0)
	if f.Health() != 10 {
		t.Errorf("expected default HP 10 got %d", f.Health())
	}
	if f.Damage != 1 {
		t.Errorf("expected default damage 1 got %d", f.Damage)
	}
	if f.Speed != 50 {
		t.Errorf("expected default speed 50 got %f", f.Speed)
	}
}

func TestFootmanDamageKills(t *testing.T) {
	f := ally.NewFootman(0, 0)
	f.ApplyDamage(5)
	if !f.Alive {
		t.Errorf("footman should still be alive")
	}
	f.ApplyDamage(5)
	if f.Alive {
		t.Errorf("footman should be dead after lethal damage")
	}
}

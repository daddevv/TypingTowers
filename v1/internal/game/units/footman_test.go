package game

import "testing"

func TestFootmanMovement(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 10
	f.Update(1.0)
	x, _ := f.Position()
	if x <= 0 {
		t.Errorf("expected footman to move right, got %f", x)
	}
}

func TestFootmanDefaults(t *testing.T) {
	f := NewFootman(0, 0)
	if f.Health() != 10 {
		t.Errorf("expected default HP 10 got %d", f.Health())
	}
	if f.damage != 1 {
		t.Errorf("expected default damage 1 got %d", f.damage)
	}
	if f.speed != 50 {
		t.Errorf("expected default speed 50 got %f", f.speed)
	}
}

func TestFootmanDamageKills(t *testing.T) {
	f := NewFootman(0, 0)
	f.Damage(5)
	if !f.Alive() {
		t.Errorf("footman should still be alive")
	}
	f.Damage(5)
	if f.Alive() {
		t.Errorf("footman should be dead after lethal damage")
	}
}

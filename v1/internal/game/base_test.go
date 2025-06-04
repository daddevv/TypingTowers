package game

import "testing"

func TestBaseDamageAndAlive(t *testing.T) {
	b := NewBase(0, 0, 2)
	if b.Health() != 2 {
		t.Fatalf("expected initial health 2 got %d", b.Health())
	}
	b.Damage(1)
	if b.Health() != 1 {
		t.Errorf("expected health 1 got %d", b.Health())
	}
	if !b.Alive() {
		t.Errorf("base should be alive")
	}
	b.Damage(1)
	if b.Alive() {
		t.Errorf("base should be dead")
	}
}

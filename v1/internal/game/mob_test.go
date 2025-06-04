package game

import "testing"

func TestMobUpdateVelocity(t *testing.T) {
	b := NewBase(100, 0, 1)
	m := NewMob(0, 0, b, 1, 1)
	vx, vy := m.Velocity()
	if vx != 0 || vy != 0 {
		t.Errorf("expected zero velocity before update")
	}
	m.Update(0.016)
	vx, _ = m.Velocity()
	if vx <= 0 {
		t.Errorf("expected positive vx after update got %v", vx)
	}
}

func TestArmoredMobDamage(t *testing.T) {
	b := NewBase(0, 0, 1)
	m := NewArmoredMob(0, 0, b, 10, 2, 1)
	m.Damage(1)
	if m.health != 10 {
		t.Errorf("armor should absorb damage")
	}
	m.Damage(3)
	if m.health != 9 {
		t.Errorf("expected health 9 got %d", m.health)
	}
}

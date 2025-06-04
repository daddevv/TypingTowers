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

func TestFastMobBurst(t *testing.T) {
	b := NewBase(0, 0, 1)
	m := NewFastMob(0, 0, b, 10, 1, 3)
	// Initial burstCD should be 2, so no burst at first
	m.Update(0.016)
	if m.burstCD <= 0 {
		t.Errorf("expected burstCD > 0 at start")
	}
	// Simulate enough time to trigger burst
	for i := 0; i < 130; i++ { // ~2 seconds
		m.Update(0.016)
	}
	if m.burstCD > 0 {
		t.Errorf("expected burstCD to be 0 or less after enough time")
	}
	// During burst, speed should be increased
	oldVX, oldVY := m.Velocity()
	m.Update(0.016)
	newVX, newVY := m.Velocity()
	if (newVX == oldVX && newVY == oldVY) || (newVX == 0 && newVY == 0) {
		t.Errorf("expected velocity to change during burst")
	}
}

func TestBossMobProperties(t *testing.T) {
	b := NewBase(0, 0, 1)
	m := NewBossMob(0, 0, b, 100, 0.5)
	if m.mobType != MobBoss {
		t.Errorf("expected mob type MobBoss, got %v", m.mobType)
	}
	if m.health != 100 {
		t.Errorf("expected boss health 100, got %d", m.health)
	}
	if m.speed != 0.5 {
		t.Errorf("expected boss speed 0.5, got %f", m.speed)
	}
}

func TestMobTypeInstantiation(t *testing.T) {
	b := NewBase(0, 0, 1)
	armored := NewArmoredMob(0, 0, b, 10, 2, 1)
	if armored.mobType != MobArmored {
		t.Errorf("expected MobArmored type")
	}
	if armored.armor != 2 {
		t.Errorf("expected armor 2, got %d", armored.armor)
	}
	fast := NewFastMob(0, 0, b, 5, 1, 2)
	if fast.mobType != MobFast {
		t.Errorf("expected MobFast type")
	}
	if fast.burst != 2 {
		t.Errorf("expected burst 2, got %f", fast.burst)
	}
	boss := NewBossMob(0, 0, b, 50, 0.5)
	if boss.mobType != MobBoss {
		t.Errorf("expected MobBoss type")
	}
}

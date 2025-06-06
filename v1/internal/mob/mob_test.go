package mob

// func TestMobUpdateVelocity(t *testing.T) {
// 	b := NewBase(100, 0, 1)
// 	m := NewMob(0, 0, b, 1, 1)
// 	vx, vy := m.Velocity()
// 	if vx != 0 || vy != 0 {
// 		t.Errorf("expected zero velocity before update")
// 	}
// 	m.Update(0.016)
// 	vx, _ = m.Velocity()
// 	if vx <= 0 {
// 		t.Errorf("expected positive vx after update got %v", vx)
// 	}
// }

// func TestArmoredMobDamage(t *testing.T) {
// 	b := NewBase(0, 0, 1)
// 	m := NewArmoredMob(0, 0, b, 10, 2, 1)
// 	m.Damage(1)
// 	if m.Health != 10 {
// 		t.Errorf("armor should absorb damage")
// 	}
// 	m.Damage(3)
// 	if m.Health != 9 {
// 		t.Errorf("expected health 9 got %d", m.Health)
// 	}
// }

// func TestFastMobBurst(t *testing.T) {
// 	b := NewBase(100, 0, 1) // Place base away from the mob's position
// 	m := NewFastMob(0, 0, b, 10, 1, 3)
// 	// Initial burstTimer should be ready (cooldown = 4.0), but burstActive should not be ready
// 	m.Update(0.016)
// 	if m.BurstTimer.Ready() {
// 		t.Errorf("expected burstTimer not ready at start, remaining: %f", m.BurstTimer.Remaining)
// 	}

// 	oldVX, oldVY := m.Velocity()
// 	m.Update(4.0) // Wait for burst cooldown to expire
// 	if !m.BurstTimer.Ready() {
// 		t.Errorf("expected burstTimer to be ready after enough time, remaining: %f", m.BurstTimer.Remaining)
// 	}
// 	// Now burst should happen
// 	m.Update(0.016) // Trigger burst
// 	// Check if velocity changed
// 	newVX, newVY := m.Velocity()
// 	if (newVX == oldVX && newVY == oldVY) || (newVX == 0 && newVY == 0) {
// 		t.Errorf("expected velocity to change during burst, got old=(%f,%f), new=(%f,%f)",
// 			oldVX, oldVY, newVX, newVY)
// 	}
// }

// func TestBossMobProperties(t *testing.T) {
// 	b := NewBase(0, 0, 1)
// 	m := NewBossMob(0, 0, b, 100, 0.5)
// 	if m.MobType != MobBoss {
// 		t.Errorf("expected mob type MobBoss, got %v", m.MobType)
// 	}
// 	if m.Health != 100 {
// 		t.Errorf("expected boss health 100, got %d", m.Health)
// 	}
// 	if m.Speed != 0.5 {
// 		t.Errorf("expected boss speed 0.5, got %f", m.Speed)
// 	}
// }

// func TestMobTypeInstantiation(t *testing.T) {
// 	b := NewBase(0, 0, 1)
// 	armored := NewArmoredMob(0, 0, b, 10, 2, 1)
// 	if armored.MobType != MobArmored {
// 		t.Errorf("expected MobArmored type")
// 	}
// 	if armored.Armor != 2 {
// 		t.Errorf("expected armor 2, got %d", armored.Armor)
// 	}
// 	fast := NewFastMob(0, 0, b, 5, 1, 2)
// 	if fast.MobType != MobFast {
// 		t.Errorf("expected MobFast type")
// 	}
// 	if fast.Burst != 2 {
// 		t.Errorf("expected burst 2, got %f", fast.Burst)
// 	}
// 	boss := NewBossMob(0, 0, b, 50, 0.5)
// 	if boss.MobType != MobBoss {
// 		t.Errorf("expected MobBoss type")
// 	}
// }

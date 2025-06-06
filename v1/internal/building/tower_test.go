package building

// func TestTowerApplyConfig(t *testing.T) {
// 	cfg := DefaultConfig
// 	cfg.TowerDamage = 5
// 	cfg.TowerRange = 250
// 	g := &Game{cfg: &cfg}
// 	tower := NewTower(g, 0, 0)
// 	if tower.Damage != cfg.TowerDamage || tower.RangeDst != cfg.TowerRange {
// 		t.Fatalf("tower did not apply config")
// 	}
// 	newCfg := cfg
// 	newCfg.TowerAmmoCapacity = 10
// 	tower.ApplyConfig(newCfg)
// 	if tower.AmmoCapacity != 10 {
// 		t.Errorf("expected ammo capacity 10 got %d", tower.AmmoCapacity)
// 	}
// }

// func TestTowerReloadQueue(t *testing.T) {
// 	g := &Game{cfg: &DefaultConfig, input: NewInput(), typing: NewTypingStats()}
// 	tower := NewTower(g, 0, 0)

// 	// Consume all ammo to trigger reload queue
// 	for tower.getAvailableAmmo() > 0 {
// 		tower.consumeAmmo()
// 	}

// 	// Update should fill reload queue
// 	tower.Update(0.016)

// 	reloading, currentLetter, previewQueue, _, jammed := tower.GetReloadStatus()
// 	if !reloading {
// 		t.Errorf("tower should be reloading when ammo is empty")
// 	}
// 	if jammed {
// 		t.Errorf("tower should not be jammed initially")
// 	}
// 	if currentLetter != 'f' && currentLetter != 'j' {
// 		t.Errorf("unexpected reload letter %c", currentLetter)
// 	}
// 	if len(previewQueue) == 0 {
// 		t.Errorf("preview queue should contain letters to type")
// 	}
// }

// func TestTowerReloadSequence(t *testing.T) {
// 	g := &Game{cfg: &DefaultConfig, input: NewInput(), typing: NewTypingStats()}
// 	tower := NewTower(g, 0, 0)
// 	tower.SetReloadSequence([]rune{'a', 'b', 'c'})
// 	for i := 0; i < tower.ammoCapacity; i++ {
// 		tower.consumeAmmo()
// 	}
// 	tower.Update(0.016)
// 	_, letter, _, _, _ := tower.GetReloadStatus()
// 	if letter != 'a' {
// 		t.Errorf("expected first reload letter 'a' got %c", letter)
// 	}
// }

// func TestTowerAmmoQueue(t *testing.T) {
// 	g := &Game{cfg: &DefaultConfig, input: NewInput(), typing: NewTypingStats()}
// 	tower := NewTower(g, 0, 0)

// 	// Test initial ammo state
// 	ammo, capacity := tower.GetAmmoStatus()
// 	if ammo != capacity {
// 		t.Errorf("expected full ammo initially, got %d/%d", ammo, capacity)
// 	}

// 	// Test ammo consumption
// 	initialAmmo := tower.getAvailableAmmo()
// 	if !tower.consumeAmmo() {
// 		t.Errorf("should be able to consume ammo when available")
// 	}
// 	newAmmo := tower.getAvailableAmmo()
// 	if newAmmo != initialAmmo-1 {
// 		t.Errorf("ammo not consumed properly, expected %d got %d", initialAmmo-1, newAmmo)
// 	}
// }

// func TestTowerAmmoCapacityUpgrade(t *testing.T) {
// 	g := &Game{cfg: &DefaultConfig, input: NewInput(), typing: NewTypingStats()}
// 	tower := NewTower(g, 0, 0)

// 	initialCapacity := tower.ammoCapacity
// 	initialAmmo, _ := tower.GetAmmoStatus()

// 	// Upgrade capacity
// 	tower.UpgradeAmmoCapacity(2)

// 	newAmmo, newCapacity := tower.GetAmmoStatus()
// 	if newCapacity != initialCapacity+2 {
// 		t.Errorf("expected capacity to increase by 2, got %d -> %d", initialCapacity, newCapacity)
// 	}
// 	if newAmmo != initialAmmo+2 {
// 		t.Errorf("expected ammo to increase with capacity, got %d -> %d", initialAmmo, newAmmo)
// 	}
// }

// func TestTowerJamming(t *testing.T) {
// 	g := &Game{input: NewInput(), typing: NewTypingStats()}
// 	tower := NewTower(g, 0, 0)

// 	// Consume ammo to trigger reload
// 	for tower.getAvailableAmmo() > 0 {
// 		tower.consumeAmmo()
// 	}
// 	tower.Update(0.016) // Fill reload queue

// 	// Test jamming preserves letter
// 	reloading, currentLetter, _, _, _ := tower.GetReloadStatus()
// 	if !reloading {
// 		t.Errorf("tower should be reloading")
// 	}

// 	originalLetter := currentLetter
// 	tower.jammed = true
// 	tower.jammedLetter = originalLetter

// 	// Clear jam
// 	tower.jammed = false

// 	_, clearedLetter, _, _, jammed := tower.GetReloadStatus()
// 	if jammed {
// 		t.Errorf("tower should not be jammed after clearing")
// 	}
// 	if clearedLetter != originalLetter {
// 		t.Errorf("expected letter to be preserved after jam clear")
// 	}
// }

// func TestUpgradePurchasing(t *testing.T) {
// 	g := &Game{cfg: &DefaultConfig, input: NewInput(), typing: NewTypingStats()}
// 	g.AddGold(25)
// 	tower := NewTower(g, 0, 0)
// 	g.towers = []*Tower{tower}

// 	// Test damage upgrade
// 	oldDamage := tower.damage
// 	if g.Gold() < 5 {
// 		t.Fatal("not enough gold for test")
// 	}
// 	g.SpendGold(5)
// 	tower.damage++
// 	if tower.damage != oldDamage+1 {
// 		t.Errorf("damage upgrade failed, expected %d got %d", oldDamage+1, tower.damage)
// 	}
// 	if g.Gold() != 20 {
// 		t.Errorf("gold not deducted correctly after damage upgrade, expected 20 got %d", g.Gold())
// 	}

// 	// Test range upgrade
// 	oldRange := tower.rangeDst
// 	g.SpendGold(5)
// 	tower.rangeDst += 50
// 	if tower.rangeDst != oldRange+50 {
// 		t.Errorf("range upgrade failed, expected %f got %f", oldRange+50, tower.rangeDst)
// 	}
// 	if g.Gold() != 15 {
// 		t.Errorf("gold not deducted correctly after range upgrade, expected 15 got %d", g.Gold())
// 	}

// 	// Test fire rate upgrade
// 	oldRate := tower.rate
// 	g.SpendGold(5)
// 	if tower.rate > 10 {
// 		tower.rate -= 10
// 	}
// 	if tower.rate != oldRate-10 && oldRate > 10 {
// 		t.Errorf("fire rate upgrade failed, expected %f got %f", oldRate-10, tower.rate)
// 	}
// 	if g.Gold() != 10 {
// 		t.Errorf("gold not deducted correctly after fire rate upgrade, expected 10 got %d", g.Gold())
// 	}

// 	// Test ammo capacity upgrade
// 	oldCapacity := tower.ammoCapacity
// 	oldAmmo, _ := tower.GetAmmoStatus()
// 	g.SpendGold(10)
// 	tower.UpgradeAmmoCapacity(2)
// 	newAmmo, newCapacity := tower.GetAmmoStatus()
// 	if newCapacity != oldCapacity+2 {
// 		t.Errorf("ammo capacity upgrade failed, expected %d got %d", oldCapacity+2, newCapacity)
// 	}
// 	if newAmmo != oldAmmo+2 {
// 		t.Errorf("ammo should increase with capacity upgrade, expected %d got %d", oldAmmo+2, newAmmo)
// 	}
// 	if g.Gold() != 0 {
// 		t.Errorf("gold not deducted correctly after ammo capacity upgrade, expected 0 got %d", g.Gold())
// 	}
// }

// func TestSingleUpgradePurchase(t *testing.T) {
// 	g := &Game{cfg: &DefaultConfig, input: NewInput(), typing: NewTypingStats()}
// 	g.AddGold(100)
// 	tower := NewTower(g, 0, 0)
// 	g.towers = []*Tower{tower}

// 	// Test that upgrade only happens once per purchase action
// 	oldDamage := tower.damage
// 	oldGold := g.Gold()

// 	// Simulate single upgrade purchase
// 	if g.SpendGold(5) {
// 		tower.damage++
// 	}

// 	// Verify single upgrade occurred
// 	if tower.damage != oldDamage+1 {
// 		t.Errorf("expected single damage upgrade, got %d -> %d", oldDamage, tower.damage)
// 	}
// 	if g.Gold() != oldGold-5 {
// 		t.Errorf("expected gold to decrease by 5, got %d -> %d", oldGold, g.Gold())
// 	}

// 	// Test insufficient funds prevention
// 	g.resources.Gold.Set(3) // Not enough for 5 gold upgrade
// 	oldDamage = tower.damage
// 	oldGold = g.Gold()

// 	// Should not upgrade when insufficient funds
// 	if g.SpendGold(5) { // This condition should fail
// 		tower.damage++
// 	}

// 	// Verify no upgrade occurred
// 	if tower.damage != oldDamage {
// 		t.Errorf("upgrade should not occur with insufficient gold, damage changed from %d to %d", oldDamage, tower.damage)
// 	}
// 	if g.Gold() != oldGold {
// 		t.Errorf("gold should not change with insufficient funds, changed from %d to %d", oldGold, g.Gold())
// 	}
// }

// func TestNewTowerTypes(t *testing.T) {
// 	g := &Game{cfg: &DefaultConfig, input: NewInput(), typing: NewTypingStats()}
// 	sniper := NewTowerWithType(g, 0, 0, TowerSniper)
// 	rapid := NewTowerWithType(g, 0, 0, TowerRapid)
// 	if sniper.towerType != TowerSniper || rapid.towerType != TowerRapid {
// 		t.Fatalf("tower types not set correctly")
// 	}
// 	if sniper.rangeDst <= rapid.rangeDst {
// 		t.Errorf("sniper should have longer range")
// 	}
// 	if rapid.rate >= sniper.rate {
// 		t.Errorf("rapid tower should fire faster")
// 	}
// }

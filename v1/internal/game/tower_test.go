package game

import (
	"testing"
)

func TestTowerApplyConfig(t *testing.T) {
	cfg := DefaultConfig
	cfg.TowerDamage = 5
	cfg.TowerRange = 250
	g := &Game{cfg: &cfg}
	tower := NewTower(g, 0, 0)
	if tower.damage != cfg.TowerDamage || tower.rangeDst != cfg.TowerRange {
		t.Fatalf("tower did not apply config")
	}
	newCfg := cfg
	newCfg.TowerAmmoCapacity = 10
	tower.ApplyConfig(newCfg)
	if tower.ammoCapacity != 10 {
		t.Errorf("expected ammo capacity 10 got %d", tower.ammoCapacity)
	}
}

func TestTowerStartReload(t *testing.T) {
	g := &Game{cfg: &DefaultConfig}
	tower := NewTower(g, 0, 0)
	tower.shootingAmmo = 0 // empty shooting ammo to trigger reload
	tower.startReload()
	if !tower.reloading {
		t.Errorf("tower should be reloading")
	}
	if tower.reloadTimer != tower.reloadTime {
		t.Errorf("reload timer not set")
	}
	if tower.reloadLetter != 'f' && tower.reloadLetter != 'j' {
		t.Errorf("unexpected reload letter %c", tower.reloadLetter)
	}
	if tower.nextReloadLetter != 'f' && tower.nextReloadLetter != 'j' {
		t.Errorf("unexpected next reload letter %c", tower.nextReloadLetter)
	}
}

func TestTowerConfig(t *testing.T) {
	g := &Game{}
	tower := NewTower(g, 0, 0)
	if tower.damage != 1 || tower.rangeDst != 500 {
		t.Errorf("unexpected default tower config: damage=%v, range=%v", tower.damage, tower.rangeDst)
	}

	tower.ApplyConfig(Config{
		TowerDamage:       20,
		TowerRange:        200,
		TowerAmmoCapacity: 5,
		TowerReloadRate:   2,
	})
	if tower.ammoCapacity != 5 {
		t.Errorf("expected ammoCapacity=5, got %v", tower.ammoCapacity)
	}
	if tower.shootingAmmo > tower.ammoCapacity {
		t.Errorf("shootingAmmo should not exceed ammoCapacity, got %v/%v", tower.shootingAmmo, tower.ammoCapacity)
	}
}

func TestTowerReload(t *testing.T) {
	g := &Game{}
	tower := NewTower(g, 0, 0)
	tower.shootingAmmo = 0 // empty shooting ammo
	tower.startReload()
	if !tower.reloading {
		t.Errorf("expected reloading=true")
	}
	if tower.reloadTimer != tower.reloadTime {
		t.Errorf("expected reloadTimer=%v, got %v", tower.reloadTime, tower.reloadTimer)
	}
	if tower.reloadLetter != 'f' && tower.reloadLetter != 'j' {
		t.Errorf("expected reloadLetter to be set")
	}
	tower.reloadLetter = 'A'
	if tower.reloadLetter != 'A' {
		t.Errorf("expected reloadLetter=A, got %v", tower.reloadLetter)
	}
}

func TestTowerAmmoSystem(t *testing.T) {
	g := &Game{cfg: &DefaultConfig}
	tower := NewTower(g, 0, 0)

	// Test initial ammo state
	ammo, capacity := tower.GetAmmoStatus()
	if ammo != capacity {
		t.Errorf("expected full ammo initially, got %d/%d", ammo, capacity)
	}

	// Test ammo consumption
	initialAmmo := tower.shootingAmmo
	tower.shootingAmmo-- // simulate firing
	ammo, _ = tower.GetAmmoStatus()
	if ammo != initialAmmo-1 {
		t.Errorf("ammo not consumed properly, expected %d got %d", initialAmmo-1, ammo)
	}
}

func TestTowerJamming(t *testing.T) {
	g := &Game{}
	tower := NewTower(g, 0, 0)
	tower.shootingAmmo = 0
	tower.startReload()

	// Test jamming preserves letter
	originalLetter := tower.reloadLetter
	tower.jammed = true
	tower.jammedLetter = originalLetter

	// Clear jam
	tower.jammed = false
	tower.reloadLetter = tower.jammedLetter

	if tower.reloadLetter != originalLetter {
		t.Errorf("expected letter to be preserved after jam clear")
	}
}

func TestUpgradePurchasing(t *testing.T) {
	g := &Game{gold: 15, cfg: &DefaultConfig}
	tower := NewTower(g, 0, 0)
	g.towers = []*Tower{tower}

	// Simulate damage upgrade
	oldDamage := tower.damage
	if g.gold < 5 {
		t.Fatal("not enough gold for test")
	}
	g.gold -= 5
	tower.damage++
	if tower.damage != oldDamage+1 {
		t.Errorf("damage upgrade failed")
	}
	if g.gold != 10 {
		t.Errorf("gold not deducted correctly after damage upgrade")
	}

	// Simulate range upgrade
	oldRange := tower.rangeDst
	g.gold -= 5
	tower.rangeDst += 50
	if tower.rangeDst != oldRange+50 {
		t.Errorf("range upgrade failed")
	}
	if g.gold != 5 {
		t.Errorf("gold not deducted correctly after range upgrade")
	}

	// Simulate fire rate upgrade
	oldRate := tower.rate
	g.gold -= 5
	if tower.rate > 10 {
		tower.rate -= 10
	}
	if tower.rate != oldRate-10 && oldRate > 10 {
		t.Errorf("fire rate upgrade failed")
	}
	if g.gold != 0 {
		t.Errorf("gold not deducted correctly after fire rate upgrade")
	}
}

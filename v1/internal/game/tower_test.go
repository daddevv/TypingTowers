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
	tower.ammo = tower.ammo[:0]
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
	if len(tower.ammo) != tower.ammoCapacity {
		t.Errorf("expected ammo=%v, got %v", tower.ammoCapacity, len(tower.ammo))
	}
}

func TestTowerReload(t *testing.T) {
	g := &Game{}
	tower := NewTower(g, 0, 0)
	tower.ammo = tower.ammo[:0]
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

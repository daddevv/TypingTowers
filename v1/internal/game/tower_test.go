package game

import "testing"

func TestTowerApplyConfig(t *testing.T) {
	cfg := DefaultConfig
	cfg.TowerDamage = 5
	cfg.TowerRange = 250
	g := &Game{cfg: &cfg}
	t := NewTower(g, 0, 0)
	if t.damage != cfg.TowerDamage || t.rangeDst != cfg.TowerRange {
		t.Fatalf("tower did not apply config")
	}
	newCfg := cfg
	newCfg.TowerAmmoCapacity = 10
	t.ApplyConfig(newCfg)
	if t.ammoCapacity != 10 {
		t.Errorf("expected ammo capacity 10 got %d", t.ammoCapacity)
	}
}

func TestTowerStartReload(t *testing.T) {
	g := &Game{cfg: &DefaultConfig}
	t := NewTower(g, 0, 0)
	t.ammo = t.ammo[:0]
	t.startReload()
	if !t.reloading {
		t.Errorf("tower should be reloading")
	}
	if t.reloadTimer != t.reloadTime {
		t.Errorf("reload timer not set")
	}
	if t.reloadLetter != 'f' && t.reloadLetter != 'j' {
		t.Errorf("unexpected reload letter %c", t.reloadLetter)
	}
}

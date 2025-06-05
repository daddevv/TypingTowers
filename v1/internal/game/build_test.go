package game

import "testing"

func TestBuildTowerCostsGold(t *testing.T) {
	cfg := DefaultConfig
	cfg.TowerConstructionCost = 5
	g := NewGameWithConfig(cfg)
	g.AddGold(10)
	g.cursorX = 4
	g.cursorY = 4
	initial := len(g.towers)
	g.buildTowerAtCursorType(TowerSniper)
	if len(g.towers) != initial+1 {
		t.Fatalf("expected tower count %d got %d", initial+1, len(g.towers))
	}
	if g.Gold() != 5 {
		t.Fatalf("expected gold 5 got %d", g.Gold())
	}
	if g.towers[len(g.towers)-1].towerType != TowerSniper {
		t.Fatalf("expected sniper tower type")
	}
}

// If stubInput is used in this file, add:
func (s *stubInput) SelectTower() bool { return false }

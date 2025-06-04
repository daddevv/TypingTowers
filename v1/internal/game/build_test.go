package game

import "testing"

func TestBuildTowerCostsGold(t *testing.T) {
	cfg := DefaultConfig
	cfg.TowerConstructionCost = 5
	g := NewGameWithConfig(cfg)
	g.gold = 10
	g.cursorX = 4
	g.cursorY = 4
	initial := len(g.towers)
	g.buildTowerAtCursor()
	if len(g.towers) != initial+1 {
		t.Fatalf("expected tower count %d got %d", initial+1, len(g.towers))
	}
	if g.gold != 5 {
		t.Fatalf("expected gold 5 got %d", g.gold)
	}
}

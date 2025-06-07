package game

// import (
// 	"testing"

// 	"github.com/daddevv/type-defense/internal/assets"
// 	"github.com/daddevv/type-defense/internal/building"
// 	"github.com/daddevv/type-defense/internal/config"
// )

// // TestBuildTowerCostsGold verifies that constructing a tower deducts gold and
// // spawns the correct tower type at the cursor position.
// func TestBuildTowerCostsGold(t *testing.T) {
// 	assets.InitImages()
// 	cfg := config.DefaultConfig
// 	cfg.TowerConstructionCost = 5
// 	g := NewGameWithConfig(cfg)
// 	g.AddGold(10)
// 	g.cursorX = 4
// 	g.cursorY = 4
// 	initial := len(g.towers)
// 	g.buildTowerAtCursorType(building.TowerSniper)
// 	if len(g.towers) != initial+1 {
// 		t.Fatalf("expected tower count %d got %d", initial+1, len(g.towers))
// 	}
// 	if g.Gold() != 5 {
// 		t.Fatalf("expected gold 5 got %d", g.Gold())
// 	}
// 	if g.towers[len(g.towers)-1].TowerType != building.TowerSniper {
// 		t.Fatalf("expected sniper tower type")
// 	}
// }

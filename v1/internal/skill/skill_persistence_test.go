//go:build test

package skill

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSkillStatePersistence(t *testing.T) {
	g := NewGame()
	g.resources.AddKingsPoints(50)
	node := g.skillTree.Nodes["sharp_arrows"]
	if !g.skillTree.Unlock("sharp_arrows", &g.resources) {
		t.Fatalf("unlock failed")
	}
	g.unlockedSkills["sharp_arrows"] = true
	g.applySkillEffects(node)

	dir := t.TempDir()
	path := filepath.Join(dir, "save.json")
	g.saveGame(path)

	ng := NewGame()
	if err := ng.loadGame(path); err != nil {
		t.Fatal(err)
	}
	if !ng.unlockedSkills["sharp_arrows"] {
		t.Fatalf("skill not loaded")
	}
	if ng.towerMods.DamageMult != 1.1 {
		t.Fatalf("modifier not restored")
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("save file missing")
	}
}

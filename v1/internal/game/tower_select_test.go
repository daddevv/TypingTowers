//go:build test

package game

import (
	"strings"
	"testing"
)

func TestEnterTowerSelectMode(t *testing.T) {
	g := NewGame()
	g.phase = PhasePlaying
	g.towers = []*Tower{NewTower(g, 0, 0), NewTower(g, 10, 10), NewTower(g, 20, 20)}
	g.enterTowerSelectMode()
	if !g.towerSelectMode {
		t.Fatalf("tower selection mode not active")
	}
	if len(g.towerLabels) != len(g.towers) {
		t.Fatalf("expected %d labels got %d", len(g.towers), len(g.towerLabels))
	}
	if idx, ok := g.towerLabels["a"]; !ok || idx != 0 {
		t.Errorf("label a not set to tower 0")
	}
}

func (g *Game) processTowerSelectInput(chars []rune) {
	for _, r := range chars {
		label := strings.ToLower(string(r))
		if idx, ok := g.towerLabels[label]; ok {
			g.selectedTower = idx
			g.towerSelectMode = false
			g.upgradeMenuOpen = true
			g.upgradeCursor = 0
			break
		}
	}
}

func TestSelectTowerOpensUpgrade(t *testing.T) {
	g := NewGame()
	g.phase = PhasePlaying
	g.towers = []*Tower{NewTower(g, 0, 0), NewTower(g, 10, 10)}
	g.enterTowerSelectMode()
	g.processTowerSelectInput([]rune{'b'})
	if g.towerSelectMode {
		t.Errorf("tower selection mode should close after selection")
	}
	if !g.upgradeMenuOpen {
		t.Fatalf("upgrade menu should open after selection")
	}
	if g.selectedTower != 1 {
		t.Errorf("expected tower 1 selected got %d", g.selectedTower)
	}
}

//go:build test

package building

// func TestEnterTowerSelectMode(t *testing.T) {
// 	g := NewGame()
// 	g.phase = PhasePlaying
// 	g.towers = []*Tower{NewTower(g, 0, 0), NewTower(g, 10, 10), NewTower(g, 20, 20)}
// 	g.enterTowerSelectMode()
// 	if !g.towerSelectMode {
// 		t.Fatalf("tower selection mode not active")
// 	}
// 	if len(g.towerLabels) != len(g.towers) {
// 		t.Fatalf("expected %d labels got %d", len(g.towers), len(g.towerLabels))
// 	}
// 	if idx, ok := g.towerLabels["a"]; !ok || idx != 0 {
// 		t.Errorf("label a not set to tower 0")
// 	}
// }

// func (g *Game) processTowerSelectInput(chars []rune) {
// 	for _, r := range chars {
// 		label := strings.ToLower(string(r))
// 		if idx, ok := g.towerLabels[label]; ok {
// 			g.selectedTower = idx
// 			g.towerSelectMode = false
// 			g.upgradeMenuOpen = true
// 			g.upgradeCursor = 0
// 			break
// 		}
// 	}
// }

// func TestSelectTowerOpensUpgrade(t *testing.T) {
// 	g := NewGame()
// 	g.phase = PhasePlaying
// 	g.towers = []*Tower{NewTower(g, 0, 0), NewTower(g, 10, 10)}
// 	g.enterTowerSelectMode()
// 	g.processTowerSelectInput([]rune{'b'})
// 	if g.towerSelectMode {
// 		t.Errorf("tower selection mode should close after selection")
// 	}
// 	if !g.upgradeMenuOpen {
// 		t.Fatalf("upgrade menu should open after selection")
// 	}
// 	if g.selectedTower != 1 {
// 		t.Errorf("expected tower 1 selected got %d", g.selectedTower)
// 	}
// }

// type stubInputSelect struct{ selectTower bool }

// func (s *stubInputSelect) TypedChars() []rune { return nil }
// func (s *stubInputSelect) Update()            {}
// func (s *stubInputSelect) Reset()             { s.selectTower = false }
// func (s *stubInputSelect) Backspace() bool    { return false }
// func (s *stubInputSelect) Space() bool        { return false }
// func (s *stubInputSelect) Quit() bool         { return false }
// func (s *stubInputSelect) Reload() bool       { return false }
// func (s *stubInputSelect) Enter() bool        { return false }
// func (s *stubInputSelect) Left() bool         { return false }
// func (s *stubInputSelect) Right() bool        { return false }
// func (s *stubInputSelect) Up() bool           { return false }
// func (s *stubInputSelect) Down() bool         { return false }
// func (s *stubInputSelect) Build() bool        { return false }
// func (s *stubInputSelect) Save() bool         { return false }
// func (s *stubInputSelect) Load() bool         { return false }
// func (s *stubInputSelect) SelectTower() bool  { v := s.selectTower; s.selectTower = false; return v }
// func (s *stubInputSelect) Command() bool      { return false }
// func (s *stubInputSelect) TechMenu() bool     { return false }
// func (s *stubInputSelect) SkillMenu() bool    { return false }
// func (s *stubInputSelect) StatsPanel() bool   { return false }

// func TestSlashOpensTowerSelect(t *testing.T) {
// 	g := NewGame()
// 	g.phase = PhasePlaying
// 	g.towers = []*Tower{NewTower(g, 0, 0)}
// 	inp := &stubInputSelect{selectTower: true}
// 	g.input = inp
// 	g.lastUpdate = time.Now()
// 	if err := g.Update(); err != nil {
// 		t.Fatal(err)
// 	}
// 	if !g.towerSelectMode {
// 		t.Fatalf("expected tower selection mode to activate")
// 	}
// 	if len(g.towerLabels) == 0 {
// 		t.Fatalf("labels should be assigned when selection mode starts")
// 	}
// }

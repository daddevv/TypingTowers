package game

// func TestNewGame(t *testing.T) {
// 	g := NewGame()
// 	if g.screen.Bounds().Dx() != 1920 || g.screen.Bounds().Dy() != 1080 {
// 		t.Errorf("screen size expected 1920x1080 got %dx%d", g.screen.Bounds().Dx(), g.screen.Bounds().Dy())
// 	}
// 	if g.base == nil || g.base.Health() <= 0 {
// 		t.Errorf("base not initialized")
// 	}
// }

// func TestLetterUnlocking(t *testing.T) {
// 	g := NewGameWithConfig(config.DefaultConfig)
// 	tree := entity.DefaultTechTree()
// 	firstLetters, _, _ := tree.UnlockNext()
// 	// Manually assign the unlocked letters to the game's letter pool
// 	g.letterPool = append([]rune{}, firstLetters...)
// 	if len(g.letterPool) != len(firstLetters) {
// 		t.Fatalf("expected initial letter pool %d got %d", len(firstLetters), len(g.letterPool))
// 	}

// 	g.currentWave = 2
// 	g.startWave()
// 	tree = entity.DefaultTechTree()
// 	tree.UnlockNext() // first stage
// 	secondLetters, _, _ := tree.UnlockNext()
// 	// Manually add the next unlocked letters
// 	g.letterPool = append(g.letterPool, secondLetters...)
// 	expected := len(firstLetters) + len(secondLetters)
// 	if len(g.letterPool) != expected {
// 		t.Errorf("expected letter pool size %d after second wave got %d", expected, len(g.letterPool))
// 	}
// }

// func TestGameBackPressureDamage(t *testing.T) {
// 	g := NewGame()
// 	g.phase = core.PhasePlaying // Ensure main update logic runs
// 	// Fill the queue to the threshold for backpressure
// 	for i := 0; i < 6; i++ {
// 		g.Queue().Enqueue(word.Word{Text: "w"})
// 	}
// 	// Simulate enough time passing for damage to occur
// 	g.lastUpdate = time.Now().Add(-2 * time.Second)
// 	g.Update()
// 	expected := entity.BaseStartingHealth - 1
// 	if g.base.Health() != expected {
// 		t.Fatalf("expected base health %d got %d", expected, g.base.Health())
// 	}
// }

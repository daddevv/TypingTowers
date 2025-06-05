//go:build test

package game

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// TestSaveVersionMismatch ensures loading a file with the wrong version fails.
func TestSaveVersionMismatch(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "slot1.json")
	sg := savedGame{Version: SaveVersion + 1}
	b, _ := json.Marshal(sg)
	os.WriteFile(path, b, 0644)

	g := NewGame()
	g.saveDir = dir
	if err := g.loadGame(path); err == nil {
		t.Fatalf("expected version error")
	}
}

// TestAutoSaveCreatesFile verifies a save file is created after a wave.
func TestAutoSaveCreatesFile(t *testing.T) {
	dir := t.TempDir()
	g := NewGame()
	g.saveDir = dir
	inp := &stubInput{}
	g.input = inp

	// simulate until first wave completes
	for i := 0; i < 600; i++ {
		if w, ok := g.Queue().Peek(); ok {
			inp.typed = []rune(w.Text)
		}
		if err := g.Step(0.1); err != nil {
			t.Fatal(err)
		}
		if g.shopOpen {
			break
		}
	}
	path := filepath.Join(dir, "save_slot1.json")
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected save file: %v", err)
	}
}

// TestSetSaveSlotChangesPath ensures SetSaveSlot picks the correct slot file.
func TestSetSaveSlotChangesPath(t *testing.T) {
	g := NewGame()
	g.saveDir = "/tmp"
	g.SetSaveSlot(2)
	expect := filepath.Join("/tmp", "save_slot2.json")
	if g.currentSavePath() != expect {
		t.Fatalf("unexpected path %s", g.currentSavePath())
	}
}

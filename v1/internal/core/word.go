package core

import "github.com/daddevv/type-defense/internal/assets"

// Word represents a queued typing challenge produced by a building.
type Word struct {
	Text   string // text the player must type
	Source string // name of the building that generated the word
	Family string // building family for color coding
}

// Colorize returns the word text wrapped with the ANSI color for its family.
func (w *Word) Colorize() string {
	if c, ok := assets.FamilyPalette[w.Family]; ok {
		return c + w.Text + assets.ColorReset
	}
	return w.Text
}

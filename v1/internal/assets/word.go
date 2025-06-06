package assets

// Word represents a queued typing challenge produced by a building.
type Word struct {
	Text   string // text the player must type
	Source string // name of the building that generated the word
	Family string // building family for colour coding
}

// Colorize returns the word text wrapped with the ANSI colour for its family.
func (w *Word) Colorize() string {
	if c, ok := FamilyPalette[w.Family]; ok {
		return c + w.Text + ColorReset
	}
	return w.Text
}

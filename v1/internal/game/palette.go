package game

// ANSI colour codes for building families.
const (
	ColorReset = "\033[0m"
)

var FamilyPalette = map[string]string{
	"Gathering": "\033[32m", // green
	"Military":  "\033[31m", // red
}

// Colorize returns the word text wrapped with the ANSI colour for its family.
func Colorize(w Word) string {
	if c, ok := FamilyPalette[w.Family]; ok {
		return c + w.Text + ColorReset
	}
	return w.Text
}

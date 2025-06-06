package assets

import "image/color"

// ANSI colour codes for building families.
const (
	ColorReset = "\033[0m"
)

var FamilyPalette = map[string]string{
	"Gathering": "\033[32m", // green
	"Military":  "\033[31m", // red
}

// FamilyColors maps building families to on-screen colours used by the HUD.
var FamilyColors = map[string]color.RGBA{
	"Gathering": {0, 255, 0, 255}, // green
	"Military":  {255, 0, 0, 255}, // red
}

// FamilyColor returns the colour for the given building family.
func FamilyColor(f string) color.RGBA {
	if c, ok := FamilyColors[f]; ok {
		return c
	}
	return color.RGBA{255, 255, 255, 255}
}

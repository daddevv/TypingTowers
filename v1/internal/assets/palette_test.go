package assets

import "testing"

func TestFamilyPaletteValues(t *testing.T) {
	if FamilyPalette["Gathering"] != "\033[32m" {
		t.Errorf("Gathering color incorrect: %q", FamilyPalette["Gathering"])
	}
	if FamilyPalette["Military"] != "\033[31m" {
		t.Errorf("Military color incorrect: %q", FamilyPalette["Military"])
	}
}

func TestFamilyColor(t *testing.T) {
	c := FamilyColor("Gathering")
	if c.R != 0 || c.G != 255 || c.B != 0 {
		t.Fatalf("unexpected colour for Gathering: %+v", c)
	}
	unknown := FamilyColor("Unknown")
	if unknown.R != 255 || unknown.G != 255 || unknown.B != 255 {
		t.Fatalf("unexpected default colour: %+v", unknown)
	}
}

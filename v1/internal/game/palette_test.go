package game

import "testing"

func TestFamilyPaletteValues(t *testing.T) {
	if FamilyPalette["Gathering"] != "\033[32m" {
		t.Errorf("Gathering color incorrect: %q", FamilyPalette["Gathering"])
	}
	if FamilyPalette["Military"] != "\033[31m" {
		t.Errorf("Military color incorrect: %q", FamilyPalette["Military"])
	}
}

func TestColorize(t *testing.T) {
	w := Word{Text: "foo", Family: "Gathering"}
	got := Colorize(w)
	expected := "\033[32mfoo\033[0m"
	if got != expected {
		t.Fatalf("Colorize mismatch: got %q want %q", got, expected)
	}
}

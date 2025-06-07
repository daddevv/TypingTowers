package assets

import (
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	fontSource *text.GoTextFaceSource
	NormalFont = Font("Game-Regular", 16)
	BoldFont   = Font("Game-Bold", 16)
)

func Font(name string, size float64) *text.GoTextFace {
	fontSource, err := getFontSource(name)
	if err != nil {
		log.Fatal("failed to get font source:", err)
	}
	return &text.GoTextFace{Source: fontSource, Size: size}
}

func getFontSource(font string) (*text.GoTextFaceSource, error) {
	// Read the .ttf file into a byte slice.
	switch font {
	case "Game-Regular":
		loadFont("assets/fonts/OpenDyslexicNerdFont-Regular.otf")
	case "Game-Bold":
		loadFont("assets/fonts/OpenDyslexicNerdFont-Bold.otf")
	case "Mob":
		loadFont("assets/fonts/ShureTechMonoNerdFontMono-Regular.ttf")
	default:
		return nil, errors.New("font not found: " + font)
	}
	return fontSource, nil
}

func loadFont(relPath string) {
	// Get the directory of this source file (font.go)
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to determine caller path for font.go")
	}
	baseDir := filepath.Dir(filename)
	fontPath := filepath.Join(baseDir, "..", "..", relPath) // relative to v1/

	if _, err := os.Stat(fontPath); os.IsNotExist(err) {
		log.Fatalf("font file not found: %s", fontPath)
	}
	// Read the font file into a byte slice.
	data, err := os.ReadFile(fontPath)
	if err != nil {
		log.Fatal("unable to read font file:", err)
	}
	fontSource, err = text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		log.Fatal("failed to parse font:", err)
	}
}

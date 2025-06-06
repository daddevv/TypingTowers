package assets

import (
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"

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

func loadFont(path string) {
	//if testing add ../../ to path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// If the file does not exist, try to load it from a relative path.
		fp := filepath.Base(path)
		path = "/home/bobbitt/projects/public/TypingTowers/v1/assets/fonts/" + fp
	}
	// Read the .ttf file into a byte slice.
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("unable to read font file:", err)
	}
	fontSource, err = text.NewGoTextFaceSource(bytes.NewReader(data))
	if err != nil {
		log.Fatal("failed to parse font:", err)
	}
}

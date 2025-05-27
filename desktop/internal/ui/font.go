package ui

import (
	"bytes"
	"errors"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var fontSource *text.GoTextFaceSource

func Font(name string, size float64) *text.GoTextFace {
	fontSource, err := GetFontSource(name)
	if err != nil {
		log.Fatal("failed to get font source:", err)
	}
	return &text.GoTextFace{Source: fontSource, Size: size}
}

func GetFontSource(font string) (*text.GoTextFaceSource, error) {
	// Read the .ttf file into a byte slice.
	switch font {
	case "Game-Regular":
		loadFont("assets/fonts/OpenDyslexicNerdFont-Regular.otf")
		break
	case "Game-Bold":
		loadFont("assets/fonts/OpenDyslexicNerdFont-Bold.otf")
		break
	default:
		return nil, errors.New("font not found: " + font)
	}
	return fontSource, nil
}

func loadFont(path string) {
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

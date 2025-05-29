package entity

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type LetterState int

const (
	LetterInactive LetterState = iota
	LetterActive
	LetterTarget
)

var (
	letterImageCache = map[LetterState]map[rune]*ebiten.Image{}
	cacheOnce sync.Once
)

func getLetterColor(state LetterState) color.Color {
	switch state {
	case LetterTarget:
		return color.RGBA{255, 0, 0, 255}
	case LetterActive:
		return color.White
	case LetterInactive:
		return color.RGBA{128, 128, 128, 255}
	default:
		return color.White
	}
}

func GetLetterImage(r rune, state LetterState, font *text.GoTextFace) *ebiten.Image {
	cacheOnce.Do(func() {
		for _, s := range []LetterState{LetterInactive, LetterActive, LetterTarget} {
			letterImageCache[s] = make(map[rune]*ebiten.Image)
		}
	})
	if img, ok := letterImageCache[state][r]; ok {
		return img
	}
	img := ebiten.NewImage(32, 48)
	opts := &text.DrawOptions{}
	opts.ColorScale.ScaleWithColor(getLetterColor(state))
	text.Draw(img, string(r), font, opts)
	letterImageCache[state][r] = img
	return img
}

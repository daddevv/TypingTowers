package menu

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// TextWidth estimates the width of a string for a given GoTextFace.
func TextWidth(face *text.GoTextFace, s string) float64 {
	return float64(len(s)) * face.Size * 0.6
}

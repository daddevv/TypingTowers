package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestBaseEntityBounds(t *testing.T) {
	e := &BaseEntity{pos: Point{X: 1, Y: 2}, width: 3, height: 4, frame: ebiten.NewImage(1, 1)}
	x, y := e.Position()
	if x != 1 || y != 2 {
		t.Errorf("position expected (1,2) got (%v,%v)", x, y)
	}
	bx, by, w, h := e.Bounds()
	if bx != 1 || by != 2 || w != 3 || h != 4 {
		t.Errorf("bounds mismatch")
	}
	if !e.Static() {
		t.Errorf("expected default static false")
	}
	e.static = true
	if !e.Static() {
		t.Errorf("static true not reported")
	}
}

package game

import (
	"github.com/daddevv/type-defense/internal/entity"
	"testing"
)

func TestPointDistance(t *testing.T) {
	p1 := entity.NewPoint(0, 0)
	p2 := entity.NewPoint(3, 4)
	if dist := p1.Distance(p2); dist != 5 {
		t.Errorf("expected distance 5 got %v", dist)
	}
}

func TestPointTranslate(t *testing.T) {
	p := entity.NewPoint(1, 2)
	p.Translate(3, 4)
	if p.X != 4 || p.Y != 6 {
		t.Errorf("expected (4,6) got (%v,%v)", p.X, p.Y)
	}
}

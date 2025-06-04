package game

import "testing"


func TestPointDistance(t *testing.T) {
    p1 := NewPoint(0, 0)
    p2 := NewPoint(3, 4)
    if dist := p1.Distance(p2); dist != 5 {
        t.Errorf("expected distance 5 got %v", dist)
    }
}

func TestPointTranslate(t *testing.T) {
    p := NewPoint(1, 2)
    p.Translate(3, 4)
    if p.X != 4 || p.Y != 6 {
        t.Errorf("expected (4,6) got (%v,%v)", p.X, p.Y)
    }
}

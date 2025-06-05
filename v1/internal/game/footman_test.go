package game

import "testing"

func TestFootmanMovement(t *testing.T) {
	f := NewFootman(0, 0)
	f.speed = 10
	f.Update(1.0)
	x, _ := f.Position()
	if x <= 0 {
		t.Errorf("expected footman to move right, got %f", x)
	}
}

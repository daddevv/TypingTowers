package building

import (
	"testing"

	"github.com/daddevv/type-defense/internal/assets"
)

func TestBaseDamageAndAlive(t *testing.T) {
	assets.InitImages()

	b := NewBase(0, 0, 2)
	if b.Health() != 2 {
		t.Fatalf("expected initial health 2 got %d", b.Health())
	}
	b.ApplyDamage(1)
	if b.Health() != 1 {
		t.Errorf("expected health 1 got %d", b.Health())
	}
	if !b.Alive() {
		t.Errorf("base should be alive")
	}
	b.ApplyDamage(1)
	if b.Alive() {
		t.Errorf("base should be dead")
	}
}

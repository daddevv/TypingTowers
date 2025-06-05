package game

import "testing"

func TestGold(t *testing.T) {
	var g Gold
	g.Add(10)
	if g.Amount() != 10 {
		t.Fatalf("expected 10 got %d", g.Amount())
	}
	if !g.Spend(5) {
		t.Fatalf("expected spend success")
	}
	if g.Amount() != 5 {
		t.Fatalf("expected 5 got %d", g.Amount())
	}
	if g.Spend(10) {
		t.Fatalf("spend should fail when amount insufficient")
	}
	if g.Amount() != 5 {
		t.Fatalf("amount changed unexpectedly to %d", g.Amount())
	}
}

func TestWood(t *testing.T) {
	var w Wood
	w.Add(3)
	if w.Amount() != 3 {
		t.Fatalf("expected 3 got %d", w.Amount())
	}
	if !w.Spend(2) || w.Amount() != 1 {
		t.Fatalf("unexpected spend result %d", w.Amount())
	}
	if w.Spend(5) {
		t.Fatalf("spend should fail when amount insufficient")
	}
}

func TestStone(t *testing.T) {
	var s Stone
	s.Add(8)
	if s.Amount() != 8 {
		t.Fatalf("expected 8 got %d", s.Amount())
	}
	if !s.Spend(1) || s.Amount() != 7 {
		t.Fatalf("unexpected spend result %d", s.Amount())
	}
}

func TestIron(t *testing.T) {
	var i Iron
	i.Add(4)
	if i.Amount() != 4 {
		t.Fatalf("expected 4 got %d", i.Amount())
	}
	if !i.Spend(4) || i.Amount() != 0 {
		t.Fatalf("unexpected spend result %d", i.Amount())
	}
	if i.Spend(1) {
		t.Fatalf("spend should fail when empty")
	}
}

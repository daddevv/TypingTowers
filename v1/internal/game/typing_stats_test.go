package game

import "testing"

func TestTypingStatsBasic(t *testing.T) {
	ts := NewTypingStats()
	ts.Record(true)
	ts.Record(false)
	ts.Record(true)
	if ts.Total() != 3 {
		t.Fatalf("expected total 3 got %d", ts.Total())
	}
	acc := ts.Accuracy()
	if acc <= 0.66 || acc >= 1 {
		t.Errorf("unexpected accuracy %.2f", acc)
	}
	if ts.WPM() < 0 {
		t.Errorf("wpm should not be negative")
	}
	_ = ts.RateMultiplier()
}

func TestTypingStatsCombo(t *testing.T) {
	ts := NewTypingStats()
	for i := 0; i < 3; i++ {
		ts.Record(true)
	}
	if ts.Combo() != 3 {
		t.Errorf("expected combo 3 got %d", ts.Combo())
	}
	ts.Record(false)
	if ts.Combo() != 0 {
		t.Errorf("combo should reset on miss")
	}
	if ts.MaxCombo() < 3 {
		t.Errorf("max combo should be at least 3")
	}
}

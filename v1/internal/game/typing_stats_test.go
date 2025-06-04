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

package core

import (
	"testing"
	"time"
)

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

func TestScoreMultiplierAndHistory(t *testing.T) {
	ts := NewTypingStats()
	ts.start = ts.start.Add(-10 * time.Second) // simulate 100 correct in 10 seconds (high WPM)
	for i := 0; i < 100; i++ {
		ts.Record(true)
	}
	if ts.ScoreMultiplier() <= 1 {
		t.Errorf("expected multiplier > 1 for good performance")
	}

	hist := PerformanceHistory{}
	hist.Record(ts)
	if hist.BestWPM <= 0 || hist.BestAccuracy <= 0 {
		t.Errorf("history not updated")
	}
}

func TestRollingWPM(t *testing.T) {
	ts := NewTypingStats()
	base := time.Now()
	ts.now = func() time.Time { return base }
	ts.events = append(ts.events, base.Add(-40*time.Second))
	for i := 0; i < 25; i++ {
		ts.events = append(ts.events, base.Add(-time.Duration(i)*time.Second))
	}
	wpm := ts.RollingWPM()
	if wpm < 9.9 || wpm > 10.1 {
		t.Errorf("expected ~10 WPM got %.2f", wpm)
	}

	wpm = ts.RollingWPM()
	if wpm < 9.9 || wpm > 10.1 {
		t.Errorf("old events should be ignored got %.2f", wpm)
	}
}

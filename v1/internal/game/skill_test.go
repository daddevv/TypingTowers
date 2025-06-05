package game

import (
	"testing"
	"time"
)

func TestSkillTreeAvailable(t *testing.T) {
	st := DefaultSkillTree()
	stats := NewTypingStats()
	stats.start = time.Now().Add(-time.Minute)
	for i := 0; i < 60; i++ {
		stats.Record(true)
	}

	avail := st.Available("", stats)
	if len(avail) != 1 {
		t.Fatalf("expected 1 available skill, got %d", len(avail))
	}

	if _, ok := st.Unlock(avail[0].ID, stats); !ok {
		t.Fatalf("failed to unlock first skill")
	}

	// 200 correct chars -> 40 wpm
	stats = NewTypingStats()
	stats.start = time.Now().Add(-time.Minute)
	for i := 0; i < 200; i++ {
		stats.Record(true)
	}

	avail = st.Available("", stats)
	if len(avail) != 2 {
		t.Fatalf("expected 2 available skills after training, got %d", len(avail))
	}
}

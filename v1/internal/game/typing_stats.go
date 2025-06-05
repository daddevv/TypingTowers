package game

import "time"

// TypingStats tracks typing performance metrics.
type TypingStats struct {
	start     time.Time
	correct   int
	incorrect int
	combo     int
	maxCombo  int
	events    []time.Time
	now       func() time.Time
}

// NewTypingStats initializes a TypingStats value.
func NewTypingStats() TypingStats {
	return TypingStats{start: time.Now(), now: time.Now}
}

// Record updates the stats with whether a typed letter was correct.
func (ts *TypingStats) Record(correct bool) {
	if correct {
		ts.correct++
		ts.combo++
		if ts.combo > ts.maxCombo {
			ts.maxCombo = ts.combo
		}
	} else {
		ts.incorrect++
		ts.combo = 0
	}
	ts.recordEvent(ts.now())
}

// recordEvent adds a timestamped entry and prunes old events.
func (ts *TypingStats) recordEvent(t time.Time) {
	ts.events = append(ts.events, t)
	ts.trimOld(t)
}

// trimOld removes events older than 30 seconds from the provided time.
func (ts *TypingStats) trimOld(now time.Time) {
	cutoff := now.Add(-30 * time.Second)
	idx := 0
	for i, e := range ts.events {
		if !e.Before(cutoff) {
			idx = i
			break
		}
		// If all events are old, idx should be set to len(ts.events)
		if i == len(ts.events)-1 {
			idx = len(ts.events)
		}
	}
	if idx > 0 {
		ts.events = ts.events[idx:]
	}
}

// RollingWPM returns the words per minute calculated from
// typing events within the last 30 seconds.
func (ts *TypingStats) RollingWPM() float64 {
	now := ts.now()
	ts.trimOld(now)
	return (float64(len(ts.events)) / 5.0) / 0.5 // 30s = 0.5 min
}

// Total returns the total number of recorded letters.
func (ts *TypingStats) Total() int {
	return ts.correct + ts.incorrect
}

// Accuracy returns typing accuracy as a fraction between 0 and 1.
func (ts *TypingStats) Accuracy() float64 {
	total := ts.Total()
	if total == 0 {
		return 1
	}
	return float64(ts.correct) / float64(total)
}

// WPM returns words per minute using a 5 chars per word estimate.
func (ts *TypingStats) WPM() float64 {
	mins := time.Since(ts.start).Minutes()
	if mins <= 0 {
		return 0
	}
	return (float64(ts.Total()) / 5.0) / mins
}

// Combo returns the current combo count of consecutive correct letters.
func (ts *TypingStats) Combo() int { return ts.combo }

// MaxCombo returns the highest combo achieved.
func (ts *TypingStats) MaxCombo() int { return ts.maxCombo }

// RateMultiplier returns a fire rate multiplier based on typing performance.
func (ts *TypingStats) RateMultiplier() float64 {
	acc := ts.Accuracy()
	wpm := ts.WPM()
	if wpm >= 40 && acc >= 0.9 {
		return 0.9
	}
	if wpm < 20 || acc < 0.6 {
		return 1.2
	}
	mult := 1.0
	if ts.combo >= 5 {
		mult *= 0.85
	}
	return mult
}

// ScoreMultiplier returns a multiplier for score/gold rewards based on typing
// performance. Higher WPM and accuracy grant better rewards while poor
// performance reduces them.
func (ts *TypingStats) ScoreMultiplier() float64 {
	acc := ts.Accuracy()
	wpm := ts.WPM()
	switch {
	case wpm >= 60 && acc >= 0.95:
		return 2.0
	case wpm >= 40 && acc >= 0.9:
		return 1.5
	case wpm < 20 || acc < 0.6:
		return 0.5
	default:
		return 1.0
	}
}

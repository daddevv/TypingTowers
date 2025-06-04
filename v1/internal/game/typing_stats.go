package game

import "time"

// TypingStats tracks typing performance metrics.
type TypingStats struct {
	start     time.Time
	correct   int
	incorrect int
}

// NewTypingStats initializes a TypingStats value.
func NewTypingStats() TypingStats {
	return TypingStats{start: time.Now()}
}

// Record updates the stats with whether a typed letter was correct.
func (ts *TypingStats) Record(correct bool) {
	if correct {
		ts.correct++
	} else {
		ts.incorrect++
	}
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
	return 1.0
}

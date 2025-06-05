package game

import "time"

// WordStat captures accuracy and completion time for a single word.
type WordStat struct {
	Text      string        // word text
	Correct   int           // correct letters typed
	Incorrect int           // incorrect letters typed
	Duration  time.Duration // time from first letter to completion
	start     time.Time     // internal start time
}

// Start begins timing for the word if not already started.
func (ws *WordStat) Start() {
	if ws.start.IsZero() {
		ws.start = time.Now()
	}
}

// Finish marks the word as completed and records its duration.
func (ws *WordStat) Finish() {
	if !ws.start.IsZero() {
		ws.Duration = time.Since(ws.start)
	}
}

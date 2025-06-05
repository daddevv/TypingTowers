package game

import "testing"

// TestWordDensitySimulation runs a 5 minute simulation with the default Farmer
// and Barracks to ensure word generation stays within the 0.1-0.3 words/sec
// target. Words are assumed to be completed instantly when generated.
func TestWordDensitySimulation(t *testing.T) {
	q := NewQueueManager()
	f := NewFarmer()
	b := NewBarracks()
	f.SetQueue(q)
	b.SetQueue(q)

	duration := 300.0 // seconds (5 minutes)
	dt := 0.1
	words := 0

	for elapsed := 0.0; elapsed < duration; elapsed += dt {
		if w := f.Update(dt); w != "" {
			words++
			q.TryDequeue(w)
			f.OnWordCompleted(w, nil)
		}
		if w := b.Update(dt); w != "" {
			words++
			q.TryDequeue(w)
			b.OnWordCompleted(w)
		}
	}

	rate := float64(words) / duration
	if rate < 0.1 || rate > 0.3 {
		t.Fatalf("word generation rate %.2f outside target", rate)
	}
}

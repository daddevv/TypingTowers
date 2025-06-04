package game

// Performance represents a single game's typing performance metrics.
type Performance struct {
	WPM      float64
	Accuracy float64
}

// PerformanceHistory tracks best stats and historical records.
type PerformanceHistory struct {
	BestWPM      float64
	BestAccuracy float64
	Records      []Performance
}

// Record adds a new performance entry and updates best metrics.
func (ph *PerformanceHistory) Record(ts TypingStats) {
	p := Performance{WPM: ts.WPM(), Accuracy: ts.Accuracy()}
	ph.Records = append(ph.Records, p)
	if p.WPM > ph.BestWPM {
		ph.BestWPM = p.WPM
	}
	if p.Accuracy > ph.BestAccuracy {
		ph.BestAccuracy = p.Accuracy
	}
}

package core

// CooldownTimer tracks time remaining until an action is ready.
type CooldownTimer struct {
	Interval  float64
	Remaining float64
}

// NewCooldownTimer creates a timer with the given interval.
func NewCooldownTimer(interval float64) CooldownTimer {
	return CooldownTimer{Interval: interval, Remaining: interval}
}

// Tick decreases the timer by dt seconds. It returns true when the timer reaches
// zero.
func (t *CooldownTimer) Tick(dt float64) bool {
	if t.Remaining > 0 {
		t.Remaining -= dt
		if t.Remaining < 0 {
			t.Remaining = 0
		}
	}
	return t.Remaining <= 0
}

// Reset sets the timer back to its full interval.
func (t *CooldownTimer) Reset() { t.Remaining = t.Interval }

// Ready reports whether the timer has completed.
func (t *CooldownTimer) Ready() bool { return t.Remaining <= 0 }

// SetInterval updates the interval and clamps remaining time if needed.
func (t *CooldownTimer) SetInterval(interval float64) {
	t.Interval = interval
	if t.Remaining > interval {
		t.Remaining = interval
	}
}

// Progress returns a value between 0 and 1 indicating how much of the
// interval has elapsed. 0 means the timer has just been reset, 1 means
// it has fully completed.
func (t *CooldownTimer) Progress() float64 {
	if t.Interval <= 0 {
		return 1
	}
	elapsed := t.Interval - t.Remaining
	if elapsed < 0 {
		elapsed = 0
	}
	if elapsed > t.Interval {
		elapsed = t.Interval
	}
	return elapsed / t.Interval
}

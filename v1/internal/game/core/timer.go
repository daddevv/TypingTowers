package game

// CooldownTimer tracks time remaining until an action is ready.
type CooldownTimer struct {
	interval  float64
	remaining float64
}

// NewCooldownTimer creates a timer with the given interval.
func NewCooldownTimer(interval float64) CooldownTimer {
	return CooldownTimer{interval: interval, remaining: interval}
}

// Tick decreases the timer by dt seconds. It returns true when the timer reaches
// zero.
func (t *CooldownTimer) Tick(dt float64) bool {
	if t.remaining > 0 {
		t.remaining -= dt
		if t.remaining < 0 {
			t.remaining = 0
		}
	}
	return t.remaining <= 0
}

// Reset sets the timer back to its full interval.
func (t *CooldownTimer) Reset() { t.remaining = t.interval }

// Ready reports whether the timer has completed.
func (t *CooldownTimer) Ready() bool { return t.remaining <= 0 }

// SetInterval updates the interval and clamps remaining time if needed.
func (t *CooldownTimer) SetInterval(interval float64) {
	t.interval = interval
	if t.remaining > interval {
		t.remaining = interval
	}
}

// Remaining exposes the time left on the timer.
func (t *CooldownTimer) Remaining() float64 { return t.remaining }

// Progress returns a value between 0 and 1 indicating how much of the
// interval has elapsed. 0 means the timer has just been reset, 1 means
// it has fully completed.
func (t *CooldownTimer) Progress() float64 {
	if t.interval <= 0 {
		return 1
	}
	elapsed := t.interval - t.remaining
	if elapsed < 0 {
		elapsed = 0
	}
	if elapsed > t.interval {
		elapsed = t.interval
	}
	return elapsed / t.interval
}

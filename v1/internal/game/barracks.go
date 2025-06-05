package game

import "math/rand"

// Barracks represents a Military building that trains Footman units.
type Barracks struct {
	timer       CooldownTimer // cooldown timer for word generation
	letterPool  []rune        // available letters for word generation
	wordLenMin  int
	wordLenMax  int
	lastWord    string        // last generated word (for testing/debug)
	pendingWord string        // word currently in queue (if any)
	active      bool          // is the Barracks running?
	queue       *QueueManager // optional global queue manager
	military    *Military     // optional military system to track units
}

// NewBarracks creates a new Barracks with default settings.
func NewBarracks() *Barracks {
	return &Barracks{
		// Faster cadence to match farmer and reach 1–1.5 words/sec total
		timer:      NewCooldownTimer(2.0), // 2 seconds base cooldown
		letterPool: []rune{'f', 'j'},
		// Barracks words are slightly longer for difficulty
		wordLenMin: 3,
		wordLenMax: 5,
		active:     true,
		queue:      nil,
		military:   nil,
	}
}

// Update ticks the Barracks cooldown and returns a word when ready.
func (b *Barracks) Update(dt float64) string {
	if !b.active || b.pendingWord != "" {
		return ""
	}
	if b.timer.Tick(dt) {
		word := b.generateWord()
		b.pendingWord = word
		if b.queue != nil {
			b.queue.Enqueue(Word{Text: word, Source: "Barracks", Family: "Military"})
		}
		return word
	}
	return ""
}

// generateWord creates a random word from the Barracks letter pool.
func (b *Barracks) generateWord() string {
	length := b.wordLenMin
	if b.wordLenMax > b.wordLenMin {
		length += rand.Intn(b.wordLenMax - b.wordLenMin + 1)
	}
	word := make([]rune, length)
	for i := 0; i < length; i++ {
		word[i] = b.letterPool[rand.Intn(len(b.letterPool))]
	}
	b.lastWord = string(word)
	return b.lastWord
}

// OnWordCompleted spawns a Footman if the provided word matches the pending one.
func (b *Barracks) OnWordCompleted(word string) *Footman {
	if word == b.pendingWord {
		b.pendingWord = ""
		b.timer.Reset()
		unit := NewFootman(0, 0)
		if b.military != nil {
			b.military.AddUnit(unit)
		}
		return unit
	}
	return nil
}

// SetLetterPool updates the Barracks letter pool.
func (b *Barracks) SetLetterPool(pool []rune) { b.letterPool = pool }

// SetActive enables or disables the Barracks.
func (b *Barracks) SetActive(active bool) { b.active = active }

// SetInterval changes the base cooldown interval.
func (b *Barracks) SetInterval(interval float64) { b.timer.SetInterval(interval) }

// SetCooldown sets the remaining cooldown directly (for testing).
func (b *Barracks) SetCooldown(c float64) { b.timer.remaining = c }

// SetQueue assigns a QueueManager for global word management.
func (b *Barracks) SetQueue(q *QueueManager) { b.queue = q }

// SetMilitary assigns a Military system for unit tracking.
func (b *Barracks) SetMilitary(m *Military) { b.military = m }

// CooldownProgress returns 0 when the timer was just reset and 1 when ready.
func (b *Barracks) CooldownProgress() float64 { return b.timer.Progress() }

// CooldownRemaining exposes the remaining cooldown time.
func (b *Barracks) CooldownRemaining() float64 { return b.timer.Remaining() }

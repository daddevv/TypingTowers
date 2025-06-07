package building

import (
	"math/rand"

	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/econ"
)

// Barracks represents a Military building that trains Footman units.
type Barracks struct {
	Timer       core.CooldownTimer // cooldown timer for word generation
	LetterPool  []rune             // available letters for word generation
	UnlockStage int                // stage index
	WordLenMin  int
	WordLenMax  int
	LastWord    string             // last generated word (for testing/debug)
	PendingWord string             // word currently in queue (if any)
	Active      bool               // is the Barracks running?
	Queue       *core.WordQueue // global queue manager
}

// NewBarracksWithOptions creates a Barracks with the provided options.
func NewBarracksWithOptions(options ...BarracksOption) *Barracks {
	b := &Barracks{
		Timer:       core.NewCooldownTimer(15.0), // default cooldown interval
		LetterPool:  []rune{'f', 'j'},            // default letter pool
		UnlockStage: 0,                           // default letter stage
		WordLenMin:  2,                           // default minimum word length
		WordLenMax:  3,                           // default maximum word length
		Active:      true,                        // default active state
		Queue:       nil,                         // no queue by default
	}
	// Apply all provided options to the Barracks instance
	for _, opt := range options {
		opt(b)
	}
	return b
}

// BarracksOption defines a function type for configuring Barracks options.
// It allows for flexible configuration of Barracks instances using functional options.
type BarracksOption func(*Barracks)

// WithLetterPool sets a custom letter pool for the Barracks.
func WithLetterPool(pool []rune) BarracksOption {
	return func(b *Barracks) {
		b.LetterPool = pool
	}
}

// WithUnlockStage sets the initial letter stage for the Barracks.
func WithUnlockStage(stage int) BarracksOption {
	return func(b *Barracks) {
		b.UnlockStage = stage
	}
}

// WithWordLength sets the minimum and maximum word lengths for the Barracks.
func WithWordLength(min, max int) BarracksOption {
	return func(b *Barracks) {
		b.WordLenMin = min
		b.WordLenMax = max
	}
}

// WithActive sets the active state of the Barracks.
func WithActive(active bool) BarracksOption {
	return func(b *Barracks) {
		b.Active = active
	}
}

// WithQueue sets the QueueManager for the Barracks.
func WithQueue(q *core.WordQueue) BarracksOption {
	return func(b *Barracks) {
		b.Queue = q
	}
}


// Update ticks the Barracks cooldown and returns a word when ready.
func (b *Barracks) Update(dt float64) string {
	if !b.Active || b.PendingWord != "" {
		return ""
	}
	if b.Timer.Tick(dt) {
		word := b.GenerateWord()
		b.PendingWord = word
		if b.Queue != nil {
			b.Queue.Enqueue(core.Word{Text: word, Source: "Barracks", Family: "Military"})
		}
		return word
	}
	return ""
}

// GenerateWord creates a random word from the Barracks letter pool.
func (b *Barracks) GenerateWord() string {
	length := b.WordLenMin
	if b.WordLenMax > b.WordLenMin {
		length += rand.Intn(b.WordLenMax - b.WordLenMin + 1)
	}
	word := make([]rune, length)
	for i := 0; i < length; i++ {
		word[i] = b.LetterPool[rand.Intn(len(b.LetterPool))]
	}
	b.LastWord = string(word)
	return b.LastWord
}

// SetLetterPool updates the Barracks letter pool.
func (b *Barracks) SetLetterPool(pool []rune) { b.LetterPool = pool }

// SetActive enables or disables the Barracks.
func (b *Barracks) SetActive(active bool) { b.Active = active }

// SetInterval changes the base cooldown interval.
func (b *Barracks) SetInterval(interval float64) { b.Timer.SetInterval(interval) }

// SetCooldown sets the remaining cooldown directly (for testing).
func (b *Barracks) SetCooldown(c float64) { b.Timer.Remaining = c }

// SetQueue assigns a QueueManager for global word management.
func (b *Barracks) SetQueue(q *core.WordQueue) { b.Queue = q }

// CooldownProgress returns 0 when the timer was just reset and 1 when ready.
func (b *Barracks) CooldownProgress() float64 { return b.Timer.Progress() }

// CooldownRemaining exposes the remaining cooldown time.
func (b *Barracks) CooldownRemaining() float64 { return b.Timer.Remaining }

// NextUnlockCost returns the King's Point cost for the next letter stage.
func (b *Barracks) NextUnlockCost() int {
	stage := b.UnlockStage + 1
	return econ.LetterStageCost(stage)
}

// UnlockNext attempts to unlock the next letter stage for the Barracks.
func (b *Barracks) UnlockNext(pool *econ.ResourcePool) bool {
	stage := b.UnlockStage + 1
	letters := econ.LetterStageLetters(stage)
	cost := econ.LetterStageCost(stage)
	if letters == nil || cost < 0 {
		return false
	}
	if pool != nil && pool.SpendKingsPoints(cost) {
		b.UnlockStage = stage
		b.LetterPool = append(b.LetterPool, letters...)
		return true
	}
	return false
}

// OnWordCompleted notifies the Barracks that its word was completed.
func (b *Barracks) OnWordCompleted(word string, w *core.Word) {
	if b.PendingWord == word {
		b.PendingWord = ""
	}
}

package structure

import (
	"math/rand"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/econ"
	"github.com/daddevv/type-defense/internal/word"
)

// Barracks represents a Military building that trains Footman units.
type Barracks struct {
	Timer       core.CooldownTimer // cooldown timer for word generation
	LetterPool  []rune             // available letters for word generation
	UnlockStage int                // next letter stage index
	WordLenMin  int
	WordLenMax  int
	LastWord    string             // last generated word (for testing/debug)
	PendingWord string             // word currently in queue (if any)
	Active      bool               // is the Barracks running?
	Queue       *word.QueueManager // optional global queue manager
}

// NewBarracks creates a new Barracks with default settings.
func NewBarracks() *Barracks {
	return &Barracks{
		// Slower cadence to reduce overall word rate
		Timer:       core.NewCooldownTimer(9.0), // 9 seconds base cooldown (was 2.0)
		LetterPool:  []rune{'f', 'j'},
		UnlockStage: 0,
		// Longer words for more time per word
		WordLenMin: 4, // was 3
		WordLenMax: 6, // was 5
		Active:     true,
		Queue:      nil,
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
			b.Queue.Enqueue(assets.Word{Text: word, Source: "Barracks", Family: "Military"})
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
func (b *Barracks) SetQueue(q *word.QueueManager) { b.Queue = q }

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

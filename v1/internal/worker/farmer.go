package worker

import (
	"math/rand"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/econ"
	"github.com/daddevv/type-defense/internal/word"
)

// Farmer represents a Gathering building that produces Food on cooldown.
type Farmer struct {
	Timer       core.CooldownTimer // cooldown timer for word generation
	LetterPool  []rune             // available letters for word generation
	UnlockStage int                // next letter stage index
	WordLenMin  int
	WordLenMax  int
	LastWord    string             // last generated word (for testing/debug)
	PendingWord string             // word currently in queue (if any)
	ResourceOut int                // amount of Food to output per completion
	Active      bool               // is the Farmer running?
	Queue       *word.QueueManager // optional global queue manager
}

// NewFarmer creates a new Farmer with default settings.
func NewFarmer() *Farmer {
	return &Farmer{
		// Slower cooldown for more manageable gameplay
		Timer:       core.NewCooldownTimer(7.0), // 7 seconds between words (was 5.0)
		LetterPool:  []rune{'f', 'j'},
		UnlockStage: 0,
		WordLenMin:  4, // Longer words (was 3)
		WordLenMax:  6, // Longer words (was 5)
		ResourceOut: 1,
		Active:      true,
		Queue:       nil,
	}
}

// Update ticks the Farmer's cooldown and pushes a word to the global queue.
// The queue processes the word letter by letter. Returns the generated word if
// one is ready, else "".
func (f *Farmer) Update(dt float64) string {
	if !f.Active || f.PendingWord != "" {
		return ""
	}
	if f.Timer.Tick(dt) {
		w := f.GenerateWord()
		f.PendingWord = w
		if f.Queue != nil {
			// Enqueue the full word; Game processes it letter by letter
			f.Queue.Enqueue(assets.Word{Text: w, Source: "Farmer", Family: "Gathering"})
		}
		return w
	}
	return ""
}

// GenerateWord creates a random word from the Farmer's letter pool.
func (f *Farmer) GenerateWord() string {
	length := f.WordLenMin
	if f.WordLenMax > f.WordLenMin {
		length += rand.Intn(f.WordLenMax - f.WordLenMin + 1)
	}
	word := make([]rune, length)
	for i := 0; i < length; i++ {
		word[i] = f.LetterPool[rand.Intn(len(f.LetterPool))]
	}
	f.LastWord = string(word)
	return f.LastWord
}

// OnWordCompleted should be called when the player completes the Farmer's word.
// Returns the amount of Food produced.
func (f *Farmer) OnWordCompleted(word string, pool *econ.ResourcePool) int {
	if word == f.PendingWord {
		f.PendingWord = ""
		f.Timer.Reset()
		if pool != nil {
			pool.AddGold(f.ResourceOut)
			pool.AddFood(f.ResourceOut)
		}
		return f.ResourceOut
	}
	return 0
}

// SetLetterPool allows updating the Farmer's available letters.
func (f *Farmer) SetLetterPool(pool []rune) {
	f.LetterPool = pool
}

// SetActive enables or disables the Farmer.
func (f *Farmer) SetActive(active bool) {
	f.Active = active
}

// SetInterval changes the base cooldown interval.
func (f *Farmer) SetInterval(interval float64) {
	f.Timer.SetInterval(interval)
}

// SetCooldown sets the remaining cooldown directly (for testing).
func (f *Farmer) SetCooldown(c float64) { f.Timer.Remaining = c }

// SetQueue assigns a QueueManager for global word management.
func (f *Farmer) SetQueue(q *word.QueueManager) { f.Queue = q }

// CooldownProgress returns 0 when the timer was just reset and 1 when ready.
func (f *Farmer) CooldownProgress() float64 { return f.Timer.Progress() }

// CooldownRemaining exposes the remaining cooldown time.
func (f *Farmer) CooldownRemaining() float64 { return f.Timer.Remaining }

// NextUnlockCost returns the King's Point cost for the next letter stage.
func (f *Farmer) NextUnlockCost() int {
	stage := f.UnlockStage + 1
	return econ.LetterStageCost(stage)
}

// UnlockNext attempts to unlock the next letter stage using the provided pool.
func (f *Farmer) UnlockNext(pool *econ.ResourcePool) bool {
	stage := f.UnlockStage + 1
	letters := econ.LetterStageLetters(stage)
	cost := econ.LetterStageCost(stage)
	if letters == nil || cost < 0 {
		return false
	}
	if pool != nil && pool.SpendKingsPoints(cost) {
		f.UnlockStage = stage
		f.LetterPool = append(f.LetterPool, letters...)
		return true
	}
	return false
}

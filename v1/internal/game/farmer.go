package game

import (
	"math/rand"
)

// Farmer represents a Gathering building that produces Food on cooldown.
type Farmer struct {
	timer       CooldownTimer // cooldown timer for word generation
	letterPool  []rune        // available letters for word generation
	unlockStage int           // next letter stage index
	wordLenMin  int
	wordLenMax  int
	lastWord    string        // last generated word (for testing/debug)
	pendingWord string        // word currently in queue (if any)
	resourceOut int           // amount of Food to output per completion
	active      bool          // is the Farmer running?
	queue       *QueueManager // optional global queue manager
}

// NewFarmer creates a new Farmer with default settings.
func NewFarmer() *Farmer {
	return &Farmer{
		// Much slower cooldown for more manageable gameplay
		timer:       NewCooldownTimer(5.0), // 5 seconds between words
		letterPool:  []rune{'f', 'j'},
		unlockStage: 0,
		wordLenMin:  3, // Slightly longer words
		wordLenMax:  5, // More predictable length range
		resourceOut: 1,
		active:      true,
		queue:       nil,
	}
}

// Update ticks the Farmer's cooldown and pushes a word to the global queue.
// The queue processes the word letter by letter. Returns the generated word if
// one is ready, else "".
func (f *Farmer) Update(dt float64) string {
	if !f.active || f.pendingWord != "" {
		return ""
	}
	if f.timer.Tick(dt) {
		word := f.generateWord()
		f.pendingWord = word
		if f.queue != nil {
			// Enqueue the full word; Game processes it letter by letter
			f.queue.Enqueue(Word{Text: word, Source: "Farmer", Family: "Gathering"})
		}
		return word
	}
	return ""
}

// generateWord creates a random word from the Farmer's letter pool.
func (f *Farmer) generateWord() string {
	length := f.wordLenMin
	if f.wordLenMax > f.wordLenMin {
		length += rand.Intn(f.wordLenMax - f.wordLenMin + 1)
	}
	word := make([]rune, length)
	for i := 0; i < length; i++ {
		word[i] = f.letterPool[rand.Intn(len(f.letterPool))]
	}
	f.lastWord = string(word)
	return f.lastWord
}

// OnWordCompleted should be called when the player completes the Farmer's word.
// Returns the amount of Food produced.
func (f *Farmer) OnWordCompleted(word string, pool *ResourcePool) int {
	if word == f.pendingWord {
		f.pendingWord = ""
		f.timer.Reset()
		if pool != nil {
			pool.AddGold(f.resourceOut)
			pool.AddFood(f.resourceOut)
		}
		return f.resourceOut
	}
	return 0
}

// SetLetterPool allows updating the Farmer's available letters.
func (f *Farmer) SetLetterPool(pool []rune) {
	f.letterPool = pool
}

// SetActive enables or disables the Farmer.
func (f *Farmer) SetActive(active bool) {
	f.active = active
}

// SetInterval changes the base cooldown interval.
func (f *Farmer) SetInterval(interval float64) {
	f.timer.SetInterval(interval)
}

// SetCooldown sets the remaining cooldown directly (for testing).
func (f *Farmer) SetCooldown(c float64) { f.timer.remaining = c }

// SetQueue assigns a QueueManager for global word management.
func (f *Farmer) SetQueue(q *QueueManager) { f.queue = q }

// CooldownProgress returns 0 when the timer was just reset and 1 when ready.
func (f *Farmer) CooldownProgress() float64 { return f.timer.Progress() }

// CooldownRemaining exposes the remaining cooldown time.
func (f *Farmer) CooldownRemaining() float64 { return f.timer.Remaining() }

// NextUnlockCost returns the King's Point cost for the next letter stage.
func (f *Farmer) NextUnlockCost() int {
	stage := f.unlockStage + 1
	return LetterStageCost(stage)
}

// UnlockNext attempts to unlock the next letter stage using the provided pool.
func (f *Farmer) UnlockNext(pool *ResourcePool) bool {
	stage := f.unlockStage + 1
	letters := LetterStageLetters(stage)
	cost := LetterStageCost(stage)
	if letters == nil || cost < 0 {
		return false
	}
	if pool != nil && pool.SpendKingsPoints(cost) {
		f.unlockStage = stage
		f.letterPool = append(f.letterPool, letters...)
		return true
	}
	return false
}

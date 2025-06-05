package game

import (
	"math/rand"
)

// Farmer represents a Gathering building that produces Food on cooldown.
type Farmer struct {
	cooldown    float64 // seconds until next word is pushed
	interval    float64 // base cooldown interval (seconds)
	letterPool  []rune  // available letters for word generation
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
		interval:    3.0, // 3 seconds base cooldown
		cooldown:    3.0,
		letterPool:  []rune{'f', 'j'},
		wordLenMin:  2,
		wordLenMax:  3,
		resourceOut: 1,
		active:      true,
		queue:       nil,
	}
}

// Update ticks the Farmer's cooldown and pushes a word if ready.
// Returns the generated word if one is ready, else "".
func (f *Farmer) Update(dt float64) string {
	if !f.active {
		return ""
	}
	f.cooldown -= dt
	if f.cooldown <= 0 {
		word := f.generateWord()
		f.pendingWord = word
		if f.queue != nil {
			f.queue.Enqueue(Word{Text: word, Source: "Farmer"})
		}
		f.cooldown = f.interval
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
func (f *Farmer) OnWordCompleted(word string) int {
	if word == f.pendingWord {
		f.pendingWord = ""
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

// SetQueue assigns a QueueManager for global word management.
func (f *Farmer) SetQueue(q *QueueManager) { f.queue = q }

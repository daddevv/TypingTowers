package game

import "math/rand"

// Barracks represents a Military building that trains Footman units.
type Barracks struct {
	cooldown    float64 // seconds until next word is pushed
	interval    float64 // base cooldown interval (seconds)
	letterPool  []rune  // available letters for word generation
	wordLenMin  int
	wordLenMax  int
	lastWord    string        // last generated word (for testing/debug)
	pendingWord string        // word currently in queue (if any)
	active      bool          // is the Barracks running?
	queue       *QueueManager // optional global queue manager
}

// NewBarracks creates a new Barracks with default settings.
func NewBarracks() *Barracks {
	return &Barracks{
		interval:   5.0, // 5 seconds base cooldown
		cooldown:   5.0,
		letterPool: []rune{'f', 'j'},
		wordLenMin: 2,
		wordLenMax: 3,
		active:     true,
		queue:      nil,
	}
}

// Update ticks the Barracks cooldown and returns a word when ready.
func (b *Barracks) Update(dt float64) string {
	if !b.active {
		return ""
	}
	b.cooldown -= dt
	if b.cooldown <= 0 {
		word := b.generateWord()
		b.pendingWord = word
		if b.queue != nil {
			b.queue.Enqueue(Word{Text: word, Source: "Barracks"})
		}
		b.cooldown = b.interval
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
		return NewFootman(0, 0)
	}
	return nil
}

// SetLetterPool updates the Barracks letter pool.
func (b *Barracks) SetLetterPool(pool []rune) { b.letterPool = pool }

// SetActive enables or disables the Barracks.
func (b *Barracks) SetActive(active bool) { b.active = active }

// SetQueue assigns a QueueManager for global word management.
func (b *Barracks) SetQueue(q *QueueManager) { b.queue = q }

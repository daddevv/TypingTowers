package core

import (
	"unicode"
)

// WordQueue maintains a global FIFO queue of words.
type WordQueue struct {
	queue    []Word
	timer    float64
	progress int // typed letters progress for first word
}

// NewQueueManager initializes an empty queue.
func NewQueueManager() *WordQueue {
	return &WordQueue{queue: make([]Word, 0), timer: 0, progress: 0}
}

// Enqueue adds a word to the end of the queue.
func (q *WordQueue) Enqueue(w Word) {
	q.queue = append(q.queue, w)
}

// Progress returns the completion ratio of the first word, 0-1.
func (q *WordQueue) Progress() float64 {
	if len(q.queue) == 0 {
		return 0
	}
	if len(q.queue[0].Text) == 0 {
		return 0
	}
	return float64(q.progress) / float64(len(q.queue[0].Text))
}

// Index returns the current typed letter index for the first word.
func (q *WordQueue) Index() int { return q.progress }

// ResetProgress clears the current letter index for the first word.
func (q *WordQueue) ResetProgress() {
	q.progress = 0
}

// TryLetter validates a single typed letter against the first word.
// It returns (matched, completed, word).
func (q *WordQueue) TryLetter(r rune) (bool, bool, Word) {
	if len(q.queue) == 0 {
		return false, false, Word{}
	}
	w := q.queue[0]
	expected := rune(w.Text[q.progress])
	if unicode.ToLower(r) != unicode.ToLower(expected) {
		q.progress = 0
		return false, false, Word{}
	}
	q.progress++
	if q.progress >= len(w.Text) {
		q.queue = q.queue[1:]
		q.progress = 0
		return true, true, w
	}
	return true, false, w
}

// Update applies back-pressure damage if backlog length exceeds threshold.
func (q *WordQueue) Update(dt float64, base ...interface{ ApplyDamage(int) }) {
	const threshold = 5
	if len(q.queue) > threshold && len(base) > 0 && base[0] != nil {
		// Apply 1 damage per update if backlog exceeds threshold
		base[0].ApplyDamage(1)
	}
}

// Len returns the number of words currently in the queue.
func (q *WordQueue) Len() int { return len(q.queue) }

// Peek returns the first word without removing it. ok is false if the queue is empty.
func (q *WordQueue) Peek() (w Word, ok bool) {
	if len(q.queue) == 0 {
		return Word{}, false
	}
	return q.queue[0], true
}

// TryDequeue compares input with the first word. If they match, the word is removed
// and returned with ok=true. Otherwise the queue is unchanged and ok=false.
func (q *WordQueue) TryDequeue(input string) (Word, bool) {
	if len(q.queue) == 0 {
		return Word{}, false
	}
	if q.queue[0].Text == input {
		w := q.queue[0]
		q.queue = q.queue[1:]
		return w, true
	}
	return Word{}, false
}

// Words returns the slice of queued words (read-only).
func (q *WordQueue) Words() []Word {
	return q.queue
}

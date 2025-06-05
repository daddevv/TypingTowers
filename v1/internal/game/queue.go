package game

import "unicode"

// Word represents a queued typing challenge produced by a building.
type Word struct {
	Text   string // text the player must type
	Source string // name of the building that generated the word
	Family string // building family for colour coding
}

// QueueManager maintains a global FIFO queue of words.
type QueueManager struct {
	queue    []Word
	base     *Base
	timer    float64
	progress int // typed letters progress for first word
}

// NewQueueManager initializes an empty queue.
func NewQueueManager() *QueueManager {
	return &QueueManager{queue: make([]Word, 0), base: nil, timer: 0, progress: 0}
}

// Enqueue adds a word to the end of the queue.
func (q *QueueManager) Enqueue(w Word) {
	q.queue = append(q.queue, w)
}

// Progress returns the completion ratio of the first word, 0-1.
func (q *QueueManager) Progress() float64 {
	if len(q.queue) == 0 {
		return 0
	}
	if len(q.queue[0].Text) == 0 {
		return 0
	}
	return float64(q.progress) / float64(len(q.queue[0].Text))
}

// Index returns the current typed letter index for the first word.
func (q *QueueManager) Index() int { return q.progress }

// ResetProgress clears the current letter index for the first word.
func (q *QueueManager) ResetProgress() {
	q.progress = 0
}

// TryLetter validates a single typed letter against the first word.
// It returns (matched, completed, word).
func (q *QueueManager) TryLetter(r rune) (bool, bool, Word) {
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

// SetBase assigns a Base that will take damage from backlog pressure.
func (q *QueueManager) SetBase(b *Base) { q.base = b }

// Update applies back-pressure damage if backlog length exceeds threshold.
func (q *QueueManager) Update(dt float64) {
	if q.base == nil {
		return
	}
	// With letter-level queue items, allow a moderate backlog before damage
	if len(q.queue) >= 6 { // was 20, now 6 to match test expectations
		q.timer += dt
		if q.timer >= 1 {
			q.base.Damage(1)
			q.timer = 0
		}
	} else {
		q.timer = 0
	}
}

// Len returns the number of words currently in the queue.
func (q *QueueManager) Len() int { return len(q.queue) }

// Peek returns the first word without removing it. ok is false if the queue is empty.
func (q *QueueManager) Peek() (w Word, ok bool) {
	if len(q.queue) == 0 {
		return Word{}, false
	}
	return q.queue[0], true
}

// TryDequeue compares input with the first word. If they match, the word is removed
// and returned with ok=true. Otherwise the queue is unchanged and ok=false.
func (q *QueueManager) TryDequeue(input string) (Word, bool) {
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
func (q *QueueManager) Words() []Word {
	return q.queue
}

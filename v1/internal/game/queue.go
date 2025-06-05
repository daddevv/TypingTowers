package game

// Word represents a queued typing challenge produced by a building.
type Word struct {
	Text   string // text the player must type
	Source string // name of the building that generated the word
	Family string // building family for colour coding
}

// QueueManager maintains a global FIFO queue of words.
type QueueManager struct {
	queue []Word
	base  *Base
	timer float64
}

// NewQueueManager initializes an empty queue.
func NewQueueManager() *QueueManager {
	return &QueueManager{queue: make([]Word, 0), base: nil, timer: 0}
}

// Enqueue adds a word to the end of the queue.
func (q *QueueManager) Enqueue(w Word) {
	q.queue = append(q.queue, w)
}

// SetBase assigns a Base that will take damage from backlog pressure.
func (q *QueueManager) SetBase(b *Base) { q.base = b }

// Update applies back-pressure damage if backlog length exceeds threshold.
func (q *QueueManager) Update(dt float64) {
	if q.base == nil {
		return
	}
	// With letter-level queue items, allow a larger backlog before damage
	if len(q.queue) >= 20 {
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

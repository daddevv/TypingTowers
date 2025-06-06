package game

import (
	"math/rand"

	"github.com/daddevv/type-defense/internal/econ"
)

// Miner represents a Gathering building that produces Stone and Iron on cooldown.
type Miner struct {
	timer       CooldownTimer
	letterPool  []rune
	unlockStage int
	wordLenMin  int
	wordLenMax  int
	lastWord    string
	pendingWord string
	stoneOut    int
	ironOut     int
	active      bool
	queue       *QueueManager
}

// NewMiner creates a new Miner with default settings.
func NewMiner() *Miner {
	return &Miner{
		timer:       NewCooldownTimer(10.0), // 10 seconds between words (was 1.5)
		letterPool:  []rune{'f', 'j'},
		unlockStage: 0,
		wordLenMin:  4, // Longer words (was 2)
		wordLenMax:  6, // Longer words (was 4)
		stoneOut:    1,
		ironOut:     1,
		active:      true,
	}
}

// Update ticks the Miner cooldown and pushes a word if ready.
func (m *Miner) Update(dt float64) string {
	if !m.active || m.pendingWord != "" {
		return ""
	}
	if m.timer.Tick(dt) {
		word := m.generateWord()
		m.pendingWord = word
		if m.queue != nil {
			m.queue.Enqueue(Word{Text: word, Source: "Miner", Family: "Gathering"})
		}
		return word
	}
	return ""
}

func (m *Miner) generateWord() string {
	length := m.wordLenMin
	if m.wordLenMax > m.wordLenMin {
		length += rand.Intn(m.wordLenMax - m.wordLenMin + 1)
	}
	word := make([]rune, length)
	for i := 0; i < length; i++ {
		word[i] = m.letterPool[rand.Intn(len(m.letterPool))]
	}
	m.lastWord = string(word)
	return m.lastWord
}

// OnWordCompleted should be called when the word is typed. Returns stone gained.
func (m *Miner) OnWordCompleted(word string, pool *econ.ResourcePool) (int, int) {
	if word == m.pendingWord {
		m.pendingWord = ""
		m.timer.Reset()
		if pool != nil {
			pool.AddGold(m.stoneOut)
			pool.AddStone(m.stoneOut)
			pool.AddIron(m.ironOut)
		}
		return m.stoneOut, m.ironOut
	}
	return 0, 0
}

func (m *Miner) SetLetterPool(p []rune)     { m.letterPool = p }
func (m *Miner) SetActive(a bool)           { m.active = a }
func (m *Miner) SetInterval(d float64)      { m.timer.SetInterval(d) }
func (m *Miner) SetCooldown(c float64)      { m.timer.remaining = c }
func (m *Miner) SetQueue(q *QueueManager)   { m.queue = q }
func (m *Miner) CooldownProgress() float64  { return m.timer.Progress() }
func (m *Miner) CooldownRemaining() float64 { return m.timer.Remaining() }

func (m *Miner) NextUnlockCost() int {
	stage := m.unlockStage + 1
	return LetterStageCost(stage)
}

func (m *Miner) UnlockNext(pool *econ.ResourcePool) bool {
	stage := m.unlockStage + 1
	letters := LetterStageLetters(stage)
	cost := LetterStageCost(stage)
	if letters == nil || cost < 0 {
		return false
	}
	if pool != nil && pool.SpendKingsPoints(cost) {
		m.unlockStage = stage
		m.letterPool = append(m.letterPool, letters...)
		return true
	}
	return false
}

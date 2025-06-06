package game

import (
	"math/rand"

	"github.com/daddevv/type-defense/internal/econ"
)

// Lumberjack represents a Gathering building that produces Wood on cooldown.
type Lumberjack struct {
	timer       CooldownTimer
	letterPool  []rune
	unlockStage int
	wordLenMin  int
	wordLenMax  int
	lastWord    string
	pendingWord string
	resourceOut int
	active      bool
	queue       *QueueManager
}

// NewLumberjack creates a new Lumberjack with default settings.
func NewLumberjack() *Lumberjack {
	return &Lumberjack{
		timer:       NewCooldownTimer(8.0), // 8 seconds between words (was 1.5)
		letterPool:  []rune{'f', 'j'},
		unlockStage: 0,
		wordLenMin:  4, // Longer words (was 2)
		wordLenMax:  6, // Longer words (was 4)
		resourceOut: 1,
		active:      true,
	}
}

// Update ticks the Lumberjack cooldown and pushes a word if ready.
func (l *Lumberjack) Update(dt float64) string {
	if !l.active || l.pendingWord != "" {
		return ""
	}
	if l.timer.Tick(dt) {
		word := l.generateWord()
		l.pendingWord = word
		if l.queue != nil {
			l.queue.Enqueue(Word{Text: word, Source: "Lumberjack", Family: "Gathering"})
		}
		return word
	}
	return ""
}

func (l *Lumberjack) generateWord() string {
	length := l.wordLenMin
	if l.wordLenMax > l.wordLenMin {
		length += rand.Intn(l.wordLenMax - l.wordLenMin + 1)
	}
	word := make([]rune, length)
	for i := 0; i < length; i++ {
		word[i] = l.letterPool[rand.Intn(len(l.letterPool))]
	}
	l.lastWord = string(word)
	return l.lastWord
}

// OnWordCompleted should be called when the word is typed. Returns wood gained.
func (l *Lumberjack) OnWordCompleted(word string, pool *econ.ResourcePool) int {
	if word == l.pendingWord {
		l.pendingWord = ""
		l.timer.Reset()
		if pool != nil {
			pool.AddGold(l.resourceOut)
			pool.AddWood(l.resourceOut)
		}
		return l.resourceOut
	}
	return 0
}

func (l *Lumberjack) SetLetterPool(p []rune)     { l.letterPool = p }
func (l *Lumberjack) SetActive(a bool)           { l.active = a }
func (l *Lumberjack) SetInterval(d float64)      { l.timer.SetInterval(d) }
func (l *Lumberjack) SetCooldown(c float64)      { l.timer.remaining = c }
func (l *Lumberjack) SetQueue(q *QueueManager)   { l.queue = q }
func (l *Lumberjack) CooldownProgress() float64  { return l.timer.Progress() }
func (l *Lumberjack) CooldownRemaining() float64 { return l.timer.Remaining() }

func (l *Lumberjack) NextUnlockCost() int {
	stage := l.unlockStage + 1
	return LetterStageCost(stage)
}

func (l *Lumberjack) UnlockNext(pool *econ.ResourcePool) bool {
	stage := l.unlockStage + 1
	letters := LetterStageLetters(stage)
	cost := LetterStageCost(stage)
	if letters == nil || cost < 0 {
		return false
	}
	if pool != nil && pool.SpendKingsPoints(cost) {
		l.unlockStage = stage
		l.letterPool = append(l.letterPool, letters...)
		return true
	}
	return false
}

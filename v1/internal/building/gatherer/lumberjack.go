package gatherer

import (
	"math/rand"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/core"
	"github.com/daddevv/type-defense/internal/econ"
	"github.com/daddevv/type-defense/internal/word"
)

// Lumberjack represents a Gathering building that produces Wood on cooldown.
type Lumberjack struct {
	Timer       core.CooldownTimer
	LetterPool  []rune
	UnlockStage int
	WordLenMin  int
	WordLenMax  int
	LastWord    string
	PendingWord string
	ResourceOut int
	Active      bool
	Queue       *word.QueueManager
}

// NewLumberjack creates a new Lumberjack with default settings.
func NewLumberjack() *Lumberjack {
	return &Lumberjack{
		Timer:       core.NewCooldownTimer(8.0), // 8 seconds between words (was 1.5)
		LetterPool:  []rune{'f', 'j'},
		UnlockStage: 0,
		WordLenMin:  4, // Longer words (was 2)
		WordLenMax:  6, // Longer words (was 4)
		ResourceOut: 1,
		Active:      true,
	}
}

// Update ticks the Lumberjack cooldown and pushes a word if ready.
func (l *Lumberjack) Update(dt float64) string {
	if !l.Active || l.PendingWord != "" {
		return ""
	}
	if l.Timer.Tick(dt) {
		w := l.GenerateWord()
		l.PendingWord = w
		if l.Queue != nil {
			l.Queue.Enqueue(assets.Word{Text: w, Source: "Lumberjack", Family: "Gathering"})
		}
		return w
	}
	return ""
}

func (l *Lumberjack) GenerateWord() string {
	length := l.WordLenMin
	if l.WordLenMax > l.WordLenMin {
		length += rand.Intn(l.WordLenMax - l.WordLenMin + 1)
	}
	word := make([]rune, length)
	for i := 0; i < length; i++ {
		word[i] = l.LetterPool[rand.Intn(len(l.LetterPool))]
	}
	l.LastWord = string(word)
	return l.LastWord
}

// OnWordCompleted should be called when the word is typed. Returns wood gained.
func (l *Lumberjack) OnWordCompleted(word string, pool *econ.ResourcePool) int {
	if word == l.PendingWord {
		l.PendingWord = ""
		l.Timer.Reset()
		if pool != nil {
			pool.AddGold(l.ResourceOut)
			pool.AddWood(l.ResourceOut)
		}
		return l.ResourceOut
	}
	return 0
}

func (l *Lumberjack) SetLetterPool(p []rune)        { l.LetterPool = p }
func (l *Lumberjack) SetActive(a bool)              { l.Active = a }
func (l *Lumberjack) SetInterval(d float64)         { l.Timer.SetInterval(d) }
func (l *Lumberjack) SetCooldown(c float64)         { l.Timer.Remaining = c }
func (l *Lumberjack) SetQueue(q *word.QueueManager) { l.Queue = q }
func (l *Lumberjack) CooldownProgress() float64     { return l.Timer.Progress() }
func (l *Lumberjack) CooldownRemaining() float64    { return l.Timer.Remaining }

func (l *Lumberjack) NextUnlockCost() int {
	stage := l.UnlockStage + 1
	return econ.LetterStageCost(stage)
}

func (l *Lumberjack) UnlockNext(pool *econ.ResourcePool) bool {
	stage := l.UnlockStage + 1
	letters := econ.LetterStageLetters(stage)
	cost := econ.LetterStageCost(stage)
	if letters == nil || cost < 0 {
		return false
	}
	if pool != nil && pool.SpendKingsPoints(cost) {
		l.UnlockStage = stage
		l.LetterPool = append(l.LetterPool, letters...)
		return true
	}
	return false
}

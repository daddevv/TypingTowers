package entity

import (
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Draw(screen *ebiten.Image)
	Update() error
	SetPosition(x, y float64)
	GetPosition() ui.Location
}

func EmptyList() []Entity {
	return make([]Entity, 0)
}

// Mob is an interface for all mob types (expandable for future mobs)
type Mob interface {
	Entity
	GetLetters() []Letter
	IsDead() bool
	IsPendingDeath() bool
	StartDeath()
	IncrementPendingProjectiles()
	GetPendingProjectiles() int
}

// MobLetterController is an optional interface for mobs that support advancing letter state directly (for rapid typing feedback).
type MobLetterController interface {
	AdvanceLetterState(char rune)
}

// LetterPool defines an interface for managing available letters for mobs
type LetterPool interface {
	GetPossibleLetters() []string
	AddLetter(letter string)
	AddLetters(letters []string)
}

// DefaultLetterPool is a basic implementation of LetterPool
// It expands the set of possible letters as the score increases
type DefaultLetterPool struct {
	allLetters   []string
	possible     []string
	unlockOrder  []string // order in which letters are unlocked
	unlockedCount int // how many letters have been unlocked
}

func NewDefaultLetterPool() *DefaultLetterPool {
	// Example: unlock vowels at 0, then common consonants, then rare
	indexFingers := []string{"f","j","g","h","r","u","t","y","v","m","b","n"}
	midFingers := []string{"d","k","e","i","c",","}
	ringFingers := []string{"s","l","w","o","x","."}
	pinkyFingers := []string{"a",";","q","p","z","/"}
	// Combine all fingers into a single unlock order
	unlockOrder := append(indexFingers, midFingers...)
	unlockOrder = append(unlockOrder, ringFingers...)
	unlockOrder = append(unlockOrder, pinkyFingers...)
	return &DefaultLetterPool{
		allLetters:   unlockOrder,
		possible:     []string{},
		unlockOrder:  unlockOrder,
		unlockedCount: 0,
	}
}

func (lp *DefaultLetterPool) GetPossibleLetters() []string {
	return lp.possible
}

func (lp *DefaultLetterPool) AddLetter(letter string) {
	if !contains(lp.possible, letter) {
		lp.possible = append(lp.possible, letter)
	}
}

func (lp *DefaultLetterPool) AddLetters(letters []string) {
	for _, l := range letters {
		lp.AddLetter(l)
	}
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

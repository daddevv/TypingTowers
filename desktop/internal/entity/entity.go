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
	Update(score int)
	AddLetter(letter string)
	AddLetters(letters []string)
}

// DefaultLetterPool is a basic implementation of LetterPool
// It expands the set of possible letters as the score increases
type DefaultLetterPool struct {
	allLetters   []string
	possible     []string
	unlockOrder  []string // order in which letters are unlocked
	unlockScores []int    // score thresholds for unlocking
}

func NewDefaultLetterPool() *DefaultLetterPool {
	// Example: unlock vowels at 0, then common consonants, then rare
	unlockOrder := []string{"a","e","i","o","u","n","s","t","r","l","d","c","m","p","b","g","h","f","y","w","k","v","x","z","j","q"}
	unlockScores := []int{0,0,0,0,0,10,20,30,40,50,60,70,80,90,100,110,120,130,140,150,160,170,180,190,200,210}
	return &DefaultLetterPool{
		allLetters:   unlockOrder,
		possible:     []string{"a","e","i","o","u"},
		unlockOrder:  unlockOrder,
		unlockScores: unlockScores,
	}
}

func (lp *DefaultLetterPool) GetPossibleLetters() []string {
	return lp.possible
}

func (lp *DefaultLetterPool) Update(score int) {
	// Expand possible letters based on score
	for i, letter := range lp.unlockOrder {
		if lp.unlockScores[i] <= score && !contains(lp.possible, letter) {
			lp.possible = append(lp.possible, letter)
		}
	}
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

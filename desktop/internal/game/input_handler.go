package game

import (
	"math"
	"td/internal/entity"
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InputHandler handles player input for typing letters and targeting mobs.
type InputHandler struct {
	playerPosition ui.Location // Position to shoot projectiles from
}

// NewInputHandler creates a new input handler.
func NewInputHandler(playerPos ui.Location) *InputHandler {
	return &InputHandler{
		playerPosition: playerPos,
	}
}

// SetPlayerPosition updates the player position for projectile origin.
func (ih *InputHandler) SetPlayerPosition(pos ui.Location) {
	ih.playerPosition = pos
}

// ProcessInput handles keyboard input and checks for letter matches with mobs.
// Returns a slice of projectiles to add to the game.
func (ih *InputHandler) ProcessInput(mobs []entity.Entity) []*entity.Projectile {
	var projectiles []*entity.Projectile
	
	// Check for any key presses
	for key := ebiten.Key(0); key <= ebiten.KeyMax; key++ {
		if inpututil.IsKeyJustPressed(key) {
			// Convert key to character
			char := ih.keyToChar(key)
			if char == 0 {
				continue // Not a valid letter key
			}
			
			// Find the closest mob with a matching target letter
			targetMob := ih.findClosestTargetMob(mobs, char)
			if targetMob != nil {
				// Create projectile aimed at the mob
				projectile := entity.NewProjectile(ih.playerPosition, targetMob.GetPosition())
				projectiles = append(projectiles, projectile)
				
				// Advance the mob's letter targeting
				ih.advanceMobLetters(targetMob)
			}
		}
	}
	
	return projectiles
}

// keyToChar converts an ebiten key to a lowercase character.
func (ih *InputHandler) keyToChar(key ebiten.Key) rune {
	switch key {
	case ebiten.KeyA:
		return 'a'
	case ebiten.KeyB:
		return 'b'
	case ebiten.KeyC:
		return 'c'
	case ebiten.KeyD:
		return 'd'
	case ebiten.KeyE:
		return 'e'
	case ebiten.KeyF:
		return 'f'
	case ebiten.KeyG:
		return 'g'
	case ebiten.KeyH:
		return 'h'
	case ebiten.KeyI:
		return 'i'
	case ebiten.KeyJ:
		return 'j'
	case ebiten.KeyK:
		return 'k'
	case ebiten.KeyL:
		return 'l'
	case ebiten.KeyM:
		return 'm'
	case ebiten.KeyN:
		return 'n'
	case ebiten.KeyO:
		return 'o'
	case ebiten.KeyP:
		return 'p'
	case ebiten.KeyQ:
		return 'q'
	case ebiten.KeyR:
		return 'r'
	case ebiten.KeyS:
		return 's'
	case ebiten.KeyT:
		return 't'
	case ebiten.KeyU:
		return 'u'
	case ebiten.KeyV:
		return 'v'
	case ebiten.KeyW:
		return 'w'
	case ebiten.KeyX:
		return 'x'
	case ebiten.KeyY:
		return 'y'
	case ebiten.KeyZ:
		return 'z'
	default:
		return 0 // Invalid key
	}
}

// findClosestTargetMob finds the closest mob that has the specified character as its current target.
func (ih *InputHandler) findClosestTargetMob(mobs []entity.Entity, char rune) entity.Entity {
	var closestMob entity.Entity
	closestDistance := math.Inf(1)
	
	for _, mob := range mobs {
		// Check if this mob has the target character
		if ih.mobHasTargetChar(mob, char) {
			// Calculate distance from player to mob
			mobPos := mob.GetPosition()
			dx := mobPos.X - ih.playerPosition.X
			dy := mobPos.Y - ih.playerPosition.Y
			distance := math.Sqrt(dx*dx + dy*dy)
			
			// Update closest if this is closer
			if distance < closestDistance {
				closestDistance = distance
				closestMob = mob
			}
		}
	}
	
	return closestMob
}

// mobHasTargetChar checks if the mob's current target letter matches the character.
func (ih *InputHandler) mobHasTargetChar(mob entity.Entity, char rune) bool {
	// We need to access the mob's letters. Since we're working with the interface,
	// we need to type assert to BeachballMob to access the Letters field.
	if beachballMob, ok := mob.(*entity.BeachballMob); ok {
		for _, letter := range beachballMob.Letters {
			if letter.State == entity.LetterTarget {
				return letter.Character == char
			}
		}
	}
	return false
}

// advanceMobLetters advances the mob's letter targeting to the next letter.
func (ih *InputHandler) advanceMobLetters(mob entity.Entity) {
	if beachballMob, ok := mob.(*entity.BeachballMob); ok {
		// Find current target letter and mark it as inactive
		targetIndex := -1
		for i, letter := range beachballMob.Letters {
			if letter.State == entity.LetterTarget {
				beachballMob.Letters[i].State = entity.LetterInactive
				beachballMob.Letters[i].Sprite = entity.GetLetterImage(
					letter.Character, 
					entity.LetterInactive, 
					ui.Font("Mob", 32),
				)
				targetIndex = i
				break
			}
		}
		
		// Set next letter as target if available
		if targetIndex >= 0 && targetIndex+1 < len(beachballMob.Letters) {
			nextIndex := targetIndex + 1
			beachballMob.Letters[nextIndex].State = entity.LetterTarget
			beachballMob.Letters[nextIndex].Sprite = entity.GetLetterImage(
				beachballMob.Letters[nextIndex].Character, 
				entity.LetterTarget, 
				ui.Font("Mob", 32),
			)
		}
	}
}

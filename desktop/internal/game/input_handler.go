package game

import (
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
// Now immediately advances letter states for rapid typing, firing projectiles for visual feedback.
func (ih *InputHandler) ProcessInput(mobs []entity.Entity, projectiles []*entity.Projectile) []*entity.Projectile {
	var newProjectiles []*entity.Projectile

	// Track mobs already targeted this frame to avoid double-firing
	targeted := make(map[*entity.BeachballMob]bool)

	// Get all keys pressed this frame, in order
	keys := inpututil.AppendJustPressedKeys(nil)
	for _, key := range keys {
		char := ih.keyToChar(key)
		if char == 0 {
			continue
		}
		// Find the closest mob whose current target letter matches this char and hasn't been targeted yet
		var closestMob *entity.BeachballMob
		var closestX float64 = 1e9
		for _, mob := range mobs {
			beachballMob, ok := mob.(*entity.BeachballMob)
			if !ok || targeted[beachballMob] {
				continue
			}
			for _, letter := range beachballMob.Letters {
				if letter.State == entity.LetterTarget && letter.Character == char {
					mobPos := beachballMob.GetPosition()
					if mobPos.X < closestX {
						closestX = mobPos.X
						closestMob = beachballMob
					}
					break
				}
			}
		}
		if closestMob != nil {
			// IMMEDIATELY advance letter state for rapid typing
			ih.advanceLetterState(closestMob, char)
			
			// Increment pending projectiles counter
			closestMob.PendingProjectiles++
			
			// Fire projectile for visual feedback
			mobPos := closestMob.GetPosition()
			centeredTarget := ui.Location{
				X: mobPos.X + 48.0*3.0/2.0,
				Y: mobPos.Y + 48.0*3.0,
			}
			projectile := entity.NewProjectile(ih.playerPosition, centeredTarget, closestMob)
			projectile.TargetChar = char
			newProjectiles = append(newProjectiles, projectile)
			targeted[closestMob] = true
		}
	}
	return newProjectiles
}

// advanceLetterState immediately advances the letter state for the given mob and character.
// This allows rapid typing without waiting for projectile collisions.
func (ih *InputHandler) advanceLetterState(mob *entity.BeachballMob, char rune) {
	// Find the target letter and mark it as inactive
	targetIndex := -1
	for i, letter := range mob.Letters {
		if letter.State == entity.LetterTarget && letter.Character == char {
			// Mark this letter as inactive
			mob.Letters[i].State = entity.LetterInactive
			mob.Letters[i].Sprite = entity.GetLetterImage(
				letter.Character, 
				entity.LetterInactive, 
				ui.Font("Mob", 32),
			)
			targetIndex = i
			break
		}
	}
	
	// Set next letter as target if available
	if targetIndex >= 0 && targetIndex+1 < len(mob.Letters) {
		nextIndex := targetIndex + 1
		mob.Letters[nextIndex].State = entity.LetterTarget
		mob.Letters[nextIndex].Sprite = entity.GetLetterImage(
			mob.Letters[nextIndex].Character, 
			entity.LetterTarget, 
			ui.Font("Mob", 32),
		)
	}
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

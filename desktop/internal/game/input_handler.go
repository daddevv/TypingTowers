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
// Now takes projectiles as argument to allow reservation logic.
func (ih *InputHandler) ProcessInput(mobs []entity.Entity, projectiles []*entity.Projectile) []*entity.Projectile {
	var newProjectiles []*entity.Projectile

	// Build a map of reserved (mob, letter index) pairs for all projectiles in flight
	reserved := make(map[*entity.BeachballMob]map[int]struct{})
	for _, p := range projectiles {
		if !p.IsActive() {
			continue
		}
		// Only consider projectiles with a valid mob and target char
		if mob, ok := p.TargetMob.(*entity.BeachballMob); ok && p.TargetChar != 0 {
			if reserved[mob] == nil {
				reserved[mob] = make(map[int]struct{})
			}
			// Find which letter index this projectile is for
			for i, letter := range mob.Letters {
				if letter.State == entity.LetterTarget && letter.Character == p.TargetChar {
					reserved[mob][i] = struct{}{}
					break
				}
			}
		}
	}

	// Check for any key presses
	for key := ebiten.Key(0); key <= ebiten.KeyMax; key++ {
		if inpututil.IsKeyJustPressed(key) {
			char := ih.keyToChar(key)
			if char == 0 {
				continue
			}
			// Find the closest mob/letter that is not reserved
			targetMob, letterIdx := ih.findClosestUnreservedTargetMob(mobs, char, reserved)
			if targetMob != nil && letterIdx >= 0 {
				// Aim at the center of the mob sprite (sprite is 48x48, scaled by 3)
				mobPos := targetMob.GetPosition()
				centeredTarget := ui.Location{
					X: mobPos.X + 48.0*3.0/2.0,
					Y: mobPos.Y + 48.0*3.0, // Aim 10px lower than center
				}
				projectile := entity.NewProjectile(ih.playerPosition, centeredTarget, targetMob)
				projectile.TargetChar = char
				newProjectiles = append(newProjectiles, projectile)
				// Mark this letter as reserved for subsequent keys in this frame
				if reserved[targetMob] == nil {
					reserved[targetMob] = make(map[int]struct{})
				}
				reserved[targetMob][letterIdx] = struct{}{}
			}
		}
	}
	return newProjectiles
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

// findClosestUnreservedTargetMob finds the closest mob and letter index for the given char that is not reserved.
func (ih *InputHandler) findClosestUnreservedTargetMob(mobs []entity.Entity, char rune, reserved map[*entity.BeachballMob]map[int]struct{}) (*entity.BeachballMob, int) {
	var closestMob *entity.BeachballMob
	closestIdx := -1
	closestDistance := math.Inf(1)

	for _, mob := range mobs {
		beachballMob, ok := mob.(*entity.BeachballMob)
		if !ok {
			continue
		}
		for i, letter := range beachballMob.Letters {
			if letter.State == entity.LetterTarget && letter.Character == char {
				// Skip if reserved
				if reserved[beachballMob] != nil {
					if _, taken := reserved[beachballMob][i]; taken {
						continue
					}
				}
				// Calculate distance from player to mob
				mobPos := beachballMob.GetPosition()
				dx := mobPos.X - ih.playerPosition.X
				dy := mobPos.Y - ih.playerPosition.Y
				distance := math.Sqrt(dx*dx + dy*dy)
				if distance < closestDistance {
					closestDistance = distance
					closestMob = beachballMob
					closestIdx = i
				}
			}
		}
	}
	return closestMob, closestIdx
}

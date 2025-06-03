package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// HUD displays simple text information about the tower state.
type HUD struct {
	game *Game
}

// NewHUD creates a new HUD bound to the given game.
func NewHUD(g *Game) *HUD {
	return &HUD{game: g}
}

// Draw renders ammo count and reload prompts.
func (h *HUD) Draw(screen *ebiten.Image) {
	if len(h.game.towers) == 0 {
		return
	}
	t := h.game.towers[0]
	ammo := fmt.Sprintf("Ammo: %d/%d", t.ammo, t.ammoCapacity)
	ebitenutil.DebugPrintAt(screen, ammo, 10, 40)
	if t.reloading {
		prompt := fmt.Sprintf("Reload in: %d", t.reloadTimer)
		if t.reloadTimer <= 0 {
			prompt = fmt.Sprintf("Type '%c'", t.reloadLetter)
		}
		ebitenutil.DebugPrintAt(screen, prompt, 10, 52)
	}
}

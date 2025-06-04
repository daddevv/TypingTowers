package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// HUD displays placeholder UI elements with basic game information.
type HUD struct {
	game *Game
}

// NewHUD creates a new HUD bound to the given game.
func NewHUD(g *Game) *HUD {
	return &HUD{game: g}
}

// Draw renders ammo count, tower stats, reload prompts, and shop interface.
func (h *HUD) Draw(screen *ebiten.Image) {
	var lines []string
	textX := 10
	initialY := 30 // Start HUD lower to avoid overlap with mouse/tile debug info
	lineHeight := 18
	padding := 5.0

	if h.game.shopOpen {
		gold := h.game.gold
		lines = append(lines, "-- SHOP --")
		lines = append(lines, fmt.Sprintf("Gold: %d", gold))

		var foresight int
		if len(h.game.towers) > 0 {
			foresight = h.game.towers[h.game.selectedTower].foresight
		}

		options := []string{
			"[1] Upgrade Damage (+1): 5 gold",
			"[2] Upgrade Range (+50): 5 gold",
			"[3] Upgrade Fire Rate (faster): 5 gold",
			"[4] Upgrade Ammo Capacity (+2): 10 gold",
			fmt.Sprintf("[5] Foresight (+2 letters) [%d]", foresight),
			"Start Next Wave",
		}

		for i, opt := range options {
			prefix := "  "
			if i == h.game.shopCursor {
				prefix = "> "
			}
			lines = append(lines, prefix+opt)
		}
	} else if h.game.buildMenuOpen {
		cost := h.game.cfg.TowerConstructionCost
		if cost == 0 {
			cost = DefaultConfig.TowerConstructionCost
		}
		lines = append(lines, "-- BUILD --")
		lines = append(lines, fmt.Sprintf("Gold: %d", h.game.gold))
		options := []string{
			"[1] Basic Tower",
			"[2] Sniper Tower",
			"[3] Rapid Tower",
			"Cancel",
		}
		for i, opt := range options {
			prefix := "  "
			if i == h.game.buildCursor {
				prefix = "> "
			}
			if i < 3 {
				lines = append(lines, fmt.Sprintf("%s%s (%d gold)", prefix, opt, cost))
			} else {
				lines = append(lines, prefix+opt)
			}
		}
	} else {
		if len(h.game.towers) > 0 {
			t := h.game.towers[h.game.selectedTower]
			lines = append(lines, fmt.Sprintf("Selected Tower: %d", h.game.selectedTower+1))
			currentAmmo, maxAmmo := t.GetAmmoStatus()
			lines = append(lines, fmt.Sprintf("Ammo: %d/%d", currentAmmo, maxAmmo))
			lines = append(lines, fmt.Sprintf("Damage: %d", t.damage))
			lines = append(lines, fmt.Sprintf("Range: %.0f", t.rangeDst))

			sps := 0.0
			if t.rate > 0 {
				sps = 1.0 / t.rate
			}
			lines = append(lines, fmt.Sprintf("Fire Speed: %.2f/s (Cooldown: %.2fs)", sps, t.rate))

			reloading, currentLetter, previewQueue, timer, jammed := t.GetReloadStatus()
			if jammed {
				lines = append(lines, "Jammed! Press Backspace")
				lines = append(lines, fmt.Sprintf("Stuck on: '%c'", currentLetter))
			} else if reloading {
				if timer <= 0 {
					queueStr := ""
					for i, letter := range previewQueue {
						if i == 0 {
							queueStr += fmt.Sprintf("[%c]", letter)
						} else {
							queueStr += fmt.Sprintf(" %c", letter)
						}
					}
					lines = append(lines, fmt.Sprintf("Type: %s", queueStr))
				} else {
					lines = append(lines, fmt.Sprintf("Reload in: %.2fs", timer))
					if len(previewQueue) > 0 {
						queueStr := ""
						for i, letter := range previewQueue {
							if i == 0 {
								queueStr += fmt.Sprintf("[%c]", letter)
							} else {
								queueStr += fmt.Sprintf(" %c", letter)
							}
						}
						lines = append(lines, fmt.Sprintf("Next: %s", queueStr))
					}
				}
			}
		}

		if h.game.base != nil {
			lines = append(lines, fmt.Sprintf("Base HP: %d", h.game.base.Health()))
		}
		lines = append(lines, fmt.Sprintf("Wave %d | Gold %d | Score %d | Mobs %d", h.game.currentWave, h.game.gold, h.game.score, len(h.game.mobs)))
		acc := h.game.typing.Accuracy() * 100
		wpm := h.game.typing.WPM()
		lines = append(lines, fmt.Sprintf("Accuracy: %.0f%% | WPM: %.1f", acc, wpm))
		cost := h.game.cfg.TowerConstructionCost
		if cost == 0 {
			cost = DefaultConfig.TowerConstructionCost
		}
		lines = append(lines, fmt.Sprintf("[h/j/k/l] move cursor | [b] build (%d gold)", cost))
	}

	if len(lines) == 0 {
		return
	}

	// Define background properties
	bgX := float64(textX) - padding
	bgY := float64(initialY) - padding
	// Dynamically calculate width based on longest line, fallback to min width
	minWidth := 520.0 // Increased from 420.0 for more space
	bgWidth := minWidth
	for _, line := range lines {
		w := float64(len(line)) * 13.0 // crude estimate: 13px per char, was 10.0
		if w+padding*2 > bgWidth {
			bgWidth = w + padding*2
		}
	}
	// Calculate height based on number of lines and line height, plus padding
	bgHeight := float64(len(lines)*lineHeight) + (padding * 2.0) + 18.0 // add extra height for clarity

	// Draw background rectangle
	bgColor := color.RGBA{0, 0, 0, 180} // Semi-transparent black
	vector.DrawFilledRect(screen, float32(bgX), float32(bgY), float32(bgWidth), float32(bgHeight), bgColor, false)

	// Draw text lines using the game's font
	currentY := initialY
	for _, line := range lines {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(float64(textX), float64(currentY))
		opts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, line, NormalFont, opts)
		currentY += lineHeight
	}
}

// drawMenu renders a simple vertical list of lines with a background box.
func drawMenu(screen *ebiten.Image, lines []string, textX, initialY int) {
	if len(lines) == 0 {
		return
	}
	lineHeight := 18
	padding := 5.0

	bgX := float64(textX) - padding
	bgY := float64(initialY) - padding
	bgWidth := 320.0
	for _, line := range lines {
		w := float64(len(line)) * 13.0
		if w+padding*2 > bgWidth {
			bgWidth = w + padding*2
		}
	}
	bgHeight := float64(len(lines)*lineHeight) + (padding * 2.0) + 18.0
	vector.DrawFilledRect(screen, float32(bgX), float32(bgY), float32(bgWidth), float32(bgHeight), color.RGBA{0, 0, 0, 180}, false)

	currentY := initialY
	for _, line := range lines {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(float64(textX), float64(currentY))
		opts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, line, NormalFont, opts)
		currentY += lineHeight
	}
}

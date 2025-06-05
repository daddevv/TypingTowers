package game

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// HUD displays placeholder UI elements with basic game information.
type HUD struct {
	game *Game
}

// progressBar returns a simple ASCII progress bar of the given width.
func progressBar(progress float64, width int) string {
	if progress < 0 {
		progress = 0
	}
	if progress > 1 {
		progress = 1
	}
	filled := int(progress * float64(width))
	if filled > width {
		filled = width
	}
	return "[" + strings.Repeat("#", filled) + strings.Repeat("-", width-filled) + "]"
}

// NewHUD creates a new HUD bound to the given game.
func NewHUD(g *Game) *HUD {
	return &HUD{game: g}
}

// drawConveyorBelt renders a simple conveyor belt animation behind the queue.
// totalWidth is the total width of the queued words to ensure the belt spans
// the text. Slanted stripes move with the conveyor offset to give an illusion
// of motion.
func (h *HUD) drawConveyorBelt(screen *ebiten.Image, totalWidth float64) {
	beltHeight := 24.0
	beltY := h.game.wordProcessY - 18
	beltX := h.game.wordProcessX - h.game.conveyorOffset - 10
	beltW := totalWidth + 20

	// Draw base belt rectangle
	vector.DrawFilledRect(screen, float32(beltX), float32(beltY), float32(beltW), float32(beltHeight), color.RGBA{50, 50, 50, 180}, false)

	// Draw slanted stripes to indicate movement
	stripeSpacing := 12.0
	offset := math.Mod(h.game.conveyorOffset, stripeSpacing)
	for x := -offset; x < beltW; x += stripeSpacing {
		vector.StrokeLine(screen,
			float32(beltX+x), float32(beltY),
			float32(beltX+x+beltHeight/2), float32(beltY+beltHeight),
			1, color.RGBA{80, 80, 80, 200}, false)
	}
}

// drawResourceIcons renders resource amounts as letter icons at the top left.
func (h *HUD) drawResourceIcons(screen *ebiten.Image) {
	type icon struct {
		label  string
		amount int
		clr    color.RGBA
	}

	icons := []icon{
		{"G", h.game.resources.GoldAmount(), color.RGBA{255, 215, 0, 255}},
		{"W", h.game.resources.WoodAmount(), color.RGBA{139, 69, 19, 255}},
		{"S", h.game.resources.StoneAmount(), color.RGBA{128, 128, 128, 255}},
		{"I", h.game.resources.IronAmount(), color.RGBA{169, 169, 169, 255}},
		{"M", 0, color.RGBA{75, 0, 130, 255}},
	}

	size := 20.0
	x := 10.0
	y := 10.0

	for _, ic := range icons {
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(size), float32(size), ic.clr, false)

		opts := &text.DrawOptions{}
		opts.GeoM.Translate(x+4, y+4)
		opts.ColorScale.ScaleWithColor(color.Black)
		text.Draw(screen, ic.label, BoldFont, opts)

		numStr := strconv.Itoa(ic.amount)
		opts = &text.DrawOptions{}
		opts.GeoM.Translate(x+size+4, y+14)
		opts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, numStr, BoldFont, opts)

		x += size + float64(len(numStr))*13.0 + 16
	}
}

// drawQueue renders the global typing queue at the top center of the screen.
func (h *HUD) drawQueue(screen *ebiten.Image) {
	if h.game.queue == nil {
		return
	}
	words := h.game.queue.Words()
	if len(words) == 0 {
		return
	}

	spacing := 20.0
	total := 0.0
	for _, w := range words {
		total += float64(len(w.Text))*13.0 + spacing
	}
	total -= spacing
	h.drawConveyorBelt(screen, total)
	x := h.game.wordProcessX - h.game.conveyorOffset
	y := h.game.wordProcessY

	for _, w := range words {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(x, y)
		opts.ColorScale.ScaleWithColor(FamilyColor(w.Family))
		text.Draw(screen, w.Text, BoldFont, opts)
		x += float64(len(w.Text))*13.0 + spacing
	}

	if h.game.queueJam {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(x+10, y)
		opts.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
		text.Draw(screen, "[JAM]", BoldFont, opts)
	}
}

// drawTowerSelectionOverlay draws letter labels and highlight boxes over each
// tower when tower selection mode is active.
func (h *HUD) drawTowerSelectionOverlay(screen *ebiten.Image) {
	if !h.game.towerSelectMode {
		return
	}
	for label, idx := range h.game.towerLabels {
		if idx < 0 || idx >= len(h.game.towers) {
			continue
		}
		t := h.game.towers[idx]
		bx, by, bw, bh := t.Bounds()
		vector.StrokeRect(screen, float32(bx-4), float32(by-4), float32(bw+8), float32(bh+8), 2, color.RGBA{255, 255, 0, 200}, false)

		opts := &text.DrawOptions{}
		opts.GeoM.Translate(float64(bx)+float64(bw)/2-6, float64(by)-20)
		opts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, label, BoldFont, opts)
	}
}

// drawTechMenu renders the tech purchase overlay when active.
func (h *HUD) drawTechMenu(screen *ebiten.Image) {
	if !h.game.techMenuOpen {
		return
	}
	nodes := h.game.filteredTechNodes()
	lines := []string{"-- TECH --", "Search: " + h.game.searchBuffer}
	for i, n := range nodes {
		letters := strings.Builder{}
		for _, r := range n.Letters {
			letters.WriteRune(r)
		}
		line := fmt.Sprintf("%s [%s] - %s", n.Name, letters.String(), n.Achievement)
		prefix := "  "
		if i == h.game.techCursor {
			prefix = "> "
		}
		lines = append(lines, prefix+line)
	}
	drawMenu(screen, lines, 760, 300)
}

// drawSkillTreeOverlay renders the global skill tree when active.
func (h *HUD) drawSkillTreeOverlay(screen *ebiten.Image) {
	if !h.game.skillMenuOpen {
		return
	}
	nodes := h.game.skillMenuNodes()
	if nodes == nil {
		nodes = h.game.skillTree.NodesByCategory(h.game.skillCategory)
	}
	lines := []string{fmt.Sprintf("-- SKILLS: %s --", h.game.skillCategory.String())}
	for i, n := range nodes {
		status := "Locked"
		if h.game.unlockedSkills[n.ID] {
			status = "Unlocked"
		}
		prefix := "  "
		if i == h.game.skillCursor {
			prefix = "> "
		}
		lines = append(lines, fmt.Sprintf("%s%s - %s", prefix, n.Name, status))
	}
	drawMenu(screen, lines, 760, 300)
}

// Draw renders ammo count, tower stats, reload prompts, and shop interface.
func (h *HUD) Draw(screen *ebiten.Image) {
	h.drawResourceIcons(screen)
	h.drawQueue(screen)
	h.drawTowerSelectionOverlay(screen)
	h.drawTechMenu(screen)
	h.drawSkillTreeOverlay(screen)
	if h.game.commandMode {
		drawMenu(screen, []string{":" + h.game.commandBuffer}, 860, 1020)
		return
	}
	var lines []string
	textX := 10
	initialY := 30 // Start HUD lower to avoid overlap with mouse/tile debug info
	lineHeight := 18
	padding := 5.0

	if h.game.shopOpen {
		gold := h.game.Gold()
		lines = append(lines, "-- SHOP --")
		lines = append(lines, fmt.Sprintf("Gold: %d", gold))
		lines = append(lines, fmt.Sprintf("Food: %d", h.game.resources.FoodAmount()))
		lines = append(lines, fmt.Sprintf("KP: %d", h.game.resources.KingsAmount()))

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
			fmt.Sprintf("[6] Unlock Farmer Letters (%d KP)", h.game.farmer.NextUnlockCost()),
			fmt.Sprintf("[7] Unlock Barracks Letters (%d KP)", h.game.barracks.NextUnlockCost()),
			"Start Next Wave",
		}

		for i, opt := range options {
			prefix := "  "
			if i == h.game.shopCursor {
				prefix = "> "
			}
			lines = append(lines, prefix+opt)
		}
	} else if h.game.upgradeMenuOpen {
		lines = append(lines, "-- UPGRADE TOWER --")
		lines = append(lines, fmt.Sprintf("Gold: %d", h.game.Gold()))
		options := []string{
			"[1] Upgrade Damage (+1): 5 gold",
			"[2] Upgrade Range (+50): 5 gold",
			"[3] Upgrade Fire Rate (faster): 5 gold",
			"[4] Upgrade Ammo Capacity (+2): 10 gold",
			"[5] Foresight (+2 letters)",
			"Cancel",
		}
		for i, opt := range options {
			prefix := "  "
			if i == h.game.upgradeCursor {
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
		lines = append(lines, fmt.Sprintf("Gold: %d", h.game.Gold()))
		lines = append(lines, fmt.Sprintf("Food: %d", h.game.resources.FoodAmount()))
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
		lines = append(lines, fmt.Sprintf("Wave %d | Gold %d | Food %d | KP %d | Score %d | Mobs %d", h.game.currentWave, h.game.Gold(), h.game.resources.FoodAmount(), h.game.resources.KingsAmount(), h.game.score, len(h.game.mobs)))
		acc := h.game.typing.Accuracy() * 100
		wpm := h.game.typing.WPM()
		lines = append(lines, fmt.Sprintf("Accuracy: %.0f%% | WPM: %.1f", acc, wpm))
		if h.game.farmer != nil {
			prog := h.game.farmer.CooldownProgress()
			lines = append(lines, "Farmer "+progressBar(prog, 10))
		}
		if h.game.barracks != nil {
			prog := h.game.barracks.CooldownProgress()
			lines = append(lines, "Barracks "+progressBar(prog, 10))
		}
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

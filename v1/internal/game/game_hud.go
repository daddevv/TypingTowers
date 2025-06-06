package game

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/daddevv/type-defense/internal/assets"
	"github.com/daddevv/type-defense/internal/skill"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// HUD displays placeholder UI elements with basic game information.
type HUD struct {
	game *Game
}

// ProgressBar returns a simple ASCII progress bar of the given width.
func ProgressBar(progress float64, width int) string {
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
func NewHUD(game *Game) *HUD {
	return &HUD{game: game}
}

// NewHUDFromGame creates a new HUD from an existing game instance.
// func NewHUDFromGame(g *Game) *HUD {
// 	return &HUD{game: g}
// }

// DrawConveyorBelt renders a simple conveyor belt animation behind the queue.
// totalWidth is the total width of the queued words to ensure the belt spans
// the text. Slanted stripes move with the conveyor offset to give an illusion
// of motion.
func (h *HUD) DrawConveyorBelt(screen *ebiten.Image, totalWidth float64) {
	// beltHeight := 24.0
	// beltY := h.game.wordProcessY - 18
	// beltX := h.game.wordProcessX - h.game.conveyorOffset - 10
	// beltW := totalWidth + 20

	// // Draw base belt rectangle
	// vector.DrawFilledRect(screen, float32(beltX), float32(beltY), float32(beltW), float32(beltHeight), color.RGBA{50, 50, 50, 180}, false)

	// // Draw slanted stripes to indicate movement
	// stripeSpacing := 12.0
	// offset := math.Mod(h.game.conveyorOffset, stripeSpacing)
	// for x := -offset; x < beltW; x += stripeSpacing {
	// 	vector.StrokeLine(screen,
	// 		float32(beltX+x), float32(beltY),
	// 		float32(beltX+x+beltHeight/2), float32(beltY+beltHeight),
	// 		1, color.RGBA{80, 80, 80, 200}, false)
	// }
}

// DrawResourceIcons renders resource amounts as letter icons at the top left.
func (h *HUD) DrawResourceIcons(screen *ebiten.Image) {
	// type icon struct {
	// 	label  string
	// 	amount int
	// 	clr    color.RGBA
	// }

	// icons := []icon{
	// 	{"G", h.game.resources.GoldAmount(), color.RGBA{255, 215, 0, 255}},
	// 	{"W", h.game.resources.WoodAmount(), color.RGBA{139, 69, 19, 255}},
	// 	{"S", h.game.resources.StoneAmount(), color.RGBA{128, 128, 128, 255}},
	// 	{"I", h.game.resources.IronAmount(), color.RGBA{169, 169, 169, 255}},
	// 	{"M", 0, color.RGBA{75, 0, 130, 255}},
	// }

	// size := 20.0
	// x := 10.0
	// y := 10.0

	// for _, ic := range icons {
	// 	vector.DrawFilledRect(screen, float32(x), float32(y), float32(size), float32(size), ic.clr, false)

	// 	opts := &text.DrawOptions{}
	// 	opts.GeoM.Translate(x+4, y+4)
	// 	opts.ColorScale.ScaleWithColor(color.Black)
	// 	text.Draw(screen, ic.label, BoldFont, opts)

	// 	numStr := strconv.Itoa(ic.amount)
	// 	opts = &text.DrawOptions{}
	// 	opts.GeoM.Translate(x+size+4, y+14)
	// 	opts.ColorScale.ScaleWithColor(color.White)
	// 	text.Draw(screen, numStr, BoldFont, opts)

	// 	x += size + float64(len(numStr))*13.0 + 16
	// }
}

// drawQueue renders the global typing queue at the top center of the screen.
func (h *HUD) DrawQueue(screen *ebiten.Image) {
	// if h.game.queue == nil {
	// 	return
	// }
	// words := h.game.queue.Words()
	// if len(words) == 0 {
	// 	return
	// }

	// spacing := 20.0
	// total := 0.0
	// for _, w := range words {
	// 	total += float64(len(w.Text))*13.0 + spacing
	// }
	// total -= spacing
	// h.drawConveyorBelt(screen, total)
	// x := h.game.wordProcessX - h.game.conveyorOffset
	// y := h.game.wordProcessY

	// for i, w := range words {
	// 	opts := &text.DrawOptions{}
	// 	opts.GeoM.Translate(x, y)
	// 	if i == 0 {
	// 		typed := h.game.queue.Index()
	// 		if typed > 0 {
	// 			done := w.Text[:typed]
	// 			rem := w.Text[typed:]
	// 			opts.ColorScale.ScaleWithColor(color.RGBA{160, 160, 160, 255})
	// 			text.Draw(screen, done, BoldFont, opts)
	// 			tw := float64(len(done)) * 13.0
	// 			opts = &text.DrawOptions{}
	// 			opts.GeoM.Translate(x+tw, y)
	// 			opts.ColorScale.ScaleWithColor(FamilyColor(w.Family))
	// 			text.Draw(screen, rem, BoldFont, opts)
	// 			x += float64(len(w.Text))*13.0 + spacing
	// 			continue
	// 		}
	// 	}
	// 	opts.ColorScale.ScaleWithColor(FamilyColor(w.Family))
	// 	text.Draw(screen, w.Text, BoldFont, opts)
	// 	x += float64(len(w.Text))*13.0 + spacing
	// }

	// if h.game.queueJam {
	// 	opts := &text.DrawOptions{}
	// 	opts.GeoM.Translate(x+10, y)
	// 	opts.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
	// 	text.Draw(screen, "[JAM]", BoldFont, opts)
	// }
}

// DrawTowerSelectionOverlay draws letter labels and highlight boxes over each
// tower when tower selection mode is active.
func (h *HUD) DrawTowerSelectionOverlay(screen *ebiten.Image) {
	// if !h.game.towerSelectMode {
	// 	return
	// }
	// for label, idx := range h.game.towerLabels {
	// 	if idx < 0 || idx >= len(h.game.towers) {
	// 		continue
	// 	}
	// 	t := h.game.towers[idx]
	// 	bx, by, bw, bh := t.Bounds()
	// 	vector.StrokeRect(screen, float32(bx-4), float32(by-4), float32(bw+8), float32(bh+8), 2, color.RGBA{255, 255, 0, 200}, false)

	// 	opts := &text.DrawOptions{}
	// 	opts.GeoM.Translate(float64(bx)+float64(bw)/2-6, float64(by)-20)
	// 	opts.ColorScale.ScaleWithColor(color.White)
	// 	text.Draw(screen, label, BoldFont, opts)
	// }
}

// DrawTechMenu renders the tech purchase overlay when active.
func (h *HUD) DrawTechMenu(screen *ebiten.Image) {
	// if !h.game.techMenuOpen {
	// 	return
	// }
	// nodes := h.game.filteredTechNodes()
	// lines := []string{"-- TECH --", "Search: " + h.game.searchBuffer}
	// for i, n := range nodes {
	// 	letters := strings.Builder{}
	// 	for _, r := range n.Letters {
	// 		letters.WriteRune(r)
	// 	}
	// 	line := fmt.Sprintf("%s [%s] - %s", n.Name, letters.String(), n.Achievement)
	// 	prefix := "  "
	// 	if i == h.game.techCursor {
	// 		prefix = "> "
	// 	}
	// 	lines = append(lines, prefix+line)
	// }
	// drawMenu(screen, lines, 760, 300)
}

// DrawSkillMenu renders the global skill tree overlay when active.
func (h *HUD) DrawSkillMenu(screen *ebiten.Image) {
	// if !h.game.skillMenuOpen {
	// 	return
	// }
	// categories := []string{"Offense", "Defense", "Typing", "Automation", "Utility"}
	// cat := categories[h.game.skillCategory]
	// nodes := h.game.skillNodesByCategory(SkillCategory(h.game.skillCategory))
	// lines := []string{"-- SKILLS --", "Category: " + cat}
	// for i, n := range nodes {
	// 	nodeStatus := "Locked"
	// 	if h.game.unlockedSkills[n.ID] {
	// 		nodeStatus = "Unlocked"
	// 	}
	// 	prefix := "  "
	// 	if i == h.game.skillCursor {
	// 		prefix = "> "
	// 	}
	// 	lines = append(lines, fmt.Sprintf("%s%s - %s", prefix, n.Name, nodeStatus))
	// }
	// drawMenu(screen, lines, 760, 300)
}

// DrawSlotMenu renders the save/load slot selection overlay when active.
func (h *HUD) DrawSlotMenu(screen *ebiten.Image) {
	// if !h.game.slotMenuOpen {
	// 	return
	// }
	// title := "-- SAVE SLOT --"
	// if !h.game.slotModeSave {
	// 	title = "-- LOAD SLOT --"
	// }
	// lines := []string{title}
	// for i := 0; i < 3; i++ {
	// 	prefix := "  "
	// 	if i == h.game.slotCursor {
	// 		prefix = "> "
	// 	}
	// 	lines = append(lines, fmt.Sprintf("%sSlot %d", prefix, i+1))
	// }
	// drawMenu(screen, lines, 860, 480)
}

// DrawWordStats displays the last completed word's accuracy and time.
func (h *HUD) DrawWordStats(screen *ebiten.Image) {
	// hist := h.game.WordHistory()
	// if len(hist) == 0 {
	// 	return
	// }
	// stat := hist[len(hist)-1]
	// acc := 1.0
	// total := stat.Correct + stat.Incorrect
	// if total > 0 {
	// 	acc = float64(stat.Correct) / float64(total)
	// }
	// line := fmt.Sprintf("%s %.0f%% %.1fs", stat.Text, acc*100, stat.Duration.Seconds())
	// opts := &text.DrawOptions{}
	// opts.GeoM.Translate(10, 80)
	// opts.ColorScale.ScaleWithColor(color.White)
	// text.Draw(screen, line, BoldFont, opts)

	// wpmLine := fmt.Sprintf("WPM: %.1f", h.game.typing.RollingWPM())
	// opts = &text.DrawOptions{}
	// opts.GeoM.Translate(10, 100)
	// opts.ColorScale.ScaleWithColor(color.White)
	// text.Draw(screen, wpmLine, BoldFont, opts)
}

// DrawSkillTreeOverlay renders the global skill tree when active.
func (h *HUD) DrawSkillTreeOverlay(screen *ebiten.Image) {
	if !h.game.skillMenuOpen {
		return
	}
	nodes := h.game.skillNodesByCategory(skill.SkillCategory(h.game.skillCategory))
	if nodes == nil {
		return
	}
	lines := []string{fmt.Sprintf("-- SKILLS: %s --", skill.SkillCategory(h.game.skillCategory).String())}
	for _, n := range nodes {
		status := "Locked"
		if h.game.unlockedSkills[n.ID] {
			status = "Unlocked"
		}
		lines = append(lines, fmt.Sprintf("  %s - %s", n.Name, status))
	}
	if len(nodes) > 0 {
		sel := nodes[h.game.skillCursor]
		effs := []string{}
		for k, v := range sel.Effects {
			effs = append(effs, fmt.Sprintf("%s=%.1f", k, v))
		}
		lines = append(lines, "")
		lines = append(lines, strings.Join(effs, ", "))
	}
	DrawMenu(screen, lines, 720, 260)
}

// DrawMenu renders a vertical list of strings at (x, y) with spacing and shadow.
func DrawMenu(screen *ebiten.Image, lines []string, x, y int) {
	const lineSpacing = 32
	for i, line := range lines {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(float64(x), float64(y+i*lineSpacing))
		opts.ColorScale.ScaleWithColor(color.Black)
		opts.GeoM.Translate(2, 2)
		text.Draw(screen, line, assets.BoldFont, opts)
		opts = &text.DrawOptions{}
		opts.GeoM.Translate(float64(x), float64(y+i*lineSpacing))
		opts.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, line, assets.BoldFont, opts)
	}
}

// DrawStatsPanel renders a panel showing recent typing stats when open.
func (h *HUD) DrawStatsPanel(screen *ebiten.Image) {
	// if !h.game.statsPanelOpen {
	// 	return
	// }
	// lines := []string{"-- STATS --"}
	// lines = append(lines, fmt.Sprintf("WPM: %.1f", h.game.typing.RollingWPM()))
	// lines = append(lines, fmt.Sprintf("Accuracy: %.0f%%", h.game.typing.Accuracy()*100))
	// lines = append(lines, "")
	// hist := h.game.WordHistory()
	// start := len(hist) - 5
	// if start < 0 {
	// 	start = 0
	// }
	// for _, ws := range hist[start:] {
	// 	acc := 1.0
	// 	total := ws.Correct + ws.Incorrect
	// 	if total > 0 {
	// 		acc = float64(ws.Correct) / float64(total)
	// 	}
	// 	lines = append(lines, fmt.Sprintf("%s %.0f%% %.1fs", ws.Text, acc*100, ws.Duration.Seconds()))
	// }
	// drawMenu(screen, lines, 720, 480)
}

// Draw renders the HUD elements on screen
func (h *HUD) Draw(screen *ebiten.Image) {
	h.DrawResourceIcons(screen)
	h.DrawWordStats(screen)
	h.DrawQueue(screen)
	h.DrawTowerSelectionOverlay(screen)
	h.DrawTechMenu(screen)
	h.DrawSkillMenu(screen)
	h.DrawSlotMenu(screen)
	h.DrawStatsPanel(screen)
}

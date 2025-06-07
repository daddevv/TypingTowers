package game

// import (
// 	"image/color"

// 	"github.com/daddevv/type-defense/internal/assets"
// 	"github.com/daddevv/type-defense/internal/core"
// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/text/v2"
// )

// // MainMenu handles the title screen UI and behavior.
// type MainMenu struct {
// 	options        []string
// 	cursor         int
// 	animOffset     float64
// 	inSettings     bool
// 	settingsCursor int
// }

// // NewMainMenu creates a default MainMenu instance.
// func NewMainMenu() *MainMenu {
// 	return &MainMenu{options: []string{"Start Game", "Settings", "Quit"}}
// }

// // Update processes input for the main menu. Returns ebiten.Termination when
// // the user chooses to quit.
// func (m *MainMenu) Update(g *Game, dt float64) error {
// 	m.animOffset += dt * 30
// 	if m.inSettings {
// 		if g.input.Down() {
// 			m.settingsCursor = (m.settingsCursor + 1) % 2
// 		}
// 		if g.input.Up() {
// 			m.settingsCursor = (m.settingsCursor - 1 + 2) % 2
// 		}
// 		if g.input.Enter() {
// 			switch m.settingsCursor {
// 			case 0:
// 				g.settings.Mute = !g.settings.Mute
// 				if g.sound != nil {
// 					g.sound.ToggleMute()
// 				}
// 			case 1:
// 				m.inSettings = false
// 			}
// 		}
// 		return nil
// 	}
// 	if g.input.Down() {
// 		m.cursor = (m.cursor + 1) % len(m.options)
// 	}
// 	if g.input.Up() {
// 		m.cursor = (m.cursor - 1 + len(m.options)) % len(m.options)
// 	}
// 	if g.input.Enter() {
// 		switch m.cursor {
// 		case 0:
// 			g.phase = core.PhasePreGame
// 			g.preGame = NewPreGame()
// 			if g.sound != nil {
// 				g.sound.StopMusic()
// 				g.sound.PlayBeep()
// 			}
// 		case 1:
// 			m.inSettings = true
// 			m.settingsCursor = 0
// 		case 2:
// 			if g.sound != nil {
// 				g.sound.StopMusic()
// 			}
// 			return ebiten.Termination
// 		}
// 	}
// 	return nil
// }

// // Draw renders the menu to the given screen.
// func (m *MainMenu) Draw(g *Game, screen *ebiten.Image) {
// 	screen.Clear()
// 	drawScrollingBackground(screen, m.animOffset)
// 	titleOpts := &text.DrawOptions{}
// 	titleOpts.GeoM.Translate(760, 200)
// 	titleOpts.ColorScale.ScaleWithColor(color.White)
// 	text.Draw(screen, "TypingTowers", assets.BoldFont, titleOpts)

// 	var lines []string
// 	if m.inSettings {
// 		mute := "Off"
// 		if g.settings.Mute {
// 			mute = "On"
// 		}
// 		opts := []string{"Toggle Mute: " + mute, "Back"}
// 		for i, opt := range opts {
// 			prefix := "  "
// 			if i == m.settingsCursor {
// 				prefix = "> "
// 			}
// 			lines = append(lines, prefix+opt)
// 		}
// 		DrawMenu(screen, append([]string{"-- SETTINGS --"}, lines...), 860, 480)
// 		return
// 	}
// 	for i, opt := range m.options {
// 		prefix := "  "
// 		if i == m.cursor {
// 			prefix = "> "
// 		}
// 		lines = append(lines, prefix+opt)
// 	}
// 	DrawMenu(screen, append([]string{"-- MAIN MENU --"}, lines...), 860, 480)
// }

// // drawScrollingBackground renders a vertically scrolling background.
// func drawScrollingBackground(screen *ebiten.Image, offset float64) {
// 	if assets.ImgBackgroundBasicTiles == nil {
// 		return
// 	}
// 	h := assets.ImgBackgroundBasicTiles.Bounds().Dy()
// 	oy := int(offset) % h
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(0, float64(oy))
// 	screen.DrawImage(assets.ImgBackgroundBasicTiles, op)
// 	op2 := &ebiten.DrawImageOptions{}
// 	op2.GeoM.Translate(0, float64(oy-h))
// 	screen.DrawImage(assets.ImgBackgroundBasicTiles, op2)
// }

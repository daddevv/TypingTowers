package game

import (
	"unicode"

	"github.com/hajimehoshi/ebiten/v2"
)

// PreGame handles the pre-game setup flow including character selection,
// difficulty selection, a brief tutorial, a typing test and mode selection.
type PreGame struct {
	step int

	charOptions []string
	charCursor  int

	diffOptions []string
	diffCursor  int

	modeOptions []string
	modeCursor  int

	typed string
}

// NewPreGame returns a PreGame initialized with default options.
func NewPreGame() *PreGame {
	return &PreGame{
		charOptions: []string{"Knight", "Archer"},
		diffOptions: []string{"Easy", "Normal", "Hard"},
		modeOptions: []string{"Classic", "Endless"},
	}
}

// Update processes input for the pre-game flow.
func (p *PreGame) Update(g *Game, dt float64) error {
	switch p.step {
	case 0: // character selection
		if g.input.Down() {
			p.charCursor = (p.charCursor + 1) % len(p.charOptions)
		}
		if g.input.Up() {
			p.charCursor = (p.charCursor - 1 + len(p.charOptions)) % len(p.charOptions)
		}
		if g.input.Enter() {
			p.step = 1
		}
	case 1: // difficulty selection
		if g.input.Down() {
			p.diffCursor = (p.diffCursor + 1) % len(p.diffOptions)
		}
		if g.input.Up() {
			p.diffCursor = (p.diffCursor - 1 + len(p.diffOptions)) % len(p.diffOptions)
		}
		if g.input.Enter() {
			p.step = 2
		}
	case 2: // tutorial message
		if g.input.Enter() {
			p.step = 3
			p.typed = ""
		}
	case 3: // typing test
		word := "ready"
		for _, r := range g.input.TypedChars() {
			expected := rune(word[len(p.typed)])
			if unicode.ToLower(r) == expected {
				p.typed += string(expected)
				if len(p.typed) == len(word) {
					p.step = 4
				}
			} else if r != 0 {
				p.typed = ""
			}
		}
	case 4: // mode selection
		if g.input.Down() {
			p.modeCursor = (p.modeCursor + 1) % len(p.modeOptions)
		}
		if g.input.Up() {
			p.modeCursor = (p.modeCursor - 1 + len(p.modeOptions)) % len(p.modeOptions)
		}
		if g.input.Enter() {
			g.phase = PhasePlaying
			g.startWave()
			if g.sound != nil {
				g.sound.PlayBeep()
			}
		}
	}
	return nil
}

// Draw renders the pre-game setup UI.
func (p *PreGame) Draw(g *Game, screen *ebiten.Image) {
	var lines []string
	switch p.step {
	case 0:
		for i, opt := range p.charOptions {
			prefix := "  "
			if i == p.charCursor {
				prefix = "> "
			}
			lines = append(lines, prefix+opt)
		}
		drawMenu(screen, append([]string{"-- SELECT CHARACTER --"}, lines...), 860, 480)
	case 1:
		for i, opt := range p.diffOptions {
			prefix := "  "
			if i == p.diffCursor {
				prefix = "> "
			}
			lines = append(lines, prefix+opt)
		}
		drawMenu(screen, append([]string{"-- SELECT DIFFICULTY --"}, lines...), 860, 480)
	case 2:
		lines = []string{"-- TUTORIAL --", "Type words to attack and build.", "Press Enter to continue"}
		drawMenu(screen, lines, 860, 480)
	case 3:
		lines = []string{"-- TYPING TEST --", "Type: ready", p.typed}
		drawMenu(screen, lines, 860, 480)
	case 4:
		for i, opt := range p.modeOptions {
			prefix := "  "
			if i == p.modeCursor {
				prefix = "> "
			}
			lines = append(lines, prefix+opt)
		}
		drawMenu(screen, append([]string{"-- SELECT MODE --"}, lines...), 860, 480)
	}
}

package menu

import (
	"image/color"
	"os"
	"td/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	lua "github.com/yuin/gopher-lua"
)

func textWidth(face *text.GoTextFace, s string) float64 {
	return float64(len(s)) * face.Size * 0.6
}

type MainMenuOption int

const (
	StartGameOption MainMenuOption = iota
	OptionsOption
	QuitOption
)

type MainMenu struct {
	Options     []MainMenuOption
	ActiveOption int
	Selected    bool
	L           *lua.LState
}

// NewMainMenu creates a new MainMenu instance with default options and Lua overrides.
func NewMainMenu(L *lua.LState) *MainMenu {
	// Initialize the main menu with default options
	menu := &MainMenu{
		Options:  []MainMenuOption{StartGameOption, OptionsOption, QuitOption},
		ActiveOption: int(StartGameOption),
		Selected: false,
		L:        L,
	}
	// Lua Override: Check if the Lua state has a MainMenu table with options
	if tbl := L.GetGlobal("MainMenu"); tbl.Type() == lua.LTTable {
		options := tbl.(*lua.LTable).RawGetString("options")
		if optTbl, ok := options.(*lua.LTable); ok {
			menu.Options = nil
			optTbl.ForEach(func(_, v lua.LValue) {
				if entry, ok := v.(*lua.LTable); ok {
					label := entry.RawGetString("label").String()
					var opt MainMenuOption
					switch label {
					case "StartGame":
						opt = StartGameOption
					case "Options":
						opt = OptionsOption
					case "Quit":
						opt = QuitOption
					default:
						// Unknown label, skip or handle as needed
						return
					}
					menu.Options = append(menu.Options, opt)
				}
			})
		}
	}
	return menu
}

func (m *MainMenu) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		m.ActiveOption--
		if m.ActiveOption < 0 {
			m.ActiveOption = len(m.Options) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.ActiveOption++
		if m.ActiveOption >= len(m.Options) {
			m.ActiveOption = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		m.Selected = true
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}
	return nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	canvasWidth := 1920
	// Clear the screen
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw the game title
	gameTitle := "Type Defense"
	gameTitleFont := ui.Font("Game-Bold", 80)
	gameTitleWidth := textWidth(gameTitleFont, gameTitle)
	gameTitleOpts := &text.DrawOptions{}
	gameTitleOpts.GeoM.Translate(float64(canvasWidth)/2-gameTitleWidth/2, 100)
	gameTitleOpts.ColorScale.ScaleWithColor(color.RGBA{0, 200, 255, 255}) // Cyan/blue for game title
	text.Draw(screen, gameTitle, gameTitleFont, gameTitleOpts)

	// Draw the menu title (section label)
	title := "Main Menu"
	titleFont := ui.Font("Game-Bold", 48)
	titleWidth := textWidth(titleFont, title)
	titleOpts := &text.DrawOptions{}
	titleOpts.GeoM.Translate(float64(canvasWidth)/2-titleWidth/2, 200)
	titleOpts.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
	text.Draw(screen, title, titleFont, titleOpts)

	// Draw the menu options
	optionFont := ui.Font("Game-Regular", 48)
	optionYStart := 350.0
	optionSpacing := 80.0
	for i, opt := range m.Options {
		var optText string
		switch opt {
		case StartGameOption:
			optText = "Start Game"
		case OptionsOption:
			optText = "Options"
		case QuitOption:
			optText = "Quit"
		}
		optWidth := textWidth(optionFont, optText)
		optOpts := &text.DrawOptions{}
		optOpts.GeoM.Translate(float64(canvasWidth)/2-optWidth/2, optionYStart+float64(i)*optionSpacing)
		if i == m.ActiveOption {
			optOpts.ColorScale.ScaleWithColor(color.RGBA{255, 215, 0, 255}) // Gold/yellow for active
		} else {
			optOpts.ColorScale.ScaleWithColor(color.RGBA{200, 200, 200, 255}) // Light gray for inactive
		}
		text.Draw(screen, optText, optionFont, optOpts)
	}
}

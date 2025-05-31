package ui

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	lua "github.com/yuin/gopher-lua"
)

type MainMenuOption string

const (
	MainMenuStartGame MainMenuOption = "Start Game"
	MainMenuOptions   MainMenuOption = "Options"
	MainMenuQuit      MainMenuOption = "Quit"
)

func textWidth(face *text.GoTextFace, s string) float64 {
	return float64(len(s)) * face.Size * 0.6
}

type MainMenu struct {
	menu BaseMenu
	L    *lua.LState
}

// NewMainMenu creates a new MainMenu instance with default options and Lua overrides.
func NewMainMenu(L *lua.LState) *MainMenu {
	// Initialize the main menu with default options
	m := &MainMenu{
		menu: BaseMenu{
			screen:     NewBaseScreen(),
			options:    []string{
				string(MainMenuStartGame),
				string(MainMenuOptions),
				string(MainMenuQuit),
			},
			activeOption: 0,
		},
		L: L,
	}
	// Lua Override: Check if the Lua state has a MainMenu table with options
	if tbl := L.GetGlobal("MainMenu"); tbl.Type() == lua.LTTable {
		options := tbl.(*lua.LTable).RawGetString("options")
		if optTbl, ok := options.(*lua.LTable); ok {
			m.menu.options = nil
			optTbl.ForEach(func(_, v lua.LValue) {
				if entry, ok := v.(*lua.LTable); ok {
					label := entry.RawGetString("label").String()
					m.menu.options = append(m.menu.options, label)
				}
			})
		}
	}
	return m
}

func (m *MainMenu) Update() (*MainMenuOption, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		m.menu.activeOption--
		if m.menu.activeOption < 0 {
			m.menu.activeOption = len(m.menu.options) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.menu.activeOption++
		if m.menu.activeOption >= len(m.menu.options) {
			m.menu.activeOption = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
	switch m.menu.activeOption {
	case 0: // Start Game
		opt := MainMenuStartGame
		return &opt, nil
	case 1: // Options
		opt := MainMenuOptions
		return &opt, nil
	case 2: // Quit
		os.Exit(0) 
	}
}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}
	return nil, nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	canvasWidth := 1920
	// Clear the screen
	screen.Fill(color.RGBA{0, 0, 0, 255})

	// Draw the game title
	gameTitle := "Type Defense"
	gameTitleFont := Font("Game-Bold", 80)
	gameTitleWidth := textWidth(gameTitleFont, gameTitle)
	gameTitleOpts := &text.DrawOptions{}
	gameTitleOpts.GeoM.Translate(float64(canvasWidth)/2-gameTitleWidth/2, 100)
	gameTitleOpts.ColorScale.ScaleWithColor(color.RGBA{0, 200, 255, 255}) // Cyan/blue for game title
	text.Draw(screen, gameTitle, gameTitleFont, gameTitleOpts)

	// Draw the menu title (section label)
	title := "Main Menu"
	titleFont := Font("Game-Bold", 48)
	titleWidth := textWidth(titleFont, title)
	titleOpts := &text.DrawOptions{}
	titleOpts.GeoM.Translate(float64(canvasWidth)/2-titleWidth/2, 200)
	titleOpts.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
	text.Draw(screen, title, titleFont, titleOpts)

	// Draw the menu options
	optionFont := Font("Game-Regular", 48)
	optionYStart := 350.0
	optionSpacing := 80.0
	for i, opt := range m.menu.options {
		optWidth := textWidth(optionFont, opt)
		optOpts := &text.DrawOptions{}
		optOpts.GeoM.Translate(float64(canvasWidth)/2-optWidth/2, optionYStart+float64(i)*optionSpacing)
		if i == m.menu.activeOption {
			optOpts.ColorScale.ScaleWithColor(color.RGBA{255, 215, 0, 255}) // Gold/yellow for active
		} else {
			optOpts.ColorScale.ScaleWithColor(color.RGBA{200, 200, 200, 255}) // Light gray for inactive
		}
		text.Draw(screen, opt, optionFont, optOpts)
	}
}

// Selected returns whether the menu option is selected.
func (m *MainMenu) Selected() bool {
	return m.menu.isSelected
}

// SetSelected sets the selection state of the menu.
func (m *MainMenu) SetSelected(selected bool) {
	m.menu.isSelected = selected
}

// ActiveOption returns the currently active menu option index.
func (m *MainMenu) ActiveOption() int {
	return m.menu.activeOption
}

// Options returns the list of menu options.
func (m *MainMenu) Options() []string {
	return m.menu.options
}
package menu

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	lua "github.com/yuin/gopher-lua"
)

type MainMenuOption string

const (
	StartGameOption MainMenuOption = "Start Game"
	OptionsOption   MainMenuOption = "Options"
	QuitOption      MainMenuOption = "Quit"
)

type MainMenu struct {
	Options  []MainMenuOption
	Selected int
	L        *lua.LState
}

func NewMainMenu(L *lua.LState) *MainMenu {
	menu := &MainMenu{
		Options:  []MainMenuOption{StartGameOption, OptionsOption, QuitOption},
		Selected: 0,
		L:        L,
	}
	// If Lua table MainMenu exists, override options
	if tbl := L.GetGlobal("MainMenu"); tbl.Type() == lua.LTTable {
		options := tbl.(*lua.LTable).RawGetString("options")
		if optTbl, ok := options.(*lua.LTable); ok {
			menu.Options = nil
			optTbl.ForEach(func(_, v lua.LValue) {
				if entry, ok := v.(*lua.LTable); ok {
					label := entry.RawGetString("label").String()
					menu.Options = append(menu.Options, MainMenuOption(label))
				}
			})
		}
	}
	return menu
}

func (m *MainMenu) Update() (string, error) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		fmt.Println("Escape pressed, exiting...")
		os.Exit(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		m.Selected--
		if m.Selected < 0 {
			m.Selected = len(m.Options) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		m.Selected++
		if m.Selected >= len(m.Options) {
			m.Selected = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		switch m.Selected {
		case 0:
			return "Start Game", nil
		case 1:
			return "Options", nil
		case 2:
			return "Quit", nil
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}
	return "", nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	canvasWidth := 1920
	// Clear the screen
	screen.Fill(color.RGBA{0, 0, 0, 255})
	// Draw the menu title
	title := "Main Menu"
	titleWidth := len(title) * 10 // Assuming each character is 10 pixels wide
	titleX := (canvasWidth - titleWidth) / 2
	titleY := 50
	ebitenutil.DebugPrintAt(screen, title, titleX, titleY)

	// Draw the menu options
	optionY := 100
	for i, option := range m.Options {
		if i == m.Selected {
			optionX := (canvasWidth - len(string(option))*10 - 30) / 2
			// Highlight the selected option
			ebitenutil.DebugPrintAt(screen, "->"+string(option), optionX, optionY)
		} else {
			optionX := (canvasWidth - len(string(option))*10) / 2
			ebitenutil.DebugPrintAt(screen, string(option), optionX, optionY)
		}
		optionY += 30 // Move down for the next option
	}
}

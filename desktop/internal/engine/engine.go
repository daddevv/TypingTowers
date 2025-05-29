package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"td/internal/game"
	"td/internal/menu"
	"td/internal/state"
	"td/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	lua "github.com/yuin/gopher-lua"
)

type Engine struct {
	WindowWidth   int
	WindowHeight  int
	Version string
	State   state.EngineState // enum: current screen to display
	Screen  *ebiten.Image // Internal buffer for drawing at 1920x1080
	Menu    *menu.MainMenu
	Lobby   *menu.LobbyMenu
	Game    *game.Game
	L      *lua.LState
}

func NewEngine(width, height int, version string) *Engine {
	// Load all content configs at engine startup
	err := LoadContentConfigs()
	if err != nil {
		panic("Failed to load content configs: " + err.Error())
	}

	// Initialize Lua state
	l := lua.NewState()
	l.SetGlobal("example", l.NewFunction(func(L *lua.LState) int {
		fmt.Println("'example' called with parameters:", L.ToString(1), L.ToNumber(2), L.ToNumber(3))
		return 0
	}))

	m := menu.NewMainMenu(l)
	return &Engine{
		Version:      version,
		WindowHeight: 1080, // Always use 1920x1080 for internal canvas
		WindowWidth:  1920,
		Screen:      ebiten.NewImage(1920, 1080),
		State:       state.MAIN_MENU,
		Menu:       m,
		Lobby:      nil,
		Game:       nil,
		L:         l,
	}
}

func (e *Engine) Update() error {
	switch e.State {
	case state.MAIN_MENU:
		currentState, err := e.Menu.Update()
		if err != nil {
			return err
		}
		switch currentState {
		case "Start Game":
			if e.Game == nil {
				// Load first level from content config
				levelCfg := LoadedLevels[0]
				// Load world background
				var bgImg *ebiten.Image
				for _, w := range LoadedWorlds {
					if w.Name == levelCfg.World {
						img, _, err := ebitenutil.NewImageFromFile(w.Background)
						if err == nil {
							bgImg = img
						}
					}
				}
				// Convert waves
				waves := make([]struct {
					WaveNumber      int
					PossibleLetters []string
					EnemyCount      int
					MobChances      []struct {
						Type   string
						Chance float64
					}
				}, len(levelCfg.Waves))
				for i, w := range levelCfg.Waves {
					mobChances := make([]struct {
						Type   string
						Chance float64
					}, len(w.MobChances))
					for j, m := range w.MobChances {
						mobChances[j] = struct {
							Type   string
							Chance float64
						}{m.Type, m.Chance}
					}
					waves[i] = struct {
						WaveNumber      int
						PossibleLetters []string
						EnemyCount      int
						MobChances      []struct {
							Type   string
							Chance float64
						}
					}{
						WaveNumber:      w.WaveNumber,
						PossibleLetters: w.PossibleLetters,
						EnemyCount:      w.EnemyCount,
						MobChances:      mobChances,
					}
				}
				level := world.Level{
					Name:               levelCfg.Name,
					WorldNumber:        levelCfg.WorldNumber,
					LevelNumber:        levelCfg.LevelNumber,
					World:              levelCfg.World,
					StartingLetters:    levelCfg.StartingLetters,
					Waves:              waves,
					LevelCompleteScore: levelCfg.LevelCompleteScore,
					Background:         bgImg,
				}
				e.Game = game.NewGame(game.GameOptions{
					Level:      level,
					GameMode:   game.ENDLESS,
					MobConfigs: LoadedMobs,
				})
			}
			e.State = state.GAME_PLAYING
		case "Quit":
			os.Exit(0)
		}
		return nil
	case state.GAME_PLAYING:
		err := e.Game.Update()
		if err != nil {
			if err.Error() == "pause" {
				e.State = state.MAIN_MENU
			} else {
				return err
			}
		}
		return nil
	default:
		// Handle default logic
	}

	return nil
}

func (g *Engine) Draw(screen *ebiten.Image) {
	// Draw everything to the internal 1920x1080 screen
	g.Screen.Clear()
	switch g.State {
	case state.MAIN_MENU:
		g.Menu.Draw(g.Screen)
	case state.GAME_PLAYING:
		g.Game.Draw(g.Screen)
	default:
		// Handle default drawing
	}
	// Now scale the internal screen to the window size
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()
	scaleX := float64(w) / 1920.0
	scaleY := float64(h) / 1080.0
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	screen.DrawImage(g.Screen, opts)
}

func (g *Engine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// Always use 1920x1080 for the internal canvas
	return 1920, 1080
}

func (e *Engine) loadLuaPlugins(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".lua" {
			script, err := os.ReadFile(filepath.Join(dir, f.Name()))
			if err == nil {
				e.L.DoString(string(script))
			}
		}
	}
}

func (e *Engine) GoToMainMenu() {
	m := &menu.MainMenu{
		Options:  []menu.MainMenuOption{menu.StartGameOption, menu.OptionsOption, menu.QuitOption},
		Selected: 0,
		L:       e.L,
	}
	// If Lua table MainMenu exists, override options
	if tbl := e.L.GetGlobal("MainMenu"); tbl.Type() == lua.LTTable {
		options := tbl.(*lua.LTable).RawGetString("options")
		if optTbl, ok := options.(*lua.LTable); ok {
			m.Options = nil
			optTbl.ForEach(func(_, v lua.LValue) {
				if entry, ok := v.(*lua.LTable); ok {
					label := entry.RawGetString("label").String()
					m.Options = append(m.Options, menu.MainMenuOption(label))
				}
			})
		}
	}
	// Reset game state
	if e.Game != nil {
		e.Game = nil
	}
	// Reset lobby state
	if e.Lobby != nil {
		e.Lobby = nil
	}
	// Set the engine state to MAIN_MENU
	e.State = state.MAIN_MENU
	e.Menu = m
}
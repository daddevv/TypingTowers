package game

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// RegisterGameAPI exposes engine/game functions to Lua scripts.
func RegisterGameAPI(L *lua.LState, g *Game) {
	L.SetGlobal("spawn_mob", L.NewFunction(func(L *lua.LState) int {
		// mobType := L.ToString(1)
		// x := L.ToNumber(2)
		// y := L.ToNumber(3)
		fmt.Println("spawn_mob called with parameters:", L.ToString(1), L.ToNumber(2), L.ToNumber(3))
		// Here you would implement the logic to spawn a mob in the game.
		// Example: spawn a mob of type at (x, y)
		// You'd need to implement this logic
		// g.SpawnMob(mobType, float64(x), float64(y))
		return 0
	}))
	L.SetGlobal("get_player_health", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(g.Player.Health))
		return 1
	}))
	L.SetGlobal("call_hello", L.NewFunction(func(L *lua.LState) int {
		name := L.ToString(1)
		// Call the Lua function 'say_hello' defined in hello.lua
		ret, err := g.CallLuaFunction("say_hello", lua.LString(name))
		if err != nil {
			L.Push(lua.LString("error: " + err.Error()))
		} else {
			L.Push(ret)
		}
		return 1
	}))
	// Add more functions as needed
}

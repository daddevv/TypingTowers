package game

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// RegisterGameAPI exposes engine/game functions to Lua scripts.
func RegisterGameAPI(L *lua.LState) {
	L.SetGlobal("action_spawn_mob", L.NewFunction(LuaActionSpawnMob))
	L.SetGlobal("get_window_size", L.NewFunction(LuaGetWindowSize))
}

func LuaGetWindowSize(L *lua.LState) int {
	width, height := 1920, 1080 // Replace with actual window width retrieval logic
	L.Push(lua.LNumber(width))
	L.Push(lua.LNumber(height))
	
	return 2 // Return 2 to indicate two return values (width and height)
}

func LuaActionSpawnMob(L *lua.LState) int {
	mobType := L.ToString(1)
	x := L.ToNumber(2)
	y := L.ToNumber(3)
	fmt.Printf("Spawning mob of type %s at (%f, %f)\n", mobType, x, y)
	// Here you would add the logic to spawn the mob in the game world

	return 0 // Return 0 to indicate no return values
}
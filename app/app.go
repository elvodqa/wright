package app

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Running bool = false
var WorkingGame Game = Game{}

func Run(game Game) {
	WorkingGame = game
	// TODO: Check if debug, then also show Debug info etc.
	rl.SetTraceLog(rl.LogError)
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagWindowResizable)
	rl.InitWindow(1280, 720, game.Name)
	rl.SetTargetFPS(60)
	rl.SetExitKey(0)
	Running = true
	mainLoop()
}

func mainLoop() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
	Running = false
	rl.CloseWindow()
}

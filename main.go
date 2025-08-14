package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1000, 800, "FlipFlip")
	defer rl.WindowShouldClose()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		rl.DrawRectangle(300, 300, 250, 100, rl.Black)
		rl.DrawCircle(300+25+10, 300+100+25+10, 25, rl.Black)
		rl.DrawCircle(300+250-25-10, 300+100+25+10, 25, rl.Black)

		rl.EndDrawing()
	}
}

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Car struct {
	position rl.Vector2
	width    float32
	height   float32
}

type Wheel struct {
	position rl.Vector2
	radius   float32
	padding  float32
}

func main() {
	rl.InitWindow(1000, 800, "FlipFlip")
	defer rl.WindowShouldClose()

	rl.SetTargetFPS(60)

	car := Car{
		position: rl.Vector2{X: 300, Y: 300},
		width:    250,
		height:   100,
	}

	back_wheel := Wheel{
		radius:  25,
		padding: 10,
	}
	back_wheel.position = rl.Vector2{
		X: car.position.X + car.width - back_wheel.radius - back_wheel.padding,
		Y: car.position.Y + car.height + back_wheel.radius + back_wheel.padding,
	}

	front_wheel := Wheel{
		radius:  25,
		padding: 10,
	}
	front_wheel.position = rl.Vector2{
		X: car.position.X + front_wheel.radius + front_wheel.padding,
		Y: car.position.Y + car.height + front_wheel.radius + front_wheel.padding,
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		rl.DrawRectangle(int32(car.position.X), int32(car.position.Y), int32(car.width), int32(car.height), rl.Black)
		rl.DrawCircle(int32(back_wheel.position.X), int32(back_wheel.position.Y), back_wheel.radius, rl.Black)
		rl.DrawCircle(int32(front_wheel.position.X), int32(front_wheel.position.Y), front_wheel.radius, rl.Black)

		rl.EndDrawing()
	}
}

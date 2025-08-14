package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const GRAVITY float32 = 20

type Car struct {
	position rl.Vector2
	velocity rl.Vector2
	width    float32
	height   float32
}

func car_move(car *Car, dt float32) {
	car.position.Y += car.velocity.Y
	car.position.X += car.velocity.X

	var max_y float32 = float32(rl.GetScreenHeight()) - car.height

	if car.position.Y >= max_y {
		car.velocity.Y = 0
		car.position.Y = max_y
	} else {
		car.velocity.Y += GRAVITY * dt
	}
}

type Wheel struct {
	position rl.Vector2
	velocity rl.Vector2
	radius   float32
	padding  float32
}

func wheel_move(wheel *Wheel, dt float32) {
	wheel.velocity.Y += GRAVITY * dt
	wheel.position.X += wheel.velocity.X
	wheel.position.Y += wheel.velocity.Y

	var max_y float32 = float32(rl.GetScreenHeight()) - wheel.radius

	if wheel.position.Y >= max_y {
		wheel.velocity.Y = 0
		wheel.position.Y = max_y
	} else {
		wheel.velocity.Y += GRAVITY * dt
	}
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

		var dt float32 = rl.GetFrameTime()

		car_move(&car, dt)
		wheel_move(&back_wheel, dt)
		wheel_move(&front_wheel, dt)

		rl.DrawRectangle(int32(car.position.X), int32(car.position.Y), int32(car.width), int32(car.height), rl.Black)
		rl.DrawCircle(int32(back_wheel.position.X), int32(back_wheel.position.Y), back_wheel.radius, rl.Black)
		rl.DrawCircle(int32(front_wheel.position.X), int32(front_wheel.position.Y), front_wheel.radius, rl.Black)

		rl.EndDrawing()
	}
}

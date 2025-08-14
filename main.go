package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1000, 800, "FlipFlip")
	defer rl.WindowShouldClose()

	rl.SetTargetFPS(60)

	car := Car{
		position: rl.Vector2{X: 300, Y: 300},
		width:    250,
		height:   100,

		back_wheel: Wheel{
			radius:    25,
			padding:   10,
			stiffness: 1.8,
			damping:   1.2,
		},

		front_wheel: Wheel{
			radius:    25,
			padding:   10,
			stiffness: 1.8,
			damping:   1.2,
		},
	}

	car.back_wheel.position = rl.Vector2{
		X: car.position.X - car.width/2 + car.back_wheel.radius + car.back_wheel.padding,
		Y: car.position.Y + car.height/2 + car.back_wheel.radius + car.back_wheel.padding,
	}

	car.front_wheel.position = rl.Vector2{
		X: car.position.X + car.width/2 - car.front_wheel.radius - car.front_wheel.padding,
		Y: car.position.Y + car.height/2 + car.front_wheel.radius + car.front_wheel.padding,
	}
	car.front_wheel.offset = car.width - car.back_wheel.radius - car.back_wheel.padding - car.front_wheel.padding - car.front_wheel.radius

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)

		var dt float32 = rl.GetFrameTime()

		car_move(&car, dt)
		wheel_move(&car.back_wheel, dt)
		wheel_move(&car.front_wheel, dt)

		car_apply_suspension(&car, &car.back_wheel, dt)
		car_apply_suspension(&car, &car.front_wheel, dt)

		car_control(&car, dt)
		car_draw(&car)

		rl.EndDrawing()
	}
}

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "FlipFlip")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	camera := rl.Camera2D{
		Offset: rl.NewVector2(WINDOW_WIDTH/2, WINDOW_HEIGHT/2),
		Target: rl.NewVector2(0, 0),
		Zoom:   1,
	}

	car := Car{
		position: rl.NewVector2(1200, 300),
		width:    250,
		height:   100,
		angle:    0,
		back_wheel: Wheel{
			radius:    25,
			padding:   10,
			stiffness: 0.8,
			damping:   0.6,
		},
		front_wheel: Wheel{
			radius:    25,
			padding:   10,
			stiffness: 0.8,
			damping:   0.6,
		},
	}

	car.back_wheel.position = rl.NewVector2(
		car.position.X-car.width/2+car.back_wheel.radius+car.back_wheel.padding,
		car.position.Y+car.height/2+car.back_wheel.radius+car.back_wheel.padding,
	)
	car.front_wheel.offset = car.width - car.back_wheel.radius - car.back_wheel.padding - car.front_wheel.padding - car.front_wheel.radius
	car.front_wheel.position = rl.NewVector2(
		car.position.X+car.width/2-car.front_wheel.radius-car.front_wheel.padding,
		car.position.Y+car.height/2+car.front_wheel.radius+car.front_wheel.padding,
	)

	terrain_length := 100
	terrain_count := 255
	terrain := make([]rl.Vector2, terrain_count)

	pos := float32(rl.GetRandomValue(int32(WINDOW_HEIGHT*0.7), int32(WINDOW_HEIGHT*0.95)))
	for i := 0; i < terrain_count; i++ {
		movement := float32(rl.GetRandomValue(-20, 20))
		terrain[i] = rl.NewVector2(float32(i*terrain_length), pos)
		pos += movement
	}

	for !rl.WindowShouldClose() {
		dt := rl.GetFrameTime()
		camera.Target = car.position
		zoom := 1.3 - car.velocity.X/10
		if zoom > MAX_ZOOM {
			zoom = MAX_ZOOM
		}
		if zoom < MIN_ZOOM {
			zoom = MIN_ZOOM
		}
		camera.Zoom = zoom

		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.BeginMode2D(camera)

		for i := 1; i < len(terrain); i++ {
			rl.DrawLineEx(terrain[i-1], terrain[i], 5, rl.Black)
		}

		car_control(&car, dt)
		car_move(&car, terrain, terrain_length, dt)
		car_rotate(&car, dt)
		wheel_move(&car.back_wheel, terrain, dt)
		wheel_move(&car.front_wheel, terrain, dt)
		car_apply_suspension(&car, &car.back_wheel, dt)
		car_apply_suspension(&car, &car.front_wheel, dt)
		car_draw(&car)
		update_terrain(&terrain, terrain_length)

		rl.EndMode2D()
		rl.EndDrawing()
	}
}

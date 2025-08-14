package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Car struct {
	position rl.Vector2
	velocity rl.Vector2
	width    float32
	height   float32
	angle    float32

	back_wheel  Wheel
	front_wheel Wheel
}

func car_draw(car *Car) {
	rl.DrawRectanglePro(rl.Rectangle{
		X:      car.position.X,
		Y:      car.position.Y,
		Width:  car.width,
		Height: car.height,
	}, rl.Vector2{
		X: car.width / 2,
		Y: car.height / 2,
	}, car.angle, TRANSPARENT_BLACK)

	rl.DrawCircle(int32(car.back_wheel.position.X), int32(car.back_wheel.position.Y), car.back_wheel.radius, TRANSPARENT_BLACK)
	rl.DrawCircle(int32(car.front_wheel.position.X), int32(car.front_wheel.position.Y), car.front_wheel.radius, TRANSPARENT_BLACK)
}

func car_control(car *Car, dt float32) {
	if rl.IsKeyDown(rl.KeyLeft) {
		car.angle += ROTATION_SPEED * dt
	} else if rl.IsKeyDown(rl.KeyRight) {
		car.angle -= ROTATION_SPEED * dt
	}
}

func car_move(car *Car, dt float32) {
	car.position.Y += car.velocity.Y
	car.position.X += car.velocity.X

	var max_y float32 = float32(rl.GetScreenHeight()) - car.height/2

	if car.position.Y >= max_y {
		car.velocity.Y = 0
		car.position.Y = max_y
	} else {
		car.velocity.Y += GRAVITY * dt
	}
}

func car_apply_suspension(car *Car, wheel *Wheel, dt float32) {
	bottom_direction := rl.Vector2Rotate(rl.Vector2{0, 1}, car.angle*rl.Deg2rad)
	var attachment_point rl.Vector2 = rl.Vector2Rotate(rl.Vector2{
		X: -car.width/2 + wheel.padding + wheel.radius + wheel.offset,
		Y: 0,
	},
		car.angle*rl.Deg2rad,
	)

	attachment_point = rl.Vector2Add(attachment_point, car.position)

	rl.DrawCircleV(attachment_point, 10, rl.Green)

	length := rl.Vector2Distance(wheel.position, attachment_point)
	resting_length := car.height/2 + wheel.padding + wheel.radius
	stretch := length - resting_length

	wheel.position = rl.Vector2Add(attachment_point, rl.Vector2Scale(bottom_direction, length))

	spring_force := stretch * wheel.stiffness * dt
	relative_velocity := rl.Vector2Subtract(car.velocity, wheel.velocity)
	damping_force := rl.Vector2Scale(relative_velocity, wheel.damping*dt)

	force := rl.Vector2Subtract(rl.Vector2Scale(bottom_direction, spring_force), damping_force)

	car.velocity = rl.Vector2Add(car.velocity, force)
	wheel.velocity = rl.Vector2Subtract(wheel.velocity, rl.Vector2Scale(force, 0.7))
}

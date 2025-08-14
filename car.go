package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func car_apply_suspension(car *Car, wheel *Wheel, dt float32) {
	attachment_point := car.position.Y + car.height/2
	length := wheel.position.Y - attachment_point
	resting_length := car.height/2 + wheel.padding + wheel.radius
	stretch := length - resting_length

	spring_force := stretch * wheel.stiffness * dt
	relative_velocity := car.velocity.Y - wheel.velocity.Y
	damping_force := relative_velocity * wheel.damping * dt

	car.velocity.Y += spring_force - damping_force
	wheel.velocity.Y -= spring_force - damping_force
}

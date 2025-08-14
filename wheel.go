package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Wheel struct {
	position  rl.Vector2
	velocity  rl.Vector2
	radius    float32
	padding   float32
	stiffness float32
	damping   float32
	offset    float32
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

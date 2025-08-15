package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Wheel struct {
	position  rl.Vector2
	velocity  rl.Vector2
	radius    float32
	padding   float32
	stiffness float32
	damping   float32
	offset    float32
	on_ground bool
	texture   rl.Texture2D
}

func wheel_move(wheel *Wheel, terrain []rl.Vector2, dt float32) {
	wheel.position.X += wheel.velocity.X
	wheel.position.Y += wheel.velocity.Y
	wheel.on_ground = false

	for i := 1; i < len(terrain); i++ {
		point1 := terrain[i-1]
		point2 := terrain[i]
		var collision rl.Vector2
		bottom := rl.NewVector2(wheel.position.X, wheel.position.Y+wheel.radius)
		if is_point_below_line(point1, point2, bottom, &collision) {
			wheel.velocity.Y = 0
			wheel.position.Y = collision.Y - wheel.radius + 1
			wheel.on_ground = true
		}
	}

	if !wheel.on_ground {
		wheel.velocity.Y += GRAVITY * dt
	}
}

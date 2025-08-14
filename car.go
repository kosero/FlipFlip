package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Car struct {
	position    rl.Vector2
	velocity    rl.Vector2
	width       float32
	height      float32
	angle       float32
	back_wheel  Wheel
	front_wheel Wheel
}

func car_control(car *Car, dt float32) {
	if !car.back_wheel.on_ground && !car.front_wheel.on_ground {
		if rl.IsKeyDown(rl.KeyLeft) {
			car.angle += ROTATION_SPEED * dt
		} else if rl.IsKeyDown(rl.KeyRight) {
			car.angle -= ROTATION_SPEED * dt
		}
	}

	if rl.IsKeyDown(rl.KeyRight) {
		if car.back_wheel.on_ground {
			car.velocity.X += CAR_SPEED * dt
		}
		if car.front_wheel.on_ground {
			car.velocity.X += CAR_SPEED * dt
		}
	} else if rl.IsKeyDown(rl.KeyLeft) {
		if car.back_wheel.on_ground {
			car.velocity.X -= CAR_SPEED * dt
		}
		if car.front_wheel.on_ground {
			car.velocity.X -= CAR_SPEED * dt
		}
	}
}

func car_rotate(car *Car, dt float32) {
	angle := rl.Vector2LineAngle(car.back_wheel.position, car.front_wheel.position) * rl.Rad2deg
	diff := angle - car.angle
	car.angle += diff * RORTATE_BACK_SPEED * dt
}

func car_move(car *Car, terrain []rl.Vector2, terrain_length int, dt float32) {
	car.position.X += car.velocity.X
	car.position.Y += car.velocity.Y

	for _, wheel := range []*Wheel{&car.back_wheel, &car.front_wheel} {
		if wheel.on_ground {
			index := int(math.Floor(float64(wheel.position.X) / float64(terrain_length)))
			if index >= len(terrain)-1 {
				index = len(terrain) - 2
			}
			point1 := terrain[index]
			point2 := terrain[index+1]

			angle := rl.Vector2LineAngle(point1, point2) * rl.Rad2deg
			car.velocity.X += angle * HILL_SPEED * dt
			car.velocity.X -= car.velocity.X * FRICTION * dt
		}
	}

	max_y := float32(WINDOW_HEIGHT) - car.height/2
	if car.position.Y >= max_y {
		car.velocity.Y = 0
		car.position.Y = max_y
	} else {
		car.velocity.Y += GRAVITY * dt
	}
}

func car_draw(car *Car) {
	rl.DrawRectanglePro(
		rl.Rectangle{
			X:      car.position.X,
			Y:      car.position.Y,
			Width:  car.width,
			Height: car.height,
		},
		rl.Vector2{X: car.width / 2, Y: car.height / 2},
		car.angle,
		TRANSPARENT_BLACK,
	)
	rl.DrawCircleV(car.back_wheel.position, car.back_wheel.radius, TRANSPARENT_BLACK)
	rl.DrawCircleV(car.front_wheel.position, car.front_wheel.radius, TRANSPARENT_BLACK)
}

func car_apply_suspension(car *Car, wheel *Wheel, dt float32) {
	bottom_direction := rl.Vector2Rotate(rl.NewVector2(0, 1), car.angle*rl.Deg2rad)
	attachment := rl.Vector2Rotate(
		rl.NewVector2(-car.width/2+wheel.padding+wheel.radius+wheel.offset, 0),
		car.angle*rl.Deg2rad,
	)
	attachment = rl.Vector2Add(attachment, car.position)

	length := rl.Vector2Distance(wheel.position, attachment)
	rest_length := car.height/2 + wheel.padding + wheel.radius
	stretch := length - rest_length

	wheel.position = rl.Vector2Add(attachment, rl.Vector2Scale(bottom_direction, length))
	spring_force := stretch * wheel.stiffness * dt
	rel_vel := rl.Vector2Subtract(car.velocity, wheel.velocity)
	damping_force := rl.Vector2Scale(rel_vel, wheel.damping*dt)
	force := rl.Vector2Subtract(rl.Vector2Scale(bottom_direction, spring_force), damping_force)

	car.velocity = rl.Vector2Add(car.velocity, force)
	wheel.velocity = rl.Vector2Subtract(wheel.velocity, rl.Vector2Scale(force, 0.7))
}

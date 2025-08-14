package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func is_point_below_line(a, b, point rl.Vector2, collision_point *rl.Vector2) bool {
	if math.Abs(float64(b.X-a.X)) > 1e-6 {
		t := (point.X - a.X) / (b.X - a.X)
		if t < 0.0 || t > 1.0 {
			return false
		}
		collision_point.X = point.X
		collision_point.Y = a.Y + t*(b.Y-a.Y)
		return point.Y > collision_point.Y
	} else {
		if math.Abs(float64(point.X-a.X)) > 1e-6 {
			return false
		}
		collision_point.X = a.X
		collision_point.Y = float32(math.Max(float64(a.Y), float64(b.Y)))
		return point.Y > collision_point.Y
	}
}

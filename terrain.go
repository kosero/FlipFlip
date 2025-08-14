package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func update_terrain(terrain *[]rl.Vector2, terrain_length int) {
	last := (*terrain)[len(*terrain)-1]
	movement := float32(rl.GetRandomValue(-20, 20))
	new_point := rl.NewVector2(last.X+float32(terrain_length), last.Y+movement)
	*terrain = append(*terrain, new_point)

	for len(*terrain) > 2 && (*terrain)[1].X < 0 {
		*terrain = (*terrain)[1:]
	}
}

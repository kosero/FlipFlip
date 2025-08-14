package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	WINDOW_WIDTH       = 1500
	WINDOW_HEIGHT      = 1000
	GRAVITY            = 7.0
	FRICTION           = 0.2
	ROTATION_SPEED     = 30.0
	RORTATE_BACK_SPEED = 3.0
	CAR_SPEED          = 3.0
	HILL_SPEED         = 0.2

	MIN_ZOOM = 0.8
	MAX_ZOOM = 1.3
)

var (
	TRANSPARENT_BLACK = rl.NewColor(0, 0, 0, 100)
)

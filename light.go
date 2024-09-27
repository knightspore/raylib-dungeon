package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PointLight struct {
	Pos    rl.Vector2
	Color  rl.Color
	Radius float32
}

func NewPointLight(x, y, radius float32, color rl.Color) *PointLight {
	return &PointLight{rl.NewVector2(x, y), color, radius}
}

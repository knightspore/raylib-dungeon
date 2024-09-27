package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PointLight struct {
	Pos    rl.Vector2
	Radius float32
	Color  rl.Color
}

func NewPointLight(x, y, radius float32, color rl.Color) *PointLight {
	return &PointLight{rl.NewVector2(x, y), radius, color}
}

func (pl *PointLight) Draw(g *Game) {
	rl.DrawCircleV(pl.Pos, pl.Radius, pl.Color)
}

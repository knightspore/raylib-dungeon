package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PointLight struct {
	pos    rl.Vector2
	colour rl.Color
	radius float32
	sprite *Sprite
}

func NewLight(x, y, radius float32, color rl.Color) *PointLight {
	return &PointLight{
		rl.NewVector2(x, y),
		color,
		radius,
		NewSprite(radius, x-radius/2, y-radius/2),
	}
}

func (l *PointLight) Setup() {
	l.sprite.Setup("textures/light.png", "", 1, map[string]rl.Shader{"light": rl.LoadShader("", "shaders/light.fs")})
}

func (l *PointLight) Cleanup() {
	l.sprite.Cleanup()
}

func (l *PointLight) Draw() {
	rl.BeginShaderMode(l.sprite.Shaders["light"])
	l.sprite.Draw()
	rl.EndShaderMode()
}

type Lights struct {
	Lights []*PointLight
}

func (l *Lights) Setup() {
	for _, light := range l.Lights {
		light.Setup()
	}
}

func (l *Lights) Cleanup() {
	for _, light := range l.Lights {
		light.Cleanup()
	}
}

func (l *Lights) Draw() {
	rl.BeginBlendMode(rl.BlendAddColors)
	for _, light := range l.Lights {
		light.Draw()
	}
	rl.EndBlendMode()
}

func (l *Lights) Add(x, y, radius float32, color rl.Color) {
	l.Lights = append(l.Lights, NewLight(x, y, radius, color))
}

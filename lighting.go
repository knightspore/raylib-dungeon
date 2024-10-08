package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PointLight struct {
	pos     rl.Vector2
	colour  rl.Color
	radius  float32
	sprite  *Sprite
	emitter *Emitter
}

func NewLight(x, y, radius float32, color rl.Color) *PointLight {
	return &PointLight{
		rl.NewVector2(x, y),
		color,
		radius,
		NewSprite(radius, x-radius/2, y-radius/2),
		NewEmitter(10, rl.NewRectangle(x-radius, y-radius, radius*2, radius*2), 5),
	}
}

func (l *PointLight) Setup() {
	l.sprite.Setup("textures/light.png", "", 1, map[string]rl.Shader{"light": rl.LoadShader("shaders/light.vs", "shaders/light.fs")})
	l.emitter.Setup()
}

func (l *PointLight) Cleanup() {
	l.sprite.Cleanup()
	l.emitter.Cleanup()
}

func (l *PointLight) Update() {
	l.emitter.Update()
	l.sprite.UpdateShaderValue("light", "u_time", []float32{float32(rl.GetTime())}, rl.ShaderUniformFloat)
}

func (l *PointLight) Draw() {
	rl.BeginShaderMode(l.sprite.Shaders["light"])
	l.sprite.Draw()
	l.emitter.Draw()
	rl.EndShaderMode()
}

func (l *PointLight) DrawNormal() {
	rl.BeginShaderMode(l.sprite.Shaders["light"])
	l.sprite.Draw()
	l.emitter.Draw()
	rl.EndShaderMode()
}

func (l *PointLight) DrawDebug() {
	l.sprite.DrawDebug()
	l.emitter.DrawDebug()
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

func (l *Lights) Update() {
	for _, light := range l.Lights {
		light.Update()
	}
}

func (l *Lights) Draw() {
	rl.BeginBlendMode(rl.BlendAddColors)
	for _, light := range l.Lights {
		light.Draw()
	}
	rl.EndBlendMode()
}

func (l *Lights) DrawNormal() {
	rl.BeginBlendMode(rl.BlendAdditive)
	for _, light := range l.Lights {
		light.DrawNormal()
	}
	rl.EndBlendMode()
}

func (l *Lights) DrawDebug() {
	for _, light := range l.Lights {
		light.DrawDebug()
	}
}

func (l *Lights) Add(x, y, radius float32, color rl.Color) {
	l.Lights = append(l.Lights, NewLight(x, y, radius, color))
}

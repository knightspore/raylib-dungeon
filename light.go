package main

import (
	"fmt"
	"log"

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
	l.sprite.Draw(l.sprite.Color)
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

func (l *Lights) Update(g *Game) {
	normLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_normal")
	rl.SetShaderValueTexture(g.Shaders.Lighting, normLoc, g.Textures.NormalPass.Texture)

	resLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_resolution")
	rl.SetShaderValue(g.Shaders.Lighting, resLoc, []float32{float32(g.Width), float32(g.Height)}, rl.ShaderUniformVec2)

	zoomLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_zoom")
	rl.SetShaderValue(g.Shaders.Lighting, zoomLoc, []float32{2.0 / g.Cam.Cam.Zoom}, rl.ShaderUniformFloat)

	ambientLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_ambient")
	rl.SetShaderValue(g.Shaders.Lighting, ambientLoc, []float32{0.1}, rl.ShaderUniformFloat)

	for i, light := range l.Lights {
		key := fmt.Sprintf("[%d]", i)
		posLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_lightPos"+key)
		colorLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_lightColor"+key)
		if posLoc == -1 || colorLoc == -1 {
			log.Fatalf("Failed to get shader location for %s", key)
		}
		pos := rl.GetWorldToScreen2D(rl.NewVector2(light.pos.X, light.pos.Y), *g.Cam.Cam)
		rl.SetShaderValue(g.Shaders.Lighting, posLoc, []float32{pos.X, pos.Y}, rl.ShaderUniformVec2)
		rl.SetShaderValue(g.Shaders.Lighting, colorLoc, []float32{float32(light.colour.R), float32(light.colour.G), float32(light.colour.B)}, rl.ShaderUniformVec3)
	}
}

func (l *Lights) Draw(g *Game) {
	rl.BeginBlendMode(rl.BlendAddColors)
	for _, light := range l.Lights {
		light.Draw()
	}
	rl.EndBlendMode()
}

func (l *Lights) Add(x, y, radius float32, color rl.Color) {
	l.Lights = append(l.Lights, NewLight(x, y, radius, color))
}

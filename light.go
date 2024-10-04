package main

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PointLight struct {
	Pos    rl.Vector2
	Color  rl.Color
	Radius float32
}

func NewLight(x, y, radius float32, color rl.Color) *PointLight {
	return &PointLight{rl.NewVector2(x, y), color, radius}
}

func (l *Lights) Update(cursorCenter rl.Vector2) {
}

func (l *Lights) Draw(g *Game) {
	for _, light := range l.Lights {
		rl.DrawTextureEx(g.Textures.Light, rl.Vector2{X: light.Pos.X - float32(BASE_SIZE)/2, Y: light.Pos.Y - float32(BASE_SIZE)/2}, 0, 2, light.Color)
	}
}

type Lights struct {
	Lights []*PointLight
}

func (l *Lights) Add(x, y, radius float32, color rl.Color) {
	l.Lights = append(l.Lights, NewLight(x, y, radius, color))
}

func (l *Lights) UpdateShader(g *Game) {
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
		pos := rl.GetWorldToScreen2D(rl.NewVector2(light.Pos.X, light.Pos.Y), *g.Cam.Cam)
		rl.SetShaderValue(g.Shaders.Lighting, posLoc, []float32{pos.X, pos.Y}, rl.ShaderUniformVec2)
		rl.SetShaderValue(g.Shaders.Lighting, colorLoc, []float32{float32(light.Color.R), float32(light.Color.G), float32(light.Color.B)}, rl.ShaderUniformVec3)
	}
}

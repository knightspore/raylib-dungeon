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

func NewPointLight(x, y, radius float32, color rl.Color) *PointLight {
	return &PointLight{rl.NewVector2(x, y), color, radius}
}

type Lights struct {
	Lights []*PointLight
}

func (l *Lights) Add(x, y, radius float32, color rl.Color) {
	l.Lights = append(l.Lights, NewPointLight(x, y, radius, color))
}

func (l *Lights) Setup(g *Game) {
	offset := float32(g.Map.tileSize) * 2.5
	fullWidth := float32(g.Map.sizeX * g.Map.tileSize)
	rad := float32(50)

	l.Add(offset, offset, rad, rl.NewColor(255, 0, 255, 255))
	l.Add(fullWidth-offset, fullWidth-offset, rad, rl.NewColor(0, 0, 255, 255))
	l.Add(fullWidth-offset, offset, rad, rl.NewColor(255, 255, 255, 255))
	l.Add(offset, fullWidth-offset, rad, rl.NewColor(255, 255, 0, 255))
}

func (l *Lights) UpdateShader(g *Game) {
	normLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_normal")
	rl.SetShaderValueTexture(g.Shaders.Lighting, normLoc, g.Textures.NormalPass.Texture)

	resLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_resolution")
	rl.SetShaderValue(g.Shaders.Lighting, resLoc, []float32{float32(g.Width), float32(g.Height)}, rl.ShaderUniformVec2)

	zoomLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_zoom")
	rl.SetShaderValue(g.Shaders.Lighting, zoomLoc, []float32{1.0 / g.Cam.Cam.Zoom}, rl.ShaderUniformFloat)

	ambientLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_ambient")
	rl.SetShaderValue(g.Shaders.Lighting, ambientLoc, []float32{0.15}, rl.ShaderUniformFloat)

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

func (l *Lights) Draw(g *Game) {
	for _, light := range l.Lights {
		rl.DrawTexture(g.Textures.Light, int32(light.Pos.X-float32(g.Textures.Light.Width)/2), int32(light.Pos.Y-float32(g.Textures.Light.Height)/2), light.Color)
	}
}

func (l *Lights) Update(cursorCenter rl.Vector2) {
	l.Lights[4].Pos.X = cursorCenter.X
	l.Lights[4].Pos.Y = cursorCenter.Y
}

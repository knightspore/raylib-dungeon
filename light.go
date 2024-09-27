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
	l.Add(0, 0, float32(g.Map.TileSize)*5, rl.NewColor(255, 0, 255, 0))
	l.Add(float32(g.Map.SizeX*g.Map.TileSize), float32(g.Map.SizeX*g.Map.TileSize), float32(g.Map.TileSize)*5, rl.NewColor(0, 0, 255, 0))
	l.Add(float32(g.Map.SizeX*g.Map.TileSize), 0, float32(g.Map.TileSize)*5, rl.NewColor(0, 255, 0, 0))
	l.Add(0, float32(g.Map.SizeX*g.Map.TileSize), float32(g.Map.TileSize)*5, rl.NewColor(255, 255, 0, 0))
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

var up = true

func (l *Lights) Update() {
	// Demo
	for _, light := range l.Lights {
		if up {
			light.Color.R += 1
		} else {
			light.Color.R -= 1
		}
		if light.Color.R == 255 {
			up = false
		}
		if light.Color.R == 0 {
			up = true
		}
	}
}

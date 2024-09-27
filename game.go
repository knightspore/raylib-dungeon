package main

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Width        int32
	Height       int32
	BaseSize     int32
	FrameTimer   float32
	CurrentFrame int
	Map          *Map
	Cam          *Camera
	Player       *Player
	Textures     *Textures
	Shaders      *Shaders
	Lights       []*PointLight
}

func NewGame(tiles []int, width int32, height int32, baseSize int32) *Game {
	Map := NewMap(tiles, baseSize)
	Player := NewPlayer(Map.Center(), baseSize)
	Cam := NewCamera(rl.NewVector2(float32(width/2), float32(height/2)), Player.Center())
	return &Game{
		Width:        width,
		Height:       height,
		BaseSize:     baseSize,
		FrameTimer:   0,
		CurrentFrame: 0,
		Map:          Map,
		Cam:          Cam,
		Player:       Player,
		Textures:     &Textures{},
		Shaders:      &Shaders{},
		Lights: []*PointLight{
			NewPointLight(0, 0, float32(Map.TileSize)*5, rl.NewColor(255, 0, 255, 0)),
			NewPointLight(float32(Map.SizeX*Map.TileSize), float32(Map.SizeX*Map.TileSize), float32(Map.TileSize)*5, rl.NewColor(0, 0, 255, 0)),
			NewPointLight(float32(Map.SizeX*Map.TileSize), 0, float32(Map.TileSize)*5, rl.NewColor(0, 255, 0, 0)),
			NewPointLight(0, float32(Map.SizeX*Map.TileSize), float32(Map.TileSize)*5, rl.NewColor(255, 255, 0, 0)),
		},
	}
}

func (g *Game) Setup() {
	g.Textures.Setup(g)
	g.Shaders.Setup()
}

func (g *Game) Update() {
	// Debug
	if rl.IsKeyPressed(rl.KeyBackSlash) {
		DEBUG = !DEBUG
	}

	// Frame timer
	g.FrameTimer += rl.GetFrameTime()
	if g.FrameTimer >= 0.25 {
		g.CurrentFrame = (g.CurrentFrame + 1) % 4
		g.FrameTimer = 0
	}

	// Game Logic
	g.Player.Update(g)
	g.Cam.Update(g)
}

func (g *Game) DrawNormalPass() {
	rl.BeginTextureMode(g.Textures.NormalPass)
	rl.BeginMode2D(*g.Cam.Cam)
	rl.ClearBackground(rl.NewColor(0, 0, 0, 0))

	g.Map.Draw(g, true)
	g.Player.Draw(g, true)

	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *Game) DrawColourPass() {
	rl.BeginTextureMode(g.Textures.ColorPass)
	rl.BeginMode2D(*g.Cam.Cam)
	rl.ClearBackground(rl.NewColor(0, 0, 0, 0))

	g.Map.Draw(g, false)
	g.Player.Draw(g, false)

	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *Game) DrawDeferredLightingPass() {
	rl.BeginShaderMode(g.Shaders.Lighting)

	normLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_normal")
	rl.SetShaderValueTexture(g.Shaders.Lighting, normLoc, g.Textures.NormalPass.Texture)

	resLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_resolution")
	rl.SetShaderValue(g.Shaders.Lighting, resLoc, []float32{float32(g.Width), float32(g.Height)}, rl.ShaderUniformVec2)

	zoomLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_zoom")
	rl.SetShaderValue(g.Shaders.Lighting, zoomLoc, []float32{1.0 / g.Cam.Cam.Zoom}, rl.ShaderUniformFloat)

	ambientLoc := rl.GetShaderLocation(g.Shaders.Lighting, "u_ambient")
	rl.SetShaderValue(g.Shaders.Lighting, ambientLoc, []float32{0.1}, rl.ShaderUniformFloat)

	for i, light := range g.Lights {
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

	rl.DrawTextureRec(g.Textures.ColorPass.Texture, rl.NewRectangle(0, 0, float32(g.Width), -float32(g.Height)), rl.NewVector2(0, 0), rl.RayWhite)

	rl.EndShaderMode()
}

func (g *Game) Draw() {
	// Clear the window
	rl.ClearBackground(rl.Black)

	// Draw to GBuffer
	g.DrawColourPass()
	g.DrawNormalPass()

	// Render Lighting
	rl.BeginDrawing()
	g.DrawDeferredLightingPass()

	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}

func (g *Game) Cleanup() {
	g.Textures.Cleanup()
	g.Shaders.Cleanup()
	rl.CloseWindow()
}

func (g *Game) Run() {
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagMsaa4xHint)
	rl.InitWindow(WIDTH, HEIGHT, "karoo")
	rl.SetTargetFPS(90)
	rl.DisableCursor()
	g.Setup()
	for !rl.WindowShouldClose() {
		g.Update()
		g.Draw()
	}
	g.Cleanup()
}

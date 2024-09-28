package main

import (
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
	Lights       *Lights
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
		Lights:       &Lights{},
	}
}

func (g *Game) Setup() {
	g.Textures.Setup(g)
	g.Shaders.Setup()
	g.Lights.Setup(g)

	g.Lights.Add(100, 100, 100, rl.NewColor(255, 255, 255, 255))
}

func (g *Game) Update() {
	// Debug
	if rl.IsKeyPressed(rl.KeyBackSlash) {
		DEBUG = !DEBUG
	}

	g.Lights.Lights[4].Pos.X = g.Player.Cursor.Dest.X
	g.Lights.Lights[4].Pos.Y = g.Player.Cursor.Dest.Y

	// Frame timer
	g.FrameTimer += rl.GetFrameTime()
	if g.FrameTimer >= 0.25 {
		g.CurrentFrame = (g.CurrentFrame + 1) % 4
		g.FrameTimer = 0
	}

	// Game Logic
	g.Player.Update(g)
	g.Cam.Update(g)
	g.Lights.Update()
}

func (g *Game) DrawNormalPass() {
	rl.BeginTextureMode(g.Textures.NormalPass)
	rl.BeginMode2D(*g.Cam.Cam)
	rl.ClearBackground(rl.Blank)

	g.Map.Draw(g, true)
	g.Player.Draw(g, true)

	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *Game) DrawColourPass() {
	rl.BeginTextureMode(g.Textures.ColorPass)
	rl.BeginMode2D(*g.Cam.Cam)
	rl.ClearBackground(rl.Blank)

	g.Map.Draw(g, false)
	g.Player.Draw(g, false)

	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *Game) DrawLightingPass() {
	g.Lights.UpdateShader(g)
	rl.BeginTextureMode(g.Textures.LightingPass)
	rl.BeginShaderMode(g.Shaders.Lighting)
	rl.DrawTextureRec(g.Textures.ColorPass.Texture, rl.NewRectangle(0, 0, float32(g.Width), -float32(g.Height)), rl.NewVector2(0, 0), rl.RayWhite)
	rl.EndShaderMode()
	rl.EndTextureMode()
}

func (g *Game) Draw() {
	// Clear the window
	rl.ClearBackground(rl.Black)

	// Draw to GBuffer
	g.DrawColourPass()
	g.DrawNormalPass()
	g.DrawLightingPass()

	// Draw to screen
	rl.BeginDrawing()

	// Render Lighting
	rl.BeginShaderMode(g.Shaders.PostProcess)
	rl.DrawTextureRec(g.Textures.LightingPass.Texture, rl.NewRectangle(0, 0, float32(g.Width), -float32(g.Height)), rl.NewVector2(0, 0), rl.RayWhite)
	rl.EndShaderMode()

	// UI
	rl.BeginMode2D(*g.Cam.Cam)
	g.Player.DrawCursor(g, false)
	rl.EndMode2D()

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

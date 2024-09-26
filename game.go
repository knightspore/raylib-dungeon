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
	}
}

func (g *Game) Setup() {
	g.Textures.Setup()
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

func (g *Game) RenderNormalPass() {
	rl.BeginTextureMode(g.Textures.NormalPass)
	rl.BeginMode2D(*g.Cam.Cam)
	rl.ClearBackground(rl.NewColor(0, 0, 0, 255))

	g.Map.Draw(g, true)
	g.Player.Draw(g, true)

	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *Game) RenderDiffusePass() {
	rl.BeginTextureMode(g.Textures.RenderPass)
	rl.BeginMode2D(*g.Cam.Cam)
	rl.ClearBackground(rl.NewColor(0, 0, 0, 255))

	g.Map.Draw(g, false)
	g.Player.Draw(g, false)

	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *Game) Draw() {
	// Clear the window
	rl.ClearBackground(rl.Black)

	// Draw to GBuffer
	g.RenderDiffusePass()
	g.RenderNormalPass()

	// Start drawing to the screen
	rl.BeginDrawing()
	rl.BeginShaderMode(g.Shaders.Render)

	normLoc := rl.GetShaderLocation(g.Shaders.Render, "u_normal")
	rl.SetShaderValueTexture(g.Shaders.Render, normLoc, g.Textures.NormalPass.Texture)

	resLoc := rl.GetShaderLocation(g.Shaders.Render, "u_resolution")
	rl.SetShaderValue(g.Shaders.Render, resLoc, []float32{float32(g.Width), float32(g.Height)}, rl.ShaderUniformVec2)

	lightLoc := rl.GetShaderLocation(g.Shaders.Render, "u_lightPos")
	lightPos := rl.GetWorldToScreen2D(g.Player.CursorCenter(), *g.Cam.Cam)
	pos := []float32{
		lightPos.X - float32(float32(g.Player.Size)*g.Cam.Cam.Zoom)/2,
		lightPos.Y - float32(float32(g.Player.Size)*g.Cam.Cam.Zoom)/2,
	}
	rl.SetShaderValue(g.Shaders.Render, lightLoc, pos, rl.ShaderUniformVec2)

	zoomLoc := rl.GetShaderLocation(g.Shaders.Render, "u_zoom")
	rl.SetShaderValue(g.Shaders.Render, zoomLoc, []float32{g.Cam.Cam.Zoom}, rl.ShaderUniformFloat)

	rl.DrawTextureRec(g.Textures.RenderPass.Texture, rl.NewRectangle(0, 0, float32(g.Width), -float32(g.Height)), rl.NewVector2(0, 0), rl.RayWhite)

	rl.EndShaderMode()

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

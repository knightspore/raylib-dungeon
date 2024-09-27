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
		Lights:       []*PointLight{},
	}
}

func (g *Game) Setup() {
	g.Textures.Setup()
	g.Shaders.Setup()
	g.Lights = append(g.Lights, NewPointLight(float32(g.Map.TileSize), float32(g.Map.TileSize), float32(g.Map.TileSize)*5, rl.NewColor(140, 12, 60, 255)))
}

func (g *Game) Update() {
	// Debug
	if rl.IsKeyPressed(rl.KeyBackSlash) {
		DEBUG = !DEBUG
	}

	if rl.IsKeyPressed(rl.KeyOne) {
		g.Lights[0].Color.R += 10
		if g.Lights[0].Color.R > 255 {
			g.Lights[0].Color.R = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyTwo) {
		g.Lights[0].Color.G += 10
		if g.Lights[0].Color.G > 255 {
			g.Lights[0].Color.G = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyThree) {
		g.Lights[0].Color.B += 10
		if g.Lights[0].Color.B > 255 {
			g.Lights[0].Color.B = 0
		}
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		g.Lights[0].Pos.X -= g.Player.Speed
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		g.Lights[0].Pos.X += g.Player.Speed
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		g.Lights[0].Pos.Y -= g.Player.Speed
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		g.Lights[0].Pos.Y += g.Player.Speed
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
	rl.ClearBackground(rl.NewColor(0, 0, 0, 0))

	g.Map.Draw(g, true)
	g.Player.Draw(g, true)

	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *Game) RenderDiffusePass() {
	rl.BeginTextureMode(g.Textures.RenderPass)
	rl.BeginMode2D(*g.Cam.Cam)
	rl.ClearBackground(rl.NewColor(0, 0, 0, 0))

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

	// Deferred Rendering Pass
	rl.BeginDrawing()
	rl.BeginShaderMode(g.Shaders.Render)

	normLoc := rl.GetShaderLocation(g.Shaders.Render, "u_normal")
	rl.SetShaderValueTexture(g.Shaders.Render, normLoc, g.Textures.NormalPass.Texture)

	diffLoc := rl.GetShaderLocation(g.Shaders.Render, "u_diffuse")
	rl.SetShaderValueTexture(g.Shaders.Render, diffLoc, g.Textures.RenderPass.Texture)

	resLoc := rl.GetShaderLocation(g.Shaders.Render, "u_resolution")
	rl.SetShaderValue(g.Shaders.Render, resLoc, []float32{float32(g.Width), float32(g.Height)}, rl.ShaderUniformVec2)

	lightLoc := rl.GetShaderLocation(g.Shaders.Render, "u_lightPos")
	lightPos := rl.GetWorldToScreen2D(g.Lights[0].Pos, *g.Cam.Cam)
	rl.SetShaderValue(g.Shaders.Render, lightLoc, []float32{lightPos.X, lightPos.Y}, rl.ShaderUniformVec2)

	lightColorLoc := rl.GetShaderLocation(g.Shaders.Render, "u_lightColor")
	rl.SetShaderValue(g.Shaders.Render, lightColorLoc, []float32{float32(g.Lights[0].Color.R), float32(g.Lights[0].Color.G), float32(g.Lights[0].Color.B)}, rl.ShaderUniformVec3)

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

package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	_debug       bool
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
	Cam := NewCamera(raylib.NewVector2(float32(width/2), float32(height/2)), Player.Center())
	return &Game{
		Width:        width,
		Height:       height,
		BaseSize:     baseSize,
		FrameTimer:   0,
		CurrentFrame: 0,
		Map:          Map,
		Cam:          Cam,
		Player:       Player,
		Textures:     NewTextures(),
		Shaders:      NewShaders(),
	}
}

func (g *Game) Setup() {
	g.Textures.Setup()
	g.Shaders.Setup()
}

func (g *Game) Update() {
	// Debug
	if rl.IsKeyPressed(rl.KeyBackSlash) {
		g._debug = !g._debug
	}

	// Frame timer
	g.FrameTimer += raylib.GetFrameTime()
	if g.FrameTimer >= 0.25 {
		g.CurrentFrame = (g.CurrentFrame + 1) % 4
		g.FrameTimer = 0
	}

	// Game Logic
	g.Player.Update(g)
	g.Cam.Zoom()
	g.Cam.SmoothFollow(g.Player.Center())
	g.Shaders.Update()
}

func (g *Game) Draw() {
	raylib.BeginDrawing()
	raylib.ClearBackground(raylib.RayWhite)

	// Draw map to render texture
	raylib.BeginTextureMode(g.Textures.Map_RenderTexture)
	raylib.BeginMode2D(*g.Cam.Cam)

	g.Map.Draw(g)
	g.Player.Draw(g)

	raylib.EndMode2D()
	raylib.EndTextureMode()

	// Draw render texture to screen
	raylib.BeginShaderMode(g.Shaders.Render)
	raylib.DrawTextureRec(g.Textures.Map_RenderTexture.Texture, raylib.NewRectangle(0, 0, float32(g.Width), -float32(g.Height)), raylib.NewVector2(0, 0), raylib.RayWhite)
	raylib.EndShaderMode()

	// Draw over render texture
	g.Player.DrawCursor(g)

	raylib.DrawFPS(10, 10)
	raylib.EndDrawing()
}

func (g *Game) Cleanup() {
	g.Textures.Cleanup()
	g.Shaders.Cleanup()
	raylib.CloseWindow()
}

func (g *Game) Run() {
	raylib.SetConfigFlags(raylib.FlagVsyncHint | raylib.FlagMsaa4xHint)
	raylib.InitWindow(WIDTH, HEIGHT, "karoo")
	raylib.SetTargetFPS(60)
	raylib.DisableCursor()
	g.Setup()
	for !raylib.WindowShouldClose() {
		g.Update()
		g.Draw()
	}
	g.Cleanup()
}

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Width    int32
	Height   int32
	BaseSize int32
	Map      *Map
	Cam      *Camera
	Player   *Player
	Textures *Textures
	Shaders  *Shaders
	Lights   *Lights
	Cursor   *Cursor
}

func NewGame(tiles []int, width int32, height int32, baseSize int32) *Game {
	Map := NewMap(tiles, baseSize)
	return &Game{
		Width:    width,
		Height:   height,
		BaseSize: baseSize,
		Map:      Map,
		Cam:      NewCamera(rl.NewVector2(float32(width/2), float32(height/2)), rl.NewVector2(float32(width/2), float32(height/2))),
		Player:   NewPlayer(rl.NewVector2(Map.center().X, Map.center().Y), int32(baseSize)),
		Textures: &Textures{},
		Shaders:  &Shaders{},
		Lights:   &Lights{},
		Cursor:   NewCursor(float32(baseSize), float32(width/2), float32(height/2)),
	}
}

func (g *Game) Setup() {
	g.Textures.Setup(g)
	g.Shaders.Setup()
	g.Map.Setup()
	g.Player.Setup()
	g.Cursor.Setup()
	g.Lights.Setup(g)
}

func (g *Game) Cleanup() {
	g.Textures.Cleanup()
	g.Shaders.Cleanup()
	g.Map.Cleanup()
	g.Player.Cleanup()
	g.Cursor.Cleanup()
	g.Lights.Cleanup()
	rl.CloseWindow()
}

func (g *Game) Update() {
	UpdateDebug()
	g.Player.Update(g)
	g.Cursor.Update()
	g.Lights.Update(g.Cursor.Center())
	g.Cam.Update(g)
}

func (g *Game) DrawNormalPass() {
	rl.BeginTextureMode(g.Textures.NormalPass)
	rl.BeginMode2D(*g.Cam.Cam)
	rl.ClearBackground(rl.Blank)

	g.Map.DrawNormal()
	g.Player.Draw(g.Player.Sprite.Normal)
	g.Cursor.Draw(g.Cursor.Sprite.Normal)

	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *Game) DrawColourPass() {
	rl.BeginTextureMode(g.Textures.ColorPass)
	rl.BeginMode2D(*g.Cam.Cam)
	rl.ClearBackground(rl.Blank)

	g.Map.Draw()
	g.Player.Draw(g.Player.Sprite.Color)
	g.Lights.Draw(g)
	g.Cursor.Draw(g.Cursor.Sprite.Color)

	if DEBUG {
		DrawDebugLine(g.Player.Center(), g.Cursor.Center())
		targetTile := g.Map.vectorToTile(g.Cursor.Center())
		if targetTile != nil {
			DrawDebugArea(targetTile.sprite.dest, targetTile.sprite.Center(), rl.Red)
			DrawDebugLine(g.Cursor.Center(), targetTile.sprite.Center())
		}
	}

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

	// Draw from GBuffer with post processing shader
	rl.BeginDrawing()

	rl.BeginShaderMode(g.Shaders.PostProcess)
	rl.DrawTextureRec(g.Textures.LightingPass.Texture, rl.NewRectangle(0, 0, float32(g.Width), -float32(g.Height)), rl.NewVector2(0, 0), rl.RayWhite)
	rl.EndShaderMode()

	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}

func (g *Game) Run() {
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagMsaa4xHint)
	rl.InitWindow(WIDTH, HEIGHT, "karoo")
	rl.SetTargetFPS(60)
	rl.DisableCursor()
	g.Setup()
	for !rl.WindowShouldClose() {
		g.Update()
		g.Draw()
	}
	g.Cleanup()
}

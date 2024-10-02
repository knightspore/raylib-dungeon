package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Width    int32
	Height   int32
	BaseSize int32
	Textures *Textures
	Shaders  *Shaders
	Cam      *Camera
	Map      *Map
	Player   *Player
	Cursor   *Cursor
	Lights   *Lights
	Emitter  *Emitter
}

func NewGame(tiles []int, width int32, height int32, baseSize int32) *Game {
	Map := NewMap(tiles, baseSize)
	return &Game{
		Width:    width,
		Height:   height,
		BaseSize: baseSize,
		Textures: &Textures{},
		Shaders:  &Shaders{},
		Cam:      NewCamera(rl.NewVector2(float32(width/2), float32(height/2)), rl.NewVector2(float32(width/2), float32(height/2))),
		Map:      Map,
		Player:   NewPlayer(rl.NewVector2(Map.center().X, Map.center().Y), int32(baseSize)),
		Cursor:   NewCursor(float32(baseSize), float32(width/2), float32(height/2)),
		Lights:   &Lights{},
		Emitter:  NewEmitter(500, rl.NewRectangle(0, 0, float32(Map.tileSize*Map.sizeX), float32(Map.tileSize*Map.sizeY)), 10),
	}
}

func (g *Game) Setup() {
	g.Textures.Setup(g)
	g.Shaders.Setup()
	g.Map.Setup()
	g.Player.Setup()
	g.Cursor.Setup()
	g.Lights.Setup(g)
	g.Emitter.Setup()
}

func (g *Game) Cleanup() {
	g.Textures.Cleanup()
	g.Shaders.Cleanup()
	g.Map.Cleanup()
	g.Player.Cleanup()
	g.Cursor.Cleanup()
	g.Lights.Cleanup()
	g.Emitter.Cleanup()
	rl.CloseWindow()
}

func (g *Game) Update() {
	UpdateDebug()
	g.Player.Update(g)
	g.Cursor.Update()
	g.Lights.Update(g.Cursor.Center())
	g.Emitter.Update()
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
	g.Emitter.Draw()
	g.Cursor.Draw(g.Cursor.Sprite.Color)

	if DEBUG {
		DrawDebugLine(g.Player.Center(), g.Cursor.Center())
		targetTile := g.Map.vectorToTile(g.Cursor.Center())
		if targetTile != nil {
			DrawDebugArea(targetTile.sprite.dest, targetTile.sprite.Center(), rl.Green)
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

	if DEBUG {
		rl.DrawFPS(10, 10)
	}
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

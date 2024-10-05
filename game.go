package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	GBuffer *GBuffer
	Cam     *Camera
	Map     *Map
	Player  *Player
	Cursor  *Cursor
	Lights  *Lights
	Emitter *Emitter
}

func NewGame(tiles []int) *Game {
	Map := NewMap(tiles)
	return &Game{
		GBuffer: &GBuffer{},
		Cam:     NewCamera(),
		Map:     Map,
		Player:  NewPlayer(rl.NewVector2(Map.center().X, Map.center().Y)),
		Cursor:  NewCursor(),
		Lights:  &Lights{},
		Emitter: NewEmitter(1000, rl.NewRectangle(0, 0, float32(Map.tileSize*Map.sizeX), float32(Map.tileSize*Map.sizeY)), 10),
	}
}

func (g *Game) Setup() {
	g.GBuffer.Setup()
	g.Map.Setup()
	g.Player.Setup()
	g.Cursor.Setup()
	g.Lights.Setup()
	g.Emitter.Setup()
}

func (g *Game) Cleanup() {
	g.GBuffer.Cleanup()
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
	g.Emitter.Update()
	g.Cam.Update(g)
	g.GBuffer.Update(
		g.Lights,
		g.Cam.Cam,
		func() {
			g.Map.Draw()
			g.Lights.Draw()
			g.Player.Draw()
			g.Emitter.Draw()
			g.Cursor.Draw()
		},
		func() {
			g.Map.DrawNormal()
			g.Player.DrawNormal()
			g.Cursor.DrawNormal()
		},
		func() {
			DrawDebugSprite(g.Map.sprite)
			DrawDebugSprite(g.Player.Sprite)
			DrawDebugSprite(g.Cursor.Sprite)
			DrawDebugParticles(&g.Emitter.particles)
			DrawDebugLine(g.Player.Center(), g.Cursor.Center())
			targetTile := g.Map.vectorToTile(g.Cursor.Center())
			if targetTile != nil {
				DrawDebugArea(targetTile.sprite.dest, targetTile.sprite.Center(), rl.Green)
				DrawDebugLine(g.Cursor.Center(), targetTile.sprite.Center())
			}
		},
	)
}

func (g *Game) Draw() {
	g.GBuffer.Draw()
}

func (g *Game) Run() {
	rl.SetConfigFlags(rl.FlagVsyncHint | rl.FlagMsaa4xHint | rl.FlagWindowUndecorated)
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

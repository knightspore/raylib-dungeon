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
	}
}

func LoadGame(path string) *Game {

	image := rl.LoadImage(path)

	tiles := []int{}
	playerPos := rl.NewVector2(0, 0)
	endPos := rl.NewVector2(0, 0)
	lightPositions := []rl.Vector2{}

	for y := 0; y < int(image.Height); y++ {
		for x := 0; x < int(image.Width); x++ {
			color := rl.GetImageColor(*image, int32(x), int32(y))

			// Read basic map data
			if color.R == 255 && color.G == 0 && color.B == 0 { // Empty space
				tiles = append(tiles, TILE_EMPTY)
			} else if color.R == 255 && color.G == 255 && color.B == 0 { // Wall
				tiles = append(tiles, TILE_EMPTY)
			} else { // Defaulting to floor for now
				tiles = append(tiles, TILE_FLOOR)
			}

			// Read player start / end positions
			if color.R == 0 && color.G == 255 && color.B == 0 {
				playerPos = rl.NewVector2(float32(x*BASE_SIZE+BASE_SIZE/2), float32(y*BASE_SIZE+BASE_SIZE/2))
			}

			if color.R == 0 && color.G == 0 && color.B == 255 {
				endPos = rl.NewVector2(float32(x*BASE_SIZE+BASE_SIZE/2), float32(y*BASE_SIZE+BASE_SIZE/2))
			}

			// Read light positions
			if color.R == 255 && color.G == 255 && color.B == 0 {
				lightPositions = append(lightPositions, rl.NewVector2(float32(x*BASE_SIZE+BASE_SIZE/2), float32(y*BASE_SIZE+BASE_SIZE/2)))
			}
		}
	}

	lights := &Lights{}
	for _, pos := range lightPositions {
		lights.Add(pos.X, pos.Y, 50, rl.NewColor(230, 230, 100, 255))
	}
	lights.Add(endPos.X, endPos.Y, 50, rl.NewColor(255, 0, 0, 255))

	rl.UnloadImage(image)

	game := &Game{
		GBuffer: &GBuffer{},
		Cam:     NewCamera(),
		Map:     NewMap(tiles),
		Player:  NewPlayer(playerPos),
		Cursor:  NewCursor(),
		Lights:  lights,
	}

	return game
}

func (g *Game) Setup() {
	g.GBuffer.Setup()
	g.Map.Setup()
	g.Player.Setup()
	g.Cursor.Setup()
	g.Lights.Setup()
}

func (g *Game) Cleanup() {
	g.GBuffer.Cleanup()
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
	g.Lights.Update()
	g.Cam.Update(g)
	g.GBuffer.Update(
		g.Lights,
		g.Cam.Cam,
		func() {
			g.Map.Draw()
			g.Lights.Draw()
			g.Player.Draw()
			g.Cursor.Draw()
		},
		func() {
			g.Map.DrawNormal()
			g.Lights.DrawNormal()
			g.Player.DrawNormal()
			g.Cursor.DrawNormal()
		},
		func() {
			g.Map.DrawDebug()
			g.Player.DrawDebug()
			g.Cursor.DrawDebug()
			g.Lights.DrawDebug()

			// Draw debug lines for player and cursor
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

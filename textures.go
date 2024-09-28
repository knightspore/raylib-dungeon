package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TEXTURE_FLOOR         = "textures/floor_tilesheet.png"
	TEXTURE_FLOOR_NORMAL  = "textures/floor_tilesheet_n.png"
	TEXTURE_NOTILE        = "textures/notile.png"
	TEXTURE_NOTILE_NORMAL = "textures/notile_n.png"
	TEXTURE_PLAYER        = "textures/player_tilesheet.png"
	TEXTURE_PLAYER_NORMAL = "textures/player_tilesheet_n.png"
	TEXTURE_LIGHT         = "textures/light.png"
)

type Textures struct {
	Floor          rl.Texture2D
	Floor_Normal   rl.Texture2D
	NoFloor        rl.Texture2D
	NoFloor_Normal rl.Texture2D
	Player         rl.Texture2D
	Player_Normal  rl.Texture2D
	Light          rl.Texture2D
	ColorPass      rl.RenderTexture2D
	NormalPass     rl.RenderTexture2D
	LightingPass   rl.RenderTexture2D
}

func (t *Textures) Setup(g *Game) {
	t.Floor = rl.LoadTexture(TEXTURE_FLOOR)
	t.Floor_Normal = rl.LoadTexture(TEXTURE_FLOOR_NORMAL)

	t.NoFloor = rl.LoadTexture(TEXTURE_NOTILE)
	t.NoFloor_Normal = rl.LoadTexture(TEXTURE_NOTILE_NORMAL)

	t.Player = rl.LoadTexture(TEXTURE_PLAYER)
	t.Player_Normal = rl.LoadTexture(TEXTURE_PLAYER_NORMAL)

	t.Light = rl.LoadTexture(TEXTURE_LIGHT)

	t.ColorPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
	t.NormalPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
	t.LightingPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
}

func (t *Textures) Cleanup() {
	rl.UnloadTexture(t.Floor)
	rl.UnloadTexture(t.Floor_Normal)

	rl.UnloadTexture(t.NoFloor)
	rl.UnloadTexture(t.NoFloor_Normal)

	rl.UnloadTexture(t.Player)
	rl.UnloadTexture(t.Player_Normal)

	rl.UnloadTexture(t.Light)

	rl.UnloadRenderTexture(t.ColorPass)
	rl.UnloadRenderTexture(t.NormalPass)
	rl.UnloadRenderTexture(t.LightingPass)
}

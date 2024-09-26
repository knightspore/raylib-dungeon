package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TEXTURE_FLOOR         = "textures/floor_tilesheet.png"
	TEXTURE_FLOOR_NORMAL  = "textures/floor_tilesheet_n.png"
	TEXTURE_PLAYER        = "textures/player_tilesheet.png"
	TEXTURE_PLAYER_NORMAL = "textures/player_tilesheet_n.png"
	TEXTURE_CURSOR        = "textures/cursor_tilesheet.png"
	TEXTURE_CURSOR_NORMAL = "textures/cursor_tilesheet_n.png"
)

type Textures struct {
	Floor         rl.Texture2D
	Floor_Normal  rl.Texture2D
	Player        rl.Texture2D
	Player_Normal rl.Texture2D
	Cursor        rl.Texture2D
	Cursor_Normal rl.Texture2D
	RenderPass    rl.RenderTexture2D
	NormalPass    rl.RenderTexture2D
}

func (t *Textures) Setup() {
	t.Floor = rl.LoadTexture(TEXTURE_FLOOR)
	t.Floor_Normal = rl.LoadTexture(TEXTURE_FLOOR_NORMAL)

	t.Player = rl.LoadTexture(TEXTURE_PLAYER)
	t.Player_Normal = rl.LoadTexture(TEXTURE_PLAYER_NORMAL)

	t.Cursor = rl.LoadTexture(TEXTURE_CURSOR)
	t.Cursor_Normal = rl.LoadTexture(TEXTURE_CURSOR_NORMAL)

	t.RenderPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))

	t.NormalPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
}

func (t *Textures) Cleanup() {
	rl.UnloadTexture(t.Floor)
	rl.UnloadTexture(t.Floor_Normal)

	rl.UnloadTexture(t.Player)
	rl.UnloadTexture(t.Player_Normal)

	rl.UnloadTexture(t.Cursor)
	rl.UnloadTexture(t.Cursor_Normal)

	rl.UnloadRenderTexture(t.RenderPass)

	rl.UnloadRenderTexture(t.NormalPass)
}

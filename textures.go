package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TEXTURE_FLOOR            = "textures/Wood/WOODA.png"
	TEXTURE_PLAYER           = "textures/player.png"
	TEXTURE_CURSOR_TILESHEET = "textures/cursor_tilesheet.png"
)

type Textures struct {
	Floor             rl.Texture2D
	Player            rl.Texture2D
	Cursor            rl.Texture2D
	Map_RenderTexture rl.RenderTexture2D
}

func NewTextures() *Textures {
	return &Textures{}
}

func (t *Textures) Setup() {
	t.Floor = rl.LoadTexture(TEXTURE_FLOOR)
	t.Player = rl.LoadTexture(TEXTURE_PLAYER)
	t.Cursor = rl.LoadTexture(TEXTURE_CURSOR_TILESHEET)
	t.Map_RenderTexture = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
}

func (t *Textures) Cleanup() {
	rl.UnloadTexture(t.Floor)
	rl.UnloadTexture(t.Player)
	rl.UnloadTexture(t.Cursor)
	rl.UnloadRenderTexture(t.Map_RenderTexture)
}

package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TEXTURE_LIGHT = "textures/light.png"
)

type Textures struct {
	Light        rl.Texture2D
	ColorPass    rl.RenderTexture2D
	NormalPass   rl.RenderTexture2D
	LightingPass rl.RenderTexture2D
}

func (t *Textures) Setup(g *Game) {
	t.Light = rl.LoadTexture(TEXTURE_LIGHT)

	t.ColorPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
	t.NormalPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
	t.LightingPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
}

func (t *Textures) Cleanup() {
	rl.UnloadTexture(t.Light)

	rl.UnloadRenderTexture(t.ColorPass)
	rl.UnloadRenderTexture(t.NormalPass)
	rl.UnloadRenderTexture(t.LightingPass)
}

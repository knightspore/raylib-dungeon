package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Textures struct {
	Light        rl.Texture2D
	ColorPass    rl.RenderTexture2D
	NormalPass   rl.RenderTexture2D
	LightingPass rl.RenderTexture2D
}

func (t *Textures) Setup(g *Game) {
	t.ColorPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
	t.NormalPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
	t.LightingPass = rl.LoadRenderTexture(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()))
}

func (t *Textures) Cleanup() {
	rl.UnloadRenderTexture(t.ColorPass)
	rl.UnloadRenderTexture(t.NormalPass)
	rl.UnloadRenderTexture(t.LightingPass)
}

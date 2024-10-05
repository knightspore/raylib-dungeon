package main

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GBuffer struct {
	Albedo              rl.RenderTexture2D
	Normal              rl.RenderTexture2D
	Render              rl.RenderTexture2D
	PostProBloom        rl.RenderTexture2D
	PostProChroma       rl.RenderTexture2D
	Debug               rl.RenderTexture2D
	LightingShader      rl.Shader
	PostProBloomShader  rl.Shader
	PostProChromaShader rl.Shader
}

func (g *GBuffer) Setup() {
	g.Albedo = rl.LoadRenderTexture(WIDTH, HEIGHT)
	g.Normal = rl.LoadRenderTexture(WIDTH, HEIGHT)
	g.Render = rl.LoadRenderTexture(WIDTH, HEIGHT)
	g.PostProBloom = rl.LoadRenderTexture(WIDTH, HEIGHT)
	g.PostProChroma = rl.LoadRenderTexture(WIDTH, HEIGHT)
	g.Debug = rl.LoadRenderTexture(WIDTH, HEIGHT)
	g.LightingShader = rl.LoadShader("", "shaders/deferred-lighting.fs")
	g.PostProBloomShader = rl.LoadShader("", "shaders/postpro-bloom.fs")
	g.PostProChromaShader = rl.LoadShader("", "shaders/postpro-chromabberation.fs")
}

func (g *GBuffer) Cleanup() {
	rl.UnloadRenderTexture(g.Albedo)
	rl.UnloadRenderTexture(g.Normal)
	rl.UnloadRenderTexture(g.Render)
	rl.UnloadRenderTexture(g.PostProBloom)
	rl.UnloadRenderTexture(g.PostProChroma)
	rl.UnloadRenderTexture(g.Debug)
	rl.UnloadShader(g.LightingShader)
	rl.UnloadShader(g.PostProBloomShader)
	rl.UnloadShader(g.PostProChromaShader)
}

func (g *GBuffer) Update(l *Lights, cam *rl.Camera2D, drawColour func(), drawNormal func(), drawDebug func()) {
	g.RenderColourPass(cam, drawColour)
	g.RenderNormalPass(cam, drawNormal)
	if DEBUG {
		g.RenderDebugPass(cam, drawDebug)
	}

	normLoc := rl.GetShaderLocation(g.LightingShader, "u_normal")
	rl.SetShaderValueTexture(g.LightingShader, normLoc, g.Normal.Texture)

	resLoc := rl.GetShaderLocation(g.LightingShader, "u_resolution")
	rl.SetShaderValue(g.LightingShader, resLoc, []float32{float32(WIDTH), float32(HEIGHT)}, rl.ShaderUniformVec2)

	zoomLoc := rl.GetShaderLocation(g.LightingShader, "u_zoom")
	rl.SetShaderValue(g.LightingShader, zoomLoc, []float32{2.0 / cam.Zoom}, rl.ShaderUniformFloat)

	ambientLoc := rl.GetShaderLocation(g.LightingShader, "u_ambient")
	rl.SetShaderValue(g.LightingShader, ambientLoc, []float32{0.15}, rl.ShaderUniformFloat)

	for i, light := range l.Lights {
		key := fmt.Sprintf("[%d]", i)
		posLoc := rl.GetShaderLocation(g.LightingShader, "u_lightPos"+key)
		colorLoc := rl.GetShaderLocation(g.LightingShader, "u_lightColor"+key)
		if posLoc == -1 || colorLoc == -1 {
			log.Fatalf("Failed to get shader location for %s", key)
		}
		pos := rl.GetWorldToScreen2D(rl.NewVector2(light.pos.X, light.pos.Y), *cam)
		rl.SetShaderValue(g.LightingShader, posLoc, []float32{pos.X, pos.Y}, rl.ShaderUniformVec2)
		rl.SetShaderValue(g.LightingShader, colorLoc, []float32{float32(light.colour.R), float32(light.colour.G), float32(light.colour.B)}, rl.ShaderUniformVec3)
	}

	g.RenderLightingPass()
	g.RenderPostProcessPass()
}

func (g *GBuffer) RenderColourPass(cam *rl.Camera2D, draw func()) {
	rl.BeginTextureMode(g.Albedo)
	rl.BeginMode2D(*cam)
	rl.ClearBackground(BG_COLOR)
	draw()
	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *GBuffer) RenderNormalPass(cam *rl.Camera2D, draw func()) {
	rl.BeginTextureMode(g.Normal)
	rl.BeginMode2D(*cam)
	rl.ClearBackground(rl.Blank)
	draw()
	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *GBuffer) RenderDebugPass(cam *rl.Camera2D, draw func()) {
	rl.BeginTextureMode(g.Debug)
	rl.BeginMode2D(*cam)
	rl.ClearBackground(rl.Blank)
	draw()
	rl.EndMode2D()
	rl.EndTextureMode()
}

func (g *GBuffer) RenderLightingPass() {
	rl.BeginTextureMode(g.Render)
	rl.BeginShaderMode(g.LightingShader)
	rl.DrawTextureRec(g.Albedo.Texture, rl.NewRectangle(0, 0, float32(WIDTH), -float32(HEIGHT)), rl.NewVector2(0, 0), rl.RayWhite)
	rl.EndShaderMode()
	rl.EndTextureMode()
}

func (g *GBuffer) RenderPostProcessPass() {
	rl.BeginTextureMode(g.PostProBloom)
	rl.BeginShaderMode(g.PostProBloomShader)
	rl.DrawTextureRec(g.Render.Texture, rl.NewRectangle(0, 0, float32(WIDTH), -float32(HEIGHT)), rl.NewVector2(0, 0), rl.RayWhite)
	rl.EndShaderMode()
	rl.EndTextureMode()

	rl.BeginTextureMode(g.PostProChroma)
	rl.BeginShaderMode(g.PostProChromaShader)
	rl.DrawTextureRec(g.PostProBloom.Texture, rl.NewRectangle(0, 0, float32(WIDTH), -float32(HEIGHT)), rl.NewVector2(0, 0), rl.RayWhite)
	rl.EndShaderMode()
	rl.EndTextureMode()
}

func (g *GBuffer) Draw() {
	rl.ClearBackground(rl.Black)
	rl.BeginDrawing()
	rl.DrawTextureRec(g.PostProChroma.Texture, rl.NewRectangle(0, 0, float32(WIDTH), -float32(HEIGHT)), rl.NewVector2(0, 0), rl.RayWhite)

	rl.EndShaderMode()
	if DEBUG {
		rl.DrawTextureRec(g.Debug.Texture, rl.NewRectangle(0, 0, float32(WIDTH), -float32(HEIGHT)), rl.NewVector2(0, 0), rl.RayWhite)
		rl.DrawText("DEBUG", 10, 30, 20, rl.RayWhite)
	}
	rl.DrawFPS(10, 10)
	rl.EndDrawing()
}

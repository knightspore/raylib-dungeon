package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SHADER_RENDER_VS   = "shaders/render.vs"
	SHADER_RENDER_FS   = "shaders/render.fs"
	SHADER_PLAYER_VS   = "shaders/player.vs"
	SHADER_PLAYER_FS   = "shaders/player.fs"
	SHADER_LIGHTING_FS = "shaders/lighting.fs"
)

type Shaders struct {
	Render   rl.Shader
	Player   rl.Shader
	Lighting rl.Shader
}

func (s *Shaders) Setup() {
	s.Render = rl.LoadShader(SHADER_RENDER_VS, SHADER_RENDER_FS)
	s.Player = rl.LoadShader(SHADER_PLAYER_VS, SHADER_PLAYER_FS)
	s.Lighting = rl.LoadShader("", SHADER_LIGHTING_FS)
}

func (s *Shaders) Cleanup() {
	rl.UnloadShader(s.Player)
	rl.UnloadShader(s.Render)
	rl.UnloadShader(s.Lighting)
}

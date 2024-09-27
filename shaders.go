package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SHADER_PLAYER_VS      = "shaders/player.vs"
	SHADER_PLAYER_FS      = "shaders/player.fs"
	SHADER_LIGHTING_FS    = "shaders/lighting.fs"
	SHADER_POSTPROCESS_FS = "shaders/postprocess.fs"
)

type Shaders struct {
	Player      rl.Shader
	Lighting    rl.Shader
	PostProcess rl.Shader
}

func (s *Shaders) Setup() {
	s.Player = rl.LoadShader(SHADER_PLAYER_VS, SHADER_PLAYER_FS)
	s.Lighting = rl.LoadShader("", SHADER_LIGHTING_FS)
	s.PostProcess = rl.LoadShader("", SHADER_POSTPROCESS_FS)
}

func (s *Shaders) Cleanup() {
	rl.UnloadShader(s.Player)
	rl.UnloadShader(s.Lighting)
	rl.UnloadShader(s.PostProcess)
}

package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SHADER_LIGHTING_FS    = "shaders/deferred-lighting.fs"
	SHADER_POSTPROCESS_FS = "shaders/postprocess.fs"
)

type Shaders struct {
	Lighting    rl.Shader
	PostProcess rl.Shader
}

func (s *Shaders) Setup() {
	s.Lighting = rl.LoadShader("", SHADER_LIGHTING_FS)
	s.PostProcess = rl.LoadShader("", SHADER_POSTPROCESS_FS)
}

func (s *Shaders) Cleanup() {
	rl.UnloadShader(s.Lighting)
	rl.UnloadShader(s.PostProcess)
}

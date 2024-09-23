package main

import (
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SHADER_RENDER_FS = "shaders/render.fs"
	SHADER_PLAYER_VS = "shaders/player.vs"
	SHADER_PLAYER_FS = "shaders/player.fs"
	SHADER_CURSOR_FS = "shaders/cursor.fs"
	SHADER_CURSOR_VS = "shaders/cursor.vs"
)

type Shaders struct {
	Render rl.Shader
	Player rl.Shader
	Cursor rl.Shader
}

func NewShaders() *Shaders {
	return &Shaders{}
}

func (s *Shaders) Setup() {
	s.Render = rl.LoadShader("", SHADER_RENDER_FS)
	s.Player = rl.LoadShader(SHADER_PLAYER_VS, SHADER_PLAYER_FS)
	s.Cursor = rl.LoadShader(SHADER_CURSOR_VS, SHADER_CURSOR_FS)
}

func (s *Shaders) Update() {
	playerTimeLoc := rl.GetShaderLocation(s.Player, "u_time")
	if playerTimeLoc == -1 {
		log.Fatalf("Could not find u_time uniform in shader: %s", SHADER_PLAYER_FS)
	}
	time := float32(rl.GetTime())
	rl.SetShaderValue(s.Player, playerTimeLoc, []float32{time}, rl.ShaderUniformFloat)
}

func (s *Shaders) Cleanup() {
	rl.UnloadShader(s.Player)
	rl.UnloadShader(s.Render)
	rl.UnloadShader(s.Cursor)
}

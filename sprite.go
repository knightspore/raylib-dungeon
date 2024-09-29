package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	src    rl.Rectangle
	dest   rl.Rectangle
	origin rl.Vector2
	rot    float32

	fps   int32
	frame int
	timer float32

	Color   rl.Texture2D
	Normal  rl.Texture2D
	Shaders map[string]rl.Shader
}

func NewSprite(baseSize float32, x float32, y float32) *Sprite {
	return &Sprite{
		src:     rl.NewRectangle(0, 0, 0, 0),
		dest:    rl.NewRectangle(x, y, baseSize, baseSize),
		origin:  rl.NewVector2(0, 0),
		rot:     0,
		frame:   0,
		timer:   0,
		Shaders: make(map[string]rl.Shader),
	}
}

func (s *Sprite) Setup(colorTex string, normalTex string, fps int32, shaders map[string]rl.Shader) {
	s.fps = fps
	if colorTex != "" {
		s.Color = rl.LoadTexture(colorTex)
	}
	if normalTex != "" {
		s.Normal = rl.LoadTexture(normalTex)
	}
	s.src.Width = float32(s.Color.Width / s.fps)
	s.src.Height = float32(s.Color.Height)
	if shaders != nil {
		s.Shaders = shaders
	}
}

func (s *Sprite) Cleanup() {
	if s.Color.ID > 0 {
		rl.UnloadTexture(s.Color)
	}
	if s.Normal.ID > 0 {
		rl.UnloadTexture(s.Normal)
	}
	for _, shader := range s.Shaders {
		rl.UnloadShader(shader)
	}
}

func (s *Sprite) Update() {
	s.timer += rl.GetFrameTime()
	if s.timer >= (1.0 / float32(4)) {
		s.frame = (s.frame + 1) % 4
		s.timer = 0
	}
	s.src.X = float32(s.frame * int(s.Color.Width))
}

func (s *Sprite) Draw(tex rl.Texture2D) {
	rl.DrawTexturePro(tex, s.src, s.dest, s.origin, s.rot, rl.White)
	if DEBUG {
		DrawDebugSprite(s, rl.Red)
	}
}

func (s *Sprite) Pos() rl.Vector2 {
	return rl.NewVector2(s.dest.X, s.dest.Y)
}

func (s *Sprite) Center() rl.Vector2 {
	return rl.NewVector2(s.dest.X+s.dest.Width/2, s.dest.Y+s.dest.Height/2)
}

func (s *Sprite) UpdateShaderValue(name string, value string, data []float32, uniformType rl.ShaderUniformDataType) {
	loc := rl.GetShaderLocation(s.Shaders[name], value)
	rl.SetShaderValue(s.Shaders[name], loc, data, uniformType)
}

func (s *Sprite) SetDest(dest rl.Vector2) {
	s.dest.X = dest.X
	s.dest.Y = dest.Y
}

func (s *Sprite) SetOrigin(origin rl.Vector2) {
	s.origin = origin
}

func (s *Sprite) SetRot(rot float32) {
	s.rot = rot
}

package main

import rl "github.com/gen2brain/raylib-go/raylib"

type SpriteTex struct {
	color  rl.Texture2D
	normal rl.Texture2D
}

func NewSpriteTex(color, normal string) *SpriteTex {
	return &SpriteTex{
		color:  rl.LoadTexture(color),
		normal: rl.LoadTexture(normal),
	}
}

func (s *SpriteTex) Cleanup() {
	rl.UnloadTexture(s.color)
	rl.UnloadTexture(s.normal)
}

type Sprite struct {
	src    rl.Rectangle
	dest   rl.Rectangle
	origin rl.Vector2
	rot    float32
	frame  int
	timer  float32
	tex    *SpriteTex
}

func NewSprite(baseSize float32, colorTex string, normalTex string, fps int32) *Sprite {
	tex := NewSpriteTex(colorTex, normalTex)
	return &Sprite{
		src:    rl.NewRectangle(0, 0, float32(tex.color.Width/fps), float32(tex.color.Height)),
		dest:   rl.NewRectangle(0, 0, baseSize, baseSize),
		origin: rl.NewVector2(baseSize/2, baseSize/2),
		tex:    tex,
	}
}

func (s *Sprite) Center() rl.Vector2 {
	return rl.NewVector2(s.dest.X+s.dest.Width/2, s.dest.Y+s.dest.Height/2)
}

func (s *Sprite) SetDest(dest *rl.Vector2) {
	s.dest.X = dest.X
	s.dest.Y = dest.Y
}

func (s *Sprite) SetRot(rot float32) {
	s.rot = rot
}

func (s *Sprite) Update(dest *rl.Vector2, rot *float32) {
	// Animation
	s.timer += rl.GetFrameTime()
	if s.timer >= (1.0 / float32(4)) {
		s.frame = (s.frame + 1) % 4
		s.timer = 0
	}
	s.src.X = float32(s.frame * int(s.tex.color.Width))
}

func (s *Sprite) Draw() {
	rl.DrawTexturePro(s.tex.color, s.src, s.dest, s.origin, s.rot, rl.White)
	if DEBUG {
		rl.DrawCircle(int32(s.Center().X), int32(s.Center().Y), 3, rl.Red)
	}
}

func (s *Sprite) DrawNormal() {
	rl.DrawTexturePro(s.tex.normal, s.src, s.dest, s.origin, s.rot, rl.White)
}

func (s *Sprite) Cleanup() {
	s.tex.Cleanup()
}

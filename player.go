package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Pos    rl.Vector2
	Size   int32
	Speed  float32
	Target rl.Vector2
	Cursor struct {
		Src    rl.Rectangle
		Dest   rl.Rectangle
		Origin rl.Vector2
	}
}

func NewPlayer(pos rl.Vector2, size int32) *Player {
	return &Player{
		Pos:   rl.Vector2{X: pos.X - float32(size)/2, Y: pos.Y - float32(size)/2},
		Size:  size,
		Speed: 2.0,
	}
}

func (p *Player) Center() rl.Vector2 {
	return rl.NewVector2(p.Pos.X+float32(p.Size)/2, p.Pos.Y+float32(p.Size)/2)
}

func (p *Player) DrawCursor(g *Game) {
	rot := rl.GetTime() * 100
	rl.BeginShaderMode(g.Shaders.Cursor)
	rl.DrawTexturePro(g.Textures.Cursor, p.Cursor.Src, p.Cursor.Dest, p.Cursor.Origin, float32(rot), rl.White)
	rl.EndShaderMode()
}

func (p *Player) Draw(g *Game) {
	rl.BeginShaderMode(g.Shaders.Player)
	rl.DrawTextureEx(g.Textures.Player, p.Pos, 0, float32(p.Size/g.Textures.Player.Width), rl.White)
	rl.EndShaderMode()
	rl.DrawLine(int32(p.Center().X), int32(p.Center().Y), int32(p.Target.X), int32(p.Target.Y), rl.Green)
}

func (p *Player) MoveCursor(g *Game) {
	pos := rl.GetWorldToScreen2D(rl.GetMousePosition(), *g.Cam.Cam)

	p.Cursor.Src = rl.NewRectangle(
		float32(g.CurrentFrame*int(g.Textures.Cursor.Height)),
		0,
		float32(g.Textures.Cursor.Height),
		float32(g.Textures.Cursor.Height),
	)

	p.Cursor.Dest = rl.NewRectangle(
		pos.X-float32(g.BaseSize/2),
		pos.Y-float32(g.BaseSize/2),
		float32(g.BaseSize),
		float32(g.BaseSize),
	)

	p.Cursor.Origin = raylib.Vector2{X: float32(g.BaseSize / 2), Y: float32(g.BaseSize / 2)}

	p.Target = rl.GetScreenToWorld2D(rl.Vector2{X: p.Cursor.Dest.X, Y: p.Cursor.Dest.Y}, *g.Cam.Cam)

}

func (p *Player) MovePlayer(m *Map) {
	nextPos := p.Pos

	if rl.IsKeyDown(rl.KeyW) {
		nextUp := rl.NewVector2(nextPos.X, nextPos.Y-p.Speed)
		if !m.CheckCollision(nextUp, p.Size) {
			nextPos = nextUp
		}
	}

	if rl.IsKeyDown(rl.KeyS) {
		nextDown := rl.NewVector2(nextPos.X, nextPos.Y+p.Speed)
		if !m.CheckCollision(nextDown, p.Size) {
			nextPos = nextDown
		}
	}

	if rl.IsKeyDown(rl.KeyA) {
		nextLeft := rl.NewVector2(nextPos.X-p.Speed, nextPos.Y)
		if !m.CheckCollision(nextLeft, p.Size) {
			nextPos = nextLeft
		}
	}

	if rl.IsKeyDown(rl.KeyD) {
		nextRight := rl.NewVector2(nextPos.X+p.Speed, nextPos.Y)
		if !m.CheckCollision(nextRight, p.Size) {
			nextPos = nextRight
		}
	}

	if m.CheckOutOfBounds(nextPos, p.Size) {
		return
	}

	p.Pos = nextPos
}

func (p *Player) Update(g *Game) {
	p.MovePlayer(g.Map)
	p.MoveCursor(g)
}

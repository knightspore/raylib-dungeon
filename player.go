package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Pos    rl.Vector2
	Size   int32
	Speed  float32
	Source rl.Rectangle
	Cursor struct {
		Source rl.Rectangle
		Dest   rl.Rectangle
		Origin rl.Vector2
	}
}

func NewPlayer(pos rl.Vector2, size int32) *Player {
	return &Player{
		Pos:    rl.Vector2{X: pos.X - float32(size)/2, Y: pos.Y - float32(size)/2},
		Size:   size,
		Speed:  float32(size / 24),
		Source: rl.NewRectangle(0, 0, float32(size), float32(size)),
	}
}

func (p *Player) Center() rl.Vector2 {
	return rl.NewVector2(p.Pos.X+float32(p.Size/2), p.Pos.Y+float32(p.Size/2))
}

func (p *Player) CursorCenter() rl.Vector2 {
	return rl.NewVector2(p.Cursor.Dest.X+float32(p.Size/2), p.Cursor.Dest.Y+float32(p.Size/2))
}

func (p *Player) DrawCursor(g *Game, normal bool) {
	rot := rl.GetTime() * 40
	rl.DrawTexturePro(g.Textures.Cursor, p.Cursor.Source, p.Cursor.Dest, p.Cursor.Origin, float32(rot), rl.White)
}

func (p *Player) Draw(g *Game, normal bool) {
	dest := rl.NewRectangle(p.Pos.X, p.Pos.Y, float32(p.Size), float32(p.Size))

	// Draw Player
	rl.BeginShaderMode(g.Shaders.Player)
	playerTimeLoc := rl.GetShaderLocation(g.Shaders.Player, "u_time")
	time := float32(rl.GetTime())
	rl.SetShaderValue(g.Shaders.Player, playerTimeLoc, []float32{time}, rl.ShaderUniformFloat)
	if normal {
		rl.DrawTexturePro(g.Textures.Player_Normal, p.Source, dest, rl.NewVector2(0, 0), 0, rl.White)
	} else {
		rl.DrawTexturePro(g.Textures.Player, p.Source, dest, rl.NewVector2(0, 0), 0, rl.White)
	}
	rl.EndShaderMode()
}

func (p *Player) UpdateCursorPosition(g *Game) {
	pos := rl.GetMousePosition()

	p.Cursor.Source = rl.NewRectangle(
		float32(g.CurrentFrame*int(g.Textures.Cursor.Height)),
		0,
		float32(g.Textures.Cursor.Height),
		float32(g.Textures.Cursor.Height),
	)

	halfSize := float32(g.BaseSize / 2)

	p.Cursor.Dest = rl.NewRectangle(pos.X-halfSize, pos.Y-halfSize, float32(g.BaseSize), float32(g.BaseSize))
	p.Cursor.Origin = raylib.Vector2{X: halfSize, Y: halfSize}
}

func (p *Player) UpdatePlayerPosition(g *Game) {
	nextPos := p.Pos

	if rl.IsKeyDown(rl.KeyLeftShift) && p.Speed == (float32(p.Size)/24) {
		p.Speed = (float32(p.Size) / 12)
	}

	if rl.IsKeyUp(rl.KeyLeftShift) && p.Speed > (float32(p.Size)/24) {
		p.Speed -= 0.1
	}

	if rl.IsKeyDown(rl.KeyW) {
		nextUp := rl.NewVector2(nextPos.X, nextPos.Y-p.Speed)
		if !g.Map.CheckCollision(nextUp) {
			nextPos = nextUp
		}
	}

	if rl.IsKeyDown(rl.KeyS) {
		nextDown := rl.NewVector2(nextPos.X, nextPos.Y+p.Speed)
		if !g.Map.CheckCollision(nextDown) {
			nextPos = nextDown
		}
	}

	if rl.IsKeyDown(rl.KeyA) {
		nextLeft := rl.NewVector2(nextPos.X-p.Speed, nextPos.Y)
		if !g.Map.CheckCollision(nextLeft) {
			nextPos = nextLeft
		}
	}

	if rl.IsKeyDown(rl.KeyD) {
		nextRight := rl.NewVector2(nextPos.X+p.Speed, nextPos.Y)
		if !g.Map.CheckCollision(nextRight) {
			nextPos = nextRight
		}
	}

	if g.Map.CheckOutOfBounds(nextPos) {
		return
	}

	p.Pos = nextPos

	// Animation

	p.Source = rl.NewRectangle(
		float32(g.CurrentFrame*int(p.Size)),
		0,
		float32(g.Textures.Player.Height),
		float32(g.Textures.Player.Height),
	)
}

func (p *Player) Update(g *Game) {
	p.UpdatePlayerPosition(g)
	p.UpdateCursorPosition(g)
}

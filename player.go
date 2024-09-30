package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Speed  float32 // can be moved to sprite
	Sprite *Sprite
}

func NewPlayer(pos rl.Vector2, size int32) *Player {
	return &Player{
		Speed:  float32(size / 24),
		Sprite: NewSprite(float32(size), pos.X-float32(size)/2, pos.Y-float32(size)/2),
	}
}

func (p *Player) Setup() {
	p.Sprite.Setup(
		"textures/player_tilesheet.png",
		"textures/player_tilesheet_n.png",
		4,
		map[string]rl.Shader{"player": rl.LoadShader("shaders/player.vs", "shaders/player.fs")},
	)
}

func (p *Player) Cleanup() {
	p.Sprite.Cleanup()
}

func (p *Player) Draw(tex rl.Texture2D) {
	rl.BeginShaderMode(p.Sprite.Shaders["player"])
	p.Sprite.UpdateShaderValue("player", "u_time", []float32{float32(rl.GetTime())}, rl.ShaderUniformFloat)
	p.Sprite.Draw(tex)
	rl.EndShaderMode()
}

func (p *Player) HandleMovement(nextPos rl.Vector2, g *Game) (bool, rl.Vector2) {
	// This can be optimized the most
	// ie. split into sprite and player

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
		return true, nextPos
	}

	return false, nextPos
}

func (p *Player) updatePosition(g *Game) {
	if outOfBounds, nextPos := p.HandleMovement(p.Sprite.Pos(), g); !outOfBounds {
		p.Sprite.SetDest(nextPos)
	}
}

func (p *Player) Update(g *Game) {
	p.Sprite.Animate()
	p.updatePosition(g)
}

func (p *Player) Center() rl.Vector2 {
	return p.Sprite.Center()
}

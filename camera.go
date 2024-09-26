package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	Cam         *rl.Camera2D
	ZoomSpeed   float32
	FollowSpeed float32
}

func NewCamera(offset rl.Vector2, target rl.Vector2) *Camera {
	return &Camera{
		Cam: &rl.Camera2D{
			Offset:   offset,
			Target:   target,
			Rotation: 0,
			Zoom:     1,
		},
		ZoomSpeed:   0.001,
		FollowSpeed: 0.1,
	}
}

func (c *Camera) UpdateZoom() {
	if rl.IsKeyDown(rl.KeyRightBracket) {
		if c.Cam.Zoom < 2.0 {
			c.Cam.Zoom += c.ZoomSpeed
			c.ZoomSpeed += 0.001
		}
	}
	if rl.IsKeyDown(rl.KeyLeftBracket) {
		if c.Cam.Zoom > 0.5 {
			c.Cam.Zoom -= c.ZoomSpeed
			c.ZoomSpeed += 0.001
		}
	}
}

func (c *Camera) UpdateSmoothFollow(g *Game) {
	target := rl.Vector2{X: g.Player.Center().X + g.Player.CursorCenter().X, Y: g.Player.Center().Y + g.Player.CursorCenter().Y}
	c.Cam.Target.X += (target.X/2 - c.Cam.Target.X) * c.FollowSpeed
	c.Cam.Target.Y += (target.Y/2 - c.Cam.Target.Y) * c.FollowSpeed
}

func (c *Camera) Update(g *Game) {
	c.UpdateZoom()
	c.UpdateSmoothFollow(g)
}

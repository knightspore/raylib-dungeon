package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	Cam         *rl.Camera2D
	ZoomSpeed   float32
	FollowSpeed float32
}

func NewCamera() *Camera {
	return &Camera{
		Cam: &rl.Camera2D{
			Offset:   rl.Vector2{X: WIDTH / 2, Y: HEIGHT / 2},
			Target:   rl.Vector2{X: WIDTH / 2, Y: HEIGHT / 2},
			Rotation: 0,
			Zoom:     1,
		},
		ZoomSpeed:   0.001,
		FollowSpeed: 0.1,
	}
}

func (c *Camera) Update(g *Game) {
	c.updateZoom()
	c.updateSmoothFollow(g)
}

func (c *Camera) updateZoom() {
	if rl.IsKeyDown(rl.KeyRightBracket) {
		if c.Cam.Zoom < 1.2 {
			c.Cam.Zoom += c.ZoomSpeed
			c.ZoomSpeed += 0.001
		} else if DEBUG {
			c.Cam.Zoom += c.ZoomSpeed
		}
	}
	if rl.IsKeyDown(rl.KeyLeftBracket) {
		if c.Cam.Zoom > 0.8 {
			c.Cam.Zoom -= c.ZoomSpeed
			c.ZoomSpeed += 0.001
		} else if DEBUG {
			c.Cam.Zoom -= c.ZoomSpeed
		}
	}
}

func (c *Camera) updateSmoothFollow(g *Game) {
	target := rl.Vector2{X: g.Player.Center().X + g.Cursor.Center().X, Y: g.Player.Center().Y + g.Cursor.Center().Y}
	c.Cam.Target.X += (target.X/2 - c.Cam.Target.X) * c.FollowSpeed
	c.Cam.Target.Y += (target.Y/2 - c.Cam.Target.Y) * c.FollowSpeed
}

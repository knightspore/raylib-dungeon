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

func (c *Camera) Zoom() {
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

func (c *Camera) SmoothFollow(target rl.Vector2) {
	c.Cam.Target.X += (target.X - c.Cam.Target.X) * c.FollowSpeed
	c.Cam.Target.Y += (target.Y - c.Cam.Target.Y) * c.FollowSpeed
}

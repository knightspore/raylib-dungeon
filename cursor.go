package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Cursor struct {
	Sprite *Sprite
}

func NewCursor(baseSize, x, y float32) *Cursor {
	return &Cursor{
		Sprite: NewSprite(baseSize, x, y),
	}
}

func (c *Cursor) Setup() {
	c.Sprite.Setup(
		"textures/cursor_tilesheet.png",
		"textures/cursor_tilesheet_n.png",
		4,
		nil,
	)
	c.Sprite.SetOrigin(rl.NewVector2(c.Sprite.dest.Width/2, c.Sprite.dest.Height/2))
}

func (c *Cursor) Cleanup() {
	c.Sprite.Cleanup()
}

func (c *Cursor) Update() {
	c.Sprite.Update()
	c.Sprite.SetDest(rl.Vector2{X: float32(rl.GetMousePosition().X - c.Sprite.dest.Width/2), Y: float32(rl.GetMousePosition().Y - c.Sprite.dest.Height/2)})
	c.Sprite.SetRot(float32(rl.GetTime() * 40))
}

func (c *Cursor) Draw(tex rl.Texture2D) {
	c.Sprite.Draw(tex)
}

func (c *Cursor) Center() rl.Vector2 {
	return c.Sprite.Pos()
}

package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Cursor struct {
	Sprite *Sprite
}

func NewCursor() *Cursor {
	return &Cursor{
		Sprite: NewSprite(BASE_SIZE, WIDTH/2, HEIGHT/2),
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
	c.Sprite.Animate()
	c.Sprite.SetDest(rl.Vector2{X: float32(rl.GetMousePosition().X - c.Sprite.dest.Width/2), Y: float32(rl.GetMousePosition().Y - c.Sprite.dest.Height/2)})
	c.Sprite.SetRot(float32(rl.GetTime() * 40))
}

func (c *Cursor) Draw() {
	c.Sprite.Draw()
}

func (c *Cursor) DrawNormal() {
	c.Sprite.DrawNormal()
}

func (c *Cursor) DrawDebug() {
	DrawDebugSprite(c.Sprite)
}

func (c *Cursor) Center() rl.Vector2 {
	return c.Sprite.Pos()
}

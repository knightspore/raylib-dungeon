package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	texture_color = "textures/cursor_tilesheet.png"
	texure_normal = "textures/cursor_tilesheet_n.png"
	animation_fps = 4
)

type Cursor struct {
	source       rl.Rectangle
	dest         rl.Rectangle
	origin       rl.Vector2
	rotation     float32
	textures     map[string]rl.Texture2D
	frameTimer   float32
	currentFrame int
}

func NewCursor(baseSize float32) *Cursor {
	return &Cursor{
		source:   rl.NewRectangle(0, 0, 0, 0),
		dest:     rl.NewRectangle(0, 0, baseSize, baseSize),
		origin:   rl.NewVector2(baseSize/2, baseSize/2),
		textures: make(map[string]rl.Texture2D),
	}
}

func (c *Cursor) Setup() {
	c.textures["color"] = rl.LoadTexture(texture_color)
	c.textures["normal"] = rl.LoadTexture(texure_normal)

	c.source.Width = float32(c.textures["color"].Width / animation_fps)
	c.source.Height = float32(c.textures["color"].Height)
}

func (c *Cursor) Cleanup() {
	rl.UnloadTexture(c.textures["color"])
	rl.UnloadTexture(c.textures["normal"])
}

func (c *Cursor) updateAnimation() {
	c.frameTimer += rl.GetFrameTime()
	if c.frameTimer >= (1.0 / float32(animation_fps)) {
		c.currentFrame = (c.currentFrame + 1) % animation_fps
		c.frameTimer = 0
	}
	c.source.X = float32(c.currentFrame * int(c.textures["color"].Width))
}

func (c *Cursor) updatePosition() {
	c.rotation = float32(rl.GetTime() * animation_fps * 10)
	c.dest.X, c.dest.Y = float32(rl.GetMousePosition().X-c.dest.Height/2), float32(rl.GetMousePosition().Y-c.dest.Height/2)
}

func (c *Cursor) Update() {
	c.updateAnimation()
	c.updatePosition()
}

func (c *Cursor) Draw() {
	rl.DrawTexturePro(c.textures["color"], c.source, c.dest, c.origin, c.rotation, rl.White)

	if DEBUG {
		rl.DrawCircle(int32(c.Center().X), int32(c.Center().Y), 5, rl.Red)
	}
}

func (c *Cursor) DrawNormal() {
	rl.DrawTexturePro(c.textures["normal"], c.source, c.dest, c.origin, c.rotation, rl.White)
}

func (c *Cursor) Center() rl.Vector2 {
	return rl.NewVector2(c.dest.X, c.dest.Y)
}

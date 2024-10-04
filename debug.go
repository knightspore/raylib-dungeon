package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var DEBUG = false

func DrawDebugArea(dest rl.Rectangle, center rl.Vector2, color rl.Color) {
	color = rl.Fade(color, 0.75)
	rl.DrawCircle(int32(center.X), int32(center.Y), 1, color)
	rl.DrawRectangleLinesEx(dest, 1, color)
	rl.DrawLineEx(rl.NewVector2(dest.X, dest.Y), rl.NewVector2(dest.X+dest.Width, dest.Y+dest.Height), 1, color)
}

func DrawDebugSprite(sprite *Sprite) {
	rl.DrawCircle(int32(sprite.Center().X), int32(sprite.Center().Y), 1, rl.Red)
	dest := rl.Rectangle{X: sprite.dest.X - sprite.origin.X, Y: sprite.dest.Y - sprite.origin.Y, Width: sprite.dest.Width, Height: sprite.dest.Height}
	rl.DrawRectangleLinesEx(dest, 1, rl.Red)
	rl.DrawLineEx(rl.NewVector2(dest.X, dest.Y), rl.NewVector2(dest.X+dest.Width, dest.Y+dest.Height), 1, rl.Red)
}

func DrawDebugLine(start rl.Vector2, end rl.Vector2) {
	rl.DrawLineEx(start, end, 1, rl.Blue)
}

func DrawDebugParticles(particles *map[int]*Particle) {
	for _, p := range *particles {
		rl.DrawCircleLines(int32(p.pos.X), int32(p.pos.Y), 10, rl.Red)
	}
}

func UpdateDebug() {
	if rl.IsKeyPressed(rl.KeyF1) {
		DEBUG = !DEBUG
	}
	if rl.IsKeyPressed(rl.KeyF2) {
		rl.TakeScreenshot(fmt.Sprintf("screenshots/screen_%d.png", time.Now().Unix()))
	}
	if rl.IsKeyPressed(rl.KeyF3) {
		rl.ToggleFullscreen()
	}
}

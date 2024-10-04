package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadGameFromImage(path string) *Game {

	image := rl.LoadImage(path)

	tiles := []int{}
	playerPos := rl.NewVector2(0, 0)
	endPos := rl.NewVector2(0, 0)
	lightPositions := []rl.Vector2{}

	for y := 0; y < int(image.Height); y++ {
		for x := 0; x < int(image.Width); x++ {
			color := rl.GetImageColor(*image, int32(x), int32(y))

			// Read basic map data
			if color.R == 255 && color.G == 0 && color.B == 0 { // Empty space
				tiles = append(tiles, TILE_EMPTY)
			} else { // Defaulting to floor for now
				tiles = append(tiles, TILE_FLOOR)
			}

			// Read player start / end positions
			if color.R == 0 && color.G == 255 && color.B == 0 {
				playerPos = rl.NewVector2(float32(x*BASE_SIZE+BASE_SIZE/2), float32(y*BASE_SIZE+BASE_SIZE/2))
			}

			if color.R == 0 && color.G == 0 && color.B == 255 {
				endPos = rl.NewVector2(float32(x*BASE_SIZE+BASE_SIZE/2), float32(y*BASE_SIZE+BASE_SIZE/2))
			}

			// Read light positions
			if color.R == 255 && color.G == 255 && color.B == 0 {
				lightPositions = append(lightPositions, rl.NewVector2(float32(x*BASE_SIZE+BASE_SIZE/2), float32(y*BASE_SIZE+BASE_SIZE/2)))
			}
		}
	}

	lights := &Lights{}
	for _, pos := range lightPositions {
		lights.Add(pos.X, pos.Y, 50, rl.NewColor(230, 230, 100, 255))
	}
	lights.Add(endPos.X, endPos.Y, 50, rl.NewColor(255, 0, 0, 255))

	rl.UnloadImage(image)

	game := &Game{
		Width:    WIDTH,
		Height:   HEIGHT,
		BaseSize: BASE_SIZE,
		Textures: &Textures{},
		Shaders:  &Shaders{},
		Cam:      NewCamera(rl.NewVector2(float32(WIDTH/2), float32(HEIGHT/2)), rl.NewVector2(float32(WIDTH/2), float32(HEIGHT/2))),
		Map:      NewMap(tiles, BASE_SIZE),
		Player:   NewPlayer(playerPos, BASE_SIZE),
		Cursor:   NewCursor(float32(BASE_SIZE), float32(WIDTH/2), float32(HEIGHT/2)),
		Lights:   lights,
		Emitter:  NewEmitter(200, rl.NewRectangle(0, 0, float32(BASE_SIZE*int(image.Width)), float32(BASE_SIZE*int(image.Height))), 10),
	}

	return game
}

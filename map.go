package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var BG_COLOR = rl.Color{R: 40, G: 30, B: 44, A: 255}

type Map struct {
	sizeX    int32
	sizeY    int32
	tileSize int32
	tiles    map[int]*Tile
}

func NewMap(tiles []int, baseSize int32) *Map {
	dim := int32(math.Sqrt(float64(len(tiles))))
	return &Map{
		sizeX:    dim,
		sizeY:    dim,
		tileSize: baseSize,
		tiles:    CreateTiles(tiles, baseSize, dim, dim),
	}
}

func (m *Map) Setup() {
	for _, tile := range m.tiles {
		tile.Setup()
		northIdx := tile.index - int(m.sizeX)
		if northIdx >= 0 && m.tiles[northIdx]._type == TILE_FLOOR {
			tile.sprite.src.X = float32(tile.sprite.Color.Height)
		}
	}
}

func (m *Map) Cleanup() {
	for _, tile := range m.tiles {
		tile.Cleanup()
	}
}

func (m *Map) Draw() {
	rl.ClearBackground(BG_COLOR)
	for i := 0; i < int(m.sizeX*m.sizeY); i++ {
		tile := m.tiles[i]
		if tile._type == TILE_FLOOR {
			tile.sprite.Draw(tile.sprite.Color)
		}
		if tile._type == TILE_EMPTY {
			tile.sprite.Draw(tile.sprite.Color)
		}
	}
}

func (m *Map) DrawNormal() {
	for i := 0; i < int(m.sizeX*m.sizeY); i++ {
		tile := m.tiles[i]
		if tile._type == TILE_FLOOR {
			tile.sprite.Draw(tile.sprite.Normal)
		}
		if tile._type == TILE_EMPTY {
			tile.sprite.Draw(tile.sprite.Normal)
		}
	}
}

func (m *Map) vec2Tile(x, y float32) int {
	X := int32(x) / m.tileSize
	Y := int32(y) / m.tileSize
	return int(X + Y*m.sizeX)
}

func (m *Map) getBoundingTiles(v rl.Vector2) []*Tile {
	adjustedSize := float32(m.tileSize - 1) // Figure out a margin that works
	// regardless of the speed:tileSize ratio

	nw := m.tiles[m.vec2Tile(v.X, v.Y)]
	ne := m.tiles[m.vec2Tile(v.X+adjustedSize, v.Y)]
	sw := m.tiles[m.vec2Tile(v.X, v.Y+adjustedSize)]
	se := m.tiles[m.vec2Tile(v.X+adjustedSize, v.Y+adjustedSize)]

	return []*Tile{nw, ne, sw, se}
}

func (m *Map) checkCollision(v rl.Vector2) bool {
	if v.X < 0 || v.Y < 0 || v.X >= float32(m.tileSize)*float32(m.sizeX) || v.Y >= float32(m.tileSize)*float32(m.sizeY) {
		return true
	}

	for _, tile := range m.getBoundingTiles(v) {
		if tile == nil || tile._type == TILE_EMPTY {
			return true
		}
	}

	return false
}

package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Map struct {
	SizeX     int32
	SizeY     int32
	TileSize  int32
	TileCount int
	Tiles     []int
}

func NewMap(tiles []int, baseSize int32) *Map {
	dim := int32(math.Sqrt(float64(len(tiles))))
	return &Map{
		SizeX:     dim,
		SizeY:     dim,
		TileSize:  baseSize,
		TileCount: len(tiles),
		Tiles:     tiles,
	}
}

func (m *Map) Center() rl.Vector2 {
	return rl.NewVector2(float32(m.SizeX*m.TileSize)/2, float32(m.SizeY*m.TileSize)/2)
}

func (m *Map) Vec2Tile(x, y float32) int32 {
	X := int32(x) / m.TileSize
	Y := int32(y) / m.TileSize
	return X + Y*m.SizeX
}

func (m *Map) GetEntityTiles(v rl.Vector2, size int32) []int32 {
	adjustedSize := float32(size - 1)

	nw := m.Vec2Tile(v.X, v.Y)
	ne := m.Vec2Tile(v.X+adjustedSize, v.Y)
	sw := m.Vec2Tile(v.X, v.Y+adjustedSize)
	se := m.Vec2Tile(v.X+adjustedSize, v.Y+adjustedSize)

	return []int32{nw, ne, sw, se}
}

func (m *Map) CheckBounds(v rl.Vector2) bool {
	return v.X < 0 || v.Y < 0 || v.X > float32(m.SizeX*m.TileSize) || v.Y > float32(m.SizeY*m.TileSize)
}

func (m *Map) CheckOutOfBounds(v rl.Vector2, size int32) bool {
	tiles := m.GetEntityTiles(v, size)
	for _, tile := range tiles {
		if tile < 0 || tile >= int32(m.TileCount) {
			return true
		}
	}
	return false
}

func (m *Map) CheckCollision(v rl.Vector2, size int32) bool {
	tiles := m.GetEntityTiles(v, size)
	if m.CheckOutOfBounds(v, size) {
		return true
	}
	for _, tile := range tiles {
		if m.Tiles[tile] == 1 {
			return true
		}
	}
	return false
}

func (m *Map) DrawTile(tile int, x, y float32, floor rl.Texture2D) {
	rl.ClearBackground(rl.SkyBlue)
	if tile == 0 {
		rl.DrawTextureEx(floor, rl.NewVector2(x, y), 0, float32(m.TileSize)/float32(floor.Width), rl.White)
	}
}

func (m *Map) Draw(g *Game) {
	for i := 0; i < m.TileCount; i++ {
		tile := m.Tiles[i]
		x := float32((int32(i) % m.SizeX) * m.TileSize)
		y := float32((int32(i) / m.SizeY) * m.TileSize)
		m.DrawTile(tile, x, y, g.Textures.Floor)
	}
}

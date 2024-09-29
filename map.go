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

func (m *Map) Vec2TileV(v rl.Vector2) int32 {
	return m.Vec2Tile(v.X, v.Y)
}

func (m *Map) Vec2Tile(x, y float32) int32 {
	X := int32(x) / m.TileSize
	Y := int32(y) / m.TileSize
	return X + Y*m.SizeX
}

func (m *Map) Tile2Sprite(tile int32) *Sprite {
	x := (tile % m.SizeX) * m.TileSize
	y := (tile / m.SizeY) * m.TileSize
	sprite := NewSprite(float32(m.TileSize), float32(x), float32(y))
	sprite.Setup("", "", 1, nil)
	return sprite
}

func (m *Map) GetEntityTiles(v rl.Vector2) []int32 {
	adjustedSize := float32(m.TileSize - 1)

	nw := m.Vec2Tile(v.X, v.Y)
	ne := m.Vec2Tile(v.X+adjustedSize, v.Y)
	sw := m.Vec2Tile(v.X, v.Y+adjustedSize)
	se := m.Vec2Tile(v.X+adjustedSize, v.Y+adjustedSize)

	return []int32{nw, ne, sw, se}
}

func (m *Map) CheckOutOfBounds(v rl.Vector2) bool {
	tiles := m.GetEntityTiles(v)

	if v.X < 0 || v.Y < 0 || v.X >= float32(m.SizeX*m.TileSize) || v.Y >= float32(m.SizeY*m.TileSize) {
		return true
	}

	for _, tile := range tiles {
		if tile < 0 || tile >= int32(m.TileCount) {
			return true
		}
	}

	return false
}

func (m *Map) CheckCollision(v rl.Vector2) bool {
	if m.CheckOutOfBounds(v) {
		return true
	}

	tiles := m.GetEntityTiles(v)
	for _, tile := range tiles {
		if m.Tiles[tile] == 1 {
			return true
		}
	}

	return false
}

func (m *Map) Draw(g *Game, normal bool) {
	rl.ClearBackground(rl.NewColor(40, 30, 44, 255))
	for i := 0; i < m.TileCount; i++ {
		tile := m.Tiles[i]
		x := float32((int32(i) % m.SizeX) * m.TileSize)
		y := float32((int32(i) / m.SizeY) * m.TileSize)
		src := rl.NewRectangle(0, 0, float32(g.Textures.Floor.Height), float32(g.Textures.Floor.Height))
		dest := rl.NewRectangle(x, y, float32(m.TileSize), float32(m.TileSize))
		origin := rl.NewVector2(0, 0)
		if tile == 0 { // Draw floor tile
			if normal {
				rl.DrawTexturePro(g.Textures.Floor_Normal, src, dest, origin, 0, rl.White)
			} else {
				rl.DrawTexturePro(g.Textures.Floor, src, dest, origin, 0, rl.White)
			}

			if DEBUG && m.Vec2Tile(g.Cursor.Center().X, g.Cursor.Center().Y) == int32(i) {
				DrawDebugSprite(m.Tile2Sprite(int32(i)), rl.Red)
			}

		} else if tile == 1 { // Draw empty tiles
			northIdx := i - int(m.SizeX)
			if northIdx >= 0 && m.Tiles[northIdx] == 0 {
				src := rl.NewRectangle(float32(g.Textures.NoFloor.Height), 0, float32(g.Textures.NoFloor.Height), float32(g.Textures.NoFloor.Height))
				rl.DrawTexturePro(g.Textures.NoFloor, src, dest, origin, 0, rl.White)
			} else {
				rl.DrawTexturePro(g.Textures.NoFloor, src, dest, origin, 0, rl.White)
			}
		}
	}
}

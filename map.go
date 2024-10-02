package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var BG_COLOR = rl.Color{R: 40, G: 30, B: 44, A: 255}

type Map struct {
	sizeX     int32
	sizeY     int32
	tileSize  int32
	tileSet   map[int]*Tile
	tileOrder []int
	sprite    *Sprite
}

func NewMap(tileOrder []int, baseSize int32) *Map {
	dim := int32(math.Sqrt(float64(len(tileOrder))))
	return &Map{
		sizeX:     dim,
		sizeY:     dim,
		tileSize:  baseSize,
		tileSet:   CreateTiles(float32(baseSize)),
		tileOrder: tileOrder,
	}
}

func (m *Map) Setup() {
	for _, tile := range m.tileSet {
		tile.Setup()
		northIdx := tile.index - int(m.sizeX)
		if northIdx >= 0 && m.tileSet[northIdx]._type == TILE_FLOOR {
			tile.sprite.src.X = float32(tile.sprite.Color.Height)
		}
	}

	mapSize := float32(math.Max(float64(m.sizeX), float64(m.sizeY)) * float64(m.tileSize))

	m.sprite = NewSprite(mapSize, 0, 0)
	colorTex := rl.LoadRenderTexture(int32(mapSize), int32(mapSize))
	normalTex := rl.LoadRenderTexture(int32(mapSize), int32(mapSize))

	rl.BeginTextureMode(colorTex)
	rl.ClearBackground(BG_COLOR)
	for i := 0; i < len(m.tileOrder); i++ {
		tile := m.tileSet[m.tileOrder[i]]
		falloff := i > int(m.sizeX) && tile._type == TILE_EMPTY &&
			m.tileSet[m.tileOrder[i-int(m.sizeX)]]._type == TILE_FLOOR
		tile.Draw(m.getTileDest(i), tile.sprite.Color, falloff)
	}
	rl.EndTextureMode()

	rl.BeginTextureMode(normalTex)
	rl.ClearBackground(rl.Blank)
	for i := 0; i < int(m.sizeX*m.sizeY); i++ {
		tile := m.tileSet[m.tileOrder[i]]
		falloff := i > int(m.sizeX) && tile._type == TILE_EMPTY &&
			m.tileSet[m.tileOrder[i-int(m.sizeX)]]._type == TILE_FLOOR
		tile.Draw(m.getTileDest(i), tile.sprite.Normal, falloff)
	}
	m.DrawNormal()
	rl.EndTextureMode()

	m.sprite.Color = colorTex.Texture
	m.sprite.Normal = normalTex.Texture
	m.sprite.src.Width = mapSize
	m.sprite.src.Height = -mapSize
}

func (m *Map) Cleanup() {
	for _, tile := range m.tileSet {
		tile.Cleanup()
	}
}

func (m *Map) Draw() {
	rl.ClearBackground(BG_COLOR)
	m.sprite.Draw(m.sprite.Color)
}

func (m *Map) DrawNormal() {
	m.sprite.Draw(m.sprite.Normal)
}

func (m *Map) getTileDest(i int) rl.Rectangle {
	return rl.NewRectangle(
		float32(i%int(m.sizeX)*int(m.tileSize)),
		float32(i/int(m.sizeY)*int(m.tileSize)),
		float32(m.tileSize),
		float32(m.tileSize),
	)
}

func (m *Map) vec2Tile(x, y float32) int {
	X := int32(x) / m.tileSize
	Y := int32(y) / m.tileSize
	return int(X + Y*m.sizeX)
}

func (m *Map) getBoundingTiles(v rl.Vector2) []*Tile {
	adjustedSize := float32(m.tileSize - 1) // TODO: Figure out a margin that works regardless of the speed:tileSize ratio

	nw := m.tileSet[m.tileOrder[m.vec2Tile(v.X, v.Y)]]
	ne := m.tileSet[m.tileOrder[m.vec2Tile(v.X+adjustedSize, v.Y)]]
	sw := m.tileSet[m.tileOrder[m.vec2Tile(v.X, v.Y+adjustedSize)]]
	se := m.tileSet[m.tileOrder[m.vec2Tile(v.X+adjustedSize, v.Y+adjustedSize)]]

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

func (m *Map) center() rl.Vector2 {
	return rl.NewVector2(float32(m.sizeX*m.tileSize)/2, float32(m.sizeY*m.tileSize)/2)
}

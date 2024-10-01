package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	TILE_FLOOR = iota
	TILE_EMPTY
)

var TILE_TYPE = map[int]string{
	TILE_FLOOR: "floor",
	TILE_EMPTY: "empty",
}

type Tile struct {
	_type  int
	size   float32
	sprite *Sprite
	index  int
}

func NewTile(index, t int, baseSize, x, y float32) *Tile {
	return &Tile{
		_type:  t,
		size:   baseSize,
		sprite: NewSprite(baseSize, x, y),
		index:  index,
	}
}

func CreateTiles(mapSlice []int, baseSize, sizeX, sizeY int32) map[int]*Tile {
	tiles := make(map[int]*Tile)
	for i, t := range mapSlice {
		x := float32(int32(i%int(sizeX)) * baseSize)
		y := float32(int32(i/int(sizeY)) * baseSize)
		tiles[i] = NewTile(i, t, float32(baseSize), x, y)
	}
	return tiles
}

func (t *Tile) Setup() {
	switch t._type {
	case TILE_FLOOR:
		t.sprite.Setup(
			"textures/floor_tilesheet.png",
			"textures/floor_tilesheet_n.png",
			1,
			nil)
	case TILE_EMPTY:
		t.sprite.Setup(
			"textures/notile.png",
			"textures/notile_n.png",
			1,
			nil,
		)
	}
	t.sprite.src = rl.NewRectangle(0, 0, float32(t.sprite.Color.Height), float32(t.sprite.Color.Height))
	t.sprite.dest = rl.NewRectangle(t.sprite.dest.X, t.sprite.dest.Y, t.size, t.size)
}

func (t *Tile) Cleanup() {
	t.sprite.Cleanup()
}

func (t *Tile) Draw(tex rl.Texture2D) {
	t.sprite.Draw(tex)
}

func (t *Tile) Center() rl.Vector2 {
	return t.sprite.Center()
}

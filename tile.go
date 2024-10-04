package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func NewTile(index, t int, baseSize float32) *Tile {
	return &Tile{
		_type:  t,
		size:   baseSize,
		sprite: NewSprite(baseSize, 0, 0),
		index:  index,
	}
}

func CreateTiles(baseSize float32) map[int]*Tile {
	tiles := make(map[int]*Tile)
	// should be an iteration
	tiles[TILE_FLOOR] = NewTile(0, TILE_FLOOR, baseSize)
	tiles[TILE_EMPTY] = NewTile(1, TILE_EMPTY, baseSize)
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

func (t *Tile) Draw(dest rl.Rectangle, tex rl.Texture2D, isFalloff bool) {
	if isFalloff {
		t.sprite.src.X = float32(t.sprite.Color.Height)
	} else {
		t.sprite.src.X = 0
	}
	t.sprite.dest = dest
	t.sprite.Draw()
}

func (t *Tile) Center() rl.Vector2 {
	return t.sprite.Center()
}

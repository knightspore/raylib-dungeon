package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
	pos      rl.Vector2
	velocity rl.Vector2
	life     float32
	size     float32
}

func NewParticle(size float32) *Particle {
	return &Particle{
		pos:      rl.NewVector2(0, 0),
		velocity: rl.NewVector2(0, 0),
		life:     0,
		size:     size,
	}
}

func (p *Particle) draw() {
	rl.DrawPixel(int32(p.pos.X), int32(p.pos.Y), rl.Fade(rl.Orange, p.life/p.size/2))
}

func (p *Particle) drawNormal() {
	rl.DrawPixel(int32(p.pos.X), int32(p.pos.Y), rl.Fade(rl.Orange, p.life/p.size))
}

func (p *Particle) Setup(area rl.Rectangle) {
	// position
	p.pos.X = float32(rl.GetRandomValue(int32(area.X), int32(area.X+area.Width)))
	p.pos.Y = float32(rl.GetRandomValue(int32(area.Y), int32(area.Y+area.Height)))
	// velocity
	p.velocity.X = 0.1 * float32(rl.GetRandomValue(-1, 1))
	p.velocity.Y = 0.1 * float32(rl.GetRandomValue(-1, 1))
	// life
	p.life = float32(rl.GetRandomValue(10, 30))
}

func (p *Particle) Update() {
	p.pos.X += p.velocity.X
	p.pos.Y += p.velocity.Y
	p.velocity.X -= 0.05 * float32(rl.GetRandomValue(-1, 1))
	p.velocity.Y -= 0.05 * float32(rl.GetRandomValue(-1, 1))
	p.life -= rl.GetFrameTime() * 10
}

type Emitter struct {
	rect      rl.Rectangle
	particles map[int]*Particle
}

func NewEmitter(particleCount int, area rl.Rectangle, size float32) *Emitter {
	particles := make(map[int]*Particle)
	for i := range particleCount {
		particles[i] = NewParticle(size)
	}
	return &Emitter{
		rect:      area,
		particles: particles,
	}
}

func (e *Emitter) Setup() {
	for i := range e.particles {
		rect := rl.Rectangle{X: e.rect.X - e.rect.Width/2, Y: e.rect.Y - e.rect.Height/2, Width: e.rect.Width * 2, Height: e.rect.Height * 2}
		e.particles[i].Setup(rect)
	}
}

func (e *Emitter) Cleanup() {
	for i := range e.particles {
		e.particles[i] = nil
	}
}

func (e *Emitter) Update() {
	for i := range e.particles {
		e.particles[i].Update()
		if e.particles[i].life <= 0 {
			e.particles[i].Setup(e.rect)
		}
	}
}

func (e *Emitter) Draw() {
	for _, particle := range e.particles {
		particle.draw()
	}
}

func (e *Emitter) DrawNormal() {
	for _, particle := range e.particles {
		particle.drawNormal()
	}
}

func (e *Emitter) SetSize(width, height float32) {
	e.rect.Width = width
	e.rect.Height = height
}

func (e *Emitter) SetPosition(x, y float32) {
	e.rect.X = x
	e.rect.Y = y
}

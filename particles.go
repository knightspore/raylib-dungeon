package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
	pos      rl.Vector2
	velocity rl.Vector2
	life     float32
}

func (p *Particle) Setup(area rl.Rectangle) {
	// position
	p.pos.X = float32(rl.GetRandomValue(int32(area.X), int32(area.X+area.Width)))
	p.pos.Y = float32(rl.GetRandomValue(int32(area.Y), int32(area.Y+area.Height)))
	// velocity
	p.velocity.X = 0.1 * float32(rl.GetRandomValue(-1, 1))
	p.velocity.Y = 0.1 * float32(rl.GetRandomValue(-1, 1))
	// life
	p.life = float32(rl.GetRandomValue(15, 30))
}

func (p *Particle) Update() {
	p.pos.X += p.velocity.X
	p.pos.Y += p.velocity.Y
	p.velocity.X -= 0.05 * float32(rl.GetRandomValue(-1, 1))
	p.velocity.Y -= 0.05 * float32(rl.GetRandomValue(-1, 1))
	p.life -= rl.GetFrameTime()
}

type Emitter struct {
	rect      rl.Rectangle
	particles map[int]*Particle
}

func NewEmitter(particleCount int, area rl.Rectangle) *Emitter {
	particles := make(map[int]*Particle)
	for i := range particleCount {
		particles[i] = &Particle{}
	}
	return &Emitter{
		rect:      area,
		particles: particles,
	}
}

func (e *Emitter) Setup() {
	for i := range e.particles {
		// old: e.particles[i].Setup(e.rect)

		// to avoid clumping, we'll spread the particles out a bit
		e.particles[i].Setup(rl.NewRectangle(e.rect.X+float32(i), e.rect.Y+float32(i), e.rect.Width, e.rect.Height))
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
		rl.DrawPixel(int32(particle.pos.X), int32(particle.pos.Y), rl.Yellow)
	}
}

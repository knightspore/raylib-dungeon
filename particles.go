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

func (p *Particle) Setup(area rl.Rectangle) {
	// position
	p.pos.X = float32(rl.GetRandomValue(int32(area.X), int32(area.X+area.Width)))
	p.pos.Y = float32(rl.GetRandomValue(int32(area.Y), int32(area.Y+area.Height)))
	// velocity
	p.velocity.X = 0.1 * float32(rl.GetRandomValue(-1, 1))
	p.velocity.Y = 0.1 * float32(rl.GetRandomValue(-1, 1))
	// life
	p.life = float32(rl.GetRandomValue(5, 50))
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
		rect := rl.Rectangle{X: e.rect.X - e.rect.Width, Y: e.rect.Y - e.rect.Height, Width: e.rect.Width * 2, Height: e.rect.Height * 2}
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
	vecs := []rl.Vector2{}
	for _, particle := range e.particles {
		rl.DrawRectanglePro(rl.NewRectangle(particle.pos.X, particle.pos.Y, particle.life/particle.size, particle.life/particle.size), rl.NewVector2(particle.size/2, particle.size/2), 0, rl.Fade(rl.Orange, particle.life/particle.size))
		vecs = append(vecs, particle.pos)
	}

	if DEBUG {
		rl.DrawLineStrip(vecs, rl.NewColor(255, 0, 0, 255))
	}
}

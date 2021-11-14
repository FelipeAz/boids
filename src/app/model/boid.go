package model

import (
	"math/rand"
	"time"
)

type Boid struct {
	id int
	position Vector
	velocity Vector
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position.Add(b.velocity)
	if next.x >= ScreenWidth || next.x < 0 {
		b.velocity = Vector{
			x: -b.velocity.x,
			y: b.velocity.y,
		}
	}

	if next.y >= ScreenHeight || next.y < 0 {
		b.velocity = Vector{
			x: b.velocity.x,
			y: -b.velocity.y,
		}
	}
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func CreateBoid(bid int) {
	b := Boid{
		id:       bid,
		position: Vector{
			x: rand.Float64() * ScreenWidth,
			y: rand.Float64() * ScreenHeight,
		},
		velocity: Vector{
			x: (rand.Float64() * 2) - 1.0,
			y: (rand.Float64() * 2) - 1.0,
		},
	}

	boids[bid] = &b
	go b.start()
}

package model

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	id       int
	position Vector
	velocity Vector
}

func (b *Boid) calcAcceleration() Vector {
	upper, lower := b.position.AddV(ViewRadius), b.position.AddV(-ViewRadius)
	avgVelocity := Vector{
		x: 0,
		y: 0,
	}
	count := 0.0

	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, ScreenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, ScreenHeight); j++ {
			if otherBoidId := BoidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < ViewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
				}
			}
		}
	}

	accel := Vector{
		x: 0,
		y: 0,
	}
	if count > 0 {
		avgVelocity = avgVelocity.DivideV(count)
		accel = avgVelocity.Subtract(b.velocity).MultiplyV(AdjRate)
	}

	return accel
}

func (b *Boid) moveOne() {
	b.velocity = b.velocity.Add(b.calcAcceleration()).limit(-1, 1)
	BoidMap[int(b.position.x)][int(b.position.y)] = -1

	b.position = b.position.Add(b.velocity)
	BoidMap[int(b.position.x)][int(b.position.y)] = b.id

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
		id: bid,
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
	BoidMap[int(b.position.x)][int(b.position.y)] = b.id
	go b.start()
}

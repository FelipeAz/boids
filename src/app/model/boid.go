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
	avgPosition := Vector{
		x: 0,
		y: 0,
	}
	avgVelocity := Vector{
		x: 0,
		y: 0,
	}
	separation := Vector{
		x: 0,
		y: 0,
	}
	count := 0.0

	rWlock.RLock()
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, ScreenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, ScreenHeight); j++ {
			if otherBoidId := BoidMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < ViewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					separation = separation.Add(b.position.Subtract(boids[otherBoidId].position).DivideV(dist))
				}
			}
		}
	}
	rWlock.RUnlock()

	accel := Vector{
		x: b.borderBounce(b.position.x, ScreenWidth),
		y: b.borderBounce(b.position.y, ScreenHeight),
	}
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivideV(count), avgVelocity.DivideV(count)
		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyV(AdjRate)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyV(AdjRate)
		accelSeparation := separation.MultiplyV(AdjRate)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}

	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < ViewRadius {
		return 1 / pos
	}
	if pos > maxBorderPos-ViewRadius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne() {
	acceleration := b.calcAcceleration()
	rWlock.Lock()
	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	BoidMap[int(b.position.x)][int(b.position.y)] = -1

	b.position = b.position.Add(b.velocity)
	BoidMap[int(b.position.x)][int(b.position.y)] = b.id

	rWlock.Unlock()
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

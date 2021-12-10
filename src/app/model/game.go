package model

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenHeight = 360
	ScreenWidth  = 640
	BoidCount    = 500
	ViewRadius   = 13
	AdjRate      = 0.015
)

var (
	green = color.RGBA{
		R: 10,
		G: 255,
		B: 50,
		A: 255,
	}
	boids   [BoidCount]*Boid
	BoidMap [ScreenWidth + 1][ScreenHeight + 1]int
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, boid := range boids {
		screen.Set(int(boid.position.x+1), int(boid.position.y), green)
		screen.Set(int(boid.position.x-1), int(boid.position.y), green)
		screen.Set(int(boid.position.x), int(boid.position.y-1), green)
		screen.Set(int(boid.position.x), int(boid.position.y+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return ScreenWidth, ScreenHeight
}

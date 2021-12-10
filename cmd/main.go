package main

import (
	"log"

	game "github.com/FelipeAz/boid/src/app/model"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	for i, row := range game.BoidMap {
		for j := range row {
			game.BoidMap[i][j] = -1
		}
	}
	for i := 0; i < game.BoidCount; i++ {
		game.CreateBoid(i)
	}
	ebiten.SetWindowSize(game.ScreenWidth*2, game.ScreenHeight*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&game.Game{}); err != nil {
		log.Fatal(err)
	}
}

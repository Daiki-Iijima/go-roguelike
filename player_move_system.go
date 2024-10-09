package main

import "github.com/hajimehoshi/ebiten/v2"

func TryMovePlayer(g *Game) {

	x, y := 0, 0

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x = 1
	}

	players := g.WorldTags["players"]
	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*Position)
		pos.X += x
		pos.Y += y
	}
}

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

	level := g.Map.CurrentLevel

	players := g.WorldTags["players"]
	for _, result := range g.World.Query(players) {
		pos := result.Components[position].(*Position)
		//	移動先のタイルを取得
		nextTile := level.GetIndexFromXY(pos.X+x, pos.Y+y)

		//	移動先のタイルの情報を取得
		tile := level.Tiles[nextTile]

		if tile.Blocked != true {
			pos.X += x
			pos.Y += y
		}
	}
}

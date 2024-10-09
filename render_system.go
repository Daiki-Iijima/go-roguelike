package main

import "github.com/hajimehoshi/ebiten/v2"

func ProcessRenderables(g *Game, level Level, screen *ebiten.Image) {
	//	描画対象になる条件に一致するエンティティを取得
	for _, result := range g.World.Query(g.WorldTags["renderables"]) {
		//	エンティティのコンポーネント情報を取得
		pos := result.Components[position].(*Position)
		img := result.Components[renderable].(*Renderable).Image

		//	タイル座標を計算して取得
		index := level.GetIndexFromXY(pos.X, pos.Y)
		tile := level.Tiles[index]

		//	描画
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
		screen.DrawImage(img, op)
	}
}

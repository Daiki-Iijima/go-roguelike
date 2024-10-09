package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

// Game Gameが必要なデータをすべて持っている
type Game struct {
	Tiles []MapTile
}

// Update 毎フレーム実行される
func (g Game) Update() error {
	return nil
}

// Draw 描画ループ毎に呼ばれ、ここでBlitがおこなわれる
func (g Game) Draw(screen *ebiten.Image) {
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := g.Tiles[GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

// Layout 画面サイズを返す
func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 800
}

func NewGame() *Game {
	g := &Game{}
	g.Tiles = CreateTiles()
	return g
}

func main() {
	g := NewGame()
	//	画面をリサイズできるか
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	//	ウィンドウのタイトル
	ebiten.SetWindowTitle("Tower")

	//	ログ
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

// Game Gameが必要なデータをすべて持っている
type Game struct {
}

// Update 毎フレーム実行される
func (g Game) Update() error {
	return nil
}

// Draw 描画ループ毎に呼ばれ、ここでBlitがおこなわれる
func (g Game) Draw(screen *ebiten.Image) {
}

// Layout 画面サイズを返す
func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 800
}

func NewGame() *Game {
	g := &Game{}
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

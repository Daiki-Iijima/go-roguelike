package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

// Game Gameが必要なデータをすべて持っている
type Game struct {
	Map GameMap
}

// Update 毎フレーム実行される
func (g Game) Update() error {
	return nil
}

// Draw 描画ループ毎に呼ばれ、ここでBlitがおこなわれる
func (g Game) Draw(screen *ebiten.Image) {
	//	Mapを描画
	gd := NewGameData()
	//	とりあえず最初のダンジョンの最初のレベルを表示
	level := g.Map.Dungeons[0].Levels[0]
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

// Layout 画面サイズを返す
func (g Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TIleHeight * gd.ScreenHeight
}

type GameMap struct {
	Dungeons []Dungeon
}

func NewGameMap() GameMap {
	//	レベルの作成
	l := NewLevel()
	levels := make([]Level, 0)
	levels = append(levels, l)

	//	ダンジョンの作成
	d := Dungeon{
		Name:   "default",
		Levels: levels,
	}
	dungeons := make([]Dungeon, 0)
	dungeons = append(dungeons, d)

	//	ゲームマップの作成
	gm := GameMap{Dungeons: dungeons}

	return gm
}

func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
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

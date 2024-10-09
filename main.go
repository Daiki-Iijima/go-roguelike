package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

// Game Gameが必要なデータをすべて持っている
type Game struct {
	Map       GameMap
	World     *ecs.Manager
	WorldTags map[string]ecs.Tag
}

// Update 毎フレーム実行される
func (g *Game) Update() error {
	TryMovePlayer(g)
	return nil
}

// Draw 描画ループ毎に呼ばれ、ここでBlitがおこなわれる
func (g *Game) Draw(screen *ebiten.Image) {
	//	とりあえず最初のダンジョンの最初のレベルを表示
	level := g.Map.Dungeons[0].Levels[0]
	level.DrawLevel(screen)

	ProcessRenderables(g, level, screen)
}

// Layout 画面サイズを返す
func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
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
	world, tags := InitializeWorld()

	g.Map = NewGameMap()
	g.World = world
	g.WorldTags = tags

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

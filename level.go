package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// Level マップタイルの配列を持つ構造体
type Level struct {
	Tiles []MapTile
	Rooms []Rect
}

func NewLevel() Level {
	l := Level{}
	rooms := make([]Rect, 0)
	l.Rooms = rooms
	l.GenerateLevelTiles()
	return l
}

// MapTile 1つのタイルを表現する構造体
type MapTile struct {
	//	左上の座標(X,Y)
	PixelX int
	PixelY int
	//	プレイヤーをブロックするか
	Blocked bool
	//	ebitenイメージへのポインタ
	Image *ebiten.Image
}

// GetIndexFromXY マップタイルの構造を1次元のスライスで表現しているのでスライス上のインデックスをX,Yから取得する
func (level *Level) GetIndexFromXY(x, y int) int {
	gb := NewGameData()
	return (y * gb.ScreenWidth) + x
}

// CreateTiles すべてのタイルを壁として生成する
func (level *Level) CreateTiles() []MapTile {
	gd := NewGameData()
	//	画面のサイズ分のタイル配列を確保
	tiles := make([]MapTile, gd.ScreenWidth*gd.ScreenHeight)

	index := 0

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			index = level.GetIndexFromXY(x, y)
			wall, _, err := ebitenutil.NewImageFromFile("assets/wall.png")
			if err != nil {
				log.Fatal(err)
			}

			tile := MapTile{
				PixelX:  x * gd.TileWidth,
				PixelY:  y * gd.TIleHeight,
				Blocked: true,
				Image:   wall,
			}

			tiles[index] = tile
		}
	}

	return tiles
}

// DrawLevel レベルを描画する
func (level *Level) DrawLevel(screen *ebiten.Image) {
	//	Mapを描画
	gd := NewGameData()
	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			tile := level.Tiles[level.GetIndexFromXY(x, y)]
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.PixelX), float64(tile.PixelY))
			screen.DrawImage(tile.Image, op)
		}
	}
}

func (level *Level) GenerateLevelTiles() {
	MIN_SIZE := 6
	MAX_SIZE := 10
	MAX_ROOMS := 30

	gd := NewGameData()
	//	すべてのタイルを壁にする
	tiles := level.CreateTiles()
	level.Tiles = tiles

	//	部屋の最大数分回す
	for idx := 0; idx < MAX_ROOMS; idx++ {
		w := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		h := GetRandomBetween(MIN_SIZE, MAX_SIZE)
		//	-1しているのは画面の一番端は壁になるので
		x := GetDiceRoll(gd.ScreenWidth - w - 1)
		y := GetDiceRoll(gd.ScreenHeight - h - 1)

		new_room := NewRect(x, y, w, h)
		okToAdd := true

		//	既存のRectと重なるかのチェック
		for _, otherRoom := range level.Rooms {
			if new_room.Intersect(otherRoom) {
				okToAdd = false
				break
			}
		}

		if okToAdd {
			//	Rectの領域を床タイルマップにする
			level.createRoom(new_room)
			//	作成した部屋として保持
			level.Rooms = append(level.Rooms, new_room)
		}
	}
}

// createRoom 指定されたRect構造体の領域をlevelが持っているTiles配列を対象に床で埋める
func (level *Level) createRoom(room Rect) {
	//	四角形のサイズ分ループさせる
	for y := room.Y1 + 1; y < room.Y2; y++ {
		for x := room.X1 + 1; x < room.X2; x++ {
			index := level.GetIndexFromXY(x, y)
			//	一度壁として埋めていた場所を床にするので対象の箇所を書き換える
			level.Tiles[index].Blocked = false
			floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
			if err != nil {
				log.Fatal(err)
			}
			level.Tiles[index].Image = floor
		}
	}
}

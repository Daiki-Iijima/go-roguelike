package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// GameData グローバルゲームデータを保存する構造体
type GameData struct {
	//	画面幅
	ScreenWidth int
	//	画面高さ
	ScreenHeight int
	//	タイル幅
	TileWidth int
	//	タイル高さ
	TIleHeight int
}

// NewGameData GameDataのコンストラクタ関数
func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TIleHeight:   16,
	}

	return g
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
func GetIndexFromXY(x, y int) int {
	gb := NewGameData()
	return (y * gb.ScreenWidth) + x
}

// CreateTiles マップの生成と
func CreateTiles() []MapTile {
	gd := NewGameData()
	//	0個の要素を持ったMapTile型のスライスを生成
	tiles := make([]MapTile, 0)

	for x := 0; x < gd.ScreenWidth; x++ {
		for y := 0; y < gd.ScreenHeight; y++ {
			//	最外周(マップの端)の場合
			if x == 0 || x == gd.ScreenWidth-1 || y == 0 || y == gd.ScreenHeight-1 {
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

				tiles = append(tiles, tile)
			} else {
				floor, _, err := ebitenutil.NewImageFromFile("assets/floor.png")
				if err != nil {
					log.Fatal(err)
				}
				tile := MapTile{
					PixelX:  x * gd.TileWidth,
					PixelY:  y * gd.TIleHeight,
					Blocked: false,
					Image:   floor,
				}
				tiles = append(tiles, tile)
			}
		}
	}

	return tiles
}

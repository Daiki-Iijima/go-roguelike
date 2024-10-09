package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// Level マップタイルの配列を持つ構造体
type Level struct {
	Tiles []MapTile
}

func NewLevel() Level {
	l := Level{}
	tiles := l.CreateTiles()
	l.Tiles = tiles
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

// CreateTiles マップの生成と
func (level *Level) CreateTiles() []MapTile {
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

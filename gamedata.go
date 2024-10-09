package main

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

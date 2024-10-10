package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// すべてのエンティティのコンテナ

func InitializeWorld(startingLevel Level) (*ecs.Manager, map[string]ecs.Tag) {
	tags := make(map[string]ecs.Tag)
	manager := ecs.NewManager()

	player := manager.NewComponent()
	position = manager.NewComponent()
	renderable = manager.NewComponent()
	movable := manager.NewComponent()

	playerImg, _, err := ebitenutil.NewImageFromFile("assets/player.png")
	if err != nil {
		log.Fatal(err)
	}

	//	レベル内に生成された部屋の中心を取得
	startingRoom := startingLevel.Rooms[0]
	x, y := startingRoom.Center()

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: x,
			Y: y,
		})

	//	プレイヤーとして認識されるのはこの2つのコンポーネントを持っているもの
	players := ecs.BuildTag(player, position)

	//	描画可能と認識されるのは、この2つのコンポーネントを持っているもの
	renderables := ecs.BuildTag(renderable, position)

	tags["players"] = players
	tags["renderables"] = renderables

	return manager, tags
}

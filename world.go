package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

// すべてのエンティティのコンテナ

func InitializeWorld() (*ecs.Manager, map[string]ecs.Tag) {
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

	manager.NewEntity().
		AddComponent(player, Player{}).
		AddComponent(renderable, &Renderable{
			Image: playerImg,
		}).
		AddComponent(movable, Movable{}).
		AddComponent(position, &Position{
			X: 40,
			Y: 25,
		})

	//	プレイヤーとして認識されるのはこの2つのコンポーネントを持っているもの
	players := ecs.BuildTag(player, position)

	//	描画可能と認識されるのは、この2つのコンポーネントを持っているもの
	renderables := ecs.BuildTag(renderable, position)

	tags["players"] = players
	tags["renderables"] = renderables

	return manager, tags
}

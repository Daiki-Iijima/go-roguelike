package main

import (
	"github.com/bytearena/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

var position *ecs.Component
var renderable *ecs.Component

type Player struct{}

type Position struct {
	X int
	Y int
}

type Renderable struct {
	Image *ebiten.Image
}

type Movable struct{}

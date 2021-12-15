package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Image    *ebiten.Image
	x, y     int
	position int
}

func NewPlayer(color color.RGBA, x, y int) *Player {
	player := &Player{
		Image: ebiten.NewImage(tileSize, tileSize),
		x:     x,
		y:     y,
	}
	player.Image.Fill(color)
	return player
}

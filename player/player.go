package player

import (
	"image/color"
	"monopoly/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Image    *ebiten.Image
	X, Y     int
	position int
}

func NewPlayer(color color.RGBA, x, y int) *Player {
	player := &Player{
		Image: ebiten.NewImage(constants.TileSize, constants.TileSize),
		X:     x,
		Y:     y,
	}
	player.Image.Fill(color)
	return player
}

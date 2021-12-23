package board

import (
	"image/color"
	"monopoly/constants"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Tile struct {
	X, Y  int
	Field Field
	Image *ebiten.Image
}

func NewTile(color color.RGBA, x, y int, field Field) *Tile {
	tile := &Tile{
		Image: ebiten.NewImage(constants.TileSize, constants.TileSize),
		X:     x,
		Y:     y,
	}
	tile.Image.Fill(color)
	return tile
}

func MakeMatrix(fields []Field, margin int) (out [][]*Tile) {
	k := 0
	numOfFields := len(fields)
	var side = numOfFields/4 + numOfFields%4
	var length = side + margin*2
	for i := 0; i <= length; i++ {
		out = append(out, []*Tile{})
		for j := 0; j <= length; j++ {
			if ((i == side+margin || i == margin) && (j <= side+margin && j >= margin)) ||
				((j == side+margin || j == margin) && (i <= side+margin && i >= margin)) {
				out[i] = append(out[i], NewTile(colornames.Blue, j, i, fields[k]))
				k++
			} else {
				out[i] = append(out[i], NewTile(colornames.Green, j, i, nil))
			}
		}
	}
	return
}

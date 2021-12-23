package board

import (
	"image/color"
	"monopoly/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type Field interface {
	Title() string
	LandOn(player *Player) error
}

type ChanceField struct {
	title string
}

func (this *ChanceField) LandOn(player *Player) error { return nil }
func (this *ChanceField) Title() string               { return "" }

type OwnableField struct {
	title string
	group int
	price int
}

func (this *OwnableField) LandOn(player *Player) error { return nil }
func (this *OwnableField) Title() string               { return "" }

type TaxField struct {
	title      string
	percentage int
}

func (this *TaxField) LandOn(player *Player) error { return nil }
func (this *TaxField) Title() string               { return "" }

type ParkingField struct {
	title string
	price int
}

func (this *ParkingField) LandOn(player *Player) error { return nil }
func (this *ParkingField) Title() string               { return "" }

type VisitField struct {
	title string
}

func (this *VisitField) LandOn(player *Player) error { return nil }
func (this *VisitField) Title() string               { return "" }

type PrisonField struct {
	title string
}

func (this *PrisonField) LandOn(player *Player) error { return nil }
func (this *PrisonField) Title() string               { return "" }

type StartField struct {
	title string
}

func (this *StartField) LandOn(player *Player) error { return nil }
func (this *StartField) Title() string               { return "" }

type Player struct {
	Image    *ebiten.Image
	x, y     int
	position int
}

func NewPlayer(color color.RGBA, x, y int) *Player {
	player := &Player{
		Image: ebiten.NewImage(constants.TileSize, constants.TileSize),
		x:     x,
		y:     y,
	}
	player.Image.Fill(color)
	return player
}

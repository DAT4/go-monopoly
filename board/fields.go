package board

import (
	"image/color"
	"monopoly/player"

	"golang.org/x/image/colornames"
)

type Field interface {
	Title() string
	LandOn(player *player.Player) error
	Color() color.RGBA
}

func NewChanceField() *ChanceField {
	return &ChanceField{
		title: "?",
		color: colornames.Black,
	}
}

type ChanceField struct {
	title string
	color color.RGBA
}

func (this *ChanceField) LandOn(player *player.Player) error { return nil }
func (this *ChanceField) Title() string                      { return "" }
func (this *ChanceField) Color() color.RGBA                  { return this.color }

func NewOwnableField(title string, price int, group int, color color.RGBA) *OwnableField {
	return &OwnableField{
		title: title,
		price: price,
		color: color,
		group: group,
	}
}

type OwnableField struct {
	title string
	group int
	price int
	color color.RGBA
}

func (this *OwnableField) LandOn(player *player.Player) error { return nil }
func (this *OwnableField) Title() string                      { return "" }
func (this *OwnableField) Color() color.RGBA                  { return this.color }

func NewTaxField(price int) *TaxField {
	return &TaxField{
		title:      "TaxField",
		percentage: price,
		color:      colornames.White,
	}
}

type TaxField struct {
	title      string
	percentage int
	color      color.RGBA
}

func (this *TaxField) LandOn(player *player.Player) error { return nil }
func (this *TaxField) Title() string                      { return "" }
func (this *TaxField) Color() color.RGBA                  { return this.color }

func NewParkingField() *ParkingField {
	return &ParkingField{
		title: "ParkingField",
		price: 0,
		color: colornames.White,
	}
}

type ParkingField struct {
	title string
	price int
	color color.RGBA
}

func (this *ParkingField) LandOn(player *player.Player) error { return nil }
func (this *ParkingField) Title() string                      { return "" }
func (this *ParkingField) Color() color.RGBA                  { return this.color }

func NewVisitField() *VisitField {
	return &VisitField{
		title: "VisitField",
		color: colornames.Black,
	}
}

type VisitField struct {
	title string
	color color.RGBA
}

func (this *VisitField) LandOn(player *player.Player) error { return nil }
func (this *VisitField) Title() string                      { return "" }
func (this *VisitField) Color() color.RGBA                  { return this.color }

func NewPrisonField() *PrisonField {
	return &PrisonField{
		title: "PrisonField",
		color: colornames.Black,
	}
}

type PrisonField struct {
	title string
	color color.RGBA
}

func (this *PrisonField) LandOn(player *player.Player) error { return nil }
func (this *PrisonField) Title() string                      { return "" }
func (this *PrisonField) Color() color.RGBA                  { return this.color }

func NewStartField() *StartField {
	return &StartField{
		title: "StartField",
		color: colornames.Black,
	}
}

type StartField struct {
	title string
	color color.RGBA
}

func (this *StartField) LandOn(player *player.Player) error { return nil }
func (this *StartField) Title() string                      { return "" }
func (this *StartField) Color() color.RGBA                  { return this.color }

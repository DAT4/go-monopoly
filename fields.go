package main

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

var Fields = []Field{
	&StartField{},
	&OwnableField{},
	&ChanceField{},
	&OwnableField{},
	&TaxField{},
	&OwnableField{},
	&OwnableField{},
	&ChanceField{},
	&OwnableField{},
	&OwnableField{},
	&VisitField{},
	&OwnableField{},
	&OwnableField{},
	&OwnableField{},
	&OwnableField{},
	&OwnableField{},
	&OwnableField{},
	&ChanceField{},
	&OwnableField{},
	&OwnableField{},
	&ParkingField{},
	&OwnableField{},
	&ChanceField{},
	&OwnableField{},
	&OwnableField{},
	&OwnableField{},
	&OwnableField{},
	&OwnableField{},
	&OwnableField{},
	&OwnableField{},
	&PrisonField{},
	&OwnableField{},
	&OwnableField{},
	&ChanceField{},
	&OwnableField{},
	&OwnableField{},
	&ChanceField{},
	&OwnableField{},
	&TaxField{},
	&OwnableField{},
}
